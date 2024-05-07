setup.sh

```bash
#!/usr/bin/env bash

set -e

mkdir -p data
touch data/.runner
mkdir -p data/.cache

chown -R 1001:1001 data/.runner
chown -R 1001:1001 data/.cache
chmod 775 data/.runner
chmod 775 data/.cache
chmod g+s data/.runner
chmod g+s data/.cache
```

docker-compose.yml

```yaml
version: '3'

networks:
  forgejo:
    external: false

services:
  forgejo:
    image: codeberg.org/forgejo/forgejo:7
    container_name: forgejo
    networks:
      - forgejo
    environment:
      - USER_UID=1000
      - USER_GID=1000
    restart: always
    volumes:
      - ./forgejo:/data
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
    ports:
      - '3000:3000'
      - '222:22'
  docker-in-docker:
    image: docker:dind
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - forgejo
    container_name: 'docker_dind'
    privileged: 'true'
    command: ['dockerd', '-H', 'tcp://0.0.0.0:2375', '--tls=false']
    restart: 'unless-stopped'

  runner:
    image: 'code.forgejo.org/forgejo/runner:3.4.1'
    links:
      - docker-in-docker
    depends_on:
      docker-in-docker:
        condition: service_started
    networks:
      - forgejo
    container_name: 'runner'
    environment:
      DOCKER_HOST: tcp://docker-in-docker:2375
    # User without root privileges, but with access to `./data`.
    user: 1001:1001
    volumes:
      - ./data:/data
      - /var/run/docker.sock:/var/run/docker.sock
    restart: 'unless-stopped'
    command: '/bin/sh -c "while : ; do sleep 1 ; done ;"'
    # command: '/bin/sh -c "sleep 5; forgejo-runner daemon"'
```

```
docker exec -it --user=root runner /bin/bash
apk add docker
```

```bash
docker exec -it runner /bin/bash
```

```bash

forgejo-runner register
or
forgejo-runner register --no-interactive --token {TOKEN} --name runner --instance https://next.forgejo.org
ubuntu-22.04:docker://ghcr.io/catthehacker/ubuntu:act-22.04
ubuntu-latest:docker://ghcr.io/catthehacker/ubuntu:act-22.04
```

```yaml
command: '/bin/sh -c "sleep 5; forgejo-runner daemon"'
```

#### Only Runner

```yaml
version: '3.8'

networks:
  forgejo:
    driver: bridge
    ipam:
     config:
       - subnet: 10.5.0.0/16
         gateway: 10.5.0.1
services:
  docker-in-docker:
    image: docker:dind
    container_name: 'docker_dind'
    privileged: true
    command: [ "dockerd", "-H", "tcp://0.0.0.0:2375", "--tls=false" ]
    restart: 'unless-stopped'
    networks:
      forgejo:
        ipv4_address: 10.5.0.2

  gitea:
    image: 'code.forgejo.org/forgejo/runner:3.3.0'
    links:
      - docker-in-docker
    depends_on:
      docker-in-docker:
        condition: service_started
    container_name: 'runner'
    environment:
      DOCKER_HOST: tcp://docker-in-docker:2375
    # A user without root privileges, but with access to `./data`.
    user: 1001:1001
    volumes:
      - ./data:/data
    restart: 'unless-stopped'
    networks:
      forgejo:
        ipv4_address: 10.5.0.3
    command: '/bin/sh -c "sleep 5; forgejo-runner daemon"'
```

#### Example Action

```yaml
name: build

on:
  workflow_dispatch:
  push:
    branches:
      - 'main'


jobs:
  docker:
    runs-on: ubuntu-22.04
    env:
      DOCKER_HOST: tcp://172.18.0.3:2375
    steps:
      -
        name: find ip
        run: |
          docker ps
          docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' docker_dind          
      -
        name: Set up QEMU
        uses: docker/setup-qemu-action@v3
      -
        name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      -
        name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          registry: forgejo.madolell.com
          username: ${{ secrets.DOCKER_USER }}
          password: ${{ secrets.DOCKER_TOKEN }}
      - 
        name: Pull repository
        uses: actions/checkout@v4
      -
        name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: forgejo.madolell.com/rafa/test:latest
```

