```yaml
image: docker:latest
services:
  - docker:dind


variables:
  DOCKER_DRIVER: overlay
  SPRING_PROFILES_ACTIVE: gitlab-ci
  MAVEN_CLI_OPTS: "-s .m2/settings.xml --batch-mode"
  MAVEN_OPTS: "-Dmaven.repo.local=.m2/repository"


cache: 
    paths:
    - .m2/repository/
    - target/


stages:
  - build
  - test
  - codecoverage
  - codeQuality
  - package
  - deploy


include:
   - template: SAST.gitlab-ci.yml
   - template: Code-Quality.gitlab-ci.yml


build:
    image: maven:3.5-jdk-8-alpine
    stage: build 
    script:   
        - mvn compile
        - ls target/generated-sources
    artifacts: 
        paths: 
         - target/*.war
        

unittests:
    image: maven:3.5-jdk-8-alpine
    stage: test 
    script: 
        - mvn $MAVEN_CLI_OPTS test 
    artifacts:
        reports:
            junit:
                - target/surefire-reports/TEST-*.xml


code_quality:
   stage: codeQuality
   artifacts:
     reports:
       codequality: gl-code-quality-report.json
   after_script:
     - cat gl-code-quality-report.json


spotbugs-sast:
  dependencies:
    - build
  script:
    - /analyzer run -compile=false
  variables:
    MAVEN_REPO_PATH: ./.m2/repository
  artifacts:
    reports:
      sast: gl-sast-report.json


codecoverage:
    image: maven:3.5-jdk-8-alpine
    stage: codecoverage
    script:
        - mvn clean verify
    artifacts:
        paths:
            - target/site/jacoco/index.html


docker-build:
  stage: package
  script:
    - docker login -u gitlab-ci-token -p $CI_BUILD_TOKEN registry.gitlab.com
    - docker build -t registry.gitlab.com/ishitasinghal/messageboard . --build-arg BUILDKIT_INLINE_CACHE=1 
    - docker push registry.gitlab.com/ishitasinghal/messageboard


k8s-deploy:
 image: google/cloud-sdk:latest
 stage: deploy
 script:
 - echo "$GOOGLE_KEY" > key.json
 - gcloud auth activate-service-account ishita-singhal@my-project-ishita.iam.gserviceaccount.com --key-file key.json
 - gcloud container clusters get-credentials cluster-1-ishita --zone us-central1-c --project my-project-ishita
 - kubectl delete secret $DPRK_SECRET_KEY_NAME 2>/dev/null || echo "secret does not exist"
 - kubectl create secret docker-registry $DPRK_SECRET_KEY_NAME --docker-username="$DPRK_DOCKERHUB_INTEGRATION_USERNAME" --docker-password="$DPRK_DOCKERHUB_INTEGRATION_PASSWORD" --docker-email="$DPRK_DOCKERHUB_INTEGRATION_EMAIL" --docker-server="$DPRK_DOCKERHUB_INTEGRATION_URL"/
 - kubectl apply -f deployment.yml
```

