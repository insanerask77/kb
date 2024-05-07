# Dockers Best Practices



### SECURING DOCKER DAEMON AND CORE COMPONENTS



##### DAEMON ROOTLESS MODE



1. Get and run the installation script:

   `  curl -fsSL https://get.docker.com/rootless | sh`

2.  Add the following environment variables to ~/.bashrc:

   ` export XDG_RUNTIME_DIR=/tmp/docker-1000 ` 

   `export PATH=/home/non-root/bin:$PATH export `

   `DOCKER_HOST=unix:///tmp/docker-1000/docker.sock`

3. Run the Docker daemon in rootless mode

​	`/home/non-root/bin/dockerd-rootless.sh --experimental --storage-driver vfs &` 

4. Check if Docker is actually running in rootless mode, for that, run „docker info“ command and  check the „Security Options“:

   `docker info`

​	

##### SECURING DOCKER SOCKET

##### SECURING API ENDPOINT

##### SECURING DOCKER CONTAINERS