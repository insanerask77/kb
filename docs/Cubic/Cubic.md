# Custom Distribution Docs

# How to create Ubuntu 22.04 custom with ALG.

What do we need?

The first thing to know is that this tool has been designed to be used on an Ubuntu operating system. So we must have a version of Ubuntu installed on our system.

·     **Cubic** ([Github Official Repo](https://github.com/PJ-Singh-001/Cubic))

o  Guide (https://ostechnix.com/how-to-create-a-custom-ubuntu-live-iso-image-with-cubic/)

·     **ISO Ubuntu 22.04** (https://www.releases.ubuntu.com/22.04/ubuntu-22.04.2-desktop-amd64.iso)

·     **Ubuntu 22.04 LTS (Jammy Jellyfish) complete sources.list (**https://gist.github.com/hakerdefo/9c99e140f543b5089e32176fe8721f5f **)**

·     **Installation steps (**[Vanilla Ubuntu setup]() **)**

Once we have all these elements at our disposal, we will proceed to create our distribution.

**Step 0**

·     We will open our previously installed **Cubic** application.

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image001.gif)

**Step 1**

The first screen we are shown is the place where we are going to save our project and ISO image.

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image003.png)

Once you have selected the folder where our information will be stored, click on the top button that says next.

**Step 2**

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image004.gif)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image006.png)

As we can see in this screen we have two types of information which we can modify. On the left side we have the information contained in the official Ubuntu image, and on the right side the information that will be saved in our custom image. It is advisable to change some of these attributes to be able to differentiate the images in the future, as it is likely that different versions will be released. The fields that most differentiate the images are: **Version, Release and Filename**. Once we have modified the necessary fields we proceed with the following button.

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image008.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image010.png)

**Step 3:**

We proceed to modify the source.list since the default one is limited. We will modify the content of the file here `/etc/apt/sources.list`

```
sudo nano /etc/apt/sources.list
```

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image012.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image014.png)

With this step done, we will start the complete step-by-step installation of ALG until we reach point 4.

**Step 4**

**Note**: In the installation you will see this message each time you use `sudo` “sudo: unable to resolve host cubic: Temporary failure in name resolution“ No matter, everything is fine.

From step 4 onwards the installation varies a bit. I leave the modified steps below.

# 4 Install Vimba, Teamviewer, FFMPEG, Git, wmctrl and Ansible

```
# Download vimba software from https://www.alliedvision.com/en/products/vimba-sdk/#c1497
cd /etc/skel
curl -L https://downloads.alliedvision.com/Vimba64_v6.0_Linux.tgz > Vimba64_v6.0_Linux.tgz
tar -xzf Vimba64_v6.0_Linux.tgz -C /etc/skel
 
# TEAMVIEWER
cd /tmp
wget https://download.teamviewer.com/download/linux/signature/TeamViewer2017.asc
sudo apt-key add TeamViewer2017.asc
sudo sh -c 'echo "deb http://linux.teamviewer.com/deb stable main" >> /etc/apt/sources.list.d/teamviewer.list'
sudo apt update
sudo apt install -y teamviewer
 
# Install FFMPEG
sudo apt install -y ffmpeg
sudo apt install -y net-tools
sudo apt update && sudo apt upgrade -y
 
# Install Git
sudo apt install -y git
 
# Install wmctrl
sudo apt install -y wmctrl
```

 

# 5 Enable xhost +

```
# Execute only once
sudo echo "xhost +" >> /etc/profile
```

# 6 Repositories setup

```
# We need to add credentials in first instalation.
## Set vault
sudo git config --global credential.helper store
## Set vault in root folder
sudo git config --global credential.helper "store --file /root/.git-credentials"
# Repositories setup
mkdir /etc/skel/alg-repositories
cd /etc/skel/alg-repositories
 
# input username and password for gitlab
# Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user
sudo git clone https://username:password@git.jolibrain.com/cartier/cartier_lg_web.git
 
# input username and password for gitlab when asked
# Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user
sudo git clone https://username:password@gitlab.com/cartier-lab/megacomputer-scripts.git
 
# input username and password for gitlab when asked
# Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user
sudo git clone https://username:password@gitlab.com/cartier-lab/mc-playbooks.git
```

# 7 Disable Automatic Updates via Command Line

Update preferences are stored in the `/etc/apt/apt.conf.d/20auto-upgrades` file. Open it with nano or your favorite text editor to make some changes to it.

```
sudo nano /etc/apt/apt.conf.d/20auto-upgrades
```

To disable automatic updates completely, make sure all these directives are set to “0”. When done, save your changes and exit the file.

```
APT::Periodic::Update-Package-Lists "0";
APT::Periodic::Download-Upgradeable-Packages "0";
APT::Periodic::AutocleanInterval "0";
APT::Periodic::Unattended-Upgrade "0";
```

# 8 Create script “First Start“

```
cd /etc/skel
touch first_start.sh
sudo chmod +x first_start.sh
nano first_start.sh
# Add content that appear below.
#!/bin/bash
 
# Check if the current user is root or has sudo privileges
if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root or with sudo privileges"
    exit 1
fi
set -e
# Sequence of commands you want to run
# Install Vimba
vimba=$(find /home -type d -name VimbaUSBTL)
cd $vimba
sudo ./Uninstall.sh
sudo ./Install.sh
# Login in registries
sudo docker login registry.gitlab.com
sudo docker login https://cartierdocker.deepdetect.com
# Give permission to Docker
sudo usermod -aG docker $SUDO_USER
sudo chown -R $SUDO_USER /var/run/docker.sock
sudo chown -R $SUDO_USER /usr/local/bin/docker-compose
# Disable ipv6
sudo sysctl net.ipv6.conf.all.disable_ipv6=1
#find repos
#alg_backend=$(find /home -type d -name alg-backend)
megacomputer_scripts=$(find /home -type d -name megacomputer-scripts)
cartier_lg_web=$(find /home -type d -name cartier_lg_web)
standalone=$(find /home -type d -name standalone_setup_realtime)
alg=$(find /home -type d -name alg-repositories)
# Add safe directory
sudo git config --global --add safe.directory $megacomputer_scripts
#sudo git config --global --add safe.directory $alg_backend
sudo git config --global --add safe.directory $cartier_lg_web
 
## Setup default version alg
# Permission to alg directory
sudo chown -R $USER $alg
# Define VARS
export DATA_ROOT_PATH=/data/cartier/cartier_realtime_data
export COMPOSE_UID=$(id -u)
export COMPOSE_GUID=$(id -g)
cd $standalone
sudo docker-compose up -d
# Copy alg files into jolibrain
cp $megacomputer_scripts/config/nginx.conf $cartier_lg_web/docker/standalone_setup_realtime/config/nginx
cp $megacomputer_scripts/config/docker-compose.yml $cartier_lg_web/docker/standalone_setup_realtime
cp $megacomputer_scripts/config/.env_api_backend $cartier_lg_web/docker/standalone_setup_realtime
cp $megacomputer_scripts/config/.env_postgres $cartier_lg_web/docker/standalone_setup_realtime
cp $megacomputer_scripts/config/.env_elastic $cartier_lg_web/docker/standalone_setup_realtime
cp $megacomputer_scripts/config/filebeat.yml $cartier_lg_web/docker/standalone_setup_realtime/config
cp $megacomputer_scripts/config/logstash.conf $cartier_lg_web/docker/standalone_setup_realtime/config
# Next start to apply alg changes
sudo docker-compose up -d
# Install
cd $megacomputer_scripts
sudo ./install.sh MEGA_COMPUTER 192.168.121.101 root root /data/cartier/cartier_realtime_data
sudo apt autoremove -y
# Programed restart
echo "Press any key to cancel restart"
sleep 30
 
if read -t 0; then
  echo "Restart canceled"
else
  echo "Restarting..."
  shutdown -r now
fi
```

When we have completed the full installation and all our files are inside our custom distribution, click next, don't worry if you forget any step you can always go back to modify the project.

**Step 5**

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image016.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image018.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image020.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image022.png)

Finally, click the **generate** button and the build process of your distribution will start.

![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image024.png)![img](file:///C:/Users/rafab/AppData/Local/Temp/msohtmlclip1/01/clip_image025.gif)

To create an USB Live with generated ISO follow the steps on this tutorial:https://www.makeuseof.com/create-bootable-usb-drive-with-etcher-linux/