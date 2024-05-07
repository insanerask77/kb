# Installation as a Docker container

*Last modified on 17-May-2023*[Edit this page on GitHub](https://github.com/Checkmk/checkmk-docs/edit/master/en/introduction_docker.asciidoc)



Related Articles



## 1. The basics

There are numerous reasons why many users would want to operate software in a Docker container. Checkmk can also be used in a Docker environment. One application scenario may be to monitor a dynamically-created container group and to make Checkmk a part of this group. Should the container group no longer be needed the Checkmk site can also be removed.

**Important**: Even it is possible and very easy to integrate Checkmk into a containerized infrastructure, it is not always the best solution. Since you get a reduced performance with every virtualization and your monitoring in general should have a minimum of physical dependencies, it is not a good solution to use a Checkmk container to monitor your complete infrastructure. But it may be a good option to integrate a Checkmk container in a self-contained container cluster, because in this case you would be able to monitor this cluster from the inside. So, especially in this case verify whether the Docker/Container tool is the best solution for your actual requirements.

In order to make the setting-up as easy as possible for you, we supply each [Checkmk edition](https://docs.checkmk.com/latest/en/intro_setup.html#editions) inclusive of its own specific image, which contains the Linux operating system Ubuntu in addition to Checkmk:

| ![CRE](https://docs.checkmk.com/latest/images/icons/CRE.png) **Checkmk Raw Edition** | [Docker Hub](https://hub.docker.com/r/checkmk/check-mk-raw/) or [Checkmk download page](https://checkmk.com/download?method=docker&edition=cre&version=stable) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Commercial editions                                          | [Checkmk download page](https://checkmk.com/download?method=docker&edition=cce&version=stable) (for the Cloud Edition), [Checkmk customer portal](https://portal.checkmk.com/) |

**Note:** Deployment in Docker Hub allows you to download and install with a single command, as we will show in the [chapter on installing the Raw Edition](https://docs.checkmk.com/latest/en/introduction_docker.html#install_cre).

In this article we will guide you through the installation of Checkmk in Docker, and show a few tricks that will make life with Checkmk in Docker easier. Further information can be found in the article [Checkmk server in a Docker container](https://docs.checkmk.com/latest/en/managing_docker.html).

## 2. Prerequisites

To execute the commands presented in this article you will need a working [Docker Engine](https://docs.docker.com/engine/install/) installation and basic knowledge of its use.

## 3. Installation of the Raw Edition

![CRE](https://docs.checkmk.com/latest/images/CRE.svg)Getting started with the ![CRE](https://docs.checkmk.com/latest/images/icons/CRE.png) **Checkmk Raw Edition** in Docker is easy. You can get a suitable image directly from the Docker Hub. This is done with just a single command on the command line. With this command not only will a Docker container with Checkmk be created, but also a monitoring site named `cmk` is set up and started. This site will be immediately available for a login as the `cmkadmin` user.

```
root@linux# docker container run -dit -p 8080:5000 -p 8000:8000 --tmpfs /opt/omd/sites/cmk/tmp:uid=1000,gid=1000 -v monitoring:/omd/sites --name monitoring -v /etc/localtime:/etc/localtime:ro --restart always checkmk/check-mk-raw:2.2.0-latest
Unable to find image 'checkmk/check-mk-raw:2.2.0-latest' locally
2.2.0-latest: Pulling from checkmk/check-mk-raw
6552179c3509: Pull complete
5a25c1702974: Pull complete
1c5c13e7b7c2: Pull complete
665c930705aa: Pull complete
597c556025e7: Pull complete
fe710cb5a7c0: Pull complete
Digest: sha256:21be9d9ded25c498834009bcb890ef678fbcdabbc5b31a168f56ab4051b9a813
Status: Downloaded newer image for checkmk/check-mk-raw:2.2.0-latest
b75e2e7d8039e73cc6e57ff09d28b4e6bccc91dbc7a85e55a3062e5fd14f596a
```

Some more information on the available options:

| Option                                             | Description                                                  |
| -------------------------------------------------- | ------------------------------------------------------------ |
| `-p 8080:5000`                                     | By default the container’s web server listens on port 5000. In this example port 8080 of the Docker node will be published to the port of the container so that it is accessible from outside. If you do not have another container or process using the standard HTTP port 80, you can also tie the container to it. In such a case the option will look like this: `-p 80:5000`. The use of HTTPS will be explained in the article [Checkmk server in a Docker container](https://docs.checkmk.com/latest/en/managing_docker.html). |
| `-p 8000:8000`                                     | Since Checkmk 2.1.0 you also have to publish the port for the Agent Receiver in order to be able to [register the agent controller](https://docs.checkmk.com/latest/en/agent_linux.html#registration). |
| `--tmpfs /opt/omd/sites/cmk/tmp:uid=1000,gid=1000` | For optimal performance, you can use a temporary file system directly in the RAM of the Docker node. The path for this file system is specified with this option. If you change the site ID this path must also be edited accordingly. |
| `-v monitoring:/omd/sites`                         | This option binds the data from the site in this container to a persistent location in the Docker node’s file system. The data is not lost if the container is deleted. The code before the colon determines the name — in this way you can clearly identify the storage location later, for example, with the `docker volume ls` command. |
| `--name monitoring`                                | This defines the name of the container. This name must be unique and may not be used again on the Docker node. |
| `-v /etc/localtime:/etc/localtime:ro`              | This option allows you to use the same time zone in the container as that used in the Docker node — at the same time the file is integrated as read only (`ro`). |
| `--restart always`                                 | A container does not normally restart automatically after it has been stopped. With this option you can ensure that it always starts again automatically. However, if you manually stop a container, it will only be restarted if the Docker daemon restarts or the container itself is restarted manually. |
| `checkmk/check-mk-raw:2.2.0-latest`                | The Checkmk image label in the `<repository>:<tag>` format. The exact labels can be read out with the command `docker images`. |

After all needed files have been loaded and the container has been started, you should access the Checkmk GUI via `http://localhost:8080/cmk/check_mk/`:

![Checkmk login dialog.](https://docs.checkmk.com/latest/images/login.png)

You can now for the first time [log in and try Checkmk out](https://docs.checkmk.com/latest/en/intro_gui.html). You will find the provisional password for the `cmkadmin` account in the logs that are written for this container (the output is abbreviated to the essential information here in this example):

```
root@linux# docker container logs monitoring
Created new site cmk with version 2.2.0p1.cre.

  The site can be started with omd start cmk.
  The default web UI is available at http://73a86e310b60/cmk/

  The admin user for the web applications is cmkadmin with password: 2JLysBmv
  For command line administration of the site, log in with 'omd su cmk'.
  After logging in, you can change the password for cmkadmin with 'cmk-passwd cmkadmin'.
```

**Note:** The URL displayed in the log for accessing the web interface with the ID of the container is only recognized within the container and is not suitable for access from outside in the web browser.

### 3.1. Short-lived containers

If you are sure that the data in the Checkmk container site should only be available in this special container, you can either refrain from assigning a persistent data storage to the container, or you can automatically remove this storage when the container is stopped.

To go without persistent storage, simply omit the `-v monitoring:/omd/sites` option:

```
root@linux# docker container run -dit -p 8080:5000 -p 8000:8000 --tmpfs /opt/omd/sites/cmk/tmp:uid=1000,gid=1000 --name monitoring -v /etc/localtime:/etc/localtime:ro --restart always checkmk/check-mk-raw:2.2.0-latest
```

To create a persistent storage and remove it automatically when the container stops, use the following command

```
root@linux# docker container run --rm -dit -p 8080:5000 -p 8000:8000 --tmpfs /opt/omd/sites/cmk/tmp:uid=1000,gid=1000 -v /omd/sites --name monitoring -v /etc/localtime:/etc/localtime:ro checkmk/check-mk-raw:2.2.0-latest
```

This command — unlike the previous one — has only two other options:

- Use the `--rm` option at the start to pass the command that the data storage for the container should also be removed when the container stops. This saves you having to tidy-up manually if you have many short-lived Checkmk containers.

  **Important**: When stopping, the container itself is completely removed!

- The `-v /omd/sites` option is altered compared to the above. It no longer contains a self-assigned name, otherwise the data storage will not be deleted correctly.

## 4. Installation of the commercial editions

![CEE](https://docs.checkmk.com/latest/images/CEE.svg)You can also run the commercial editions in a Docker container. The images of the commercial editions are not freely-available through Docker Hub. Download the desired edition and version from the [Checkmk download page](https://checkmk.com/download?method=docker&edition=cce&version=stable) (for the Cloud Edition) or from the [Checkmk customer portal](https://portal.checkmk.com/).

Load the image from the downloaded tar archive file into Docker:

```
root@linux# docker load -i check-mk-enterprise-docker-2.2.0p1.tar.gz
9e77dbc9fb80: Loading layer [==================================================>]  2.048kB/2.048kB
333c8e825260: Loading layer [==================================================>]  305.6MB/305.6MB
23a76a052da6: Loading layer [==================================================>]  175.7MB/175.7MB
f8583c4a8a97: Loading layer [==================================================>]  758.4MB/758.4MB
789d4e45d714: Loading layer [==================================================>]  6.656kB/6.656kB
Loaded image: checkmk/check-mk-enterprise:2.2.0p1
```

You can then start the container with a very similar command to that described above. Just take care that you use the name of the `Loaded image` from the previous command output in the following start command, so in this example `checkmk/check-mk-enterprise:2.2.0p1`:

```
root@linux# docker container run -dit -p 8080:5000 -p 8000:8000 --tmpfs /opt/omd/sites/cmk/tmp:uid=1000,gid=1000 -v monitoring:/omd/sites --name monitoring -v /etc/localtime:/etc/localtime:ro --restart always checkmk/check-mk-enterprise:2.2.0p1
f00d10fcb16313d3539065933b90c4dec9f81745f3d7283d794160f4f9b28df1
```

After starting the container, you can log in to the Checkmk web interface as described at the [Installation of the Raw Edition](https://docs.checkmk.com/latest/en/introduction_docker.html#login).

## 5. Update

How to update Checkmk in a Docker container is described in the [Updates and Upgrades](https://docs.checkmk.com/latest/en/update.html#updatedocker) article.

## 6. Uninstallation

When uninstalling, remove the Docker container and optionally the data created when the container was created.

Have the Docker containers listed:

```
root@linux# docker container ls -a
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS                    PORTS                              NAMES
9a82ddbabc6e   checkmk/check-mk-enterprise:2.2.0p1   "/docker-entrypoint.…"   57 minutes ago   Up 53 minutes (healthy)   6557/tcp, 0.0.0.0:8080->5000/tcp   monitoring
```

Take over the displayed `CONTAINER ID` from the command output for the next commands.

First stop the container and then remove it:

```
root@linux# docker container stop 9a82ddbabc6e
9a82ddbabc6e
root@linux# docker container rm 9a82ddbabc6e
9a82ddbabc6e
```

If you created the container with the `-v monitoring:/omd/sites` option, you can also remove the Docker volume created by this: `docker volume ls` displays the volumes and `docker volume rm <VOLUME NAME>` deletes the volume.

Finally, you can remove the image in a similar way: with `docker images` you get the list of images and `docker rmi <IMAGE ID>` removes the selected image.

```
version: '3.6'

services:
  checkmk:
    container_name: checkmk
    image: checkmk/check-mk-raw:1.6.0-latest
    tmpfs:
      - /opt/omd/sites/cmk/tmp:uid=1000,gid=1000
    ulimits:
      nofile: 1024
    volumes:
      - ./monitoring:/omd/sites
      - /etc/localtime:/etc/localtime:ro
    ports:
      - "8080:5000"
    restart: unless-stopped
    networks:
      checkmk_network:

networks:
    checkmk_network:
```

