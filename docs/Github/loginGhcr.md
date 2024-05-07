# Login Docker registry

```
#!/bin/bash
# docker login ghcr.io -p <token>

echo "Define your PUBLIC_URL to build it. e.g https://centex.centexteam.com"
read PUBLIC_URL
# Vars
# docker pull ghcr.io/bts-centex/centex-strapi:875cb539
# docker login ghcr.io -u developer -p <TOKEN>
REGISTRY=ghcr.io
ORG=bts-centex
USER=developer
IMAGE_NAME=centex-strapi
IMAGE_TAG=localbuild_$(hostname)


check_login() {
    docker login $REGISTRY
}
build_image() {
    docker build -t $REGISTRY/$ORG/$IMAGE_NAME:$IMAGE_TAG --build-arg PUBLIC_URL=$PUBLIC_URL .
}
push_image() {
    docker push $REGISTRY/$ORG/$IMAGE_NAME:$IMAGE_TAG
}

check_login
build_image
push_image

```

