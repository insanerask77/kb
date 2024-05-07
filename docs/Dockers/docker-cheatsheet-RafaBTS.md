#  DOCKER 

## Enviornment Variables

```sh
docker build -t cdxocrdev.azurecr.io/te-nlp_worker:104 --build-arg GIT_PASS=$AZ_PAS --build-arg CDXO_COMMON=master .
docker build -t cdxo/stats-service:sqlfix --build-arg GIT_PASS=$GIT_PASS --build-arg CDXO_COMMON=master .
```





### Multi Line Command

```yaml
version: '3.1'
services:
  db:
    image: postgres
  web:
    build: .
    command:
      - /bin/bash
      - -c
      - |
        python manage.py migrate
        python manage.py runserver 0.0.0.0:8000

    volumes:
      - .:/code
    ports:
      - "8000:8000"
    links:
      - db
```



## Run Image

docker run -it --entrypoint sh <image>
docker run -it --entrypoint sh  cdxocr.azurecr.io/ci/stg_img:latest
docker run -it --entrypoint /bin/bash cdxocr.azurecr.io/common/rabbitmq-autoscaler:2689

## Aumentar Memoria Virtual "shell"

```sh
sysctl -w vm.max_map_count=262144
```

```sh
sysctl -p
```

**CMD espera**

```sh
tail -f /dev/null
```

#### gitlab

/etc/gitlab/initial_root_password

##### socket compartido

volumes:

   \- /var/run/docker.sock:/var/run/docker.sock

###### Permisos socker

`chown -R <user> /var/run/docker.sock`

##### Linux Version

`lsb_release -d`

## Export image content

docker export container_name -o output.tar

## ACR Login

$spPassword="sgj+QlIYo1Zp0M12JKz8JoWQNDsH4Vwv"
echo $spPassword | helm registry login zentcash.azurecr.io --username zentcash --password-stdin
docker run -it --entrypoint sh zentcash.azurecr.io/main:latest
docker run -it --entrypoint sh 

## Tagear y Pushear

az acr login --name cdxocr -u cdxocr -p <PASSWORD>
docker pull cdxocr.azurecr.io/metrics:2831-staging
docker tag cdxocr.azurecr.io/metrics:2831-staging cdxoclients.azurecr.io/metrics:2831-staging
az acr login --name cdxoclients -u cdxoclients -p  <PASSWORD>
docker push cdxoclients.azurecr.io/metrics:2831-staging

az acr login --name cdxocr -u cdxocr -p  <PASSWORD>
docker pull cdxocr.azurecr.io/common/rabbitmq-autoscaler:latest
docker tag cdxocr.azurecr.io/common/rabbitmq-autoscaler:latest 
az acr login --name cdxoclients -u cdxoclients -p  <PASSWORD>
docker push cdxoclients.azurecr.io/common/rabbitmq-autoscaler:latest

**BUILD**

sudo docker build --rm -f "Dockerfile" -t main:latest "."
sudo docker run -p 3000:3000 --name client -d fe_qa_tst:latest

docker run --name -v $PWD:/target trufflehog dxa4481/trufflehog:latest --regex --entropy=False https://anything: <PASSWORD>@dev.azure.com/cdxo-devops/TEXT-EXTRACT_PROJECT/_git/text-extract_repo 
docker run --rm -ti --name trufflehog -v F:\PROJECTS\trufflehog:/target dxa4481/trufflehog:latest --regex --entropy=False https://anything: <PASSWORD>@dev.azure.com/cdxo-devops/TEXT-EXTRACT_PROJECT/_git/text-extract_repo 

## Remove all

docker rm -f $(docker ps -a -q)

docker rm -f $(docker ps -a -q --filter "name=k8s")

### Delete dangling images

docker rmi -f $(docker images -f "dangling=true" -q)



**Clear All**

docker system prune -a -f

**Other commands:**

docker container prune -f
docker image prune -f
docker volume prune -f
docker network prune -f
docker system prune -a -f

## Docker Network

docker network connect zent-network --host=localhost


docker run -it --rm bitnami/postgresql:10 psql -h pg-0 -U postgres

docker run -it --rm adminer:latest adminer
docker run -p 8080:8080 --name adminer -d adminer:latest

psql -U postgres -W $PGPOOL_POSTGRES_PASSWORD -d tx_history < transaction_db.sql
psql -U postgres -W $PGPOOL_POSTGRES_PASSWORD -d users < user_db.sql

Network Compose

```
networks:
  host:
    name: host
    external: true
```

Or

```
version: "3"
services:
  web:
    image: conatinera:latest
    network_mode: "host"        
    restart: on-failure
```



## Get IP

docker inspect textextract_worker1_1 | grep '"IPAddress"'

## MOVE DOCKER FOLDER

sudo mkdir /datadrive/DOCKER
sudo service docker stop
sudo mv /var/lib/docker /datadrive/DOCKER
ln -s /datadrive/DOCKER /var/lib/docker
sudo service docker start

## DOCKER Recreate containers

sudo docker-compose up -d  --force-recreate

## Ver el contenido de una imagen

 docker run -it --entrypoint sh cdxocr.azurecr.io/tst/rabbit-scaler

## DOCKER Logs of Docker Compose

cd /datadrive/browser-v2
sudo docker-compose logs -f
sudo docker-compose logs > browserv2.log

## DOCKER Add packages
apk add nano

## DOCKER: Restart Neo4j

/var/lib/neo4j/bin/neo4j restart

## DOCKER ACCESS NEO4J

sudo docker exec -it cdxo_neo4j bash
cypher-shell -u neo4j -p neo4j

## DOCKER Logs

docker logs cdxo_neo4j -f

## DOCKER COPY FILES

docker cp /datadrive/neo4j_old/certificates/neo4j.cert cdxo_neo4j:/var/lib/neo4j/certificates/neo4j.cert
docker cp .\keycloak-admin-cli-13.0.1.jar relaxed_panini:/scripts/keycloak/client/keycloak-admin-cli-13.0.1.jar
echo " ----- Change Docker Images folder"
sudo mkdir /datadrive/DOCKER
sudo su -c "echo DOCKER_OPTS='--dns 8.8.8.8 --dns 8.8.4.4 -g /datadrive/DOCKER' >> /etc/default/docker"
sudo service docker restart

### IF EXIST IMG

```bash
Get-CDXOImport() {
    sudo docker ps -a -q --filter ancestor="cdxoimport_server"
}
if [ $(Get-CDXOImport) ]
then
    sudo docker rm $(docker stop $(docker ps -a -q --filter ancestor="cdxoimport_server" --format="{{.ID}}"));
else
    echo " ----- cdxo-import Container container don´t exist ----- ";
fi
```

### IF EXIST CONTAINER

```bash
sudo docker ps -a -q --filter name="client"
```

## Install Docker

```bash
sudo apt-get update
sudo apt install apt-transport-https ca-certificates curl software-properties-common -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
sudo apt install docker-ce libpcre3 libpcre3-dev python3-pip python3-venv libicu-dev nginx -y
sudo -H pip3 install PyICU virtualenv

echo " ----- Install Docker Compose"
sudo curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

## Forza recreate Docker compose

```sh
--force-recreate 
```

## Install Azure CLI

```bash
sudo usermod -aG docker ${USER}
curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
```