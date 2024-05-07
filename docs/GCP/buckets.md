## Docker + Bucket + Fuse



**DOCS:**

https://cloud.google.com/storage/docs/gcsfuse-quickstart-mount-bucket?hl=es-419

Lo primero que necesitamos es contruir gcsfuse dentro de nuestra imagen para ello seguremos los pasos de este repositorio oficial de GCP.

https://github.com/GoogleCloudPlatform/gcsfuse

Agregaremos este stage a nuestro dockerfile adaptandolo para poder compilar y mover el binario dentro de nuestra imagen.

```yaml
# Build an alpine image with gcsfuse installed from the source code:
#  > docker build . -t gcsfuse
# Mount the gcsfuse to /mnt/gcs:
#  > docker run --privileged --device /fuse -v /mnt/gcs:/gcs:rw,rshared gcsfuse

FROM golang:1.20.4-alpine as builder

RUN apk add git

ARG GCSFUSE_REPO="/run/gcsfuse/"
ADD . ${GCSFUSE_REPO}
WORKDIR ${GCSFUSE_REPO}
RUN go install ./tools/build_gcsfuse
RUN build_gcsfuse . /tmp $(git log -1 --format=format:"%H")

FROM alpine:3.13

RUN apk add --update --no-cache bash ca-certificates fuse

COPY --from=builder /tmp/bin/gcsfuse /usr/local/bin/gcsfuse
COPY --from=builder /tmp/sbin/mount.gcsfuse /usr/sbin/mount.gcsfuse
ENTRYPOINT ["gcsfuse", "-o", "allow_other", "--foreground", "--implicit-dirs", "/gcs"]
```

#### Install GCP SDK intro docker (Optional)

```dockerfile
RUN curl -sSL https://sdk.cloud.google.com > /tmp/gcl && bash /tmp/gcl --install-dir=~/gcloud --disable-prompts
ENV PATH=$PATH:~/gcloud/google-cloud-sdk/bin
```

#### Define Docker Compose

When we mount GCP fuse into container we need to define the volume with the flag **:shared**

```yaml
version: '3.1'
services:
  db:
    container_name: alg_postgres
    image: postgres:14
    restart: always
    healthcheck:
      test: [ "CMD", "pg_isready", "-q", "-d", "postgres", "-U", "postgres" ]
      timeout: 45s
      interval: 10s
      retries: 10
    ports:
      - "5432:5432"
    env_file:
      - .env
    volumes:
      - db-postgres:/var/lib/postgresql/data

  alg_backend:
    container_name: alg_backend
    user: root
    image: registry.gitlab.com/cartier-lab/alg-backend:latest
    restart: always
    ports:
      - "3000:3000"
    env_file:
      - .env
    privileged: true
    volumes:
      - ${DATA_ROOT_PATH}:/data:shared 

volumes:
  db-postgres:
  alg-backend:
```

#### Define entrypoint

To mount GCSfuse intro container when it starts, we need to define it into entrypoint script.

```bash
#!/bin/sh
if [ "$BUCKET_ENABLED" = "true" ]; then
    gcsfuse --key-file=/usr/auth/bucket.json $BUCKET_NAME /data
fi

npm start
```

![image-20230725175931981](C:\Users\rafab\OneDrive\Documentos\GitHub\KB\GCP\image-20230725175931981.png)
