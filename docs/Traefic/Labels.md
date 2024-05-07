```yaml
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.app-secure.entrypoints=websecure"
      - "traefik.http.routers.app-secure.rule=Host(`bitsecret.duckdns.org`)"
      - "traefik.http.routers.nginx-secure.service=api@internal"
      - "traefik.http.services.nginx.loadbalancer.server.port=8080"
      - "traefik.http.routers.nginx.service=nginx"


```

```
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.app-secure.entrypoints=websecure"
      - "traefik.http.routers.app-secure.rule=Host(`bitsecret.duckdns.org`)"
      - "traefik.http.routers.nginx-secure.service=api@internal"
      - "traefik.http.services.nginx.loadbalancer.server.port=8080"
      - "traefik.http.routers.nginx.service=nginx"

```

```
  sonarqube:
    image: docker.io/bitnami/sonarqube:9
    volumes:
      - 'sonarqube_data:/bitnami/sonarqube'
    depends_on:
      - postgresql
    environment:
      # ALLOW_EMPTY_PASSWORD is recommended only for development.
      - ALLOW_EMPTY_PASSWORD=yes
      - SONARQUBE_DATABASE_HOST=postgresql
      - SONARQUBE_DATABASE_PORT_NUMBER=5432
      - SONARQUBE_DATABASE_USER=bn_sonarqube
      - SONARQUBE_DATABASE_NAME=bitnami_sonarqube
      - SONARQUBE_PORT_NUMBER=4200
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.sonar.entrypoints=websecure"
      - "traefik.http.routers.sonar.rule=Host(`sqmybts.duckdns.org`)"
      - "traefik.http.services.sonar.loadbalancer.server.port=4200"
    networks:
      - public
      - proxy
    expose:
      - 4200
networks:
    public:
    proxy:
      external: true
```

