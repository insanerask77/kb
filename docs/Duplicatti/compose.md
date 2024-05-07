```yaml
#Docker compose with superadmin user PID=0 less privileges PID=1000
# duplicati:v2.0.4.23-2.0.4.23_beta_2019-07-14-ls48 -- S3 AWS working
version: "2.1"
services:
  duplicati:
    image: linuxserver/duplicati:v2.0.4.23-2.0.4.23_beta_2019-07-14-ls48
    container_name: duplicati
    environment:
      - PUID=0
      - PGID=0
      - TZ=Europe/London
      - CLI_ARGS= #optional
    volumes:
      - ./to/appdata/config:/config
      - ./to/backups:/backups
      - /:/source
    ports:
      - 8200:8200
    restart: unless-stopped

```

