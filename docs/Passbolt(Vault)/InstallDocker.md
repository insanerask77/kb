# Docker passbolt installation

**Important:** Installing Passbolt with Docker is considered a somewhat advanced method. Using this method assumes you are familiar with Docker and have run other applications with Docker. If you do not have experience working with Docker we recommend you use another of our installation methods.

## System requirements

- docker: https://docs.docker.com/get-docker/
- docker-compose: https://docs.docker.com/compose/install/
- [A Linux user able to run docker commands without sudo](https://docs.docker.com/engine/install/linux-postinstall/)
- a working SMTP server for email notifications
- a working NTP service to avoid GPG authentication issues

FAQ pages:

- [Set up NTP](https://help.passbolt.com/faq/hosting/set-up-ntp)
- [Firewall rules](https://help.passbolt.com/faq/hosting/firewall-rules)

## docker-compose

The easiest and recommended way to deploy your passbolt stack is to use docker-compose.

**Step 1.** Download our docker-compose.yml example file

```
wget https://download.passbolt.com/ce/docker/docker-compose-ce.yaml
wget https://github.com/passbolt/passbolt_docker/releases/latest/download/docker-compose-ce-SHA512SUM.txt
```

**Step 2.** Ensure the file has not been corrupted by verifying its shasum

```
$ sha512sum -c docker-compose-ce-SHA512SUM.txt
```

Must return:

```
docker-compose-ce.yaml: OK
```

**Warning:** If the *shasum* command output is not correct, the downloaded file has been corrupted. Retry step 1 or ask for support on [our community forum](https://community.passbolt.com/).

**Step 3.** Configure environment variables in docker-compose-ce.yaml file to customize your instance.

**Notice:** By default the docker-compose.yaml file is set to **latest**. We strongly recommend changing that to the [tag](https://hub.docker.com/r/passbolt/passbolt/tags) for the version you want to install.

The `APP_FULL_BASE_URL` environment variable is set by default to [https://passbolt.local](https://passbolt.local/), using a self-signed certificate.

Update this variable with the server name you plan to use. You will find at the bottom of this documentation links about how to set your own SSL certificate.

You must configure also SMTP settings to be able to receive notifications and recovery emails. Please find below the most used environment variables for this purpose:

| Variable name                    | Description                    | Default value     |
| :------------------------------- | :----------------------------- | :---------------- |
| EMAIL_DEFAULT_FROM_NAME          | From email username            | `'Passbolt'`      |
| EMAIL_DEFAULT_FROM               | From email address             | `'you@localhost'` |
| EMAIL_TRANSPORT_DEFAULT_HOST     | Server hostname                | `'localhost'`     |
| EMAIL_TRANSPORT_DEFAULT_PORT     | Server port                    | `25`              |
| EMAIL_TRANSPORT_DEFAULT_USERNAME | Username for email server auth | `null`            |
| EMAIL_TRANSPORT_DEFAULT_PASSWORD | Password for email server auth | `null`            |
| EMAIL_TRANSPORT_DEFAULT_TLS      | Set tls                        | `null`            |

For more information on which environment variables are available on passbolt, please check the [passbolt environment variable reference](https://help.passbolt.com/configure/environment/reference.html).

**Step 4.** Start your containers

```
docker-compose -f docker-compose-ce.yaml up -d
```

**Step 5.** Create first admin user

```
$ docker-compose -f docker-compose-ce.yaml exec passbolt su -m -c "/usr/share/php/passbolt/bin/cake \
                                passbolt register_user \
                                -u <your@email.com> \
                                -f <yourname> \
                                -l <surname> \
                                -r admin" -s /bin/sh www-data
```

It will output a link similar to the below one that can be pasted on the browser to finalize user registration:

```
https://my.domain.tld/setup/install/1eafab88-a17d-4ad8-97af-77a97f5ff552/f097be64-3703-41e2-8ea2-d59cbe1c15bc
```

At this point, you should have a working docker setup running on the **latest** tag. However, it is recommended that users [pull the tags pointing to specific passbolt versions](https://hub.docker.com/r/passbolt/passbolt/tags) when running in environments other than testing.