# Full Environment dockerized



##### Requirements:

- Repositories
  - bts_internship_2019_fe_app - Branch: **Docker/master**
    - git@gitlab.bluetrail.software:bts-platform/bts_internship_2019_fe_app.git
  - mybts_fe - Branch: **DT-301**
    - git@gitlab.bluetrail.software:mybts-platform/mybts_fe.git
  - mybts_be - Branch: **DT-283**
    - git@gitlab.bluetrail.software:mybts-platform/mybts_be.git
- Docker installed
- Keep free ports 80,3000,4200,5050,5433
- PgAdmin installed

#### Steps:

Step 1: 

- Download the repositories and checkout to specified branch.

Step 2:

- Add the env files inside specified folders.

  - **environment.prod.ts** on bts_internship_2019_fe_app\src\environments\environment.prod.ts

  ```
  export const environment = {
    production: false,
    apiUrl: 'api',
    url: 'http://localhost:3000/',
    citiesApiUrl: 'https://autocomplete.travelpayouts.com',
    managerIdRole: 1,
    s3BucketUrl: 'https://s3.amazonaws.com/cdn.platform.bluetrail.software/prod/',
  };
  ```

  - .**env.default** on mybts_be/.env.default

  ```
  # Database config
  DB_HOST=postgres
  DB_USER=postgres
  DB_PASSWORD=postgres
  DB_NAME=postgres
  DB_PORT=5433
  DB_SCHEMA=public
  TOKEN_SECRET=secret
  TIME_DURATION=24
  TIME_INTERVAL=hours
  TIME_EXPIRATION=24:00:00
  SECURITY_SALT=10
  FRONT_HOST=http://localhost
  MAX_ID_RANGE=2147483647
  
  # Email credentials
  MAIL_SENDER=no-reply@bluetrailsoft.com
  MAIL_HOST=
  MAIL_PORT=465
  MAIL_USER=
  MAIL_PASS=
  
  APP_PORT=3000
  API_URL=http://back:3000
  
  # AWS S3 Credentials
  #S3_BUCKET=cdn.platform.bluetrail.software
  #S3_FOLDER=dev
  #AWS_ACCESS_KEY_ID=
  #AWS_SECRET_ACCESS_KEY=
  #AWS_REGION=us-east-1
  #S3_FOLDER_BACKUP=platformBackups/dev_backups
  #S3_PROJECT_FILES_FOLDER=dev
  
  
  S3_BUCKET=mybts-s3
  S3_FOLDER=dev
  AWS_ACCESS_KEY_ID=
  AWS_SECRET_ACCESS_KEY=
  AWS_REGION=us-east-2
  S3_FOLDER_BACKUP=platformBackups/dev_backups
  S3_PROJECT_FILES_FOLDER=projectFiles
  
  V2_TIME_DURATION=24
  V2_TIME_INTERVAL=hours
  V2_TIME_EXPIRATION=24:00:00
  
  # Weather API
  WEATHER_API_KEY = 134e7fd1f8e8a5365e270b8ad776300b
  WEATHER_UPDATE_FREQUENCY = 1
  
  # PROTECTED ID
  MANAGER_ID_ROLE= 1
  REDIRECT_URIS=["http://mybts.bluetrail.software/api/V2/module0/login/redirect"]
   # GOOGLE OAUTH 2.0 Credentials
  CLIENT_ID=
  PROJECT_ID=bts-platform-v2
  AUTH_URI=https://accounts.google.com/o/oauth2/auth
  TOKEN_URI=https://oauth2.googleapis.com/token
  AUTH_PROVIDER_X509_CERT_URL=https://www.googleapis.com/oauth2/v3/certs
  CLIENT_SECRET=
  REDIRECT_URIS=["http://mybts.bluetrail.software/api/V2/module0/login/redirect"]
  
  JAVASCRIPT_ORIGINS=["http://localhost:4200","http://localhost","https://dev.platform.bluetrail.software","https://dev.platform.bluetrail.software:80","http://dev.platform.bluetrail.software","http://dev.platform.bluetrail.software:80","http://mybts-qa.bluetrail.software","https://mybts-qa.bluetrail.software",http://mybts.bluetrail.software]
  
  ```

  .**env.production** on mybts_fe/.env.production

  ```
  REACT_APP_CLIENT_ID=
  REACT_APP_BACKEND=http://mybts.bluetrail.software
  REACT_APP_PORT=80
  REACT_APP_BACKEND_VERSION=/api/V2
  REACT_APP_S3_FOLDER=dev
  REACT_APP_S3_BUCKET=""
  REACT_APP_S3_ROOT=https://mybts-s3.s3.amazonaws.com
  ```

  

Step 3:

- Launch build image from dockerfile in all repositories with specified names
  - Legacy FE no needs build.
  - BE build name: **platform_be_local**
    - With the terminal on the root repository folder execute: `docker build -t platform_be_local .`
  - FE build name: **mybts**
    - With the terminal on the root repository folder execute: `docker build -t mybts .`

Step 4:

- On bts_internship_2019_fe_app - Branch: **Docker/master**
  - Launch docker-compose.yml with `docker-compose up -d`

Step 5:

- Register your local server on PgAdmin
  - Register
    - Name: "local"
  - Connection
    - host: 127.0.0.1 or localhost
    - port: 5433
    - username: postgres
    - password: postgres
    - database: postgres

- Restore backup with pgadmin on schema public.
  - Go to database postgres
  - Drop cascade schema public
  - Create new schema public
  - Restore UAT backup
- Go to permissions table and change name of the column **description** to **name**.

Step 6:

- Add this lines on */etc/hosts* or */system32/drivers/etc/hosts*

```
127.0.0.1 mybts.bluetrail.software
127.0.0.1:3000 legacy.test
localhost mybts.bluetrail.software
```

Step 7:

- Test environment
  - FE legacy -  http://legacy.test doesn't work ATM (Because database are not adapted)
  - FE mybts - http://mybts.bluetrail.software/ Working