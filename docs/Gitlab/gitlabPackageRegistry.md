## Gitlab Package Registry



##### Agregar registro vía terminal

Primero agregaremos el registro del package en el package.json

```json
{
"publishConfig": { "@foo:registry":" https://gitlab.example.com/api/v4/projects/<your_project_id>/packages/npm/" }
}
```

Ejemplo:

```json
  "publishConfig": {
    "@rafaelm:registry": "https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/"
  },
```

![image-20220912110456169](C:\Users\rafab\AppData\Roaming\Typora\typora-user-images\image-20220912110456169.png)

Para publicar nuestros paquetes debemos crear un archivo de configuracion llamado `.npmrc`

este contendrá el url del registro y su token de autentificacion.

```
.npmrc
@foo:registry=https://gitlab.example.com/api/v4/projects/${CI_PROJECT_ID}/packages/npm/
//gitlab.example.com/api/v4/projects/${CI_PROJECT_ID}/packages/npm/:_authToken=${CI_JOB_TOKEN}
```

Ejemplo:

```
//gitlab.bluetrail.software/api/v4/projects/486/packages/npm/:_authToken=LhPRciy4nBNeuUS_g_kY
@rafaelm:registry https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/
registry=https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/
```

Para agregar estas credenciales vía terminal agregaremos estos dos comandos.

```
npm config set -- '//gitlab.example.com/api/v4/projects/<your_project_id>/packages/npm/:_authToken' "${NPM_TOKEN}"
npm config set -- '//gitlab.example.com/api/v4/packages/npm/:_authToken' "${NPM_TOKEN}"
```

Una vez tengamos esta configuración preparada solo necesitamos lanzar el siguiente comando.

`npm publish`

`NPM_TOKEN=<your_token> npm publish`

`npm publish --registry=https://gitlab.bluetrail.software/api/v4/projects/486/packages/npm/`

> Es importante agregar el .npmrc al gitignore.



##### Lanza registro vía pipeline.

Debemos crear un token de despliegue y agregarlo a las variables de GITLAB.

Agregaremos un stage a nuestra pipeline con la siguiente configuración.

```yaml
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
```

