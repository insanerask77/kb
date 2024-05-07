```
version: "3"
services:
  docker-registry:
    image: registry:2
    container_name: docker-registry
    volumes:
      - ./auth:/auth
    environment:
      - REGISTRY_STORAGE_DELETE_ENABLED=TRUE
      #- REGISTRY_AUTH=htpasswd
      #- REGISTRY_AUTH_HTPASSWD_REALM="Registry Realm"
      #- REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd
    ports:
      - 5000:5000
    restart: always
    volumes:
      - ./volume:/var/lib/registry

  docker-registry-ui:
    image: konradkleine/docker-registry-frontend:v2
    container_name: docker-registry-ui
    ports:
      - 9088:80
    environment:
      - ENV_DOCKER_REGISTRY_HOST=io.madolell.tk
      #- ENV_DOCKER_REGISTRY_PORT=5000

```

```
htpasswd -c -B -b </path/to/users.htpasswd> <user_name> <password>
```

