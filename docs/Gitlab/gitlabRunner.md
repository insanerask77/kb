# Gitlab Cheatsheet

#### Gitlab Runner

create user

```sh
sudo gitlab-runner install --working-directory /home/ec2-user --user ec2-user --syslog --config /home/ec2-user/.gitlab-runner/config.toml
```

On Ubuntu 20.04 gitlab ruuner show a error if you don't delete a:

```bash
sudo rm /home/gitlab-runner/.bash_logout
```

Permissions

```sh
sudo chown -R gitlab-runner /var/run/docker.sock
```

`/etc/gitlab-runner/conf.toml`

```
concurrent = 1
check_interval = 0

[[runners]]
  name = "#####"
  url = "#####"
  token = "#####"
  executor = "docker"
  [runners.docker]
    tls_verify = false
    image = "docker:latest"
    privileged = false
    disable_cache = false
    cache_dir = "cache"
    volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/cache"]
  [runners.cache]
    Insecure = false
```

docker:dind ---> conf.toml in `/etc/gitlab-runner`

```
volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/cache"]
  [runners.cache]
    Insecure = false
```

```
volumes = ["/cache", "/var/run/docker.sock:/var/run/docker.sock"] # Ubuntu Jammy 22.04
```

## Grant sudo permissions

You can grant sudo permissions to the `gitlab-runner` user as this is who is executing the build script.

```bash
$ sudo usermod -a -G sudo gitlab-runner
```

You now have to remove the password restriction for `sudo` for the `gitlab-runner` user.

Start the sudo editor with

```bash
$ sudo visudo
```

Now add the following to the bottom of the file

```bash
gitlab-runner ALL=(ALL) NOPASSWD:ALL
```

> Do not do this for gitlab runners that can be executed by untrusted users.

![Fig.01: How to execute sudo without password for tom user](https://www.cyberciti.biz/media/new/faq/2015/05/execute-sudo-without-Password.jpg)
