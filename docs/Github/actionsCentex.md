## Actions Example Centex

```yaml
name: Docker build and deploy Prod

on:
  push:
    tags:
      - '[0-9]+\.[0-9]+\.[0-9]'
      - '[0-9]+\.[0-9]+\.[0-9]+'
      - '[0-9]+\.[0-9]+\.[0-9]+\-prod'


jobs:
  build-image:
      runs-on: ubuntu-latest
      env:
        CI_REGISTRY: ghcr.io
        CI_REGISTRY_IMAGE: ghcr.io/bts-centex/centex-strapi
        ACTIONS_ALLOW_UNSECURE_COMMANDS: true
        SHA8: ${GITHUB_SHA::8}
      steps:
        - name: Extract branch name on push
          if: github.event_name != 'pull_request'
          shell: bash
          run: echo "::set-env name=CI_IMAGE_TAG::${{github.ref_name}}"
          id: extract_branch

        - name: Login to Github
          uses: docker/login-action@v3
          with:
            registry: ${{ env.CI_REGISTRY }}
            username: ${{ secrets.GHCR_USER }}
            password: ${{ secrets.GHCR_TOKEN }}
            
        - name: Checkout repository
          uses: actions/checkout@v3

        - name: Build and Push Docker Image
          run: |
            docker build . --tag $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG --tag $CI_REGISTRY_IMAGE:latest --build-arg PUBLIC_URL=${{ secrets.PUBLIC_URL_PROD }}
            docker push $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG
            docker push $CI_REGISTRY_IMAGE:latest
  deploy:
    runs-on: ubuntu-latest
    needs: build-image
    env:
      DOCKER_TAG: ${{github.ref_name}}
      ACTIONS_ALLOW_UNSECURE_COMMANDS: 'true'
    steps:
        - uses: trstringer/manual-approval@v1.9.0
          timeout-minutes: 5
          with:
            secret: ${{ secrets.APROVAL }}
            approvers: devops
            minimum-approvals: 1
            issue-title: "Deploying ${{github.ref_name}} to prod"
            issue-body: "Please approve or deny the deployment of version ${{github.ref_name}}"

        - name : Set environment variables based on tag 
          run : |
            if [[ "${{github.ref_name}}" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
              echo "::set-env name=SERVER::production"
              echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
              echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST_PROD }}"
              echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY_PROD }}"
              echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_PROD }}"
            else 
              echo "Invalid tag format specified."
              exit 1;
            fi
      
        - name: install ssh keys
          run: |
            install -m 600 -D /dev/null ~/.ssh/id_rsa
            echo "${{ secrets.SSH_PRIVATE_KEY_PROD }}" > ~/.ssh/id_rsa
            ssh-keyscan -H ${{ env.SERVER_IP }} > ~/.ssh/known_hosts
        - name: connect and pull
          run: |
            echo "Using tag: ${{github.ref_name}}"
            echo "Deploying in ${{ env.SERVER }} server"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "docker login ghcr.io -u ${{ secrets.GHCR_USER }} -p ${{ secrets.GHCR_TOKEN_ONLY_READ }} && exit"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && export TAG="${{github.ref_name}}" && docker compose pull && docker compose up -d --force-recreate && exit"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && docker system prune -af && exit"
    
        - name: cleanup
          run: rm -rf ~/.ssh
        

```

```yaml
name: Docker build and deploy test/staging

on:
  push:
    tags:
      - '[0-9]+\.[0-9]+\.[0-9]+\-test'
      - '[0-9]+\.[0-9]+\.[0-9]+\-rc'
      - '[0-9]+\.[0-9]+\.[0-9]+\-RC'

jobs:
  build-image:
      runs-on: ubuntu-latest
      env:
        CI_REGISTRY: ghcr.io
        CI_REGISTRY_IMAGE: ghcr.io/bts-centex/centex-strapi
        ACTIONS_ALLOW_UNSECURE_COMMANDS: true
        SHA8: ${GITHUB_SHA::8}
      steps:
        - name: Extract branch name on push
          if: github.event_name != 'pull_request'
          shell: bash
          run: echo "::set-env name=CI_IMAGE_TAG::${{github.ref_name}}"
          id: extract_branch

        - name: Login to Github
          uses: docker/login-action@v3
          with:
            registry: ${{ env.CI_REGISTRY }}
            username: ${{ secrets.GHCR_USER }}
            password: ${{ secrets.GHCR_TOKEN }}
            
        - name: Checkout repository
          uses: actions/checkout@v3
        
        - name : Set environment variables based on tag 
          env:
            DOCKER_TAG: ${{github.ref_name}}
            ACTIONS_ALLOW_UNSECURE_COMMANDS: 'true'
          run : |
            if [[ "${{github.ref_name}}" =~ ^[0-9]+\.[0-9]+\.[0-9]+-(rc|RC)$ ]]; then
              echo "::set-env name=SERVER::staging"
              echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
              echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
              echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
              echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_STAGING }}"
              echo "::set-env name=PUBLIC_URL::${{ secrets.PUBLIC_URL_STG }}"
            elif [[ "${{github.ref_name}}" =~ ^[0-9]+\.[0-9]+\.[0-9]+(-test)$ ]]; then
              echo "::set-env name=SERVER::test"
              echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
              echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
              echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
              echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_TEST }}"
              echo "::set-env name=PUBLIC_URL::${{ secrets.PUBLIC_URL_TEST }}"
            else 
              echo "Invalid tag format specified."
              exit 1;
            fi

        - name: Build and Push Docker Image
          run: |
            docker build . --tag $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG --tag $CI_REGISTRY_IMAGE:latest --build-arg PUBLIC_URL=${{ env.PUBLIC_URL }}
            docker push $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG
            docker push $CI_REGISTRY_IMAGE:latest
      
        - name: install ssh keys
          run: |
            install -m 600 -D /dev/null ~/.ssh/id_rsa
            echo "${{ secrets.SSH_PRIVATE_KEY }}" > ~/.ssh/id_rsa
            ssh-keyscan -H ${{ env.SERVER_IP }} > ~/.ssh/known_hosts
        - name: connect and pull
          run: |
            echo "Using tag: ${{github.ref_name}}"
            echo "Deploying in ${{ env.SERVER }} server"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "docker login ghcr.io -u ${{ secrets.GHCR_USER }} -p ${{ secrets.GHCR_TOKEN_ONLY_READ }} && exit"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && export TAG="${{github.ref_name}}" && docker compose pull && docker compose up -d --force-recreate && exit"
            ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && docker system prune -af && exit"
    
        - name: cleanup
          run: rm -rf ~/.ssh
        

```

```yaml
name: Manual Deploy Strapi

on:
    workflow_dispatch:
      inputs:
        docker-image-tag:
          description: 'Docker Image Tag'
          required: true
          default: 'latest'
        env-name:
          description: 'Select the environment in which you want to deploy'
          type: choice
          options:
            - test
            - staging
          required: true

jobs:
  manual-deploy:
    name: Deploying Strapi Manually
    runs-on: ubuntu-latest
    steps:
    - name: Code Checkout
      uses: actions/checkout@v3
    
    - name : Set environment variables based on tag 
      env:
        DOCKER_TAG: ${{ github.event.inputs.docker-image-tag }}
        ENV_NAME: ${{ github.event.inputs.env-name }}
        ACTIONS_ALLOW_UNSECURE_COMMANDS: 'true'
      run : |
        if [[ "${{ env.ENV_NAME }}" == "staging" ]]; then
          echo "::set-env name=SERVER::staging"
          echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
          echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
          echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
          echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_STAGING }}"
          echo "::set-env name=DOCKER_TAG::${{ env.DOCKER_TAG }}"
        elif [[ "${{ env.ENV_NAME }}" == "test" ]]; then
          echo "::set-env name=SERVER::test"
          echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
          echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
          echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
          echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_TEST }}" 
          echo "::set-env name=DOCKER_TAG::${{ env.DOCKER_TAG }}"
        else 
          echo "Invalid environment format specified."
          exit 1;
        fi

    - name: install ssh keys
      run: |
        install -m 600 -D /dev/null ~/.ssh/id_rsa
        echo "${{ secrets.SSH_PRIVATE_KEY }}" > ~/.ssh/id_rsa
        ssh-keyscan -H ${{ env.SERVER_IP }} > ~/.ssh/known_hosts
    - name: connect and pull
      run: |
        echo "Using tag: ${{ env.DOCKER_TAG }}"
        echo "Deploying in ${{ env.SERVER }} server"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "docker login ghcr.io -u ${{ secrets.GHCR_USER }} -p ${{ secrets.GHCR_TOKEN_ONLY_READ }} && exit"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && export TAG="${{ env.DOCKER_TAG }}" && docker compose pull && docker compose up -d --force-recreate && exit"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && docker system prune -af && exit"

    - name: cleanup
      run: rm -rf ~/.ssh
```

```yaml
name: Manual Build & Deploy Strapi 

on:
    workflow_dispatch:
      inputs:
        docker-image-tag:
          description: 'Docker Image Tag'
          required: true
          default: 'latest'
        env-name:
          description: 'Select the environment in which you want to deploy'
          type: choice
          options:
            - test
            - staging
          required: true

jobs:
  build-image:
      name: Building APP
      runs-on: ubuntu-latest
      env:
        CI_REGISTRY: ghcr.io
        CI_REGISTRY_IMAGE: ghcr.io/bts-centex/centex-strapi
        DOCKER_TAG: ${{ github.event.inputs.docker-image-tag }}
        ACTIONS_ALLOW_UNSECURE_COMMANDS: 'true'
      steps:    
      - name: Login to Github
        uses: docker/login-action@v3
        with:
          registry: ${{ env.CI_REGISTRY }}
          username: ${{ secrets.GHCR_USER }}
          password: ${{ secrets.GHCR_TOKEN }}
          
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Build and Push Docker Image
        run: |
          docker build . --tag $CI_REGISTRY_IMAGE:$DOCKER_TAG --tag $CI_REGISTRY_IMAGE:latest-test
          docker push $CI_REGISTRY_IMAGE:$DOCKER_TAG
          docker push $CI_REGISTRY_IMAGE:latest-test
  manual-deploy:
    name: Deploying Strapi Manually
    needs: build-image
    runs-on: ubuntu-latest
    steps:
    - name: Code Checkout
      uses: actions/checkout@v3
    
    - name : Set environment variables based on tag 
      env:
        DOCKER_TAG: ${{ github.event.inputs.docker-image-tag }}
        ENV_NAME: ${{ github.event.inputs.env-name }}
        ACTIONS_ALLOW_UNSECURE_COMMANDS: 'true'
      run : |
        if [[ "${{ env.ENV_NAME }}" == "staging" ]]; then
          echo "::set-env name=SERVER::staging"
          echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
          echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
          echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
          echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_STAGING }}"
          echo "::set-env name=DOCKER_TAG::${{ env.DOCKER_TAG }}"
        elif [[ "${{ env.ENV_NAME }}" == "test" ]]; then
          echo "::set-env name=SERVER::test"
          echo "::set-env name=SSH_USER::${{ secrets.SSH_USER }}"
          echo "::set-env name=SERVER_IP::${{ secrets.SSH_HOST }}"
          echo "::set-env name=SSH_KEY::${{ secrets.SSH_PRIVATE_KEY }}"
          echo "::set-env name=WORK_DIR::${{ secrets.WORK_DIR_TEST }}" 
          echo "::set-env name=DOCKER_TAG::${{ env.DOCKER_TAG }}"
        else 
          echo "Invalid environment format specified."
          exit 1;
        fi

    - name: install ssh keys
      run: |
        install -m 600 -D /dev/null ~/.ssh/id_rsa
        echo "${{ secrets.SSH_PRIVATE_KEY }}" > ~/.ssh/id_rsa
        ssh-keyscan -H ${{ env.SERVER_IP }} > ~/.ssh/known_hosts
    - name: connect and pull
      run: |
        echo "Using tag: ${{ env.DOCKER_TAG }}"
        echo "Deploying in ${{ env.SERVER }} server"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "docker login ghcr.io -u ${{ secrets.GHCR_USER }} -p ${{ secrets.GHCR_TOKEN_ONLY_READ }} && exit"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && export TAG="${{ env.DOCKER_TAG }}" && docker compose pull && docker compose up -d --force-recreate && exit"
        ssh ${{ env.SSH_USER }}@${{ env.SERVER_IP }} "cd ${{ env.WORK_DIR }} && docker system prune -af && exit"

    - name: cleanup
      run: rm -rf ~/.ssh
```

```yaml
name: Build main code in PR

on:
  pull_request:
    branches:
      - main

jobs:
  build-image:
      runs-on: ubuntu-latest
      env:
        CI_REGISTRY: ghcr.io
        CI_REGISTRY_IMAGE: ghcr.io/bts-centex/centex-strapi
        ACTIONS_ALLOW_UNSECURE_COMMANDS: true
        SHA8: ${GITHUB_SHA::8}
      steps:
        - name: Extract branch name on push
          if: github.event_name != 'pull_request'
          shell: bash
          run: echo "::set-env name=CI_IMAGE_TAG::${{github.ref_name}}"
          id: extract_branch

        - name: Extract branch name on pull request
          if: github.event_name == 'pull_request'
          run: echo "::set-env name=CI_IMAGE_TAG::${{ env.SHA8 }}" 

        - name: Login to Github
          uses: docker/login-action@v3
          with:
            registry: ${{ env.CI_REGISTRY }}
            username: ${{ secrets.GHCR_USER }}
            password: ${{ secrets.GHCR_TOKEN }}
            
        - name: Checkout repository
          uses: actions/checkout@v3

        - name: Build and Push Docker Image
          run: |
            docker build . --tag $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG --tag $CI_REGISTRY_IMAGE:latest
            docker push $CI_REGISTRY_IMAGE:$CI_IMAGE_TAG
            docker push $CI_REGISTRY_IMAGE:latest

```

```yaml
name: 'Bump Version'

on:
  push:
    paths-ignore:
      - '.github'
    branches:
      - 'main'

jobs:
  bump-version:
    name: 'Bump Version on master'
    runs-on: ubuntu-latest
    permissions:
      contents: write

    steps:
      - name: 'Checkout source code'
        uses: 'actions/checkout@v2'
        with:
          ref: ${{ github.ref }}
      - name: 'cat package.json'
        run: cat ./package.json
      - name: 'Automated Version Bump'
        id: version-bump
        uses: 'phips28/gh-action-bump-version@master'
        with:
          version-type: 'patch'
          tag-suffix: '-beta'
          commit-message: 'CI: bumps version to {{version}} [skip ci]'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      - name: 'cat package.json'
        run: cat ./package.json
      - name: 'Output Step'
        env:
          NEW_TAG: ${{ steps.version-bump.outputs.newTag }}
        run: echo "new tag $NEW_TAG"

```

