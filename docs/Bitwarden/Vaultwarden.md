# Instalación de Vaultwarden en Docker

*Posted on November 13, 2022*

[Vaultwarden](https://github.com/dani-garcia/vaultwarden) es una implementación alternativa de la [API de Bitwarden](https://bitwarden.com/help/public-api/), un gestor de contraseñas de código abierto.

Al estar escrita en Rust, necesita mucho menos recursos que la instalación de [Bitwarden en Docker](https://bitwarden.com/help/install-on-premise-linux/) por lo que puede funcionar incluso en una Raspberry Pi 4.

La instalación de [Vaultwarden en Docker](https://hub.docker.com/r/vaultwarden/server) se realiza usando la imagen `vaultwarden/server`.

La forma más sencilla de hacerlo es usando un fichero `docker-compose.yml` con el siguiente contenido:

```
version: '3'
services:
  vaultwarden:
    image: vaultwarden/server:latest
    container_name: vaultwarden
    environment:
      TZ: 'Europe/Madrid'
      DOMAIN: 'https://vault.mydomain.com/'
    volumes:
      - ~/volumes/vaultwarden:/data
    restart: unless-stopped
networks:
  default:
    external: true
    name: nginxpm_network
```

> **Nota**: Nótese que no se expone ningún puerto al *Docker host* y que se utiliza la misma red que [Nginx Proxy Manager](https://www.manelrodero.com/blog/instalacion-de-nginx-proxy-manager-en-docker).

Para iniciar el contenedor, se puede usar `docker-compose up -d` o usar el contenido del fichero anterior en **Portainer**.

# Pre-requisitos

Antes de poder acceder a Vaultwarden para configurarlo es necesario:

- Crear un [registro DNS de tipo A en Cloudflare](https://www.manelrodero.com/blog/instalacion-de-nginx-proxy-manager-en-docker#registro-dns-en-cloudflare): `vault.mydomain.com` → 192.168.1.180

![DNS Type A](https://www.manelrodero.com/assets/img/blog/2022-11-13_image_1.png)

- Crear un [certificado SSL de Let’s Encrypt](https://www.manelrodero.com/blog/instalacion-de-nginx-proxy-manager-en-docker#certificado-ssl) para el *hostname* anterior

![SSL Certificate](https://www.manelrodero.com/assets/img/blog/2022-11-13_image_2.png)

- Crear un [*proxy host*](https://www.manelrodero.com/blog/instalacion-de-nginx-proxy-manager-en-docker#certificado-ssl) para el *hostname* anterior: `vault.mydomain.com` → `http://vaultwarden:80`

![Proxy host](https://www.manelrodero.com/assets/img/blog/2022-11-13_image_3.png)

# Configuración

Una vez en marcha, se puede acceder a Vaultwarden a través de `HTTPS` usando Nginx Proxy Manager (en este ejemplo `https://vault.mydomain.com/`) y comenzar la configuración.

El primer paso consiste en crear una nueva cuenta de usuario para acceder al *vault* pulsando en el botón `Create Account`:

- Email Address: `vaultwarden@mydomain.com`
- Name: `Manel`
- Master Password: `*********`
- Re-type Master Password: `*********`
- Master Password Hint (optional): `Pista`

Si el usuario se crea correctamente, se podrá iniciar sesión con el mismo y ver la siguiente pantalla:

![Vaultwarden](https://www.manelrodero.com/assets/img/blog/2022-11-13_image_4.png)

> **Nota**: Es aconsejable crear todos los usuarios que se necesiten en este momento y así se puede deshabilitar el registro de usuarios para aumentar la seguridad.

## Habilitar [Admin Panel](https://github.com/dani-garcia/vaultwarden/wiki/Enabling-admin-page)

Para habilitar la [página de administración](https://github.com/dani-garcia/vaultwarden/wiki/Enabling-admin-page) hay que realizar lo siguiente:

- Ejecutar el comando `openssl rand -base64 48` para generar un *token* aleatorio:

```
U543Xx6Bhj+uiv9DhNDcAExM9Cef2C644CC7wwFGUh7Z4I6OvOEaYZaxjnrrKZYX
```

- Agregar la variable `ADMIN_TOKEN` al fichero `docker-compose.yml`:

```
    environment:
      ADMIN_TOKEN: 'U543Xx6Bhj+uiv9DhNDcAExM9Cef2C644CC7wwFGUh7Z4I6OvOEaYZaxjnrrKZYX'
```

- Acceder al Vaultwarden Admin Panel a través de `https://vault.mydomain.com/admin`:

![Admin Panel](https://www.manelrodero.com/assets/img/blog/2022-11-13_image_5.png)

> **Nota**: La primera vez que se guarde la configuración, se generará el fichero `/data/config.json`. Este fichero tiene preferencia sobre las variables indicadas en `docker-compose.yml`.

> **Nota**: También se podría haber activado el Admin Panel creando o modificando el fichero `/data/config.json` añadiendo la variable `admin_token` tal como se muestra a continuación:

```
{
  "admin_token": "U543Xx6Bhj+uiv9DhNDcAExM9Cef2C644CC7wwFGUh7Z4I6OvOEaYZaxjnrrKZYX"
}
```

Algunos valores interesantes a configurar:

- General settings
  - Domain URL: `https://vault.mydomain.com/`
  - Trash auto-delete days: `90`
  - [Allow new signups](https://github.com/dani-garcia/vaultwarden/wiki/Disable-registration-of-new-users): `Disable`
  - Org creation users: `vaultwarden@mydomain.com`
  - [Allow invitations](https://github.com/dani-garcia/vaultwarden/wiki/Disable-invitations): `Disable`
- Advanced settings
- Yubikey settings
- Global Duo settings
- SMTP Email Settings
- Email 2FA Settings
- Read-Only Config
- Backup Database

> **Nota**: Una vez se haya configurado el entorno, es recomendable deshabilitar el Admin Panel. Para ello hay que eliminar la variable de entorno `ADMIN_TOKEN` y asegurarse que no exista la variable `admin_token` en el fichero `/data/config.json`.

# Actualizar

Si ya se había instalado Vaultwarden anteriormente, se puede actualizar de la siguiente manera:

```
docker stop vaultwarden
docker rm vaultwarden
docker rmi vaultwarden/server
docker-compose up -d
```

# Soporte

Algunos comandos para gestionar la configuración del contenedor:

```
# Acceder al shell mientras el contenedor está ejecutándose
docker exec -it vaultwarden /bin/bash

# Monitorizar los logs del contenedor en tiempo real
docker logs -f vaultwarden

# Obtener la versión de Vaultwarden
docker exec vaultwarden /vaultwarden --version
```