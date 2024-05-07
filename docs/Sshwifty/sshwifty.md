

#### Docker Compose

```yaml
version: "2"
services:
  ssh:
    image: niruix/sshwifty
    container_name: sshwifty
    volumes:
      - postgres:/config
    networks:
      - public
      - proxy
    environment:

      - SSHWIFTY_SOCKS5_USER=Rafa
      - SSHWIFTY_SOCKS5_PASSWORD="12345"
      - SSHWIFTY_SHAREDKEY=Madolell1..
    expose:
      - 8182
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.gua.entrypoints=websecure"
      - "traefik.http.routers.gua.rule=Host(`cv.madolell.tk`)"
networks:
    public:
    proxy:
      external: true
volumes:
  postgres:
    driver: local
```



