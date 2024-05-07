```
version: '3'

services:

  sonarqube:
    container_name: sonarqube
    image: docker.io/bitnami/sonarqube:9
    ports:
      - 9000:9000
    volumes:
      - ./sonar/sonarqube_data:/bitnami/sonarqube
    depends_on:
      - db_sonar
    environment:
      # ALLOW_EMPTY_PASSWORD is recommended only for development.
      - ALLOW_EMPTY_PASSWORD=no
      - SONARQUBE_USERNAME=admin
      - SONARQUBE_PASSWORD=BTS2022
      - SONARQUBE_DATABASE_HOST=db_sonar
      - SONARQUBE_DATABASE_PORT_NUMBER=5432
      - SONARQUBE_DATABASE_USER=bn_sonarqube
      - SONARQUBE_DATABASE_PASSWORD=bn_sonarqube
      - SONARQUBE_DATABASE_NAME=bitnami_sonarqube
    networks:
      - denuncias
  db_sonar:
    container_name: sonar_db
    image: docker.io/bitnami/postgresql:13
    volumes:
      - ./sonar/postgresql_data:/bitnami/postgresql
    environment:
      # ALLOW_EMPTY_PASSWORD is recommended only for development.
      - ALLOW_EMPTY_PASSWORD=no
      - POSTGRESQL_USERNAME=bn_sonarqube
      - POSTGRESQL_PASSWORD=bn_sonarqube
      - POSTGRESQL_DATABASE=bitnami_sonarqube
    networks:
      - denuncias

# docker run --rm -e SONAR_HOST_URL= -e SONAR_LOGIN=${TOKEN} -v "${YOUR_REPO}:/usr/src"  sonarsource/sonar-scanner-cli 
# image: sonarsource/sonar-scanner-cli
  front:
    container_name: frontsq
    image: sonarsource/sonar-scanner-cli
    environment:
      - SONAR_HOST_URL=http://sonarqube:9000
      - SONAR_LOGIN=ca263994f75c64b6be0aac780a1ce7e247a16043
    volumes:
      - ./front:/usr/src
    networks:
      - denuncias
    depends_on:
      - sonarqube

networks:
  denuncias:
    
```

```
sysctl -w vm.max_map_count=524288
sysctl -w fs.file-max=131072
ulimit -n 131072
ulimit -u 8192
```

