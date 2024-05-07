```
version: "3.9"

services:
  caddy:
    image: caddy
    container_name: caddy
    restart: always
    ports:
      - "80:80"
      - "443:443"
    networks:
      - caddy
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - caddy_data:/data
      - caddy_config:/config
volumes:
  caddy_data:
  caddy_config:
networks:
  caddy:
    external: true

```

```
#Caddyfile

mongo.bluetrailsoft.tk {
    reverse_proxy mongo_express:8081
}

portainer.bluetrailsoft.tk {
    reverse_proxy portainer:9000
}

test.bluetrailsoftutils {
    reverse_proxy netdata:19999
}


```

