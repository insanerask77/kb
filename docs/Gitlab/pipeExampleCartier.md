```yaml
---
variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: ""
  PG_HOST: postgres
  PG_PORT: 5432
  POSTGRES_DB: alg_database
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: "123456"
  POSTGRES_HOST_AUTH_METHOD: trust

services:
  - postgres:13-alpine

stages:
  - mcuy
  - build
  - test
  - package
  - tagging
  
build:
  stage: build
  image: node:16-alpine3.14
  tags:
    - docker
  script:
    - npm install
    - npm run build
    - cd public
    - npm install --omit-dev
    - npm run build
  allow_failure: false
  cache:
    paths:
      - node_modules/
      - public/node_modules/
  artifacts:
    expire_in: 1 days
    when: on_success
    paths:
      - node_modules/
      - public/node_modules/
      - dist/
      - public/build/
  rules:
    - if: $CI_PIPELINE_SOURCE == 'merge_request_event'

test:
  stage: test
  needs:
    - build
  dependencies:
    - build
  image: node:16-alpine3.14
  coverage: /All files[^|]*\|[^|]*\s+([\d\.]+)/
  services:
   - postgres
  tags:
    - docker
  script:
    - npm run autotest
  allow_failure: false
  artifacts:
    when: always
    reports:
      junit:
        - junit.xml
  rules:
    - if: $CI_PIPELINE_SOURCE == 'merge_request_event'

package:
  stage: package
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
    - VERSION="$(cat ./changelog/version.md)_$CI_MERGE_REQUEST_TARGET_BRANCH_NAME"
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$VERSION
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: false
  dependencies:
    - test
  rules:
    - if: $CI_MERGE_REQUEST_TARGET_BRANCH_NAME == "main"
    - if: $CI_MERGE_REQUEST_TARGET_BRANCH_NAME == "develop"

tag:
  tags:
    - docker
  stage: tagging
  image: bitnami/git:latest
  allow_failure: false
  dependencies:
    - "package"
  script:
    - VERSION="$(cat ./changelog/version.md)"_$CI_MERGE_REQUEST_TARGET_BRANCH_NAME;
    - echo $VERSION_
    - git config --global user.name "${GITLAB_USER_NAME}"
    - git config --global user.email "${GITLAB_USER_EMAIL}"
    - >
      if [ "$(git remote -v | grep 'api-origin')" ]; then
        echo "Remote has been re-established";
        git remote remove api-origin;
        git remote add api-origin https://oauth2:${TAG_TOKEN}@gitlab.com/cartier-lab/alg-backend;
      else
        echo "Remote has been created"
        git remote add api-origin https://oauth2:${TAG_TOKEN}@gitlab.com/cartier-lab/alg-backend;
      fi
    - >
      if [ $(git tag -l "$VERSION") ]; then
        echo "Version $VERSION already exists";
        exit 1;
      else
        git tag -a $VERSION -m "Version $VERSION";
        git push api-origin $VERSION;
      fi
  rules:
    - if: $CI_MERGE_REQUEST_TARGET_BRANCH_NAME == "main"

mcuy:
  stage: mcuy
  when: manual
  tags:
    - mcuy-runner
  script:
    - whoami
    # - cd ~/alg-repositories/cartier_lg_web
# # Login might be asked to pull from this repository
#     - git stash
#     - git pull
#     - cd ~/alg-repositories/megacomputer-scripts
# # Login might be asked to pull from this repository
#     - git pull
#     - sudo ./uninstall.sh
#     - sudo ./install.sh MEGA_COMPUTER 192.168.121.101 root root /data/cartier/cartier_realtime_data

# # Replace the configuration files to include alg-backend in the docker-compose and nginx configuration
#     - cp ~/alg-repositories/megacomputer-scripts/config/nginx.conf ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config/nginx
#     - cp ~/alg-repositories/megacomputer-scripts/config/docker-compose.yml ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime
#     - cp ~/alg-repositories/megacomputer-scripts/config/.env_api_backend ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime
#     - cp ~/alg-repositories/megacomputer-scripts/config/.env_postgres ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime
#     - cp ~/alg-repositories/megacomputer-scripts/config/.env_elastic ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime
#     - cp ~/alg-repositories/megacomputer-scripts/config/filebeat.yml ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config
#     - cp ~/alg-repositories/megacomputer-scripts/config/logstash.conf ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config
#     - cd ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime
#     - docker-compose stop
# # Login might be asked to pull from docker registry
#     - docker-compose pull
#     - export  DATA_ROOT_PATH=/data/cartier/cartier_realtime_data
#     - export COMPOSE_UID=$(id -u)
#     - export COMPOSE_GUID=$(id -g)
#     - docker-compose up -d --force-recreate


build-tokyo:
  stage: build
  image: node:16-alpine3.14
  tags:
    - docker
  script:
    - npm install
    - npm run build
    - cd public
    - npm install --omit-dev
    - npm run build
  allow_failure: false
  cache:
    paths:
      - node_modules/
      - public/node_modules/
  artifacts:
    expire_in: 1 days
    when: on_success
    paths:
      - node_modules/
      - public/node_modules/
      - dist/
      - public/build/
  only:
    - tokyo

package-tokyo:
  stage: package
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:tokyo-latest --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:tokyo-latest
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: false
  dependencies:
    - build-tokyo
  only:
    - tokyo
```

```yml
---
variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: ''
  PG_HOST: postgres
  PG_PORT: 5432
  POSTGRES_DB: alg_database
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: '123456'
  POSTGRES_HOST_AUTH_METHOD: trust

services:
  - postgres:13-alpine

stages:
  - MCS
  - build
  - test
  - develop
  - ready2test
  - release
  - package
  - tagging

build:
  stage: build
  image: node:16-alpine3.14
  tags:
    - docker
  script:
    - npm install
    - npm run build
    - cd public
    - npm install --omit-dev
    - npm run build
  allow_failure: false
  cache:
    paths:
      - node_modules/
      - public/node_modules/
  artifacts:
    expire_in: 1 days
    when: on_success
    paths:
      - node_modules/
      - public/node_modules/
      - dist/
      - public/build/
  rules:
    - if: $CI_PIPELINE_SOURCE == 'merge_request_event' || $CI_COMMIT_BRANCH =~ "/^release.+/"

test:
  stage: test
  needs:
    - build
  dependencies:
    - build
  image: node:16-alpine3.14
  coverage: /All files[^|]*\|[^|]*\s+([\d\.]+)/
  services:
    - postgres
  tags:
    - docker
  script:
    - npm run autotest
  allow_failure: false
  artifacts:
    when: always
    reports:
      junit:
        - junit.xml
  rules:
    - if: $CI_PIPELINE_SOURCE == 'merge_request_event' || $CI_COMMIT_BRANCH =~ "/^release.+/"

develop:
  stage: develop
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
    - VERSION="$(cat ./changelog/version.md)_develop"
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$VERSION
    - docker push $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
  allow_failure: false
  dependencies:
    - test
  only:
    - develop

release:
  stage: release
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
    - VERSION="release_$(cat ./changelog/version.md)"
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$VERSION
  allow_failure: false
  dependencies:
    - test
  rules:
    - if: $CI_MERGE_REQUEST_TARGET_BRANCH_NAME =~ /^release.+/ || $CI_COMMIT_BRANCH =~ "/^release.+/"

package:
  stage: package
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
    - VERSION="$(cat ./changelog/version.md)_main"
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$VERSION
    - docker push $CI_REGISTRY_IMAGE:$VERSION_$CI_COMMIT_SHA
  allow_failure: false
  dependencies:
    - test
  only:
    - main

tag:
  tags:
    - docker
  stage: tagging
  image: bitnami/git:latest
  allow_failure: false
  dependencies:
    - 'package'
  script:
    - VERSION="$(cat ./changelog/version.md)";
    - echo $VERSION
    - git config --global user.name "${GITLAB_USER_NAME}"
    - git config --global user.email "${GITLAB_USER_EMAIL}"
    - >
      if [ "$(git remote -v | grep 'api-origin')" ]; then
        echo "Remote has been re-established";
        git remote remove api-origin;
        git remote add api-origin https://oauth2:${TAG_TOKEN}@gitlab.com/cartier-lab/alg-backend;
      else
        echo "Remote has been created"
        git remote add api-origin https://oauth2:${TAG_TOKEN}@gitlab.com/cartier-lab/alg-backend;
      fi
    - >
      if [ $(git tag -l "$VERSION") ]; then
        echo "Version $VERSION already exists";
        echo "Recreating..."
        git tag --delete $VERSION
        git tag -a $VERSION -m "Version $VERSION";
        git push api-origin $VERSION;
        exit 0;
      else
        git tag -a $VERSION -m "Version $VERSION";
        git push api-origin $VERSION;
      fi
  only:
    - main

MCS:
  stage: MCS
  image: bitnami/git:latest
  variables:
    BRANCH:
      value: ''
    TAG:
      value: ''
  rules:
    - if: '$BRANCH != "" && $TAG != ""'
  tags:
    - docker
  cache: {}
  when: manual
  before_script:
    - export TAG=$TAG
    - git config --global user.name "${GITLAB_USER_NAME}"
    - git config --global user.email "${GITLAB_USER_EMAIL}"
    - >
      if [ -d "./megacomputer-scripts/.git" ]; then
        cd ./megacomputer-scripts;
        git fetch --all;
        git status;
        git pull;
      else
        git clone https://0auth2:$MCS_TOKEN@gitlab.com/cartier-lab/megacomputer-scripts.git;
        cd ./megacomputer-scripts;
        git fetch --all;
        git status;
        git pull;
      fi
  script:
    - >
      if ! git checkout "$BRANCH"; then
        echo "Creando rama";
        git checkout -b $BRANCH;
        git status;
        cat $DOCKER_COMPOSE_DEV > ./config/docker-compose.yml.new;
        sed 's/$TAG/'"$TAG"'/' ./config/docker-compose.yml.new > ./config/docker-compose.yml
        rm ./config/docker-compose.yml.new
        git add *;
        git commit -m "Updated docker compose";
      else
        echo "La rama existe";
        git status;
        cat $DOCKER_COMPOSE_DEV > ./config/docker-compose.yml.new;
        sed 's/$TAG/'"$TAG"'/' ./config/docker-compose.yml.new > ./config/docker-compose.yml
        rm ./config/docker-compose.yml.new
        git add *;
        git commit -m "Updated docker compose"
      fi
    - >
      if ! git push origin $BRANCH; then
        git push -f origin $BRANCH;
      fi

build-tokyo:
  stage: build
  image: node:16-alpine3.14
  tags:
    - docker
  script:
    - npm install
    - npm run build
    - cd public
    - npm install --omit-dev
    - npm run build
  allow_failure: false
  cache:
    paths:
      - node_modules/
      - public/node_modules/
  artifacts:
    expire_in: 1 days
    when: on_success
    paths:
      - node_modules/
      - public/node_modules/
      - dist/
      - public/build/
  only:
    - tokyo

package-tokyo:
  stage: package
  image: docker:latest
  services:
    - docker:dind
  tags:
    - docker
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
  script:
    - docker build . --tag $CI_REGISTRY_IMAGE:tokyo-latest --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:tokyo-latest
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: false
  dependencies:
    - build-tokyo
  only:
    - tokyo

```

