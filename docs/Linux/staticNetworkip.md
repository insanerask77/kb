# Configure Ubuntu WiFi Adapter with Netplan

[#linux](https://dev.to/t/linux)[#ubuntu](https://dev.to/t/ubuntu)

## [Ubuntu Networking with Netplan (2 Part Series)](https://dev.to/joeneville_/series/14554)

[1Configure Ubuntu Networking with Netplan](https://dev.to/joeneville_/configure-ubuntu-networking-with-netplan-1cbc)[2Configure Ubuntu WiFi Adapter with Netplan](https://dev.to/joeneville_/configure-ubuntu-wifi-with-netplan-4je0)

👉 Here's a short explanation of how to configure an Ubuntu machine to join a wireless network, with Netplan.
👉 This is for a wireless network with WPA2 Personal authentication (you need a password).
👉 My test machine is running Ubuntu desktop 21.04.

## 1. Gather Required Details

1. Get the wireless network details.
   This is WPA2 Personal, you're going to need the usual details:

   - SSID
   - Wireless network password

2. Get your Ubuntu machine's wireless adapter name.
   You can use `ip link` or `ip add` for this.

   ```
   joe@ub1:~$ ip add
   ...
   3: wlx18d6c7116805: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
   ```

In this example, my adapter is `wlx18d6c7116805`.

## 2. Compose the Netplan file

1. Create a netplan yaml file in the /etc/netplan directory.

- In this example my file is called `mynet1.yaml`.
- Configure the wireless adapter details under `wifis`.
- The SSID for your wireless network (the name of the network) is configured under `access-points`.
- Record the wireless network password under the SSID.
- Make sure to configure any other networking adapters that you require as well, see my previous blog for examples of wired networks [here.](https://dev.to/joeneville_/configure-ubuntu-networking-with-netplan-1cbc)

*Example*

```
network:
  ethernets:
    eno1:
      addresses:
        - 10.150.15.25/24
  wifis:
    wlx18d6c7116805:
      dhcp4: yes
      dhcp6: yes
      access-points:
        "IDontLikeSand15":
          password: "Supersecure123"
  version: 2
  renderer: NetworkManager
```



*Example notes*

- The wired ethernet adapter `eno1` is configured with a static IPv4 address.
- The wireless adapter `wlx18d6c7116805` is configured for DHCPv4 and DHCPv6 address allocation.
- The SSID is `IDontLikeSand15` with a password of `Supersecure123`.

## 3. Apply the Netplan file

Run the command `netplan apply <your-netplan-file-name>`.

## 4. Verify

- Use the command `iwconfig` to check the wireless adapter state.
- Use `ip add` to view all your network adapter.

```
$ iwconfig
lo        no wireless extensions.

eno1      no wireless extensions.

wlx18d6c7116805  IEEE 802.11bgn  ESSID:"IDontLikeSand15"  Nickname:"<WIFI@REALTEK>"
          Mode:Managed  Frequency:2.412 GHz  Access Point: 24:F2:7F:D1:89:81
          Bit Rate:72.2 Mb/s   Sensitivity:0/0
          Retry:off   RTS thr:off   Fragment thr:off
          Power Management:off
          Link Quality=100/100  Signal level=100/100  Noise level=0/100
          Rx invalid nwid:0  Rx invalid crypt:0  Rx invalid frag:0
          Tx excessive retries:0  Invalid misc:0   Missed beacon:0

$ ip add
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eno1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 50:65:f3:2f:c9:a1 brd ff:ff:ff:ff:ff:ff
    altname enp0s25
    inet 10.150.15.25/24 brd 10.150.15.255 scope global noprefixroute eno1
       valid_lft forever preferred_lft forever
    inet6 2001:db8:15:0:5265:f3ff:fe2f:c9a1/64 scope global dynamic mngtmpaddr
       valid_lft 2591913sec preferred_lft 604713sec
    inet6 fe80::5265:f3ff:fe2f:c9a1/64 scope link noprefixroute
       valid_lft forever preferred_lft forever
3: wlx18d6c7116805: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 18:d6:c7:11:68:05 brd ff:ff:ff:ff:ff:ff
    inet 192.168.15.74/24 brd 192.168.15.255 scope link noprefixroute wlx18d6c7116805
       valid_lft forever preferred_lft forever
    inet6 2001:db8:15:0:a402:f49a:d7be:5049/64 scope global temporary dynamic
       valid_lft 599939sec preferred_lft 81131sec
    inet6 2001:db8:15:0:1ad6:c7ff:fe11:6805/64 scope global dynamic mngtmpaddr noprefixroute
       valid_lft 2591914sec preferred_lft 604714sec
    inet6 fe80::1ad6:c7ff:fe11:6805/64 scope link noprefixroute
       valid_lft forever preferred_lft forever
```

## Part 2

**Nota:** Quedan advertidos.

Ahora probemos como se vería con una dirección WAN:

![img](https://www.sysadminsdecuba.com/wp-content/uploads/2020/05/Screenshot_20200516_204723.png)

Ahora un ejemplo para una LAN:

![img](https://www.sysadminsdecuba.com/wp-content/uploads/2020/05/Screenshot_20200516_205328.png)

**Nota:** Para añadir rutas debemos añadir al final del archivo.

```
routes:
  - to: 192.168.44.0/24
    via: 192.168.0.1
```

Ahora, un ejemplo completo con múltiples rutas:

```
network:
  version: 2
  renderer: networkd
  ethernets:
    ens3:
      dhcp4: yes
    eno2:
      dhcp4: no
      dhcp6: no
      addresses: [10.10.0.11/24]
      gateway4: 10.10.0.1
      routes:
        - to: 192.168.1.0/24
          via: 10.10.0.1        
        - to: 192.168.10.0/24
          via: 10.10.0.1 
```

**Probando la configuración**

Antes de aplicar cualquier cambio, probaremos el archivo de configuración. Ejecute el siguiente comando:

```
sudo netplan try
```

Si usted ve el mensaje:

```
Configuration accepted
```

Entonces puede proseguir a aplicar la configuración. Lo cual se lleva a cabo con:

```
sudo netplan apply
```

En caso de error, podemos debuguear los posibles errores con:

```
sudo netplan apply
```

**Reiniciando los servicios de red**

Una vez que todas las configuraciones se hayan aplicado correctamente, reinicie el servicio Network-Manager[Ubuntu Desktop] ejecutando el siguiente comando:

```
sudo netplan --d apply
```

Si está utilizando Ubuntu Server, utilice el siguiente comando:

```
sudo systemctl restart network-manager
```

O puede reiniciar de la manera tradicional:

```
sudo systemctl restart system-networkd
```

**Verificar la dirección IP**

Ahora, para verificar si las nuevas configuraciones se aplican correctamente, ejecute el siguiente comando para verificar la dirección IP:

```
ip -a
```

Con esto creo que es suficiente. Si este post no ha sido lo suficiente explicativo[y me disculpo por ello, a veces yo mismo me enredo intentando explicar cosas T.T], acá en la página de Netplan [hay bastantes ejemplos](https://netplan.io/examples) que harán más ameno su entendimiento.