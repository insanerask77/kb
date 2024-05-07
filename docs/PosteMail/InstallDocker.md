# Docker builds a corporate mailbox, porte.io tutorial

** Metawalk ** 2023-05-20 PM

[![img](https://iweec.com/laika.jpg)](https://www.lcayun.com/aff/SWTQTESR)

# poste.io introduction

First of all, you need to know that you can have your own domain name suffix mailbox by setting up a corporate mailbox. You can define your email address, and what admin, root, and info are used at will. Poste.io Official Network：[https://poste.io/](https://iweec.com/go/aHR0cHM6Ly9wb3N0ZS5pby8=) Documentation：[https://poste.io/doc/](https://iweec.com/go/aHR0cHM6Ly9wb3N0ZS5pby9kb2Mv)

Poste.io is an email server solution designed to provide a simple and secure email system. It provides a complete mail server stack, including mail transmission agent （MTA）, mail transmission agent （IMAP/POP3） and mail filter. Poste.io's design goal is easy to deploy and manage, and it is suitable for individual users, small businesses, and small and medium-sized organizations.

Poste.io provides an intuitive Web interface that allows users to easily set up and manage their mail servers. It supports multiple domain names and user accounts, and provides functions such as user management, email filtration, spam and virus detection. In addition, Poste.io has also integrated web-based email clients to enable users to access and send emails through Web browsers.
![poste.jpg](https://iweec.com/usr/uploads/2023/05/2207498322.jpg)

# poste.io ready

1. Domain name, if not, click [Gname](https://iweec.com/go/aHR0cHM6Ly93d3cuZ25hbWUuY29tL3JlZ2lzdGVyP3R0PTg2Nzc0JiZ0dGNvZGU9TVdLODY3NzR0dGJqPWVtYWls) Purchase a com, net or ororg, and do not recommend the use of untrusted domain name suffixes such as icu.
2. vps, recommended for use [Lekayun](https://iweec.com/go/aHR0cHM6Ly93d3cubGNheXVuLmNvbS9hZmYvU1dUUVRFU1I=) ，On the Chinese page, all the purchased hosts are turned on 25. It is recommended to purchase 2G memory or above （ video presentation. I use CN2 GIA（elasticity）2 core 1G ）, but it cannot be abused. Enterprises or individuals can apply for rDNS for normal use.

As you know, the 25 vps on the market are not easy to find. The CC I introduced before is also directly supported by rDND. I can bind it in the background and register the address：[https://app.cloudcone.com.cn/?ref=7462](https://iweec.com/go/aHR0cHM6Ly9hcHAuY2xvdWRjb25lLmNvbS5jbi8|cmVmPTc0NjI=) Preferential vps can refer to the page：[https://bbs.csdn.net/topics/610404063](https://iweec.com/go/aHR0cHM6Ly9iYnMuY3Nkbi5uZXQvdG9waWNzLzYxMDQwNDA2Mw==)

# poste.io construction

This tutorial, the vps system I use is Ubuntu 20.04！

However, of course, we still conduct domain name analysis as follows：

| Host record | Record type | Recorded value  |
| ----------- | ----------- | --------------- |
| mail        | A           | Your IP address |
| smtp        | CNAME       | mail.**.com     |
| pop         | CNAME       | mail.**.com     |
| imap        | CNAME       | mail.**.com     |
| @           | MX          | mail.**.com     |
| @           | TXT         | v=spf1 mx ~all  |

1. Update the system and install docker and screen；

```lua
apt  update && apt install screen docker.io -y
```

2. Pull the mirror image；

```bash
docker pull analogic/poste.io
```

3. New mail catalog

```bash
mkdir /home/mail
```

4. Start the container in screen, pay attention to here: mail.*.com to change to your email domain name！

```bash
screen
docker run \
    --net=host \
    -e TZ=Europe/Prague \
    -v /home/mail:/data \
    --name "mailserver" \
    -h "mail.*.com" \
    -t analogic/poste.io
```

5. Visit address mail.Your domain name/admin/install/server（ is shown to be unsafe, continue to visit, set up the certificate ） next, set the domain name, administrator mailbox and password.
6. Find the label in the system settings`TLS Certificate`，Automatically apply for a certificate. After applying for the certificate, you can visit https. Then in the domain name details, click to generate`redirect`，Add domain name DKIM analysis after generation, for example：

```bash
s20230520790._domainkey.proxies.icu. IN TXT "k=rsa; p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxdSK7/g146G3kTo9KrjXBmHJr6PQA80RbL/f6iAQ1zRgGi3n9sbxmXXsBFrgXhMqOdE5BTVts2Z1z2TsWyBHxhHYJcy2uDJN6xnTMOxiLWgjLkzcl49BM//53n75VLlQIJcmmRzHrHfbowWk8g7wAKH6ClC/GRoJ7VVs8/ESZYQPd1oQdcQ1XiDCt4XI7u+CzupfOKQ+9XnEsCKFQTye4Qtjbbp/SXI8CCl0Bdv8bdRAtwHxPGf2f8fee1KnmUCHWT5Cfdw9oB3Dwd77eTPKVFRtFYz7IT5yrk2HWmQT3oBVIepWpapxMIpviOX8zJ522HTlPuhBJhoi9Ep4qmzPnQIDAQAB"
```

![Screenshot 2023-05-20 19.35.18.png](https://iweec.com/usr/uploads/2023/05/3468958996.png)

7. In the mail account, delete users can be added; in the server state, check the diagnosis, and clearly see the server port state；

By the way, let everyone test the command at the 25th port: telnet smtp.qq.com 25

8. The mailing address of the mailbox user is mail.Your domain name/webmail/, you can test and send a letter; test the health of the mailbox [https://www.mail-tester.com/](https://iweec.com/go/aHR0cHM6Ly93d3cubWFpbC10ZXN0ZXIuY29tLw==)

It can be used if the score exceeds 5, but to get a higher score, you can submit a work order to apply for rDNS, provided that you cannot send spam.

# Customer end settings

Receipt server 【IMAP】

| Set                  | content             |
| -------------------- | ------------------- |
| EMAIL                | Your mailbox        |
| password             | Your email password |
| Server 【Host Name】 | mail.*.com          |
| Port 【Port Number】 | 993                 |
| Security             | SSL                 |

Sender 【IMAP】

| Set                  | content             |
| -------------------- | ------------------- |
| EMAIL                | Your mailbox        |
| password             | Your email password |
| Server 【Host Name】 | mail.*.com          |
| Port 【Port Number】 | 587                 |
| Security             | SSL                 |

# supplement

Forgot to say, it’s better to set a hostname again, referencehttps://iweec.com/221.html Or directly

```cpp
sudo hostnamectl set-hostname mail.* .com
```

# Video tutorial

```yaml
version: '3.7'

services:
  nginx-proxy:
    container_name: rproxy
    image: jwilder/nginx-proxy
    ports:
    - "80:80"
    - "443:443"
    volumes:
    - ./nginx/certs:/etc/nginx/certs:ro
    - ./nginx/conf.d:/etc/nginx/conf.d
    - ./nginx/vhosts.d:/etc/nginx/vhost.d
    - ./nginx/html:/usr/share/nginx/html
    - /var/run/docker.sock:/tmp/docker.sock:ro
    restart: always
    environment:
    - "TZ=Europe/Madrid"
    labels:
    - com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy

  letsencrypt:
    container_name: encrypt
    image: jrcs/letsencrypt-nginx-proxy-companion
    volumes:
    - ./nginx/certs:/etc/nginx/certs:rw
    - ./nginx/conf.d:/etc/nginx/conf.d
    - ./nginx/vhosts.d:/etc/nginx/vhost.d
    - ./nginx/html:/usr/share/nginx/html
    - /var/run/docker.sock:/var/run/docker.sock:ro
    restart: always
    environment:
    - "TZ=Europe/Madrid"

  db:
    container_name: mysql
    image: mysql:5.7
    ports:
    - "3306:3306"
    command: --default-authentication-plugin=mysql_native_password --lower_case_table_names=1
    volumes:
    - ./mysql:/var/lib/mysql
    restart: unless-stopped
    environment:
      TZ: Europe/Madrid
      MYSQL_ROOT_PASSWORD: rootpass
      MYSQL_DATABASE: db_name
      MYSQL_USER: user
      MYSQL_PASSWORD: pass

  mail:
    container_name: poste
    image: analogic/poste.io
    ports:
    - "25:25"
    - "110:110"
    - "143:143"
    - "465:465"
    - "587:587"
    - "993:993"
    - "995:995"
    volumes:
    - ./poste:/data
    restart: always
    environment:
    - "TZ=Europe/Madrid"
    - "VIRTUAL_HOST=mail.madolell.com"
    - "VIRTUAL_PORT=80"
    - "LETSENCRYPT_HOST=mail.madolell.com"
    - "LETSENCRYPT_EMAIL=admin@madolell.com"
    - "HTTPS=OFF"
    - "DISABLE_CLAMAV=TRUE"
```



```yaml
version: '3'
services:                                                                                                                  
  posteio:                                                                                                            
    image: analogic/poste.io:latest                                                                                    
    container_name: poste.io                                                                                    
    hostname: mail.madolell.tk                                                                             
    restart: always
    expose: 
      - '80'
      - '25'    # SMTPS - mostly processing incoming mails
      - '465'    # SMTPS - mostly processing incoming mails
      - '110'  #  POP3 - standard protocol for accessing mailbox, STARTTLS is required before client auth   
      - '143'  #  IMAP - standard protocol for accessing mailbox, STARTTLS is required before client auth
      - '443'  #  HTTPS - access to administration or webmail client                                
      - '587'  #  MSA - SMTP port used primarily for email clients after STARTTLS and auth            
      - '993'  #  IMAPS - alternative port for IMAP encrypted since connection                        
      - '995'  #  POP3S - encrypted POP3 since connections                                            
    environment:                                                                                             
      - HTTPS=OFF                                                                                         
    volumes:                                                                                                 
      - './containers/poste.io/data:/data'                                                                
      - '/etc/localtime:/etc/localtime:ro'                                                                  
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=proxy"
      - "traefik.http.routers.mail.rule=Host(`mail.madolell.tk`)"
      - "traefik.http.routers.mail.entrypoints=websecure"
      - "traefik.http.services.mail.loadbalancer.server.port=80"
    networks:
      - public
      - proxy
networks:
    public:
    proxy:
      external: true
```

```yaml
version: '3'

services:
  nginx-proxy:
    image: jwilder/nginx-proxy
    labels:
        com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy: "true"
    container_name: nginx-proxy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - /data/nginx/conf.d:/etc/nginx/conf.d
      - /data/nginx/vhost.d:/etc/nginx/vhost.d
      - /data/nginx/html:/usr/share/nginx/html
      - /data/nginx/certs:/etc/nginx/certs:ro
      - /var/run/docker.sock:/tmp/docker.sock:ro

  nginx-letsencrypt:
    image: jrcs/letsencrypt-nginx-proxy-companion
    container_name: nginx-letsencrypt
    restart: unless-stopped
    volumes:
      - /data/nginx/conf.d:/etc/nginx/conf.d
      - /data/nginx/vhost.d:/etc/nginx/vhost.d
      - /data/nginx/html:/usr/share/nginx/html
      - /data/nginx/certs:/etc/nginx/certs:rw
      - /var/run/docker.sock:/var/run/docker.sock:ro
    environment:
      - NGINX_DOCKER_GEN_CONTAINER=nginx-proxy
      - NGINX_PROXY_CONTAINER=nginx-proxy

  mailserver:
    image: poste.io/mailserver:dev
    container_name: mailserver
    restart: unless-stopped
    ports:
      - "25:25"
      - "110:110"
      - "143:143"
      - "587:587"
      - "993:993"
      - "995:995"
      - "4190:4190"
    environment:
      - LETSENCRYPT_EMAIL=info@analogic.cz
      - LETSENCRYPT_HOST=mail.poste.io
      - VIRTUAL_HOST=mail.poste.io
      - HTTPS=OFF
      - DISABLE_CLAMAV=TRUE
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - /data/nginx/html/.well-known:/opt/www/.well-known
      - /data/mailserver:/data
```

Hola Pelado! En primer lugar gracias por el tutorial para instalar Poste.io. Ya lo tengo corriendo en mi servidor y funciona excelente 😄 🎉

[![Captura](https://camo.githubusercontent.com/24ef8648163789879368a9f0d97681169e699f1fd2187923189033bfb4ed985c/68747470733a2f2f692e696d6775722e636f6d2f797a52696765512e706e67)](https://camo.githubusercontent.com/24ef8648163789879368a9f0d97681169e699f1fd2187923189033bfb4ed985c/68747470733a2f2f692e696d6775722e636f6d2f797a52696765512e706e67)

**Para los que esten en la misma les dejo algunos consejos**:

1. **Desactivar ClamAV:** Si estan deployando en Digital Ocean van a tener que hacerlo porque se come el procesador. Basta con agregar DISABLE_CLAMAV=TRUE como variable de entorno en userdata.yaml.

Se van a quedar sin el antivirus, pero si usan Linux o tienen sentido común 🤷‍♂️

1. **Agregar SPF y DMARC:**
   - tudominio.com. TXT "v=spf1 mx ~all"
   - _dmarc.tudominio.com. TXT "v=DMARC1; p=none; rua=mailto:dmarc-reports@tudominio.com"
2. **Agregar DKIM:** Lo pueden sacar desde *Virtual Domains > tudominio.com > DKIM key*
3. **Usar un SMTP Relay:** En mi caso los e-mails no se mandaban. También hay gente que manda y a la otra persona le llega a spam. Para esto se puede configurar un relay en *System Setting > Email Processing > Default SMTP Route*. Pueden usar Sendgrid, Mailgun o cualquier otro 👍

No olviden validar el dominio asi no sale "Enviado desde sendgrid.net" o lo que sea. En el caso de Sendgrid se hace [desde acá](https://app.sendgrid.com/settings/sender_auth/domain/create).

1. Opcional: Crear un CNAME para SMTP, POP, IMAP:

    

   Algunos clientes por defecto van a buscar, por ejemplo, smtp.tudominio.com, por eso es buena idea crear un CNAME:

   - smtp.tudominio.com. CNAME mail.tudominio.com.
   - pop.tudominio.com. CNAME mail.tudominio.com.
   - imap.tudominio.com. CNAME mail.tudominio.com.

**Otras cosas que yo hice y quizá no son necesarias**:

1. **Crear un volumen:** Metí los datos en un volumen, asi si por algún motivo tengo que regenerar el servidor de e-mail no pierdo los datos. ¿Se puede poner algo como prevent_destroy para evitar que a Terraform le pinte boletearlo?
2. **Eliminé la creación del dominio:** Como ya tenia el dominio raiz generado de antes me fallaba porque intentaba regenerarlo.
3. **Puse el dominio raiz en una variable:** Para más placer. (Y por el punto 7)
4. **Cambie la interpolación de variables:** En Terraform 0.12 ya no hace falta el "${variable}". Si es solo eso se puede poner la variable directamente y me ahorro el warning.
5. **Use la imágen de Debian 10:** ¿Porqué? No hay porque. (Me gusta más nomás)

Como toda persona de bien hice los cambios con Terraform (Excepto puntos 3 y 4) 💪
Estan en un fork, por si alguno quiere chusmear o tomar algo lo puede hacer desde acá:

- https://github.com/imcosta/peladonerd/tree/master/terraform/3

Saludos





Hi, I'm trying to use something similar to this, but with separated docker-compose.yml files like this:

# For nginx-proxy and lets-encrypt

```
version: "3.8"
services:
  # nginx-proxy
  nginx-proxy:
    image: jwilder/nginx-proxy
    container_name: MyProxy
    restart: always
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - "/var/run/docker.sock:/tmp/docker.sock:ro"
      - "./certs:/etc/nginx/certs"
      - "./vhost.d:/etc/nginx/vhost.d"
      - "./html:/usr/share/nginx/html"
      - "./conf.d:/etc/nginx/conf.d"
    networks:
      - "net"
  # lets-encrypt
  letsencrypt-nginx-proxy-companion:
    image: jrcs/letsencrypt-nginx-proxy-companion
    container_name: lets-encrypt
    depends_on:
      - nginx-proxy
    restart: always
    environment:
      - NGINX_DOCKER_GEN_CONTAINER=MyProxy
      - NGINX_PROXY_CONTAINER=MyProxy
    volumes:
      - "/var/run/docker.sock:/tmp/docker.sock:ro"
      - "./certs:/etc/nginx/certs"
      - "./vhost.d:/etc/nginx/vhost.d"
      - "./html:/usr/share/nginx/html"
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
    networks:
      - net
networks:
  net:
    external: true
```

And for poste.io:

```
version: '3.8'

services:
  mailserver:
    image: analogic/poste.io:latest
    container_name: mailserver
    hostname: mail
    domainname: johandroid.com
    restart: unless-stopped
    ports:
      - "25:25"
      - "110:110"
      - "143:143"
      - "587:587"
      - "993:993"
      - "995:995"
      - "4190:4190"
      - "465:465"
    environment:
      - LETSENCRYPT_EMAIL=info@johandroid.com
      - LETSENCRYPT_HOST=mail.johandroid.com
      - VIRTUAL_HOST=mail.johandroid.com,smtp.johandroid.com,imap.johandroid.com
      - DISABLE_CLAMAV=TRUE
      - HTTPS=OFF
    volumes:
      - "/etc/localtime:/etc/localtime:ro"
      - "~/nginx/nginx-proxy/html/.well-known:/opt/www/.well-known"
      - "./data/mailserver:/data"
    networks:
      - "net"

networks:
  net:
    external: true
```

This is working nice and smooth, I can login into the admin or webclient, and even use thunderbird as mail client, I can send and recieve mail ... But, when I try to use the gmail app for android, it complains about the certificate, it says the issuer of certificate is poste and not my lets-encrypt certificate, actually the expiration date is Sep 17, 2020

How can I make it use my certificates?

Hi johandroid.
I got the same problem. Link the letsencrypt certs from the nginx proxy to the poste.io container =>

```
    volumes:
      - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/key.pem:/data/ssl/server.key:ro
      - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/fullchain.pem:/data/ssl/ca.crt:ro
      - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/cert.pem:/data/ssl/server.crt:ro
```

=> https://www.cloudrocket.at/posts/self-hosted-mail-server-with-poste.io-and-nginx/#the-tls-termination-problem



[![@fedeaguilera](https://avatars.githubusercontent.com/u/40297771?s=80&v=4)](https://gist.github.com/fedeaguilera)



### **[fedeaguilera](https://gist.github.com/fedeaguilera)** commented [on Jul 8, 2022](https://gist.github.com/analogic/51fbe91b580d7913b72320f89bf994cc?permalink_comment_id=4224921#gistcomment-4224921)

> Hi johandroid. I got the same problem. Link the letsencrypt certs from the nginx proxy to the poste.io container =>
>
> ```
>     volumes:
>       - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/key.pem:/data/ssl/server.key:ro
>       - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/fullchain.pem:/data/ssl/ca.crt:ro
>       - NGINX_PROXY_PATH/ssl/certs/mail.YOUR_DOMAIN.com/cert.pem:/data/ssl/server.crt:ro
> ```
>
> => https://www.cloudrocket.at/posts/self-hosted-mail-server-with-poste.io-and-nginx/#the-tls-termination-problem

hi guys. if you have a 2 domains. works only one certificate?





```
version: '3'

services:

  mailserver:
    image: analogic/poste.io
    container_name: mailserver-leninnet.ni
    restart: always
    network_mode: "host"
    expose:
      - "25"
      - "8181"
      - "4443"
      - "110"
      - "143"
      - "465"
      - "587"
      - "993"
      - "995"
    environment:
      - TZ=America/Managua
      - h=mail.leninnet.ni  # Direccion del servidor de mail hosting
      - HTTP_PORT=8181
      - HTTPS_PORT=4443
      - DISABLE_CLAMAV=TRUE
    volumes:
      - /mnt/mail:/data
```

```
version: '3.8'

services:
  mailserver:
    image: analogic/poste.io:latest
    container_name: mailserver
    hostname: mail
    domainname: madolell.com
    restart: unless-stopped
    network_mode: host
    expose: 
      - '80'
      - '25'    # SMTPS - mostly processing incoming mails
      - '465'    # SMTPS - mostly processing incoming mails
      - '110'  #  POP3 - standard protocol for accessing mailbox, STARTTLS is required before client auth   
      - '143'  #  IMAP - standard protocol for accessing mailbox, STARTTLS is required before client auth
      - '443'  #  HTTPS - access to administration or webmail client                                
      - '587'  #  MSA - SMTP port used primarily for email clients after STARTTLS and auth            
      - '993'  #  IMAPS - alternative port for IMAP encrypted since connection                        
      - '995'  #  POP3S - encrypted POP3 since connections
    environment:
      - h=mail.madolell.com
      - LETSENCRYPT_EMAIL=info@madolell.com
      - LETSENCRYPT_HOST=mail.madolell.com
      - VIRTUAL_HOST=mail.madolell.com
      - DISABLE_CLAMAV=TRUE
      - HTTPS=ON
    volumes:
      - "/etc/localtime:/etc/localtime:ro"
      - "~/nginx/nginx-proxy/html/.well-known:/opt/www/.well-known"
      - "./data/mailserver:/data"
```

