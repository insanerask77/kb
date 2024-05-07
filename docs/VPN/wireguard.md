### Wireguard

```
systemctl disable systemd-resolved.service
systemctl stop systemd-resolved
```

##### Compose with PiHole

```
version: '3.7'

services:
  wireguard:
    image: linuxserver/wireguard
    container_name: wireguard
    cap_add:
      - NET_ADMIN
      - SYS_MODULE
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=America/Mendoza
      - SERVERPORT=51820 #optional
      - PEERS=2 #optional
      - PEERDNS=auto #optional
      - INTERNAL_SUBNET=10.13.13.0 #optional
    volumes:
      - /root/wireguard:/config
      - /lib/modules:/lib/modules
      - /usr/src:/usr/src
    ports:
      - 51820:51820/udp
    sysctls:
      - net.ipv4.conf.all.src_valid_mark=1
    dns:
      - 172.20.0.7
    restart: unless-stopped
    networks:
      containers:
        ipv4_address: 172.20.0.6

  pihole:
    container_name: pihole
    image: pihole/pihole:v5.7
    expose:
      - "53"
      - "67"
      - "80"
      - "443"
    environment:
      TZ: 'America/Mendoza'
      WEBPASSWORD: 'peladonerd'
    volumes:
      - './etc-pihole/:/etc/pihole/'
      - './etc-dnsmasq.d/:/etc/dnsmasq.d/'
    cap_add:
      - NET_ADMIN
    restart: unless-stopped
    networks:
      containers:
        ipv4_address: 172.20.0.7

networks:
  containers:
    ipam:
      config:
        - subnet: 172.20.0.0/24
```

##### Compose Wireguard

```
version: '3.3'
services:
  wireguard:
    image: linuxserver/wireguard
    container_name: wireguard
    cap_add:
      - NET_ADMIN
      - SYS_MODULE
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/Madrid
      - SERVERURL=1.1.1.1 #optional
      - SERVERPORT=51820 #optional
      - PEERS=2 #optional
      - PEERDNS=auto #optional
      - INTERNAL_SUBNET=10.13.13.0 #optional
    volumes:
      - /path/to/appdata/config:/config
      - /lib/modules:/lib/modules
      - /usr/src:/usr/src
    ports:
      - 51820:51820/udp
    sysctls:
      - net.ipv4.conf.all.src_valid_mark=1
    restart: unless-stopped
```



##### Docker Command

```
docker run -d \
  --name=wireguard \
  --cap-add=NET_ADMIN \
  --cap-add=SYS_MODULE \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Europe/London \
  -e SERVERPORT=51820 `#optional` \
  -e PEERS=4 `#optional` \
  -e PEERDNS=auto `#optional` \
  -e INTERNAL_SUBNET=10.13.13.0 `#optional` \
  -e ALLOWEDIPS=0.0.0.0/0 `#optional` \
  -e LOG_CONFS=true `#optional` \
  -p 51820:51820/udp \
  -v /path/to/appdata/config:/config \
  -v /lib/modules:/lib/modules `#optional` \
  --sysctl="net.ipv4.conf.all.src_valid_mark=1" \
  --restart unless-stopped \
  linuxserver/wireguard
```

##### Install in linux



***Nota\*:** *Se tomó como ejemplo Ubuntu 19.10.*

**1.** En primer lugar, debe crear el archivo de configuración WireGuard® en su [oficina del usuario](https://my.keepsolid.com/products/vpn/). Para hacer esto, siga las instrucciones descritas en [este manual](https://www.vpnunlimited.com/es/help/manuals/how-to-manually-create-vpn-conf).

**2.** Cree el repositorio WireGuard®:

```
sudo add-apt-repository ppa:wireguard/wireguard
```

**3.** Instale los paquetes WireGuard®:

```
sudo apt install wireguard
```

Instale el paquete resolv.conf:

```
sudo apt install resolvconf
```

**4.** Vaya al directorio WireGuard® y cree el archivo wg0.conf:

```
cd /etc/wireguard
```

**5.** Copie la configuracion de WireGuard® que recibió en su Oficina del usuario y péguelas en el archivo **wg0.conf** usando su editor de texto:

```
nano wg0.conf
```

**6.** Encienda su conexión WireGuard® y disfrute de una navegación web rápida y fiable:

```
systemctl start wg-quick@wg0
```

**7.** Si desea mantener su conexión WireGuard® activa desde el inicio del sistema, imprima el siguiente comando:

```
systemctl enable wg-quick@wg0
```

**8.** Apague la conexión WireGuard® usando el comando:

```
systemctl stop wg-quick@wg0
```

**9.** Si desea desactivar el inicio automático, use el comando:

```
systemctl disable wg-quick@wg0
```

```
sudo wg-quick down wg0
systemctl restart wg-quick@wg0.service
wg show
```


```version: '3.3'
services:
    wg-easy:
        container_name: wg-easy
        environment:
            - WG_HOST=🚨YOUR_SERVER_IP
            - PASSWORD=🚨YOUR_ADMIN_PASSWORD
        volumes:
            - '~/.wg-easy:/etc/wireguard'
        ports:
            - '51820:51820/udp'
            - '51821:51821/tcp'
        restart: unless-stopped
        image: weejewel/wg-easy

```

```
wireguard:
    image: linuxserver/wireguard
    container_name: alg_wireguard
    cap_add:
      - NET_ADMIN
      - SYS_MODULE
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/Madrid
      - SERVERURL=1.1.1.1 #optional
      - SERVERPORT=51820 #optional
      - PEERDNS=auto #optional
      - INTERNAL_SUBNET=10.13.13.0 #optional
    volumes:
      - ./config/wireguard:/config
      - /lib/modules:/lib/modules
      - /usr/src:/usr/src
    network_mode: host
    restart: unless-stopped
```

### wg-easy

```
version: "3.8"
services:
  wg-easy:
    environment:
      # ⚠️ Required:
      # Change this to your host's public address
      - WG_HOST= 
      # Optional:
      - PASSWORD= [ADMIN PASSWORD]
      # - WG_PORT=51820
      # - WG_DEFAULT_ADDRESS= 10.8.0.x
      # - WG_DEFAULT_DNS= 
      # - WG_MTU=1420
      - WG_ALLOWED_IPS= [SUBNET]/[MASK]
      # - WG_PRE_UP=echo "Pre Up" > /etc/wireguard/pre-up.txt
      # - WG_POST_UP=echo "Post Up" > /etc/wireguard/post-up.txt
      # - WG_PRE_DOWN=echo "Pre Down" > /etc/wireguard/pre-down.txt
      # - WG_POST_DOWN=echo "Post Down" > /etc/wireguard/post-down.txt
    image: weejewel/wg-easy
    container_name: wg-easy
    volumes:
      - /storage/path/ofyour/choice:/etc/wireguard
    ports:
      - "51820:51820/udp"
      - "51821:51821/tcp"
    restart: unless-stopped
    cap_add:
      - NET_ADMIN
      - SYS_MODULE
    sysctls:
      - net.ipv4.ip_forward=1
      - net.ipv4.conf.all.src_valid_mark=1
```

```
version: "3.8"
services:
  wg-easy:
    environment:
      # ⚠️ Required:
      # Change this to your host's public address
      - WG_HOST=35.239.20.88

      # Optional:
      - PASSWORD=
      - WG_PORT=51820
      - WG_DEFAULT_ADDRESS=10.13.13.x
      - WG_DEFAULT_DNS=1.1.1.1
      # - WG_MTU=1420
      - WG_PERSISTENT_KEEPALIVE=25
      - WG_ALLOWED_IPS=10.13.13.0/24
      # - WG_PRE_UP=echo "Pre Up" > /etc/wireguard/pre-up.txt
      # - WG_POST_UP=echo "Post Up" > /etc/wireguard/post-up.txt
      # - WG_PRE_DOWN=echo "Pre Down" > /etc/wireguard/pre-down.txt
      # - WG_POST_DOWN=echo "Post Down" > /etc/wireguard/post-down.txt
    #image: weejewel/wg-easy:latest
    image: weejewel/wg-easy:7-nightly
    container_name: wg-easy
    volumes:
      - ./:/etc/wireguard
    ports:
      - "51820:51820/udp"
      - "51821:51821/tcp"
    restart: unless-stopped
    cap_add:
      - NET_ADMIN
      - SYS_MODULE
    sysctls:
      - net.ipv4.ip_forward=1
      - net.ipv4.conf.all.src_valid_mark=1
```

