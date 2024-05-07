[Hosting n8n](https://docs.n8n.io/hosting/)[Installation](https://docs.n8n.io/hosting/installation/)

# Docker Installation[#](https://docs.n8n.io/hosting/installation/docker/#docker-installation)

[Docker](https://www.docker.com/) offers the following advantages:

- Install n8n in a clean environment.
- Easier setup for your preferred database.
- Can avoid issues due to different operating systems, as Docker provides a consistent system.

## Prerequisites[#](https://docs.n8n.io/hosting/installation/docker/#prerequisites)

Before proceeding, install [Docker Desktop](https://docs.docker.com/get-docker/).

Linux Users

Docker Desktop is available for Mac and Windows. Linux users must install [Docker Engine](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/) individually for your distribution.

Self-hosting knowledge prerequisites

Self-hosting n8n requires technical knowledge, including:

- Setting up and configuring servers and containers
- Managing application resources and scaling
- Securing servers and applications
- Configuring n8n

n8n recommends self-hosting for expert users. Mistakes can lead to data loss, security issues, and downtime. If you aren't experienced at managing servers, n8n recommends [n8n Cloud](https://n8n.io/cloud/).

Latest and Next versions

n8n releases a new minor version most weeks. The `latest` version is for production use. `next` is the most recent release. You should treat `next` as a beta: it may be unstable. To report issues, use the [forum](https://community.n8n.io/c/questions/12).

Current `latest`: 1.15.2
Current `next`: 1.16.0

## Starting n8n[#](https://docs.n8n.io/hosting/installation/docker/#starting-n8n)

From your terminal, run:

| `1 2 3` | `docker volume create n8n_data docker run -it --rm --name n8n -p 5678:5678 -v n8n_data:/home/node/.n8n docker.n8n.io/n8nio/n8n ` |
| ------- | ------------------------------------------------------------ |
|         |                                                              |

This command will download all required n8n images and start your container, exposed on port `5678`. To save your work between container restarts, it also mounts a docker volume, `n8n_data`, to persist your data locally.

You can then access n8n by opening: [http://localhost:5678](http://localhost:5678/)

## Using alternate databases[#](https://docs.n8n.io/hosting/installation/docker/#using-alternate-databases)

By default n8n uses SQLite to save credentials, past executions and workflows. n8n also supports PostgresDB configurable using environment variables as detailed below.

It's important to still persist data in the `/home/node/.n8n` folder as it contains n8n user data and even more importantly the encryption key for credentials. It's also the name of the webhook when the n8n tunnel is used.

If no directory is found, n8n creates automatically one on startup. In this case, existing credentials saved with a different encryption key can not be used anymore.

Keep in mind

Persisting the `/home/node/.n8n` directory even when using alternate databases is the recommended best practice, but not explicitly required. The encryption key can be provided via the `N8N_ENCRYPTION_KEY` [environment variable](https://docs.n8n.io/hosting/environment-variables/environment-variables/#deployment).

### PostgresDB[#](https://docs.n8n.io/hosting/installation/docker/#postgresdb)

To use n8n with Postgres, provide the corresponding [configuration](https://docs.n8n.io/hosting/configuration/):

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13 14` | `docker volume create n8n_data docker run -it --rm \ --name n8n \ -p 5678:5678 \ -e DB_TYPE=postgresdb \ -e DB_POSTGRESDB_DATABASE=<POSTGRES_DATABASE> \ -e DB_POSTGRESDB_HOST=<POSTGRES_HOST> \ -e DB_POSTGRESDB_PORT=<POSTGRES_PORT> \ -e DB_POSTGRESDB_USER=<POSTGRES_USER> \ -e DB_POSTGRESDB_SCHEMA=<POSTGRES_SCHEMA> \ -e DB_POSTGRESDB_PASSWORD=<POSTGRES_PASSWORD> \ -v n8n_data:/home/node/.n8n \ docker.n8n.io/n8nio/n8n ` |
| ----------------------------------- | ------------------------------------------------------------ |
|                                     |                                                              |

A complete `docker-compose` file for Postgres can be found [here](https://github.com/n8n-io/n8n/blob/master/docker/compose/withPostgres/).

### MySQL[#](https://docs.n8n.io/hosting/installation/docker/#mysql)

Deprecated

n8n deprecated MySQL and MariaDB as backend databases in version 0.227.0.

n8n recommends using PostgreSQL.

Refer to [how to export and import workflows and credentials](https://docs.n8n.io/hosting/cli-commands/) for instructions.

To use n8n with MySQL, provide the corresponding [configuration](https://docs.n8n.io/hosting/configuration/):

| ` 1 2 3 4 5 6 7 8 9 10 11 12 13` | `docker volume create n8n_data docker run -it --rm \ --name n8n \ -p 5678:5678 \ -e DB_TYPE=mysqldb \ -e DB_MYSQLDB_DATABASE=<MYSQLDB_DATABASE> \ -e DB_MYSQLDB_HOST=<MYSQLDB_HOST> \ -e DB_MYSQLDB_PORT=<MYSQLDB_PORT> \ -e DB_MYSQLDB_USER=<MYSQLDB_USER> \ -e DB_MYSQLDB_PASSWORD=<MYSQLDB_PASSWORD> \ -v n8n_data:/home/node/.n8n \ docker.n8n.io/n8nio/n8n ` |
| -------------------------------- | ------------------------------------------------------------ |
|                                  |                                                              |

## Setting timezone[#](https://docs.n8n.io/hosting/installation/docker/#setting-timezone)

To define the timezone n8n should use, the environment variable `GENERIC_TIMEZONE` can be set. This gets used by schedule based nodes such as the Cron node.

The timezone of the system can also be set separately. This controls what some scripts and commands return like `$ date`. The system timezone can be set via the environment variable `TZ`.

Example using the same timezone for both:

| `1 2 3 4 5 6 7 8 9` | `docker volume create n8n_data docker run -it --rm \ --name n8n \ -p 5678:5678 \ -e GENERIC_TIMEZONE="Europe/Berlin" \ -e TZ="Europe/Berlin" \ -v n8n_data:/home/node/.n8n \ docker.n8n.io/n8nio/n8n ` |
| ------------------- | ------------------------------------------------------------ |
|                     |                                                              |

## Updating[#](https://docs.n8n.io/hosting/installation/docker/#updating)

From your Docker Desktop, navigate to the **Images** tab and select **Pull** from the context menu to download the latest n8n image:

![Docker Desktop](https://docs.n8n.io/_images/hosting/installation/docker/docker_desktop.png)

You can also use the command line to pull the latest, or a specific version:

| `1 2 3 4 5 6 7 8` | `# Pull latest (stable) version docker pull docker.n8n.io/n8nio/n8n # Pull specific version docker pull docker.n8n.io/n8nio/n8n:0.220.1 # Pull next (unstable) version docker pull docker.n8n.io/n8nio/n8n:next ` |
| ----------------- | ------------------------------------------------------------ |
|                   |                                                              |

Stop the container and start it again. You can also use the command line:

| ` 1 2 3 4 5 6 7 8 9 10 11` | `# Get the container ID docker ps -a # Stop the container with ID container_id docker stop [container_id] # Remove the container with ID container_id docker rm [container_id] # Start the container docker run --name=[container_name] [options] -d docker.n8n.io/n8nio/n8n ` |
| -------------------------- | ------------------------------------------------------------ |
|                            |                                                              |

### Docker Compose[#](https://docs.n8n.io/hosting/installation/docker/#docker-compose)

If you run n8n using a Docker Compose file, follow these steps to update n8n:

| `1 2 3 4 5 6 7 8` | `# Pull latest version docker compose pull # Stop and remove older version docker compose down # Start the container docker compose up -d ` |
| ----------------- | ------------------------------------------------------------ |
|                   |                                                              |

## Further reading[#](https://docs.n8n.io/hosting/installation/docker/#further-reading)

More information about Docker setup can be found in the README file of the [Docker Image](https://github.com/n8n-io/n8n/tree/master/docker/images/n8n).

## n8n with tunnel[#](https://docs.n8n.io/hosting/installation/docker/#n8n-with-tunnel)

Danger

This is only meant for local development and testing. Do not use it in production.

To be able to use webhooks for trigger nodes of external services like GitHub, n8n has to be reachable from the web. To make that easy, n8n has a special [tunnel service](https://github.com/localtunnel/localtunnel) which redirects requests from our servers to your local n8n instance.

Start n8n with `--tunnel` by running:

| `1 2 3 4 5 6 7 8` | `docker volume create n8n_data docker run -it --rm \ --name n8n \ -p 5678:5678 \ -v n8n_data:/home/node/.n8n \ docker.n8n.io/n8nio/n8n \ start --tunnel` |
| ----------------- | ------------------------------------------------------------ |
|                   |                                                              |