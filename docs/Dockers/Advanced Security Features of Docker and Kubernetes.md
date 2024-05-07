# Advanced Security Features of Docker and Kubernetes

### Securing Docker 

#### Docker Rootles (Limited)

######   https://docs.docker.com/engine/security/rootless/

#### Docker socket

No es recomendable compartir el socket de docker con otro docker (Jenkins) evitarla o aplicar defensa en profuncidad.

TCP socket: Por defecto esta sin ecriptar, sin cifrar y sin autenticar. (Shodan --> Products Docker 2375:2375)  

- Alternativa (build in https ecrypted docker) - create CA server key OpenSSL
- Run docker daemon TSL
  - `dockerd --tlsverify --tlscacert=ca.pem --tlscert=server-cert.pem --tlskey=serrver-key.pem -H=0.0.0.0:2376`
- Remote API certificated
  - `docker --tlsverify --tlscacert=ca.pem --tlscert=server-cert.pem --tlskey=serrver-key.pem -H=<host>:2376 version`
  - `export DOCKER_HOST=tecp://<hsot>:2376 DOCKER_TSL_VERIFY=1`

#### Securing docker containers

-  kernel capabilities (don't run --privileged)  show capabilities(ps -fC) 
  - `docker run --cap-drop=all --cap-add=cap_net_bind_service -p 80:80 nginx`
- Seccomp (default)
  - `docker run --security-opt seccomp=/pth/to/seccomp/profile.json <myapp>`
  - seccomp-gen (gitlhub)

-  apparmor

  - genuinetools/bane (github)
  - `sudo apparmor parser -r -W /path/to/your/apparmor-nginx-profile`
  - `docker run -d -security-opt "apparmor-apparmor-profile-name"-p 80:80 nginx`

- root (by default)

  - `docker run -it --user 2000 alpine sh` (user rootless)

  - user remap (docker run with root privileges but host process are runing without privileges)

    - /etc/docker/daemon.json

    - {

      "userns-remap":"default"

      }

- docker container trust (signed images)

  - Step 1: $ DOCKER_CONTENT_TRUST=1
    Step 2: $ docker trust key generate syourname>
    Step 3: $ docker trust signer add -key syour-key.pub> <your-name> your-repo>

### Securing Kubernetes Enviroments





   