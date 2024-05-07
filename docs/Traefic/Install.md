

## Traefic



https://blog.elhacker.net/2022/01/traefik-un-proxy-inverso-para-contenedores-docker.html

https://doc.traefik.io/traefik/v2.0/middlewares/basicauth/

```yaml
version: '3'

# echo $(htpasswd -nb <USER> <PASSWORD>) | sed -e s/\\$/\\$\\$/g

services:
  whoami:
    image: "containous/whoami"
    container_name: "simple-service"
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.whoami.rule=Host(`whoami.localhost`)"
      - "traefik.http.routers.whoami.entrypoints=http"
  traefik:
    image: traefik:v2.4.2
    container_name: traefik
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    networks:
      - traefik-proxy
    ports:
      - 80:80
      - 443:443
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./traefik-data/traefik.yml:/traefik.yml:ro
      - ./traefik-data/acme.json:/acme.json
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.traefik.entrypoints=http"
      - "traefik.http.routers.traefik.rule=Host(`localhost`)"
      - "traefik.http.middlewares.traefik-auth.basicauth.users=user:$$apr1$$dd0JhG/Z$$tvNh9gRhT3F/FNtyC3dPp/"
      - "traefik.http.middlewares.test-auth.basicauth.users=user:$$apr1$$dd0JhG/Z$$tvNh9gRhT3F/FNtyC3dPp/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
      - "traefik.http.middlewares.traefik-https-redirect.redirectscheme.scheme=https"
      - "traefik.http.routers.traefik.middlewares=traefik-https-redirect"
      - "traefik.http.routers.traefik-secure.entrypoints=https"
      - "traefik.http.routers.traefik-secure.rule=Host(`localhost`)"
      - "traefik.http.routers.traefik-secure.middlewares=traefik-auth"
      - "traefik.http.routers.traefik-secure.tls=true"
      - "traefik.http.routers.traefik-secure.tls.certresolver=http"
      - "traefik.http.routers.traefik-secure.service=api@internal"

networks:
  traefik-proxy:
    external: true
```

