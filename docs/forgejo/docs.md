[HUIJZER.XYZ](https://huijzer.xyz/)

[About](https://huijzer.xyz/about/) [Blog](https://huijzer.xyz/posts/) [![img](https://huijzer.xyz/assets/feed.svg)](https://huijzer.xyz/feed.xml)

# Installing Forgejo with a separate runner

*2024-03-08*

On the 15th of February 2024, Forgejo annouced that they will be decoupling (hard forking) their project further from Gitea. I think this is great since Forgejo is the only European Git forge that I know of, and a hard fork means that the project can now grow more independently. With Forgejo, it is now possible to self-host host a forge on a European cloud provider like Hetzner. This is great because it allows decoupling a bit from American Big Tech. Put differently, a self-hosted Forgejo avoids having all your eggs in one basket.

This post will go through a full step by step guide on how to set things up. This guide is based on my Gitea configuration that I ran for a year, so it works. During the year, I paid about 10 euros per month for two Hetzner servers. The two servers allow separating Forgejo from the runners. This ensures that a heavy job on the runner will not slow down the Forgejo server.

# [Creating a server](https://huijzer.xyz/posts/forgejo-setup/#creating_a_server)

On Hetzner, I went for the second cheapest x86 server with 2 VCPU, 2 GB RAM, and 40 GB SSD. This server responds much quicker to Git pushes than the cheapest 1 VCPU setting. The OS is set to Ubuntu 22.04. With backups and a IPv4 address, this costs €6.20 per month. For the firewall, ensure that TCP ports 22, 443, and 80 are open. For the server name, I would advice to give it a name that is easy to remember. In my case, I called it `arnold`.

Unfortunately, I do not pick the ARM server here. Even if Forgejo works with ARM (I'm not sure but it could be), then having an ARM will be restrictive. It's very cumbersome to have a full Forgejo instance running only to find out that some part doesn't work. Or, that it is not possible to co-host another service next to it. Maybe I'll switch later.

So, after the server called `arnold` is created, let's add it to our local SSH config at `~/.ssh/config`:

```julia
Host arnold
    HostName <IP ADDRESS>
    User root
    IdentityFile ~/.ssh/hetzner
```

Now, we can connect to the server with `ssh arnold`. As always with any new server, start with:

```sh
sudo apt update

sudo apt upgrade

sudo reboot
```

Next, because we're going to use Docker Compose, install Docker via their [apt repository](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository). And ensure that it works by running:

```julia
docker run hello-world
```

# [Caddy](https://huijzer.xyz/posts/forgejo-setup/#caddy)

Next, note that want to make our Forgejo server available to the outside world. This requires certificates so that a secure connection can be established. We'll use Caddy with Let's Encrypt to do this. By using Caddy as a reverse proxy, we will get HTTPS and can also use it to add extra services to the server later if we want.

Before we start Caddy, we need to make our server available on some domain. Assuming you have some domain, say `example.com` available, add the following A and AAAA records:

```julia
A git <IP ADDRESS>
AAAA git <IPv6 ADDRESS>
```

With a reasonably low TTL of say 15 minutes. By default, the TTL is often much higher which means that you need to wait for hours if you make a mistake. Now, `git.example.com` will point to our server. I will call this `<DOMAIN>` from here onward in this tutorial.

Now we can configure Caddy. Add a new directory on your server called `caddy` and put the following in `Caddyfile`:

```julia
{
  email <YOUR EMAIL ADDRESS>
  admin off
}

<DOMAIN> {
  reverse_proxy 127.0.0.1:3000
}
```

Also add a `docker-compose.yml` file:

```yml
version: "3.7"

services:
  caddy:
    image: "caddy:2.7.6-alpine"
    network_mode: "host"
    container_name: "caddy"
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
    volumes:
      - "./Caddyfile:/Caddyfile:ro"
      # This allows Caddy to cache the certificates.
      - "/data/caddy:/data:rw"
    command: "caddy run --config /Caddyfile --adapter caddyfile"
    restart: "unless-stopped"
```

The logging limits ensure that the logs will not grow infinitely. I've been there. Having to recover a server which ran out of disk space is not fun.

Now Caddy can be started with:

```sh
docker compose up
```

and the server should be available in the browser at the URL `https://<DOMAIN>`. It should show an empty page with status 502 Bad Gateway. This 502 is because we told Caddy that it should resolve to port 3000, but there is nothing there yet! All is good at this point, press CTRL + C to stop Caddy and start it again with:

```sh
docker compose up -d
```

Now the Caddy service should remain online even after you close the terminal. Thanks to `restart: "unless-stopped"`, the Caddy service will also automatically restart after a server reboot.

# [Forgejo](https://huijzer.xyz/posts/forgejo-setup/#forgejo)

Go back to the main directory and make a new directory called `forgejo`. Step into `forgejo/` and add a file called `app.ini`:

```ini
APP_NAME = git
RUN_USER = git
RUN_MODE = prod
WORK_PATH = /var/lib/forge

[server]
SSH_DOMAIN = localhost
HTTP_PORT = 3000
ROOT_URL = https://<DOMAIN>
DISABLE_SSH = true
; In rootless gitea container only internal ssh server is supported
START_SSH_SERVER = true
SSH_PORT = 2222
SSH_LISTEN_PORT = 2222
BUILTIN_SSH_SERVER_USER = git

[database]
DB_TYPE = sqlite3
HOST = localhost:3306
NAME = forge
USER = root
PASSWD = 

[security]
INSTALL_LOCK = true
REVERSE_PROXY_LIMIT = 1
REVERSE_PROXY_TRUSTED_PROXIES = *

[service]
DISABLE_REGISTRATION = true
REQUIRE_SIGNIN_VIEW = false

[actions]
ENABLED = true
DEFAULT_ACTIONS_URL = https://github.com
```

These are some values that I picked, but feel free to tweak them. This assumes that you want a personal Git forge which doesn't allow other people to register on it.

Also, I've set the `DEFAULT_ACTIONS_URL` to GitHub in order to have Forgejo be more of a drop-in replacement for the GitHub Actions. This works very well if also specifying the right runner *label*, see below.

Change `<DOMAIN>` to your git server's domain name. Next, add a file called `setup.sh`:

```sh
set -e

mkdir -p work
mkdir -p work/data

chown -R 1000:1000 work/data
chmod 775 work/data
chmod g+s work/data

chown 1000:1000 app.ini
chmod 775 app.ini
chmod g+s app.ini
```

This will setup the rootless work directory that Forgejo will use. Run this file with

```sh
bash setup.sh
```

Finally, add `docker-compose.yml`:

```yaml
version: '3'

networks:
  forgejo:
    external: false

services:
  gitea:
    image: 'codeberg.org/forgejo/forgejo:1.21-rootless'
    container_name: 'forgejo'
    environment:
      USER_UID: '1000'
      USER_GID: '1000'
      FORGEJO_WORK_DIR: '/var/lib/forge'
    user: '1000:1000'
    networks:
      - forgejo
    ports:
      - '3000:3000'
      - '222:22'
    volumes:
      - './app.ini:/etc/gitea/app.ini'
      - './data:/data:rw'
      - '/etc/timezone:/etc/timezone:ro'
      - '/etc/localtime:/etc/localtime:ro'
      # Depends on `FORGEJO_WORK_DIR`.
      - './work:/var/lib/forge:rw'
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "10"
    restart: 'unless-stopped'
```

and start this with `docker-compose up`.

While this is running, open another terminal to add an admin user. First, step into the running container:

```bash
docker exec -it forgejo /bin/bash
```

and then add an admin user:

```bash
forgejo admin user create --username <USERNAME> --password <PASSWORD> --email <EMAIL> --admin
```

Now Forgejo should be up and running at `https://<DOMAIN>` and signing in should work with the newly created admin account. If there are errors, try restarting the server with `reboot now`. With that, both Caddy and Forgejo restart which might solve the problem.

# [Forgejo runner](https://huijzer.xyz/posts/forgejo-setup/#forgejo_runner)

Having a forge is one thing, but in my opinion a CI runner is also a must have. For that, we setup another Hetzner server and install the Forgejo runner on that. Also here, I advice to take a x86 server as ARM will likely give problems. Set the server up in the same way as before and SSH into it again. This time, I called the server `runner` and ssh into it with `ssh runner`.

Run `update` and `upgrade`, and install Docker and reboot, like before.

Next, create `setup.sh`:

```sh
#!/usr/bin/env bash

set -e

mkdir -p data
touch data/.runner
mkdir -p data/.cache

chown -R 1001:1001 data/.runner
chown -R 1001:1001 data/.cache
chmod 775 data/.runner
chmod 775 data/.cache
chmod g+s data/.runner
chmod g+s data/.cache
```

and run with

```sh
bash setup.sh
```

Then create `docker-compose.yml` with:

```yaml
version: '3.8'

services:
  docker-in-docker:
    image: docker:dind
    container_name: 'docker_dind'
    privileged: true
    command: [ "dockerd", "-H", "tcp://0.0.0.0:2375", "--tls=false" ]
    restart: 'unless-stopped'

  gitea:
    image: 'code.forgejo.org/forgejo/runner:3.3.0'
    links:
      - docker-in-docker
    depends_on:
      docker-in-docker:
        condition: service_started
    container_name: 'runner'
    environment:
      DOCKER_HOST: tcp://docker-in-docker:2375
    # A user without root privileges, but with access to `./data`.
    user: 1001:1001
    volumes:
      - ./data:/data
    restart: 'unless-stopped'

    command: '/bin/sh -c "while : ; do sleep 1 ; done ;"'
```

This `command` doesn't start the runner yet, we first register it with the server. To do so, run:

```sh
docker compose up
```

And in another terminal, run:

```sh
docker exec -it runner /bin/bash
```

And

```sh
forgejo-runner register
```

with instance URL: `https://<DOMAIN>`.

For the runner token, browse to the following URL:

```julia
https://<DOMAIN>/user/settings/actions/runners
```

to get it.

For the runner name, I used `hetzner_runner`.

Then for the labels, and this is **very important**, use:

```julia
ubuntu-22.04:docker://ghcr.io/catthehacker/ubuntu:act-22.04
```

This label specifies to which workflows the runner will respond. In this case, it will respond to `ubuntu-22.04` workflows with a `ghcr.io/catthehacker/ubuntu:act-22.04` container. This is a 1.2 GB container with a lot of pre-installed software. It's not as feature-full as the GitHub runners, but those are about 20 GB in size, so this is a good compromise. See [`nektos/act`](https://github.com/nektos/act/blob/master/IMAGES.md) for more information about the available containers.

The runner should now be visible at

```julia
https://<DOMAIN>/user/settings/actions/runners
```

with the status `Offline`.

To fix that, exit the Docker Compose by pressing CTRL+C a few times and modify the `command` to

```yaml
command: '/bin/sh -c "sleep 5; forgejo-runner daemon"'
```

The sleep here provides the `docker-in-docker` service with some extra time to get started. Run `docker compose up -d` to start the runner in the background.

The runner should now be visible at

```julia
https://<DOMAIN>/user/settings/actions/runners
```

with the status `Idle`.

# [Testing](https://huijzer.xyz/posts/forgejo-setup/#testing)

Finally, let's test whether the runner works. To do so, create a new repository via the web interface. In the new repository enable repository Actions at

```julia
https://<DOMAIN>/<USERNAME>/<REPO>/settings
```

and click "Update Settings".

Then, either via the web interface or via cloning the repository and working there, create the following files.

`.github/workflows/ci.yml`:

```yaml
name: ci

on:
  push:
    branches:
      - main
  pull_request:
  workflow_dispatch:

jobs:
  test:
    runs-on: ubuntu-22.04
    steps:
      - uses: actions/checkout@v4
      - run: |
          curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
          source "$HOME/.cargo/env"
          echo "$PATH" >> $GITHUB_PATH
      - run: |
          rustup update stable
          rustup default stable
      - name: Cache
        uses: Swatinem/rust-cache@v2
        with:
          prefix-key: 'rust'
      - run: rustc hello.rs
      - run: ./hello
```

`hello.rs`:

```rust
fn main() {
    println!("Hello from Rust!");
}
```

Note that this workflow did have to manually install `rustup` whereas that is installed in the GitHub Runners by default. This is because our `node:20` docker image doesn't have `rustup` installed by default. `nektos/act` also has more extensive docker images, but those are multiple GB in size. The `nektos/act-environments-ubuntu:18.04-full`, for example, is 12.1 GB.

When pushing these changes to the repository, the workflow should run and print "Hello from Rust!" in the last step.

At this point, you have a fully functional personal forge with a runner 🎉🎉.