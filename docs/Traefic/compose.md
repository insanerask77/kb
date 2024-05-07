```
version: '3'

# echo $(htpasswd -nb <USER> <PASSWORD>) | sed -e s/\\$/\\$\\$/g

services:
  whoami:
    image: "containous/whoami"
    container_name: "simple-service"
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.whoami.rule=Host(`quiensoy.duckdns.org`)"
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
      - "traefik.http.routers.traefik.rule=Host(`traeffic.duckdns.org`)"
      - "traefik.http.middlewares.traefik-auth.basicauth.users=admin:$$apr1$$jQNmMH0x$$Gop4wvcDcDE3NNynkb5iZ/"
      - "traefik.http.middlewares.test-auth.basicauth.users=user:$$apr1$$qv5TEQ/X$$h9J43URfmFG0E65egLR4S1"
      - "traefik.http.middlewares.traefik-https-redirect.redirectscheme.scheme=https"
      - "traefik.http.routers.traefik.middlewares=traefik-https-redirect"
      - "traefik.http.routers.traefik-secure.entrypoints=https"
      - "traefik.http.routers.traefik-secure.rule=Host(`traeffic.duckdns.org`)"
      - "traefik.http.routers.traefik-secure.middlewares=traefik-auth"
      - "traefik.http.routers.traefik-secure.tls=true"
      - "traefik.http.routers.traefik-secure.tls.certresolver=http"
      - "traefik.http.routers.traefik-secure.service=api@internal"

networks:
  traefik-proxy:
    external: true


```

