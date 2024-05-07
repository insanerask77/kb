## Install Elasticsearch with Docker[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Elasticsearch is also available as Docker images. A list of all published Docker images and tags is available at [www.docker.elastic.co](https://www.docker.elastic.co/). The source files are in [Github](https://github.com/elastic/elasticsearch/blob/8.6/distribution/docker).

This package contains both free and subscription features. [Start a 30-day trial](https://www.elastic.co/guide/en/elasticsearch/reference/current/license-settings.html) to try out all of the features.

Starting in Elasticsearch 8.0, security is enabled by default. With security enabled, Elastic Stack security features require TLS encryption for the transport networking layer, or your cluster will fail to start.

### Install Docker Desktop or Docker Engine[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Install the appropriate [Docker application](https://docs.docker.com/get-docker/) for your operating system.

Make sure that Docker is allotted at least 4GiB of memory. In Docker Desktop, you configure resource usage on the Advanced tab in Preference (macOS) or Settings (Windows).

### Pull the Elasticsearch Docker image[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Obtaining Elasticsearch for Docker is as simple as issuing a `docker pull` command against the Elastic Docker registry.

```sh
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.6.1
```

Now that you have the Elasticsearch Docker image, you can start a [single-node](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-cli-run-dev-mode) or [multi-node](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-compose-file) cluster.

### Start a single-node cluster with Docker[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

If you’re starting a single-node Elasticsearch cluster in a Docker container, security will be automatically enabled and configured for you. When you start Elasticsearch for the first time, the following security configuration occurs automatically:

- [Certificates and keys](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#elasticsearch-security-certificates) are generated for the transport and HTTP layers.
- The Transport Layer Security (TLS) configuration settings are written to `elasticsearch.yml`.
- A password is generated for the `elastic` user.
- An enrollment token is generated for Kibana.

You can then [start Kibana](https://www.elastic.co/guide/en/kibana/8.6/docker.html) and enter the enrollment token, which is valid for 30 minutes. This token automatically applies the security settings from your Elasticsearch cluster, authenticates to Elasticsearch with the `kibana_system` user, and writes the security configuration to `kibana.yml`.

The following commands start a single-node Elasticsearch cluster for development or testing.

1. Create a new docker network for Elasticsearch and Kibana

   ```sh
   docker network create elastic
   ```

2. Start Elasticsearch in Docker. A password is generated for the `elastic` user and output to the terminal, plus an enrollment token for enrolling Kibana.

   ```sh
   docker run --name es01 --net elastic -p 9200:9200 -it docker.elastic.co/elasticsearch/elasticsearch:8.6.1
   ```

   You might need to scroll back a bit in the terminal to view the password and enrollment token.

3. Copy the generated password and enrollment token and save them in a secure location. These values are shown only when you start Elasticsearch for the first time.

   If you need to reset the password for the `elastic` user or other built-in users, run the [`elasticsearch-reset-password`](https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-password.html) tool. This tool is available in the Elasticsearch `/bin` directory of the Docker container. For example:

   ```sh
   docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-reset-password
   ```

4. Copy the `http_ca.crt` security certificate from your Docker container to your local machine.

   ```sh
   docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt .
   ```

5. Open a new terminal and verify that you can connect to your Elasticsearch cluster by making an authenticated call, using the `http_ca.crt` file that you copied from your Docker container. Enter the password for the `elastic` user when prompted.

   ```sh
   curl --cacert http_ca.crt -u elastic https://localhost:9200
   ```

### Enroll additional nodes[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

When you start Elasticsearch for the first time, the installation process configures a single-node cluster by default. This process also generates an enrollment token and prints it to your terminal. If you want a node to join an existing cluster, start the new node with the generated enrollment token.

**Generating enrollment tokens**

The enrollment token is valid for 30 minutes. If you need to generate a new enrollment token, run the [`elasticsearch-create-enrollment-token`](https://www.elastic.co/guide/en/elasticsearch/reference/current/create-enrollment-token.html) tool on your existing node. This tool is available in the Elasticsearch `bin` directory of the Docker container.

For example, run the following command on the existing `es01` node to generate an enrollment token for new Elasticsearch nodes:

```sh
docker exec -it es01 /usr/share/elasticsearch/bin/elasticsearch-create-enrollment-token -s node
```

1. In the terminal where you started your first node, copy the generated enrollment token for adding new Elasticsearch nodes.

2. On your new node, start Elasticsearch and include the generated enrollment token.

   ```sh
   docker run -e ENROLLMENT_TOKEN="<token>" --name es02 --net elastic -it docker.elastic.co/elasticsearch/elasticsearch:8.6.1
   ```

   Elasticsearch is now configured to join the existing cluster.

#### Setting JVM heap size[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

If you experience issues where the container where your first node is running exits when your second node starts, explicitly set values for the JVM heap size. To [manually configure the heap size](https://www.elastic.co/guide/en/elasticsearch/reference/current/advanced-configuration.html#set-jvm-heap-size), include the `ES_JAVA_OPTS` variable and set values for `-Xms` and `-Xmx` when starting each node. For example, the following command starts node `es02` and sets the minimum and maximum JVM heap size to 1 GB:

```sh
docker run -e ES_JAVA_OPTS="-Xms1g -Xmx1g" -e ENROLLMENT_TOKEN="<token>" --name es02 -p 9201:9200 --net elastic -it docker.elastic.co/elasticsearch/elasticsearch:8.6.1
```

#### Next steps[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

You now have a test Elasticsearch environment set up. Before you start serious development or go into production with Elasticsearch, review the [requirements and recommendations](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-prod-prerequisites) to apply when running Elasticsearch in Docker in production.

#### Security certificates and keys[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/security-files-reference.asciidoc)

When you install Elasticsearch, the following certificates and keys are generated in the Elasticsearch configuration directory, which are used to connect a Kibana instance to your secured Elasticsearch cluster and to encrypt internode communication. The files are listed here for reference.

- **`http_ca.crt`**

  The CA certificate that is used to sign the certificates for the HTTP layer of this Elasticsearch cluster.

- **`http.p12`**

  Keystore that contains the key and certificate for the HTTP layer for this node.

- **`transport.p12`**

  Keystore that contains the key and certificate for the transport layer for all the nodes in your cluster.

`http.p12` and `transport.p12` are password-protected PKCS#12 keystores. Elasticsearch stores the passwords for these keystores as [secure settings](https://www.elastic.co/guide/en/elasticsearch/reference/current/secure-settings.html). To retrieve the passwords so that you can inspect or change the keystore contents, use the [`bin/elasticsearch-keystore`](https://www.elastic.co/guide/en/elasticsearch/reference/current/elasticsearch-keystore.html) tool.

Use the following command to retrieve the password for `http.p12`:

```sh
bin/elasticsearch-keystore show xpack.security.http.ssl.keystore.secure_password
```

Use the following command to retrieve the password for `transport.p12`:

```sh
bin/elasticsearch-keystore show xpack.security.transport.ssl.keystore.secure_password
```

### Start a multi-node cluster with Docker Compose[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

To get a multi-node Elasticsearch cluster and Kibana up and running in Docker with security enabled, you can use Docker Compose.

This configuration provides a simple method of starting a secured cluster that you can use for development before building a distributed deployment with multiple hosts.

#### Prerequisites[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Install the appropriate [Docker application](https://docs.docker.com/get-docker/) for your operating system.

If you’re running on Linux, install [Docker Compose](https://docs.docker.com/compose/install/).

Make sure that Docker is allotted at least 4GB of memory. In Docker Desktop, you configure resource usage on the Advanced tab in Preferences (macOS) or Settings (Windows).

#### Prepare the environment[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Create the following configuration files in a new, empty directory. These files are also available from the [elasticsearch](https://github.com/elastic/elasticsearch/tree/master/docs/reference/setup/install/docker) repository on GitHub.

##### `.env`[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The `.env` file sets environment variables that are used when you run the `docker-compose.yml` configuration file. Ensure that you specify a strong password for the `elastic` and `kibana_system` users with the `ELASTIC_PASSWORD` and `KIBANA_PASSWORD` variables. These variable are referenced by the `docker-compose.yml` file.

Your passwords must be alphanumeric, and cannot contain special characters such as `!` or `@`. The `bash` script included in the `docker-compose.yml` file only operates on alphanumeric characters.

```txt
# Password for the 'elastic' user (at least 6 characters)
ELASTIC_PASSWORD=

# Password for the 'kibana_system' user (at least 6 characters)
KIBANA_PASSWORD=

# Version of Elastic products
STACK_VERSION=8.6.1

# Set the cluster name
CLUSTER_NAME=docker-cluster

# Set to 'basic' or 'trial' to automatically start the 30-day trial
LICENSE=basic
#LICENSE=trial

# Port to expose Elasticsearch HTTP API to the host
ES_PORT=9200
#ES_PORT=127.0.0.1:9200

# Port to expose Kibana to the host
KIBANA_PORT=5601
#KIBANA_PORT=80

# Increase or decrease based on the available host memory (in bytes)
MEM_LIMIT=1073741824

# Project namespace (defaults to the current folder name if not set)
#COMPOSE_PROJECT_NAME=myproject
```

##### `docker-compose.yml`[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

This `docker-compose.yml` file creates a three-node secure Elasticsearch cluster with authentication and network encryption enabled, and a Kibana instance securely connected to it.

**Exposing ports**

This configuration exposes port `9200` on all network interfaces. Because of how Docker handles ports, a port that isn’t bound to `localhost` leaves your Elasticsearch cluster publicly accessible, potentially ignoring any firewall settings. If you don’t want to expose port `9200` to external hosts, set the value for `ES_PORT` in the `.env` file to something like `127.0.0.1:9200`. Elasticsearch will then only be accessible from the host machine itself.

```yaml
version: "2.2"

services:
  setup:
    image: docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}
    volumes:
      - certs:/usr/share/elasticsearch/config/certs
    user: "0"
    command: >
      bash -c '
        if [ x${ELASTIC_PASSWORD} == x ]; then
          echo "Set the ELASTIC_PASSWORD environment variable in the .env file";
          exit 1;
        elif [ x${KIBANA_PASSWORD} == x ]; then
          echo "Set the KIBANA_PASSWORD environment variable in the .env file";
          exit 1;
        fi;
        if [ ! -f config/certs/ca.zip ]; then
          echo "Creating CA";
          bin/elasticsearch-certutil ca --silent --pem -out config/certs/ca.zip;
          unzip config/certs/ca.zip -d config/certs;
        fi;
        if [ ! -f config/certs/certs.zip ]; then
          echo "Creating certs";
          echo -ne \
          "instances:\n"\
          "  - name: es01\n"\
          "    dns:\n"\
          "      - es01\n"\
          "      - localhost\n"\
          "    ip:\n"\
          "      - 127.0.0.1\n"\
          "  - name: es02\n"\
          "    dns:\n"\
          "      - es02\n"\
          "      - localhost\n"\
          "    ip:\n"\
          "      - 127.0.0.1\n"\
          "  - name: es03\n"\
          "    dns:\n"\
          "      - es03\n"\
          "      - localhost\n"\
          "    ip:\n"\
          "      - 127.0.0.1\n"\
          > config/certs/instances.yml;
          bin/elasticsearch-certutil cert --silent --pem -out config/certs/certs.zip --in config/certs/instances.yml --ca-cert config/certs/ca/ca.crt --ca-key config/certs/ca/ca.key;
          unzip config/certs/certs.zip -d config/certs;
        fi;
        echo "Setting file permissions"
        chown -R root:root config/certs;
        find . -type d -exec chmod 750 \{\} \;;
        find . -type f -exec chmod 640 \{\} \;;
        echo "Waiting for Elasticsearch availability";
        until curl -s --cacert config/certs/ca/ca.crt https://es01:9200 | grep -q "missing authentication credentials"; do sleep 30; done;
        echo "Setting kibana_system password";
        until curl -s -X POST --cacert config/certs/ca/ca.crt -u "elastic:${ELASTIC_PASSWORD}" -H "Content-Type: application/json" https://es01:9200/_security/user/kibana_system/_password -d "{\"password\":\"${KIBANA_PASSWORD}\"}" | grep -q "^{}"; do sleep 10; done;
        echo "All done!";
      '
    healthcheck:
      test: ["CMD-SHELL", "[ -f config/certs/es01/es01.crt ]"]
      interval: 1s
      timeout: 5s
      retries: 120

  es01:
    depends_on:
      setup:
        condition: service_healthy
    image: docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}
    volumes:
      - certs:/usr/share/elasticsearch/config/certs
      - esdata01:/usr/share/elasticsearch/data
    ports:
      - ${ES_PORT}:9200
    environment:
      - node.name=es01
      - cluster.name=${CLUSTER_NAME}
      - cluster.initial_master_nodes=es01,es02,es03
      - discovery.seed_hosts=es02,es03
      - ELASTIC_PASSWORD=${ELASTIC_PASSWORD}
      - bootstrap.memory_lock=true
      - xpack.security.enabled=true
      - xpack.security.http.ssl.enabled=true
      - xpack.security.http.ssl.key=certs/es01/es01.key
      - xpack.security.http.ssl.certificate=certs/es01/es01.crt
      - xpack.security.http.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.enabled=true
      - xpack.security.transport.ssl.key=certs/es01/es01.key
      - xpack.security.transport.ssl.certificate=certs/es01/es01.crt
      - xpack.security.transport.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.verification_mode=certificate
      - xpack.license.self_generated.type=${LICENSE}
    mem_limit: ${MEM_LIMIT}
    ulimits:
      memlock:
        soft: -1
        hard: -1
    healthcheck:
      test:
        [
          "CMD-SHELL",
          "curl -s --cacert config/certs/ca/ca.crt https://localhost:9200 | grep -q 'missing authentication credentials'",
        ]
      interval: 10s
      timeout: 10s
      retries: 120

  es02:
    depends_on:
      - es01
    image: docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}
    volumes:
      - certs:/usr/share/elasticsearch/config/certs
      - esdata02:/usr/share/elasticsearch/data
    environment:
      - node.name=es02
      - cluster.name=${CLUSTER_NAME}
      - cluster.initial_master_nodes=es01,es02,es03
      - discovery.seed_hosts=es01,es03
      - bootstrap.memory_lock=true
      - xpack.security.enabled=true
      - xpack.security.http.ssl.enabled=true
      - xpack.security.http.ssl.key=certs/es02/es02.key
      - xpack.security.http.ssl.certificate=certs/es02/es02.crt
      - xpack.security.http.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.enabled=true
      - xpack.security.transport.ssl.key=certs/es02/es02.key
      - xpack.security.transport.ssl.certificate=certs/es02/es02.crt
      - xpack.security.transport.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.verification_mode=certificate
      - xpack.license.self_generated.type=${LICENSE}
    mem_limit: ${MEM_LIMIT}
    ulimits:
      memlock:
        soft: -1
        hard: -1
    healthcheck:
      test:
        [
          "CMD-SHELL",
          "curl -s --cacert config/certs/ca/ca.crt https://localhost:9200 | grep -q 'missing authentication credentials'",
        ]
      interval: 10s
      timeout: 10s
      retries: 120

  es03:
    depends_on:
      - es02
    image: docker.elastic.co/elasticsearch/elasticsearch:${STACK_VERSION}
    volumes:
      - certs:/usr/share/elasticsearch/config/certs
      - esdata03:/usr/share/elasticsearch/data
    environment:
      - node.name=es03
      - cluster.name=${CLUSTER_NAME}
      - cluster.initial_master_nodes=es01,es02,es03
      - discovery.seed_hosts=es01,es02
      - bootstrap.memory_lock=true
      - xpack.security.enabled=true
      - xpack.security.http.ssl.enabled=true
      - xpack.security.http.ssl.key=certs/es03/es03.key
      - xpack.security.http.ssl.certificate=certs/es03/es03.crt
      - xpack.security.http.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.enabled=true
      - xpack.security.transport.ssl.key=certs/es03/es03.key
      - xpack.security.transport.ssl.certificate=certs/es03/es03.crt
      - xpack.security.transport.ssl.certificate_authorities=certs/ca/ca.crt
      - xpack.security.transport.ssl.verification_mode=certificate
      - xpack.license.self_generated.type=${LICENSE}
    mem_limit: ${MEM_LIMIT}
    ulimits:
      memlock:
        soft: -1
        hard: -1
    healthcheck:
      test:
        [
          "CMD-SHELL",
          "curl -s --cacert config/certs/ca/ca.crt https://localhost:9200 | grep -q 'missing authentication credentials'",
        ]
      interval: 10s
      timeout: 10s
      retries: 120

  kibana:
    depends_on:
      es01:
        condition: service_healthy
      es02:
        condition: service_healthy
      es03:
        condition: service_healthy
    image: docker.elastic.co/kibana/kibana:${STACK_VERSION}
    volumes:
      - certs:/usr/share/kibana/config/certs
      - kibanadata:/usr/share/kibana/data
    ports:
      - ${KIBANA_PORT}:5601
    environment:
      - SERVERNAME=kibana
      - ELASTICSEARCH_HOSTS=https://es01:9200
      - ELASTICSEARCH_USERNAME=kibana_system
      - ELASTICSEARCH_PASSWORD=${KIBANA_PASSWORD}
      - ELASTICSEARCH_SSL_CERTIFICATEAUTHORITIES=config/certs/ca/ca.crt
    mem_limit: ${MEM_LIMIT}
    healthcheck:
      test:
        [
          "CMD-SHELL",
          "curl -s -I http://localhost:5601 | grep -q 'HTTP/1.1 302 Found'",
        ]
      interval: 10s
      timeout: 10s
      retries: 120

volumes:
  certs:
    driver: local
  esdata01:
    driver: local
  esdata02:
    driver: local
  esdata03:
    driver: local
  kibanadata:
    driver: local
```

#### Start your cluster with security enabled and configured[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

1. Modify the `.env` file and enter strong password values for both the `ELASTIC_PASSWORD` and `KIBANA_PASSWORD` variables.

   You must use the `ELASTIC_PASSWORD` value for further interactions with the cluster. The `KIBANA_PASSWORD` value is only used internally when configuring Kibana.

2. Create and start the three-node Elasticsearch cluster and Kibana instance:

   ```sh
   docker-compose up -d
   ```

3. When the deployment has started, open a browser and navigate to [http://localhost:5601](http://localhost:5601/) to access Kibana, where you can load sample data and interact with your cluster.

#### Stop and remove the deployment[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

To stop the cluster, run `docker-compose down`. The data in the Docker volumes is preserved and loaded when you restart the cluster with `docker-compose up`.

```sh
docker-compose down
```

To **delete** the network, containers, and volumes when you stop the cluster, specify the `-v` option:

```sh
docker-compose down -v
```

#### Next steps[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

You now have a test Elasticsearch environment set up. Before you start serious development or go into production with Elasticsearch, review the [requirements and recommendations](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-prod-prerequisites) to apply when running Elasticsearch in Docker in production.

### Using the Docker images in production[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The following requirements and recommendations apply when running Elasticsearch in Docker in production.

#### Set `vm.max_map_count` to at least `262144`[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The `vm.max_map_count` kernel setting must be set to at least `262144` for production use.

How you set `vm.max_map_count` depends on your platform.

##### Linux[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

To view the current value for the `vm.max_map_count` setting, run:

```sh
grep vm.max_map_count /etc/sysctl.conf
vm.max_map_count=262144
```

To apply the setting on a live system, run:

```sh
sysctl -w vm.max_map_count=262144
```

To permanently change the value for the `vm.max_map_count` setting, update the value in `/etc/sysctl.conf`.

##### macOS with [Docker for Mac](https://docs.docker.com/docker-for-mac)[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The `vm.max_map_count` setting must be set within the xhyve virtual machine:

1. From the command line, run:

   ```sh
   screen ~/Library/Containers/com.docker.docker/Data/vms/0/tty
   ```

2. Press enter and use `sysctl` to configure `vm.max_map_count`:

   ```sh
   sysctl -w vm.max_map_count=262144
   ```

3. To exit the `screen` session, type `Ctrl a d`.

##### Windows and macOS with [Docker Desktop](https://www.docker.com/products/docker-desktop)[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The `vm.max_map_count` setting must be set via docker-machine:

```sh
docker-machine ssh
sudo sysctl -w vm.max_map_count=262144
```

##### Windows with [Docker Desktop WSL 2 backend](https://docs.docker.com/docker-for-windows/wsl)[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The `vm.max_map_count` setting must be set in the docker-desktop container:

```sh
wsl -d docker-desktop
sysctl -w vm.max_map_count=262144
```

#### Configuration files must be readable by the `elasticsearch` user[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

By default, Elasticsearch runs inside the container as user `elasticsearch` using uid:gid `1000:0`.

One exception is [Openshift](https://docs.openshift.com/container-platform/3.6/creating_images/guidelines.html#openshift-specific-guidelines), which runs containers using an arbitrarily assigned user ID. Openshift presents persistent volumes with the gid set to `0`, which works without any adjustments.

If you are bind-mounting a local directory or file, it must be readable by the `elasticsearch` user. In addition, this user must have write access to the [config, data and log dirs](https://www.elastic.co/guide/en/elasticsearch/reference/current/important-settings.html#path-settings) (Elasticsearch needs write access to the `config` directory so that it can generate a keystore). A good strategy is to grant group access to gid `0` for the local directory.

For example, to prepare a local directory for storing data through a bind-mount:

```sh
mkdir esdatadir
chmod g+rwx esdatadir
chgrp 0 esdatadir
```

You can also run an Elasticsearch container using both a custom UID and GID. You must ensure that file permissions will not prevent Elasticsearch from executing. You can use one of two options:

- Bind-mount the `config`, `data` and `logs` directories. If you intend to install plugins and prefer not to [create a custom Docker image](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#_c_customized_image), you must also bind-mount the `plugins` directory.
- Pass the `--group-add 0` command line option to `docker run`. This ensures that the user under which Elasticsearch is running is also a member of the `root` (GID 0) group inside the container.

#### Increase ulimits for nofile and nproc[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Increased ulimits for [nofile](https://www.elastic.co/guide/en/elasticsearch/reference/current/setting-system-settings.html) and [nproc](https://www.elastic.co/guide/en/elasticsearch/reference/current/max-number-threads-check.html) must be available for the Elasticsearch containers. Verify the [init system](https://github.com/moby/moby/tree/ea4d1243953e6b652082305a9c3cda8656edab26/contrib/init) for the Docker daemon sets them to acceptable values.

To check the Docker daemon defaults for ulimits, run:

```sh
docker run --rm docker.elastic.co/elasticsearch/elasticsearch:{version} /bin/bash -c 'ulimit -Hn && ulimit -Sn && ulimit -Hu && ulimit -Su'
```

If needed, adjust them in the Daemon or override them per container. For example, when using `docker run`, set:

```sh
--ulimit nofile=65535:65535
```

#### Disable swapping[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Swapping needs to be disabled for performance and node stability. For information about ways to do this, see [Disable swapping](https://www.elastic.co/guide/en/elasticsearch/reference/current/setup-configuration-memory.html).

If you opt for the `bootstrap.memory_lock: true` approach, you also need to define the `memlock: true` ulimit in the [Docker Daemon](https://docs.docker.com/engine/reference/commandline/dockerd/#default-ulimits), or explicitly set for the container as shown in the [sample compose file](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-compose-file). When using `docker run`, you can specify:

```sh
-e "bootstrap.memory_lock=true" --ulimit memlock=-1:-1
```

#### Randomize published ports[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

The image [exposes](https://docs.docker.com/engine/reference/builder/#/expose) TCP ports 9200 and 9300. For production clusters, randomizing the published ports with `--publish-all` is recommended, unless you are pinning one container per host.

#### Manually set the heap size[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

By default, Elasticsearch automatically sizes JVM heap based on a nodes’s [roles](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-node.html#node-roles) and the total memory available to the node’s container. We recommend this default sizing for most production environments. If needed, you can override default sizing by manually setting JVM heap size.

To manually set the heap size in production, bind mount a [JVM options](https://www.elastic.co/guide/en/elasticsearch/reference/current/advanced-configuration.html#set-jvm-options) file under `/usr/share/elasticsearch/config/jvm.options.d` that includes your desired [heap size](https://www.elastic.co/guide/en/elasticsearch/reference/current/advanced-configuration.html#set-jvm-heap-size) settings.

For testing, you can also manually set the heap size using the `ES_JAVA_OPTS` environment variable. For example, to use 16GB, specify `-e ES_JAVA_OPTS="-Xms16g -Xmx16g"` with `docker run`. The `ES_JAVA_OPTS` variable overrides all other JVM options. We do not recommend using `ES_JAVA_OPTS` in production. The `docker-compose.yml` file above sets the heap size to 512MB.

#### Pin deployments to a specific image version[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Pin your deployments to a specific version of the Elasticsearch Docker image. For example `docker.elastic.co/elasticsearch/elasticsearch:8.6.1`.

#### Always bind data volumes[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

You should use a volume bound on `/usr/share/elasticsearch/data` for the following reasons:

1. The data of your Elasticsearch node won’t be lost if the container is killed
2. Elasticsearch is I/O sensitive and the Docker storage driver is not ideal for fast I/O
3. It allows the use of advanced [Docker volume plugins](https://docs.docker.com/engine/extend/plugins/#volume-plugins)

#### Avoid using `loop-lvm` mode[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

If you are using the devicemapper storage driver, do not use the default `loop-lvm` mode. Configure docker-engine to use [direct-lvm](https://docs.docker.com/engine/userguide/storagedriver/device-mapper-driver/#configure-docker-with-devicemapper).

#### Centralize your logs[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Consider centralizing your logs by using a different [logging driver](https://docs.docker.com/engine/admin/logging/overview/). Also note that the default json-file logging driver is not ideally suited for production use.

### Configuring Elasticsearch with Docker[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

When you run in Docker, the [Elasticsearch configuration files](https://www.elastic.co/guide/en/elasticsearch/reference/current/settings.html#config-files-location) are loaded from `/usr/share/elasticsearch/config/`.

To use custom configuration files, you [bind-mount the files](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-config-bind-mount) over the configuration files in the image.

You can set individual Elasticsearch configuration parameters using Docker environment variables. The [sample compose file](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-compose-file) and the [single-node example](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-cli-run-dev-mode) use this method. You can use the setting name directly as the environment variable name. If you cannot do this, for example because your orchestration platform forbids periods in environment variable names, then you can use an alternative style by converting the setting name as follows.

1. Change the setting name to uppercase
2. Prefix it with `ES_SETTING_`
3. Escape any underscores (`_`) by duplicating them
4. Convert all periods (`.`) to underscores (`_`)

For example, `-e bootstrap.memory_lock=true` becomes `-e ES_SETTING_BOOTSTRAP_MEMORY__LOCK=true`.

You can use the contents of a file to set the value of the `ELASTIC_PASSWORD` or `KEYSTORE_PASSWORD` environment variables, by suffixing the environment variable name with `_FILE`. This is useful for passing secrets such as passwords to Elasticsearch without specifying them directly.

For example, to set the Elasticsearch bootstrap password from a file, you can bind mount the file and set the `ELASTIC_PASSWORD_FILE` environment variable to the mount location. If you mount the password file to `/run/secrets/bootstrapPassword.txt`, specify:

```sh
-e ELASTIC_PASSWORD_FILE=/run/secrets/bootstrapPassword.txt
```

You can override the default command for the image to pass Elasticsearch configuration parameters as command line options. For example:

```sh
docker run <various parameters> bin/elasticsearch -Ecluster.name=mynewclustername
```

While bind-mounting your configuration files is usually the preferred method in production, you can also [create a custom Docker image](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#_c_customized_image) that contains your configuration.

#### Mounting Elasticsearch configuration files[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Create custom config files and bind-mount them over the corresponding files in the Docker image. For example, to bind-mount `custom_elasticsearch.yml` with `docker run`, specify:

```sh
-v full_path_to/custom_elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml
```

If you bind-mount a custom `elasticsearch.yml` file, ensure it includes the `network.host: 0.0.0.0` setting. This setting ensures the node is reachable for HTTP and transport traffic, provided its ports are exposed. The Docker image’s built-in `elasticsearch.yml` file includes this setting by default.

The container **runs Elasticsearch as user `elasticsearch` using uid:gid `1000:0`**. Bind mounted host directories and files must be accessible by this user, and the data and log directories must be writable by this user.

#### Create an encrypted Elasticsearch keystore[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

By default, Elasticsearch will auto-generate a keystore file for [secure settings](https://www.elastic.co/guide/en/elasticsearch/reference/current/secure-settings.html). This file is obfuscated but not encrypted.

To encrypt your secure settings with a password and have them persist outside the container, use a `docker run` command to manually create the keystore instead. The command must:

- Bind-mount the `config` directory. The command will create an `elasticsearch.keystore` file in this directory. To avoid errors, do not directly bind-mount the `elasticsearch.keystore` file.
- Use the `elasticsearch-keystore` tool with the `create -p` option. You’ll be prompted to enter a password for the keystore.

For example:

```sh
docker run -it --rm \
-v full_path_to/config:/usr/share/elasticsearch/config \
docker.elastic.co/elasticsearch/elasticsearch:8.6.1 \
bin/elasticsearch-keystore create -p
```

You can also use a `docker run` command to add or update secure settings in the keystore. You’ll be prompted to enter the setting values. If the keystore is encrypted, you’ll also be prompted to enter the keystore password.

```sh
docker run -it --rm \
-v full_path_to/config:/usr/share/elasticsearch/config \
docker.elastic.co/elasticsearch/elasticsearch:8.6.1 \
bin/elasticsearch-keystore \
add my.secure.setting \
my.other.secure.setting
```

If you’ve already created the keystore and don’t need to update it, you can bind-mount the `elasticsearch.keystore` file directly. You can use the `KEYSTORE_PASSWORD` environment variable to provide the keystore password to the container at startup. For example, a `docker run` command might have the following options:

```sh
-v full_path_to/config/elasticsearch.keystore:/usr/share/elasticsearch/config/elasticsearch.keystore
-e KEYSTORE_PASSWORD=mypassword
```

#### Using custom Docker images[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

In some environments, it might make more sense to prepare a custom image that contains your configuration. A `Dockerfile` to achieve this might be as simple as:

```sh
FROM docker.elastic.co/elasticsearch/elasticsearch:8.6.1
COPY --chown=elasticsearch:elasticsearch elasticsearch.yml /usr/share/elasticsearch/config/
```

You could then build and run the image with:

```sh
docker build --tag=elasticsearch-custom .
docker run -ti -v /usr/share/elasticsearch/data elasticsearch-custom
```

Some plugins require additional security permissions. You must explicitly accept them either by:

- Attaching a `tty` when you run the Docker image and allowing the permissions when prompted.
- Inspecting the security permissions and accepting them (if appropriate) by adding the `--batch` flag to the plugin install command.

See [Plugin management](https://www.elastic.co/guide/en/elasticsearch/plugins/8.6/_other_command_line_parameters.html) for more information.

#### Troubleshoot Docker errors for Elasticsearch[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

Here’s how to resolve common errors when running Elasticsearch with Docker.

#### elasticsearch.keystore is a directory[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

```txt
Exception in thread "main" org.elasticsearch.bootstrap.BootstrapException: java.io.IOException: Is a directory: SimpleFSIndexInput(path="/usr/share/elasticsearch/config/elasticsearch.keystore") Likely root cause: java.io.IOException: Is a directory
```

A [keystore-related](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-keystore-bind-mount) `docker run` command attempted to directly bind-mount an `elasticsearch.keystore` file that doesn’t exist. If you use the `-v` or `--volume` flag to mount a file that doesn’t exist, Docker instead creates a directory with the same name.

To resolve this error:

1. Delete the `elasticsearch.keystore` directory in the `config` directory.
2. Update the `-v` or `--volume` flag to point to the `config` directory path rather than the keystore file’s path. For an example, see [Create an encrypted Elasticsearch keystore](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-keystore-bind-mount).
3. Retry the command.

#### elasticsearch.keystore: Device or resource busy[edit](https://github.com/elastic/elasticsearch/edit/8.6/docs/reference/setup/install/docker.asciidoc)

```txt
Exception in thread "main" java.nio.file.FileSystemException: /usr/share/elasticsearch/config/elasticsearch.keystore.tmp -> /usr/share/elasticsearch/config/elasticsearch.keystore: Device or resource busy
```

A `docker run` command attempted to [update the keystore](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-keystore-bind-mount) while directly bind-mounting the `elasticsearch.keystore` file. To update the keystore, the container requires access to other files in the `config` directory, such as `keystore.tmp`.

To resolve this error:

1. Update the `-v` or `--volume` flag to point to the `config` directory path rather than the keystore file’s path. For an example, see [Create an encrypted Elasticsearch keystore](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-keystore-bind-mount).
2. Retry the command.