## Introduction

Docker Compose is a command-line tool that uses a specially formatted YAML file as input to assemble and then run single, or multiple, containers as applications. This allows developers to develop, test and then deliver to their users a single YAML file for their application, and use only one command to start, and stop, it reliably. This portability and reliability has made Docker Compose not only hugely popular with both users and developers, but increasingly a requirement.

### Objectives

This lab shows how to install and use both `podman-compose` and `docker-compose` with Podman, and verify they work with a simple `docker-compose.yaml` file.

### Prerequisites

- A client system with Oracle Linux installed
- Podman installed (the ‘container-tools’ package)
- Access to the Internet

## Oracle Support Disclaimer

Oracle does not provide technical support for the sequence of steps that are provided in the following instructions because these steps refer to software programs and operating systems that are not provided by Oracle. This tutorial provides optional instructions as a convenience only.

For more information about Oracle’s supported method for the development and usage of Podman-based services see https://docs.oracle.com/en/operating-systems/oracle-linux/podman/.

## Setup the Lab Environment

> **Note:** When using the free lab environment, see [Oracle Linux Lab Basics](https://luna.oracle.com/lab/c84a78db-4e92-4a58-86d1-a332bf47f99a/steps) for connection and other usage instructions.

1. Open a terminal and connect via ssh to the *ol-server* instance if not already connected.

   ```
   Copy
   ssh oracle@<ip_address_of_instance>
   ```

### Confirm Podman Works

The container-tools package in Oracle Linux provides the latest versions of Podman, Buildah, Skopeo, and associated dependencies.

1. Check the version of Podman.

   ```
   Copy
   podman -v
   ```

2. Confirm the Podman CLI is working.

   ```
   Copy
   podman run quay.io/podman/hello
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ podman run quay.io/podman/hello
   > Trying to pull quay.io/podman/hello:latest...
   > Getting image source signatures
   > Copying blob f82b04e85914 done  
   > Copying config dbd85e09a1 done  
   > Writing manifest to image destination
   > Storing signatures
   > !... Hello Podman World ...!
   > 
   >          .--"--.           
   >        / -     - \         
   >       / (O)   (O) \        
   >    ~~~| -=(,Y,)=- |         
   >     .---. /`  \   |~~      
   >  ~/  o  o \~~~~.----. ~~   
   >   | =(X)= |~  / (O (O) \   
   >    ~~~~~~~  ~| =(Y_)=-  |   
   >   ~~~~    ~~~|   U      |~~ 
   > 
   > Project:   https://github.com/containers/podman
   > Website:   https://podman.io
   > Documents: https://docs.podman.io
   > Twitter:   @Podman_io
   > ```

## Setup Podman to Work with Compose Files

Podman introduced support for Docker Compose functionality in [Podman v3.2.0](https://github.com/containers/podman/releases/tag/v3.2.0), after limited support was introduced in [Podman v3.0.0](https://github.com/containers/podman/releases/tag/v3.0.0), thereby introducing the ability to use [Docker Compose](https://github.com/docker/compose) from within Podman. More recently [Podman v4.1.0](https://github.com/containers/podman/releases/tag/v4.1.0) extended the support of Docker Compose functionality to include using Docker Compose v2.2 and higher.

The following steps describe how to install, and configure, both of `podman-compose` and `docker-compose`.

### Install the Podman Docker Package

This enables Podman to work natively with Docker commands.

1. Install the `podman-docker` package.

   ```
   Copy
   sudo dnf install -y podman-docker
   ```

### Install Docker Compose

> **Note:**
>
> [Installing Compose in a standalone manner](https://docs.docker.com/compose/install/other/) as described here requires using `docker-compose` instead of the standard syntax used when using the *Docker* utility of `docker compose`. In other words substitute the `docker compose up` syntax with `docker-compose up` instead.

1. Download and install Compose standalone.

   ```
   Copy
   sudo curl -SL https://github.com/docker/compose/releases/download/v2.15.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ sudo curl -SL https://github.com/docker/compose/releases/download/v2.15.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
   >   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
   >                                  Dload  Upload   Total   Spent    Left  Speed
   >   0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
   > 100 42.8M  100 42.8M    0     0   158M      0 --:--:-- --:--:-- --:--:--  158M
   > ```

2. Apply executable permissions to the binary.

   ```
   Copy
   sudo chmod +x /usr/local/bin/docker-compose
   ```

3. Confirm Compose standalone works.

   ```
   Copy
   docker-compose version
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ docker-compose version
   > Docker Compose version v2.15.1
   > ```

### Start the Podman Socket

The following steps are required to make Docker Compose functionality work by creating a Unix socket.

1. Configure the Podman socket with these steps.

   ```
   Copysudo systemctl enable --now podman.socket
   sudo systemctl status podman.socket
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ sudo systemctl enable --now podman.socket
   > Created symlink /etc/systemd/system/sockets.target.wants/podman.socket -> /usr/lib/systemd/system/podman.socket.
   > [oracle@ol-server ~]$ sudo systemctl status podman.socket
   > * podman.socket - Podman API Socket
   >    Loaded: loaded (/usr/lib/systemd/system/podman.socket; enabled; vendor preset)
   >    Active: active (listening) since Thu 2023-01-19 20:58:20 GMT; 16s ago
   >      Docs: man:podman-system-service(1)
   >    Listen: /run/podman/podman.sock (Stream)
   >    CGroup: /system.slice/podman.socket
   > 
   > Jan 19 20:58:20 ol-server systemd[1]: Listening on Podman API Socket.
   > ```

2. Verify that the socket works as expected using curl.

```
Copy
sudo curl -w "\n" -H "Content-Type: application/json" --unix-socket /var/run/docker.sock http://localhost/_ping
```

> Example Output:
>
> ```
> Copy[oracle@ol-server ~]$ sudo curl -w "\n" -H "Content-Type: application/json" --unix-socket /var/run/docker.sock http://localhost/_ping
> OK
> [oracle@ol-server ~]$
> ```

If the output of this command returns `OK`, then Compose functionality is successfully configured to work with docker-compose.yaml files.

### Install Podman Compose

[Podman Compose](https://github.com/containers/podman-compose) is a Python 3 library that implements the [Compose Specification](https://compose-spec.io/) to run with Podman.

1. Before installing Podman Compose, confirm that pip is updated to the latest version.

   ```
   Copy
   sudo -H pip3 install --upgrade pip
   ```

2. Install the Podman Compose package.

   ```
   Copy
   sudo pip3 install podman-compose 
   ```

   > When running the `pip3 install` as `sudo`, the *WARNING* messages can be ignored.

3. Confirm Podman Compose works.

   ```
   Copy
   podman-compose version
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ podman-compose version
   > ['podman', '--version', '']
   > using podman version: 4.2.0
   > podman-composer version  1.0.3
   > podman --version 
   > podman version 4.2.0
   > exit code: 0
   > ```

## Create a Docker Compose File

This *docker-compose.yaml* file allows the pulling and starting of the designated application.

1. Create a directory for the test.

   ```
   Copy
   mkdir -p Documents/examples/echo
   ```

2. Change into the directory.

   ```
   Copy
   cd Documents/examples/echo
   ```

3. Create the docker-compose.yaml file.

   ```
   Copycat >> docker-compose.yaml << EOF
   ---
   version: '3' 
   services: 
     web: 
       image: k8s.gcr.io/echoserver:1.4
       ports:
           - "${HOST_PORT:-8080}:8080" 
   EOF
   ```

4. Review the docker-compose.yaml just created.

   ```
   Copy
   cat docker-compose.yaml
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ cat docker-compose.yaml 
   > ---
   > version: '3' 
   > services: 
   >   web: 
   >     image: k8s.gcr.io/echoserver:1.4
   >     ports:
   >         - "8080:8080"
   > ```

## Confirm Podman Compose is Working

1. Change into the directory where the docker-compose file is located.

   > **Important:** `podman-compose` commands will **not** work unless you are in the directory where the `docker-compose.yaml` file is located.

   ```
   Copy
   cd ~/Documents/examples/echo/
   ```

2. Start the `echoserver` application.

   ```
   Copy
   podman-compose up -d
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ podman-compose up -d
   > ['podman', '--version', '']
   > using podman version: 4.2.0
   > ** excluding:  set()
   > ['podman', 'network', 'exists', 'echo_default']
   > ['podman', 'network', 'create', '--label', 'io.podman.compose.project=echo', '--label', 'com.docker.compose.project=echo', 'echo_default']
   > ['podman', 'network', 'exists', 'echo_default']
   > podman create --name=echo_web_1 --label io.podman.compose.config-hash=123 --label io.podman.compose.project=echo --label io.podman.compose.version=0.0.1 --label com.docker.compose.project=echo --label com.docker.compose.project.working_dir=/home/oracle/examples/echo --label com.docker.compose.project.config_files=docker-compose.yaml --label com.docker.compose.container-number=1 --label com.docker.compose.service=web --net echo_default --network-alias web -p 8080:8080 k8s.gcr.io/echoserver:1.4
   > Trying to pull k8s.gcr.io/echoserver:1.4...
   > Getting image source signatures
   > Copying blob d3c51dabc842 done  
   > Copying blob a3ed95caeb02 done  
   > Copying blob 6d9e6e7d968b done  
   > Copying blob 412c0feed608 done  
   > Copying blob cd23f57692f8 done  
   > Copying blob dcd34d50d5ee done  
   > Copying blob a3ed95caeb02 skipped: already exists  
   > Copying blob a3ed95caeb02 skipped: already exists  
   > Copying blob a3ed95caeb02 skipped: already exists  
   > Copying blob b4241160ce0e done  
   > Copying blob 7abee76f69c0 done  
   > Writing manifest to image destination
   > Storing signatures
   > 1b54b75ca13786d33df6708da1d83ecce14b055e78b03007c3c4e1f441e7139c
   > exit code: 0
   > ```

   > **Note:** As with Podman, any container(s) referenced in the `docker-compose.yaml` file are only pulled if not already present on the system.

3. Test the `echoserver` application is up and running.

   ```
   Copy
   curl -X POST -d "foobar" http://localhost:8080/; echo
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ curl -X POST -d "foobar" http://localhost:8080/; echo
   > CLIENT VALUES:
   > client_address=10.89.0.2
   > command=POST
   > real path=/
   > query=nil
   > request_version=1.1
   > request_uri=http://localhost:8080/
   > 
   > SERVER VALUES:
   > server_version=nginx: 1.10.0 - lua: 10001
   > 
   > HEADERS RECEIVED:
   > accept=*/*
   > content-length=6
   > content-type=application/x-www-form-urlencoded
   > host=localhost:8080
   > user-agent=curl/7.61.1
   > BODY:
   > foobar
   > ```

4. Additionally confirm success by reviewing the logs.

   ```
   Copy
   podman-compose logs
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ podman-compose logs
   > ['podman', '--version', '']
   > using podman version: 4.2.0
   > podman logs echo_web_1
   > 10.89.0.2 - - [17/Jan/2023:12:46:47 +0000] "POST / HTTP/1.1" 200 446 "-" "curl/7.61.1"
   > exit code: 0
   > ```

5. Use the Podman Compose utility to see running containers.

   ```
   Copy
   podman-compose ps
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ podman-compose ps   
   > ['podman', '--version', '']
   > using podman version: 4.2.0
   > podman ps -a --filter label=io.podman.compose.project=echo
   > CONTAINER ID  IMAGE                      COMMAND               CREATED        STATUS            PORTS                   NAMES
   > f4053947c8c1  k8s.gcr.io/echoserver:1.4  nginx -g daemon o...  2 minutes ago  Up 2 minutes ago  0.0.0.0:8080->8080/tcp  echo_web_1
   > exit code: 0
   > ```

   > **Note:** See other Podman Compose utility commands by running `podman-compose --help`. However, these additional commands are out of scope for this Lab.

6. Now it is time to stop the *echoweb* service.

   ```
   Copy
   podman-compose down
   ```

## Confirm Docker Compose is Working

1. (Optional) Change into the directory where the docker-compose file is located.

   > **Important:** `docker-compose` commands will **not** work unless you are in the directory where the `docker-compose.yaml` file is located.

   ```
   Copy
   cd ~/Documents/examples/echo/
   ```

2. Start the `echoserver` application.

   ```
   Copy
   sudo /usr/local/bin/docker-compose up -d
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ sudo /usr/local/bin/docker-compose up -d
   > [+] Running 0/0
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.0
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.1
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.2
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.3
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.4
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.5
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.6
   > [+] Running 0/1echo-web-1  Starting                                                                                               0.7
   > [+] Running 1/1echo-web-1  Starting                                                                                               0.8
   >  ��� Container echo-web-1  Started                                                                                                0.8s
   > ```

   > **NOTE:**
   >
   > If the following output is returned when executing this command:
   >
   > ```
   > Copy��� Container echo-web-1  Starting                                                                                               0.9s
   > Error response from daemon: cannot listen on the TCP port: listen tcp4 :8080: bind: address already in use
   > ```
   >
   > Don’t worry, it only means that the `podman-compose down` command wasn’t executed, and the *echoserver* container previously started using `podman-compose` is still running. Follow the previous steps to stop it.

3. Test the container is running.

   ```
   Copy
   curl -X POST -d "foobar" http://localhost:8080/; echo
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server ~]$ curl -X POST -d "foobar" http://localhost:8080/; echo
   > CLIENT VALUES:
   > client_address=10.89.0.2
   > command=POST
   > real path=/
   > query=nil
   > request_version=1.1
   > request_uri=http://localhost:8080/
   > 
   > SERVER VALUES:
   > server_version=nginx: 1.10.0 - lua: 10001
   > 
   > HEADERS RECEIVED:
   > accept=*/*
   > content-length=6
   > content-type=application/x-www-form-urlencoded
   > host=localhost:8080
   > user-agent=curl/7.61.1
   > BODY:
   > foobar
   > ```

4. As before, use the Docker Compose utility to inspect the logs and confirm a successful request was returned by this application.

   ```
   Copy
   sudo /usr/local/bin/docker-compose logs
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ sudo /usr/local/bin/docker-compose logs 
   > echo-web-1  | 10.89.0.1 - - [17/Jan/2023:14:48:56 +0000] "POST / HTTP/1.1" 200 446 "-" "curl/7.61.1"
   > ```

5. The Docker Compose utility also provides a way to review the containers started by the Compose file.

   ```
   Copy
   sudo /usr/local/bin/docker-compose ps
   ```

   > Example Output:
   >
   > ```
   > Copy
   > echo-web-1        k8s.gcr.io/echoserver:1.4   "nginx -g daemon off;"   web             12 minutes ago      Up 12 minutes       8080/tcp
   > ```

   > **Note:** See other Docker Compose utility commands by running `docker-compose --help`. However, these additional commands are out of scope for this Lab.

6. Stop the *echoweb* service.

   ```
   Copy
   sudo /usr/local/bin/docker-compose down
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ sudo /usr/local/bin/docker-compose down
   > [+] Running 0/0
   > [+] Running 0/1echo-web-1  Stopping                                                                                                   0.1
   > [+] Running 0/1echo-web-1  Stopping                                                                                                   0.2
   > [+] Running 0/1echo-web-1  Stopping                                                                                                   0.3
   > [+] Running 2/1echo-web-1  Removing                                                                                                   0.4
   >  ��� Container echo-web-1  Removed                                                                                                    0.4s
   >  ��� Network echo_default  Removed                                                                                                    0.0s
   > ```

## Important Information

If both the `podman-compose` **and** `docker-compose` executables have been installed on the same system, it is important to note that it is **not possible** to invoke them interchangeably. This means that if a process is started by `podman-docker` it cannot be queried, or stopped by using the `docker-compose` utility, or vice-versa, as shown in the example below:

1. Change into the directory where the Compose file is located.

   > **Important:** The `podman-compose` or `docker-compose` commands will **not** work unless you are in the directory where the `docker-compose.yaml` file is located.

   ```
   Copy
   cd ~/Documents/examples/echo/
   ```

2. Start by using Podman Compose.

   ```
   Copy
   podman-compose up -d
   ```

3. Try to list the running containers using Docker Compose.

   ```
   Copy
   sudo /usr/local/bin/docker-compose ps
   ```

   > Example Output
   >
   > ```
   > Copy[oracle@ol-server echo]$ sudo /usr/local/bin/docker-compose ps
   > NAME                IMAGE               COMMAND             SERVICE             CREATED             STATUS              PORTS
   > ```

4. However using Podman Compose confirms the container is running.

   ```
   Copy
   podman-compose ps
   ```

   > Example Output:
   >
   > ```
   > Copy[oracle@ol-server echo]$ podman-compose ps
   > ['podman', '--version', '']
   > using podman version: 4.2.0
   > podman ps -a --filter label=io.podman.compose.project=echo
   > CONTAINER ID  IMAGE                      COMMAND               CREATED        STATUS            PORTS                   NAMES
   > 55335e797296  k8s.gcr.io/echoserver:1.4  nginx -g daemon o...  4 minutes ago  Up 4 minutes ago  0.0.0.0:8080->8080/tcp  echo_web_1
   > exit code: 0
   > ```