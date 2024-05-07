## Actions Example Cartier

```yaml
name: Full_stage

on:
  push:
    branches:
      - develop
  pull_request:
    branches:
      - develop

jobs:
  build_and_test:
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:13
        env:
          PG_HOST: postgres
          POSTGRES_DB: alg_database_test
          POSTGRES_USER: ${{ secrets.POSTGRES_USER }}
          POSTGRES_PASSWORD: ${{ secrets.POSTGRES_PASSWORD }}
          POSTGRES_HOST_AUTH_METHOD: trust
        ports:
          - 5432:5432
        options: --name=postgres --health-cmd pg_isready --health-interval 10s --health-timeout 5s --health-retries 5

      wg-easy:
        image: weejewel/wg-easy:7
        env:
          WG_HOST: 0.0.0.0
          PASSWORD: wireguard-test-network
        ports:
          - '51820:51820/udp'
          - '51821:51821/tcp'
        options: '--privileged'

    steps:
      - name: Checkout code
        uses: actions/checkout@v3.5.3
      - name: Test user
        run: |
          sudo su - postgres
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE runner WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE root WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
      - name: Wait for PostgreSQL to be ready
        run: |
          until nc -zv localhost 5432; do
            echo "Waiting for PostgreSQL to be ready..."
            sleep 2
          done

      - name: Set up Node.js
        uses: actions/setup-node@v3.8.1
        with:
          node-version: '18.16.0'

      - name: Install dependencies and build
        run: |
          npm install --omit-dev
          npm run build
          cd public
          npm install --omit-dev
          npm run build
        working-directory: ${{ github.workspace }}
        # continue-on-error: true

      - name: Test application
        run: |
          npm run autotest

      - name: Upload test results
        uses: actions/upload-artifact@v3.1.2
        with:
          name: test-results
          path: ${{ github.workspace }}/junit.xml

  build-image:
    runs-on: ubuntu-latest
    needs: build_and_test
    env:
      CI_REGISTRY_IMAGE: registry.gitlab.com/cartier-lab/alg-backend
    steps:
      - name: Check Docker Info
        run: docker info

      - name: Docker Login
        run: docker login registry.gitlab.com -u ${{ secrets.GITLAB_USER_REGISTRY }} -p ${{ secrets.GITLAB_USER_PWD }}

      - name: Check Docker and Docker Compose Versions
        run: |
          docker --version
          docker-compose version

      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Read alg version
        id: read-image-version
        run: echo "image_version=$(cat changelog/version.md)_develop" >> $GITHUB_OUTPUT

      - name: Check version
        run: |
          echo ${{ steps.read-image-version.outputs.image_version }}

      - name: 'base64'
        run: |
          echo -n "${{ secrets.GCP_PRIVATE_KEY_RW  }}" | base64 -d > gcp_credentials-rw.json
          echo -n "${{ secrets.GCP_PRIVATE_KEY_R  }}" | base64 -d > gcp_credentials-r.json
          echo -n "${{ secrets.BUILD_ARGS  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_POSTGRES  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_ELASTIC  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_API  }}" | base64 -d >> build.args

      - name: Build and Push Docker Image
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker build . --tag $CI_REGISTRY_IMAGE:latest --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA \
            $(out=""; for i in $(cat build.args); do out="$out--build-arg $i "; done; echo "$out"; out="")
          docker push $CI_REGISTRY_IMAGE:$VERSION
          docker push $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA
          docker push $CI_REGISTRY_IMAGE:latest

      - name: Check Vars
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker run $CI_REGISTRY_IMAGE:$VERSION env

```

```yaml
name: Full_stage

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build_and_test:
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:13
        env:
          PG_HOST: postgres
          POSTGRES_DB: alg_database_test
          POSTGRES_USER: ${{ secrets.POSTGRES_USER }}
          POSTGRES_PASSWORD: ${{ secrets.POSTGRES_PASSWORD }}
          POSTGRES_HOST_AUTH_METHOD: trust
        ports:
          - 5432:5432
        options: --name=postgres --health-cmd pg_isready --health-interval 10s --health-timeout 5s --health-retries 5

      wg-easy:
        image: weejewel/wg-easy:7
        env:
          WG_HOST: 0.0.0.0
          PASSWORD: wireguard-test-network
        ports:
          - '51820:51820/udp'
          - '51821:51821/tcp'
        options: '--privileged'

    steps:
      - name: Checkout code
        uses: actions/checkout@v3.5.3
      - name: Test user
        run: |
          sudo su - postgres
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE runner WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE root WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
      - name: Wait for PostgreSQL to be ready
        run: |
          until nc -zv localhost 5432; do
            echo "Waiting for PostgreSQL to be ready..."
            sleep 2
          done

      - name: Set up Node.js
        uses: actions/setup-node@v3.8.1
        with:
          node-version: '18.16.0'

      - name: Install dependencies and build
        run: |
          npm install --omit-dev
          npm run build
          cd public
          npm install --omit-dev
          npm run build
        working-directory: ${{ github.workspace }}
        # continue-on-error: true

      - name: Test application
        run: npm run autotest

      - name: Upload test results
        uses: actions/upload-artifact@v3.1.2
        with:
          name: test-results
          path: ${{ github.workspace }}/junit.xml

  build-image:
    runs-on: ubuntu-latest
    needs: build_and_test
    env:
      CI_REGISTRY_IMAGE: registry.gitlab.com/cartier-lab/alg-backend
    steps:
      - name: Check Docker Info
        run: docker info

      - name: Docker Login
        run: docker login registry.gitlab.com -u ${{ secrets.GITLAB_USER_REGISTRY }} -p ${{ secrets.GITLAB_USER_PWD }}

      - name: Check Docker and Docker Compose Versions
        run: |
          docker --version
          docker-compose version

      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Read alg version
        id: read-image-version
        run: echo "image_version=$(cat changelog/version.md)_main" >> $GITHUB_OUTPUT

      - name: Check version
        run: |
          echo ${{ steps.read-image-version.outputs.image_version }}

      - name: 'base64'
        run: |
          echo -n "${{ secrets.GCP_PRIVATE_KEY_RW  }}" | base64 -d > gcp_credentials-rw.json
          echo -n "${{ secrets.GCP_PRIVATE_KEY_R  }}" | base64 -d > gcp_credentials-r.json
          echo -n "${{ secrets.BUILD_ARGS  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_POSTGRES  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_ELASTIC  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_API  }}" | base64 -d >> build.args

      - name: Build and Push Docker Image
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker build . --tag $CI_REGISTRY_IMAGE:latest --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA \
            $(out=""; for i in $(cat build.args); do out="$out--build-arg $i "; done; echo "$out"; out="")
          docker push $CI_REGISTRY_IMAGE:$VERSION
          docker push $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA
          docker push $CI_REGISTRY_IMAGE:latest

      - name: Check Vars
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker run $CI_REGISTRY_IMAGE:$VERSION env

```

```yaml
name: Full_stage

on:
  push:
    branches:
      - 'release*'
  pull_request:
    branches:
      - 'release*'

jobs:
  build_and_test:
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:13
        env:
          PG_HOST: postgres
          POSTGRES_DB: alg_database_test
          POSTGRES_USER: ${{ secrets.POSTGRES_USER }}
          POSTGRES_PASSWORD: ${{ secrets.POSTGRES_PASSWORD }}
          POSTGRES_HOST_AUTH_METHOD: trust
        ports:
          - 5432:5432
        options: --name=postgres --health-cmd pg_isready --health-interval 10s --health-timeout 5s --health-retries 5

      wg-easy:
        image: weejewel/wg-easy:7
        env:
          WG_HOST: 0.0.0.0
          PASSWORD: wireguard-test-network
        ports:
          - '51820:51820/udp'
          - '51821:51821/tcp'
        options: '--privileged'

    steps:
      - name: Checkout code
        uses: actions/checkout@v3.5.3
      - name: Test user
        run: |
          sudo su - postgres
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE runner WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
          PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE root WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
      - name: Wait for PostgreSQL to be ready
        run: |
          until nc -zv localhost 5432; do
            echo "Waiting for PostgreSQL to be ready..."
            sleep 2
          done

      - name: Set up Node.js
        uses: actions/setup-node@v3.8.1
        with:
          node-version: '18.16.0'

      - name: Install dependencies and build
        run: |
          npm install --omit-dev
          npm run build
          cd public
          npm install --omit-dev
          npm run build
        working-directory: ${{ github.workspace }}
        # continue-on-error: true

      - name: Test application
        run: npm run autotest

      - name: Upload test results
        uses: actions/upload-artifact@v3.1.2
        with:
          name: test-results
          path: ${{ github.workspace }}/junit.xml

  build-image:
    runs-on: ubuntu-latest
    needs: build_and_test
    env:
      CI_REGISTRY_IMAGE: registry.gitlab.com/cartier-lab/alg-backend
    steps:
      - name: Check Docker Info
        run: docker info

      - name: Docker Login
        run: docker login registry.gitlab.com -u ${{ secrets.GITLAB_USER_REGISTRY }} -p ${{ secrets.GITLAB_USER_PWD }}

      - name: Check Docker and Docker Compose Versions
        run: |
          docker --version
          docker-compose version

      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Read alg version
        id: read-image-version
        run: echo "image_version=release_$(cat changelog/version.md)" >> $GITHUB_OUTPUT

      - name: Check version
        run: |
          echo ${{ steps.read-image-version.outputs.image_version }}

      - name: 'base64'
        run: |
          echo -n "${{ secrets.GCP_PRIVATE_KEY_RW  }}" | base64 -d > gcp_credentials-rw.json
          echo -n "${{ secrets.GCP_PRIVATE_KEY_R  }}" | base64 -d > gcp_credentials-r.json
          echo -n "${{ secrets.BUILD_ARGS  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_POSTGRES  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_ELASTIC  }}" | base64 -d >> build.args
          echo -n "${{ secrets.BUILD_ARGS_API  }}" | base64 -d >> build.args

      - name: Build and Push Docker Image
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker build . --tag $CI_REGISTRY_IMAGE:latest --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA \
            $(out=""; for i in $(cat build.args); do out="$out--build-arg $i "; done; echo "$out"; out="")
          docker push $CI_REGISTRY_IMAGE:$VERSION
          docker push $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA
          docker push $CI_REGISTRY_IMAGE:latest

      - name: Check Vars
        run: |
          export VERSION=${{ steps.read-image-version.outputs.image_version }}
          docker run $CI_REGISTRY_IMAGE:$VERSION env

```

```yaml
name: Full_stage

on:
  push:
    branches:
      - test
      # - LG-1523-v2
  #     - develop
  #     - 'release/*'
  # pull_request:
  #   branches:
  #     - main
  #     - develop
  #     - 'release/*'

jobs:
  # build_and_test:
  #   runs-on: ubuntu-latest
    
  #   services:
  #     postgres:
  #       image: postgres:13
  #       env:
  #         PG_HOST: postgres
  #         POSTGRES_DB: alg_database_test
  #         POSTGRES_USER: ${{ secrets.POSTGRES_USER }}
  #         POSTGRES_PASSWORD: ${{ secrets.POSTGRES_PASSWORD }}
  #         POSTGRES_HOST_AUTH_METHOD: trust
  #       ports:
  #         - 5432:5432
  #       options: --name=postgres --health-cmd pg_isready --health-interval 10s --health-timeout 5s --health-retries 5
        
  #   steps:
  #   - name: Checkout code
  #     uses: actions/checkout@v3.5.3
  #   - name: Test user
  #     run: |
  #       sudo su - postgres
  #       PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE runner WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
  #       PGPASSWORD=${{ secrets.POSTGRES_PASSWORD }} psql -h localhost -p 5432 -U ${{ secrets.POSTGRES_USER }} -c "CREATE ROLE root WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '${{ secrets.POSTGRES_PASSWORD }}';"
  #   - name: Wait for PostgreSQL to be ready
  #     run: |
  #       until nc -zv localhost 5432; do
  #         echo "Waiting for PostgreSQL to be ready..."
  #         sleep 2
  #       done
      
  #   - name: Set up Node.js
  #     uses: actions/setup-node@v3.8.1
  #     with:
  #       node-version: '16'
        
  #   - name: Install dependencies and build
  #     run: |
  #       npm install --omit-dev
  #       npm run build
  #       cd public
  #       npm install --omit-dev
  #       npm run build
  #     working-directory: ${{ github.workspace }}
  #     # continue-on-error: true
      
  #   - name: Test application
  #     run: npm run autotest

  #   - name: Upload test results
  #     uses: actions/upload-artifact@v3.1.2
  #     with:
  #       name: test-results
  #       path: ${{ github.workspace }}/junit.xml
        
  build-image:
      runs-on: ubuntu-latest
      # needs: build_and_test
      env:
        CI_REGISTRY_IMAGE: ${{secrets.DOCKER_REPO}}/${{github.repository}}
      steps:
        - name: Check Docker Info
          run: docker info
  
        - name: Docker Login
          run: | 
            echo ${{secrets.DOCKER_REPO_PASS}} | docker login -u ${{secrets.DOCKER_REPO_USER}} --password-stdin ${{secrets.DOCKER_REPO}}
            docker login registry.gitlab.com -u ${{ secrets.GITLAB_USER_REGISTRY }} -p ${{ secrets.GITLAB_USER_PWD }}

        - name: Check Docker and Docker Compose Versions
          run: |
            docker --version
            docker-compose version

        - name: Checkout repository
          uses: actions/checkout@v3
          
        - name: Read alg version
          id: read-image-version
          run: echo "image_version=$(cat changelog/version.md)_fuse" >> $GITHUB_OUTPUT

        - name: Check version
          run: | 
            echo ${{ steps.read-image-version.outputs.image_version }}
          
        - name: Build and Push Docker Image
          run: |
            export VERSION=${{ steps.read-image-version.outputs.image_version }}
            docker build . --tag $CI_REGISTRY_IMAGE:$VERSION --tag $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA --tag registry.gitlab.com/cartier-lab/alg-backend:$VERSION \
              $(out=""; for i in $(cat ./build.args | base64 -d); do out="$out--build-arg $i "; done; echo "$out"; out="")
            docker push $CI_REGISTRY_IMAGE:$VERSION
            docker push $CI_REGISTRY_IMAGE:$VERSION_$GITHUB_SHA
            docker push registry.gitlab.com/cartier-lab/alg-backend:$VERSION
        
        - name: Check Vars
          run: |
            export VERSION=${{ steps.read-image-version.outputs.image_version }}
            docker run $CI_REGISTRY_IMAGE:$VERSION env
```

