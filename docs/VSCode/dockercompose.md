```yaml
---
version: "2.1"
services:
  code-server:
    image: lscr.io/linuxserver/code-server:latest
    container_name: code-server
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
      - PASSWORD=Madolell1.. #optional
      - HASHED_PASSWORD= #optional
      - SUDO_PASSWORD=Madolell1.. #optional
      - SUDO_PASSWORD_HASH= #optional
      - PROXY_DOMAIN=code-server.my.domain #optional
      - DEFAULT_WORKSPACE=/config/workspace #optional
    volumes:
      - /path/to/appdata/config:/config
      - /root/src/core/nginx-data/web:/mnt/mydata
    expose:
      - 8443
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.code.entrypoints=websecure"
      - "traefik.http.routers.code.rule=Host(`code.madolell.tk`)"

    restart: unless-stopped
    networks:
      - public
      - proxy
networks:
    public:
    proxy:
      external: true
```

