```yaml
stages:
  - buildQA
  - buildDEV
  - buildUAT
  - test
  - docker-build
  - deploy
  - registry
  - publish
  - package-registry

tagging:
  tags:
    - shell-devqa
  stage: publish
  only:
    - UAT
  before_script:
    - git config --global user.name "${GITLAB_USER_NAME}"
    - git config --global user.email "${GITLAB_USER_EMAIL}"
  script:
    - cd /home/gitlab-runner/UAT/mybts_fe
    - git status
    - git pull origin UAT
    - tag=$(cat tagversion.json | grep version | grep -Eo "[[:digit:]]+\.[[:digit:]]+\.[[:digit:]]")
    - git tag "$tag"
    - git push --tags
  allow_failure:
    exit_codes: 1

package-registry:
  image: node:latest
  stage: package-registry
  tags:
    - htzdoc
  script:
    - echo "//${CI_SERVER_HOST}/api/v4/projects/${CI_PROJECT_ID}/packages/npm/:_authToken=${REGISTRY_TOKEN}">>.npmrc
    - echo "@rafaelm:registry https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/">>.npmrc
    - echo "registry=https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/">>.npmrc
    - cat .npmrc
    - npm config set @rafaelm:registry https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/
    - npm config set -- '//gitlab.bluetrail.software/api/v4/projects/486/packages/npm/:_authToken' "$REGISTRY_TOKEN"
    - npm publish --registry=https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/
  environment: development
  only:
    changes:
      - /package.json
      - package.json
   

buildQA:
  stage: buildQA
  tags:
    - htzdoc
  image: node:16.14.0
  script:
    - echo "Start building App"
    - npm ci --silent
    - npm install react-scripts --silent@3.4.1 -g --silent
    - cat $ENV_QA >> .env.production
    - GENERATE_SOURCEMAP=false
    - NODE_OPTIONS=\"--max-old-space-size=2048\"
    - node --max-old-space-size=2048
  allow_failure:
    exit_codes: 1
  after_script:
    - npm run build --max-old-space-size=2048
  artifacts:
    name: "QA build"
    paths:
      - build/*
    expire_in: 1 min
    when: on_success
  only:
    - QA

buildDEV:
  stage: buildDEV
  tags:
    - htzdoc
  image: node:16.14.0
  script:
    - echo "Start building App"
    - npm ci --silent
    - npm install react-scripts --silent@3.4.1 -g --silent
    - cat $ENV_DEV >> .env.production
    - GENERATE_SOURCEMAP=false
    - NODE_OPTIONS=\"--max-old-space-size=2048\"
    - node --max-old-space-size=2048
  allow_failure:
    exit_codes: 1
  after_script:
    - npm run build --max-old-space-size=2048
  artifacts:
    name: "DEV build"
    paths:
      - build/*
    expire_in: 1 min
    when: on_success
  only:
    - DEV
    - DEVTEST

buildUAT:
  stage: buildUAT
  tags:
    - htzdoc
  image: node:16.14.0
  script:
    - echo "Start building App"
    - npm ci --silent
    - npm install react-scripts --silent@3.4.1 -g --silent
    - cat $ENV_UAT >> .env.production
    - GENERATE_SOURCEMAP=false
    - NODE_OPTIONS=\"--max-old-space-size=2048\"
    - node --max-old-space-size=2048

  allow_failure:
    exit_codes: 1
  after_script:
    - npm run build --max-old-space-size=2048
  artifacts:
    name: "UAT build"
    paths:
      - build/*
    expire_in: 1 min
    when: on_success
  only:
    - UAT

deploy:
  stage: deploy
  tags:
    - shell-devqa
  before_script:
    - echo "$PASSWORD" | sudo -S chown -R gitlab-runner /var/run/docker.sock
    - rm -rf /home/gitlab-runner/$CI_COMMIT_BRANCH/nginx/*
    - echo "Deleted Old Version"
  script:
    - cp -r build/* /home/gitlab-runner/$CI_COMMIT_BRANCH/nginx
  after_script:
    - echo "Deploy Complete!"
  only:
    - QA
    - DEV
    - UAT


# deploy:
#   stage: deploy
#   image: kroniak/ssh-client
#   before_script:
#     - echo "deploying app"
#   script:
#     - chmod 400 $SSH_PRIVATE_KEY
#     - ssh -o StrictHostKeyChecking=no -i $SSH_PRIVATE_KEY root@$PROD_SERVER_IP "docker pull registry.gitlab.com/alfredomartinezzz/budgefy"
#     - ssh -o StrictHostKeyChecking=no -i $SSH_PRIVATE_KEY root@$PROD_SERVER_IP "docker stop budgefycontainer || true && docker rm budgefycontainer || true"
#     - ssh -o StrictHostKeyChecking=no -i $SSH_PRIVATE_KEY root@$PROD_SERVER_IP "docker run -p 3001:80 -d --name budgefycontainer registry.gitlab.com/alfredomartinezzz/budgefy"
```

```yaml
variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: ""
  PG_HOST: postgres
  PG_PORT: 5432
  POSTGRES_DB: alg_database
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: "123456"
  POSTGRES_HOST_AUTH_METHOD: trust

stages:
  - build
  - test
  - package
  - tag

services:
  - postgres:13-alpine

build:
  stage: build
  image: node:16-alpine3.14
  tags:
    - gcp-improvments
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
  # rules:
  #   - if: $CI_PIPELINE_SOURCE == 'merge_request_event'
  only:
    - LG-453

test:
  stage: test
  dependencies:
    - build
  image: node:16-alpine3.14
  coverage: /All files[^|]*\|[^|]*\s+([\d\.]+)/
  services:
   - postgres
  tags:
    - gcp-improvments
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
    - gcp-improvments
  before_script:
    - docker info
    - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
    - VERSION="$(cat ./changelog/version.md)"
  script:
    # - docker build . --tag $CI_REGISTRY_IMAGE:latest-test --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker build . --tag $CI_REGISTRY_IMAGE:$VERSION
    # - docker push $CI_REGISTRY_IMAGE:latest-test
    - docker push $CI_REGISTRY_IMAGE:$VERSION
    # - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: false
  # dependencies:
  #   - test
  only:
    refs:
      - LG-453
tag:
  tags:
    - gcp-improvments
  stage: tag
  image: bitnami/git:latest
  allow_failure: true
  dependencies:
    - "package"
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
        exit 1;
      else
        git tag -a $VERSION -m "Version $VERSION";
        git push api-origin $VERSION;
      fi
  only: 
    - LG-453

# build-tokyo:
#   stage: build
#   image: node:16-alpine3.14
#   tags:
#     - docker
#   script:
#     - npm install
#     - npm run build
#     - cd public
#     - npm install --omit-dev
#     - npm run build
#   allow_failure: false
#   cache:
#     paths:
#       - node_modules/
#       - public/node_modules/
#   artifacts:
#     expire_in: 1 days
#     when: on_success
#     paths:
#       - node_modules/
#       - public/node_modules/
#       - dist/
#       - public/build/
#   only:
#     - tokyo

# package-tokyo:
#   stage: package
#   image: docker:latest
#   services:
#     - docker:dind
#   tags:
#     - docker
#   before_script:
#     - docker info
#     - echo -n $CI_REGISTRY_PASSWORD | docker login -u $CI_REGISTRY_USER --password-stdin $CI_REGISTRY
#   script:
#     - docker build . --tag $CI_REGISTRY_IMAGE:tokyo-latest --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
#     - docker push $CI_REGISTRY_IMAGE:tokyo-latest
#     - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_NAME
#     - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
#   allow_failure: false
#   dependencies:
#     - build-tokyo
#   only:
#     - tokyo

```

