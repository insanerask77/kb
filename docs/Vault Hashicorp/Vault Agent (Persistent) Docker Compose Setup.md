# Vault Agent (Persistent) Docker Compose Setup

May 01, 2022

​      [        ![vault](https://d33wubrfki0l68.cloudfront.net/5045c92d1afa5de2fdf624342bcbe019edf5fb98/4f46d/static/c3ae3a7109d48731ebbf7502f82decd8/1c72d/vault.jpg)   ](https://www.spektor.dev/static/c3ae3a7109d48731ebbf7502f82decd8/c08c5/vault.jpg)    

**TL;DR:** You can find the code in this Github [repo](https://github.com/yossisp/vault-agent-docker-compose).

Recently I needed to integrate [Hashicorp Vault](https://www.hashicorp.com/products/vault) with a Java application. For local development I wanted to use [Vault Agent](https://www.vaultproject.io/docs/agent) which can connect to the Vault server. The advantage of using Vault  Agent is that it bears the brunt of authentication complexity with Vault server (including SSL certificates). Effectively, this means that a  client application can send HTTP requests to Vault Agent without any  need to authenticate. This setup is frequently used in the real world  for example by using [Agent Sidecar Injector](https://www.vaultproject.io/docs/platform/k8s/injector) inside a Kubernetes cluster. It makes it easy for client applications  inside a K8s pod to get/put information to a Vault server without each  one having to perform the tedious authentication process.

Surprisingly, I couldn’t find much information on using Vault with  Vault Agent via docker-compose, which in my opinion is by far the  easiest method to set up a Vault playground. I did find [this](https://gitlab.com/kawsark/vault-agent-docker/-/tree/master) example which served as the inspiration for this post however it  involves a more complex setup as well as using Postgres and Nginx. I’d  like to present the most minimal setup, the bare basics needed to spin  up a Vault Agent and access it locally via `localhost`.

**WARNING:** the setup is intentionally simplified, please don’t use it in production.

First of all we’ll use the official Vault docker images for the `docker-compose.yml`:

```yml
version: '3.7'

services:
  vault-agent:
    image: hashicorp/vault:1.9.6
    restart: always
    ports:
      - "8200:8200"
    volumes:
      - ./helpers:/helpers
    environment:
      VAULT_ADDR: "http://vault:8200"
    container_name: vault-agent
    entrypoint: "vault agent -log-level debug -config=/helpers/vault-agent.hcl"
    depends_on:
      vault:
        condition: service_healthy
  vault:
    image: hashicorp/vault:1.9.6
    restart: always
    volumes:
      - ./helpers:/helpers
      - vault_data:/vault/file
    ports:
      - "8201:8200/tcp"
    cap_add:
      - IPC_LOCK
    container_name: vault
    entrypoint: "vault server -config=/helpers/vault-config.hcl"
    healthcheck:
      test: wget --no-verbose --tries=1 --spider http://localhost:8200 || exit 1
      interval: 10s
      retries: 12
      start_period: 10s
      timeout: 10s

volumes:
  vault_data: {}
```

Here we’re using the same image to start Vault server in dev mode as  well as start the Vault Agent. In addition a volume is created for `helpers` directory which will contain:

1. The policy for Vault server `admin-policy.hcl`:

   ```hcl
   path "sys/health"
   {
   capabilities = ["read", "sudo"]
   }
   path "sys/policies/acl"
   {
   capabilities = ["list"]
   }
   path "sys/policies/acl/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "auth/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "sys/auth/*"
   {
   capabilities = ["create", "update", "delete", "sudo"]
   }
   path "sys/auth"
   {
   capabilities = ["read"]
   }
   path "kv/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "secret/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "identity/entity-alias"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "identity/entity-alias/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "identity/entity"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "identity/entity/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "sys/mounts/*"
   {
   capabilities = ["create", "read", "update", "delete", "list", "sudo"]
   }
   path "sys/mounts"
   {
   capabilities = ["read"]
   }
   ```

2. The policy for Vault Agent `vault-agent.hcl`:

   ```hcl
   pid_file = "./pidfile"
   vault {
   address = "http://vault:8200"
   retry {
   num_retries = 5
   }
   }
   auto_auth {
   method {
   type = "approle"
   config = {
     role_id_file_path = "/helpers/role_id"
     secret_id_file_path = "/helpers/secret_id"
     remove_secret_id_file_after_reading = false
   }
   }
   sink "file" {
   config = {
     path = "/helpers/sink_file"
   }
   }
   }
   cache {
   use_auto_auth_token = true
   }
   listener "tcp" {
   address = "0.0.0.0:8200"
   tls_disable = true
   }
   ```

3. The `init.sh` script which will create AppRole auth method:

   ```bash
   apk add jq curl
   export VAULT_ADDR=http://localhost:8200
   root_token=$(cat /helpers/keys.json | jq -r '.root_token')
   unseal_vault() {
   export VAULT_TOKEN=$root_token
   vault operator unseal -address=${VAULT_ADDR} $(cat /helpers/keys.json | jq -r '.keys[0]')
   vault login token=$VAULT_TOKEN
   }
   if [[ -n "$root_token" ]]
   then
     echo "Vault already initialized"
     unseal_vault
   else
     echo "Vault not initialized"
     curl --request POST --data '{"secret_shares": 1, "secret_threshold": 1}' http://127.0.0.1:8200/v1/sys/init > /helpers/keys.json
     root_token=$(cat /helpers/keys.json | jq -r '.root_token')
   
     unseal_vault
   
     vault secrets enable -version=2 kv
     vault auth enable approle
     vault policy write admin-policy /helpers/admin-policy.hcl
     vault write auth/approle/role/dev-role token_policies="admin-policy"
     vault read -format=json auth/approle/role/dev-role/role-id \
       | jq -r '.data.role_id' > /helpers/role_id
     vault write -format=json -f auth/approle/role/dev-role/secret-id \
       | jq -r '.data.secret_id' > /helpers/secret_id
   fi
   printf "\n\nVAULT_TOKEN=%s\n\n" $VAULT_TOKEN
   ```

4. Below is the config for the Vault server to be saved in `vault-config.hcl` file:

```hcl
storage "file" {
  # this path is used so that volume can be enabled https://hub.docker.com/_/vault
  path = "/vault/file"
}

listener "tcp" {
  address     = "0.0.0.0:8200"
  tls_disable = "true"
}

api_addr = "http://127.0.0.1:8200"
cluster_addr = "https://127.0.0.1:8201"
ui = true
```

Next we’ll create `startVault.sh` script to start Vault:

```shell
WAIT_FOR_TIMEOUT=120 # 2 minutes
docker-compose up --detach
echo Waiting for Vault Agent container to be up
curl https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for | sh -s -- localhost:8200 -t $WAIT_FOR_TIMEOUT -- echo success
docker exec vault /bin/sh -c "source /helpers/init.sh"
docker restart vault-agent
```

```
#!/bin/sh

# The MIT License (MIT)
#
# Copyright (c) 2017 Eficode Oy
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

VERSION="2.2.3"

set -- "$@" -- "$TIMEOUT" "$QUIET" "$PROTOCOL" "$HOST" "$PORT" "$result"
TIMEOUT=15
QUIET=0
# The protocol to make the request with, either "tcp" or "http"
PROTOCOL="tcp"

echoerr() {
  if [ "$QUIET" -ne 1 ]; then printf "%s\n" "$*" 1>&2; fi
}

usage() {
  exitcode="$1"
  cat << USAGE >&2
Usage:
  $0 host:port|url [-t timeout] [-- command args]
  -q | --quiet                        Do not output any status messages
  -t TIMEOUT | --timeout=timeout      Timeout in seconds, zero for no timeout
  -v | --version                      Show the version of this tool
  -- COMMAND ARGS                     Execute command with args after the test finishes
USAGE
  exit "$exitcode"
}

wait_for() {
  case "$PROTOCOL" in
    tcp)
      if ! command -v nc >/dev/null; then
        echoerr 'nc command is missing!'
        exit 1
      fi
      ;;
    http)
      if ! command -v wget >/dev/null; then
        echoerr 'wget command is missing!'
        exit 1
      fi
      ;;
  esac

  TIMEOUT_END=$(($(date +%s) + TIMEOUT))

  while :; do
    case "$PROTOCOL" in
      tcp) 
        nc -w 1 -z "$HOST" "$PORT" > /dev/null 2>&1
        ;;
      http)
        wget --timeout=1 -q "$HOST" -O /dev/null > /dev/null 2>&1 
        ;;
      *)
        echoerr "Unknown protocol '$PROTOCOL'"
        exit 1
        ;;
    esac

    result=$?
        
    if [ $result -eq 0 ] ; then
      if [ $# -gt 7 ] ; then
        for result in $(seq $(($# - 7))); do
          result=$1
          shift
          set -- "$@" "$result"
        done

        TIMEOUT=$2 QUIET=$3 PROTOCOL=$4 HOST=$5 PORT=$6 result=$7
        shift 7
        exec "$@"
      fi
      exit 0
    fi

    if [ $TIMEOUT -ne 0 -a $(date +%s) -ge $TIMEOUT_END ]; then
      echo "Operation timed out" >&2
      exit 1
    fi

    sleep 1
  done
}

while :; do
  case "$1" in
    http://*|https://*)
    HOST="$1"
    PROTOCOL="http"
    shift 1
    ;;
    *:* )
    HOST=$(printf "%s\n" "$1"| cut -d : -f 1)
    PORT=$(printf "%s\n" "$1"| cut -d : -f 2)
    shift 1
    ;;
    -v | --version)
    echo $VERSION
    exit
    ;;
    -q | --quiet)
    QUIET=1
    shift 1
    ;;
    -q-*)
    QUIET=0
    echoerr "Unknown option: $1"
    usage 1
    ;;
    -q*)
    QUIET=1
    result=$1
    shift 1
    set -- -"${result#-q}" "$@"
    ;;
    -t | --timeout)
    TIMEOUT="$2"
    shift 2
    ;;
    -t*)
    TIMEOUT="${1#-t}"
    shift 1
    ;;
    --timeout=*)
    TIMEOUT="${1#*=}"
    shift 1
    ;;
    --)
    shift
    break
    ;;
    --help)
    usage 0
    ;;
    -*)
    QUIET=0
    echoerr "Unknown option: $1"
    usage 1
    ;;
    *)
    QUIET=0
    echoerr "Unknown argument: $1"
    usage 1
    ;;
  esac
done

if ! [ "$TIMEOUT" -ge 0 ] 2>/dev/null; then
  echoerr "Error: invalid timeout '$TIMEOUT'"
  usage 3
fi

case "$PROTOCOL" in
  tcp)
    if [ "$HOST" = "" ] || [ "$PORT" = "" ]; then
      echoerr "Error: you need to provide a host and port to test."
      usage 2
    fi
  ;;
  http)
    if [ "$HOST" = "" ]; then
      echoerr "Error: you need to provide a host to test."
      usage 2
    fi
  ;;
esac

wait_for "$@"
```

After you created the above files in the `helpers` directory, the project structure should be as follows:

```text
.
├── docker-compose.yml
├── helpers
│   ├── admin-policy.hcl
│   ├── init.sh
│   ├── vault-agent.hcl
│   └── vault-config.hcl
└── startVault.sh
```

Finally, run `source startVault.sh` to start Vault server and Vault Agent.

Now any client application can access Vault Agent over `http://localhost:8200` on the host machine, for example the following command creates a secret name `hello`:

```shell
curl --request POST -H "Content-Type: application/json"  \
--data '{"data":{"foo":"bar"}}' http://localhost:8200/v1/kv/data/hello
```

while this command retrieves the secret name `hello`:

```shell
curl http://localhost:8200/v1/kv/data/hello
```

In addition Vault web UI is available at `http://localhost:8201/ui`. In order to log into the UI use the value of `root_token` field in `./helpers/key.json` file (using token login method in the UI).

Vault server uses file storage backend which makes this a persistent  setup (a docker volume is mounted), so that tokens data will persist  after machine restart or running `docker-compose down`.