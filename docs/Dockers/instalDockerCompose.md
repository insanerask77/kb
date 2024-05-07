## [Install the plugin manually](https://docs.docker.com/compose/install/linux/#install-the-plugin-manually)

> **Note**
>
> This option requires you to manage upgrades manually. We recommend setting up Docker's repository for easier maintenance.

1. To download and install the Compose CLI plugin, run:

  ```
  sudo apt-get install docker-compose-plugin
  ```

  ```console
  $ DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
  $ mkdir -p $DOCKER_CONFIG/cli-plugins
  $ curl -SL https://github.com/docker/compose/releases/download/v2.23.0/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose
  ```
  ```
  DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
  mkdir -p $DOCKER_CONFIG/cli-plugins
  curl -SL https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose
  chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose
  docker compose version
  ```

  This command downloads the latest release of Docker Compose (from the Compose releases repository) and installs Compose for the active user under `$HOME` directory.

  To install:

  - Docker Compose for *all users* on your system, replace `~/.docker/cli-plugins` with `/usr/local/lib/docker/cli-plugins`.
  - A different version of Compose, substitute `v2.23.0` with the version of Compose you want to use.

  - For a different architecture, substitute `x86_64` with the [architecture you wantopen_in_new](https://github.com/docker/compose/releases).

2. Apply executable permissions to the binary:

   ```console
   chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose
   ```

   or, if you chose to install Compose for all users:

   ```console
   sudo chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
   ```

3. Test the installation.

   content_copy

   ```console
   docker compose version
   Docker Compose version v2.23.0
   ```