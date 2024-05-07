# Cartier Eslastic & Kibana

**Server**

<u>Server name:</u> **[elastic-search-vm-1-https](https://console.cloud.google.com/compute/instancesDetail/zones/us-central1-a/instances/elastic-search-vm-1-https?project=cartier-project&orgonly=true&supportedpurview=organizationId)**

<u>IP Address:</u> **34.28.95.74**

<u>APP SOURCE PATH:</u>

- Docker-compose (Kibana & Elastic): ***"/usr/src/cartier/apps"***
- Docker-compose (Postainer & Traefik): **"/usr/src/cartier/portainer+traefik/core"**
- Docker Volume (es_data): ***"/mnt/disks/disk1/docker/volumes/cartier_es_data/"***

**Temporary DNS**

https://elastic.carti3r.tk

https://kibana.carti3r.tk

https://portainer.carti3r.tk

https://traefik.carti3r.tk

### New deploy steps

Connect to VM by ssh.

```bash
# Get sudo permissions
sudo -i
```

Create a work directory and go in

```bash
# Create directory
mkdir /usr/src/cartier
# Change the current working directory
cd /usr/src/cartier
```

Clone this private repository with preset-ed tools.

```bash
# Clone repository
git clone git@github.com:RMadVar/CrtDeployElasticKibana.git ./portainer+traefik
```

*This repository have all configuration files ready to start but I will share with you a basic instructions to modify it*.

This is a tree of repository.

```apl
.
└── portainer+traefik/
    ├── core/
    │   ├── traefik-data/
    │   │   ├── configurations/
    │   │   │   └── dynamic.yml
    │   │   ├── traefik.yml
    │   │   └── acme.json
    │   └── docker-compose.yml
    └── apps/docker-compose.yml
```

We need to change permission for "acme.json"

```
cd /usr/src/cartier/portainer+traefik/core/traefik-data
chmod 600 acme.json
```

Traefik.yml control the entrypoint and SSL.

```yaml
# traefik.yml
api:
  dashboard: true

entryPoints:
  web:
    address: :80
    http:
      redirections:
        entryPoint:
          to: websecure

  websecure:
    address: :443
    http:
      middlewares:
        - secureHeaders@file
      tls:
        certResolver: letsencrypt

providers:
  docker:
    endpoint: "unix:///var/run/docker.sock"
    exposedByDefault: false
  file:
    filename: /configurations/dynamic.yml

certificatesResolvers:
  letsencrypt:
    acme:
      email: <your_email> <------------------------------------- #Change it
      storage: acme.json
      keyType: EC384
      httpChallenge:
        entryPoint: web
```

#### 

dynamic.yml This file contains our middlewares to make sure all our traffic is fully secure and runs over TLS. We also set the basic auth here for our Traefik dashboard, because by default it is accessible for everyone.

```yaml
# dynamic.yml
http:
  middlewares:
    secureHeaders:
      headers:
        sslRedirect: true
        forceSTSHeader: true
        stsIncludeSubdomains: true
        stsPreload: true
        stsSeconds: 31536000

    user-auth:
      basicAuth:
        users:
          - "raf:$apr1$MTqfVwiE$FKkzT5ERGFqwH9f3uipxA1" <------------- # Generate user and password with
          # "echo $(htpasswd -nb <USER> <PASSWORD>) | sed -e s/\\$/\\$\\$/g"

tls:
  options:
    default:
      cipherSuites:
        - TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
        - TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
        - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
        - TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
        - TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305
        - TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
      minVersion: VersionTLS12
```

core/docker-compose.yml

```yaml
# docker-compose.yml
version: "3"

services:
  traefik:
    image: traefik:latest
    container_name: traefik
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    networks:
      - proxy
    ports:
      - 80:80
      - 443:443
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./traefik-data/traefik.yml:/traefik.yml:ro
      - ./traefik-data/acme.json:/acme.json
      - ./traefik-data/configurations:/configurations
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.traefik-secure.entrypoints=websecure"
      - "traefik.http.routers.traefik-secure.rule=Host(`traefik.yourdomain.com`)"
      - "traefik.http.routers.traefik-secure.service=api@internal"
      - "traefik.http.routers.traefik-secure.middlewares=user-auth@file"

  portainer:
    image: portainer/portainer-ce:latest
    container_name: portainer
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    networks:
      - proxy
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./portainer-data:/data
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.portainer-secure.entrypoints=websecure"
      - "traefik.http.routers.portainer-secure.rule=Host(`portainer.yourdomain.com`)"
      - "traefik.http.routers.portainer-secure.service=portainer"
      - "traefik.http.services.portainer.loadbalancer.server.port=9000"

networks:
  proxy:
    external: true
```

apps/docker-compose.yml

```yaml
version: "3.9"
services:
  elasticsearch:
    image: elasticsearch:8.2.2
    environment:
      - discovery.type=single-node
      - ES_JAVA_OPTS=-Xms1g -Xmx1g
      - xpack.security.enabled=true
      - ELASTIC_PASSWORD=cartier419
    volumes:
      - es_data:/usr/share/elasticsearch/data
        #ports:
        #- target: 9200
        #published: 9200
    expose:
      - 9200
    networks:
      - elastic
      - proxy
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.elastic.entrypoints=websecure"
      - "traefik.http.routers.elastic.rule=Host(`elastic.carti3r.tk`)"
      - "traefik.http.services.elastic.loadbalancer.server.port=9200"
  kibana:
    image: kibana:8.2.2
    environment:
      - ELASTICSEARCH_URL=http://0.0.0.0:9200/ 
      - ELASTICSEARCH_USERNAME=alg_kibana
      - ELASTICSEARCH_PASSWORD=cartier419
        #ports:
        #- target: 5601
        #published: 5601
    expose:
      - 5601
    depends_on:
      - elasticsearch
    networks:
      - elastic      
      - proxy
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.kibana.entrypoints=websecure"
      - "traefik.http.routers.kibana.rule=Host(`kibana.carti3r.tk`)"
      - "traefik.http.services.kibana.loadbalancer.server.port=5601"
volumes:
  es_data:
    driver: local

networks:
  elastic:
    name: elastic
    driver: bridge
  proxy:
      external: true
```

Create this network

```
docker network create proxy
```

Go to ***"/usr/src/cartier/portainer+traefik/core"***

```
# Move to the work directory and execute docker compose command.
cd /usr/src/cartier/portainer+traefik/core
# Start Portainer + Traefik
docker-compose up -d
```

Now we are going to start Kibana & Elastic

```
# Move to the work directory and execute docker compose command.
cd /usr/src/cartier/portainer+traefik/apps
# Start Portainer + Traefik
docker-compose up -d
```

You can check "TUTORIAL.md" for more details.