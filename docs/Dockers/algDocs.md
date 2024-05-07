# Vanilla Ubuntu setup

- 
- 

Owned by [Jason AGYEKUM](https://cartier.atlassian.net/wiki/people/63a0d3e315d69a40aa16acc7?ref=confluence&src=profilecard)



Last updated: [about 5 hours ago](https://cartier.atlassian.net/wiki/pages/diffpagesbyversion.action?pageId=1056669766&selectedPageVersions=7&selectedPageVersions=8) by [Rafael MADOLELL](https://cartier.atlassian.net/wiki/people/712020:ad0d2cfd-e134-474a-8abb-c07f17573f8d?ref=confluence&src=profilecard)

14 min read4 people viewed

- [Xorg setup ](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#Xorg-setup)
- [0 Install Rust Desktop](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#0-Install-Rust-Desktop)
- [1 Install Nvidia Drivers](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#1-Install-Nvidia-Drivers)
- [2 Install basic packages](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#2-Install-basic-packages)
- [3 Install Docker & docker-compose](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#3-Install-Docker-&-docker-compose)
  - [3.1 Setup permissions for docker-compose](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#3.1-Setup-permissions-for-docker-compose)
  - [3.2 Install NVidia-docker2](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#3.2-Install-NVidia-docker2)
- [4 Install Vimba, Teamviewer, FFMPEG, Git, wmctrl and Ansible ](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#4-Install-Vimba,-Teamviewer,-FFMPEG,-Git,-wmctrl-and-Ansible)
- [5 Enable xhost +](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#5-Enable-xhost-+)
- [6 Repositories setup](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#6-Repositories-setup)
- [7 Disable Automatic Updates via Command Line](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#7-Disable-Automatic-Updates-via-Command-Line)
- [8 Putting everything together and running it](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#8-Putting-everything-together-and-running-it)
- [9 Install Gstreamer + Plugins](https://cartier.atlassian.net/wiki/spaces/LG/pages/1056669766/Vanilla+Ubuntu+setup#9-Install-Gstreamer-+-Plugins)

![:warning:](https://pf-emoji-service--cdn.us-east-1.prod.public.atl-paas.net/atlassian/warning_32.png)

NOTE: latest Ubuntu versions have Wayland set as a default display server. However, this doesn’t work well with TeamViewer. We strongly suggest using Ubuntu with Xorg. 



For any problem during or after the setup please refer to: [Troubleshooting articles](https://brooklynlab.atlassian.net/wiki/spaces/A/pages/205258760) 

# Xorg setup 





```
sudo nano /etc/gdm3/custom.conf
```

and then make sure that WaylandEnabled is set to false





```
WaylandEnable=false
```

Save and then logout and login again from the Ubuntu session.

# 0 Install Rust Desktop





```
## Anydesk install sudo apt update && sudo apt upgrade -y sudo apt-get install -y \    apt-transport-https \    ca-certificates \    curl \    gnupg-agent \    software-properties-common wget https://github.com/rustdesk/rustdesk/releases/download/1.2.3/rustdesk-1.2.3-x86_64.deb -O ./rustdesk-1.2.3-x86_64.deb sudo apt install -y ./rustdesk-1.2.3-x86_64.deb
```

 

# 1 Install Nvidia Drivers





```
sudo apt update && sudo apt upgrade -y sudo apt install -y build-essential libglvnd-dev pkg-config # Note: These gcc versions will change according your kernel version. # For utuntu 22.04, it's gcc-11 g++-11 and not the suggested 7. sudo apt -y install gcc-7 g++-7 gcc-8 g++-8 gcc-9 g++-9 sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 7 sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 7 sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-8 8 sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-8 8 sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-9 9 sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-9 9 # nvidia driver needs gcc 7, run the command and choose gcc version 7 sudo update-alternatives --config gcc # nvidia driver needs g++ 7, run the command and choose g++ version 7 sudo update-alternatives --config g++ # verify gcc version gcc --version # Sample output gcc (Ubuntu 7.5.0-6ubuntu2) 7.5.0 Copyright (C) 2017 Free Software Foundation, Inc. This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. # verify g++ version g++ --version # Sample output g++ (Ubuntu 7.5.0-6ubuntu2) 7.5.0 Copyright (C) 2017 Free Software Foundation, Inc. This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. # if add-apt repository takes too long, run this comand before hand sudo sysctl net.ipv6.conf.all.disable_ipv6=1 sudo add-apt-repository ppa:graphics-drivers/ppa -y sudo apt update && sudo apt upgrade -y ubuntu-drivers devices ##### == /sys/devices/pci0000:00/0000:00:03.1/0000:09:00.0 == modalias : pci:v000010DEd00002204sv00001462sd00003881bc03sc00i00 vendor   : NVIDIA Corporation driver   : nvidia-driver-515 - third-party non-free recommended driver   : nvidia-driver-470-server - distro non-free driver   : nvidia-driver-510 - distro non-free driver   : nvidia-driver-510-server - distro non-free driver   : nvidia-driver-470 - distro non-free driver   : xserver-xorg-video-nouveau - distro free builtin ##### # pick up the latest one with format nvidia-driver-### sudo apt install nvidia-driver-515 -y
```

If the machine has **secure boot** see the following instructions





















``























 





```
sudo reboot nvidia-smi
```

Check for CUDA Version and Driver Version above or equal to the ones shown here. It should also list some processes.

Sample output of `nvidia-smi`





```
##################################### # Sample output Fri May 13 12:56:11 2022        +-----------------------------------------------------------------------------+ | NVIDIA-SMI 510.68.02    Driver Version: 510.68.02    CUDA Version: 11.6     | |-------------------------------+----------------------+----------------------+ | GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC | | Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. | |                               |                      |               MIG M. | |===============================+======================+======================| |   0  NVIDIA GeForce ...  Off  | 00000000:01:00.0 Off |                  N/A | | N/A   48C    P8     9W /  N/A |      5MiB /  8192MiB |      0%      Default | |                               |                      |                  N/A | +-------------------------------+----------------------+----------------------+                                                                                +-----------------------------------------------------------------------------+ | Processes:                                                                  | |  GPU   GI   CI        PID   Type   Process name                  GPU Memory | |        ID   ID                                                   Usage      | |=============================================================================| |    0   N/A  N/A       943      G   /usr/lib/xorg/Xorg                  4MiB | +-----------------------------------------------------------------------------+ #####################################
```

NVIDIA Driver should be successfully installed.

If not, try it and reinatll NVIDIA drivers





```
sudo apt-get remove --purge '^nvidia-.*'
```

# 2 Install basic packages





```
sudo apt update && sudo apt upgrade -y curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - sudo apt-key fingerprint 0EBFCD88
```

Expected output of `sudo apt-key fingerprint 0EBFCD88`





```
##################################### # expected output pub   rsa4096 2017-02-22 [SCEA]    9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88 uid           [ unknown] Docker Release (CE deb) <docker@docker.com> sub   rsa4096 2017-02-22 [S] #####################################
```

# 3 Install Docker & docker-compose





```
# multiline command sudo add-apt-repository \ "deb [arch=amd64] https://download.docker.com/linux/ubuntu \ $(lsb_release -cs) \ stable" # end of multiline command sudo apt update && sudo apt upgrade -y sudo apt-get install -y docker-ce docker-ce-cli containerd.io # verify docker is running sudo docker run hello-world
```

Expected output of `sudo docker run hello-world`





```
##################################### # expected output Unable to find image 'hello-world:latest' locally latest: Pulling from library/hello-world 2db29710123e: Pull complete  Digest: sha256:80f31da1ac7b312ba29d65080fddf797dd76acfb870e677f390d5acba9741b17 Status: Downloaded newer image for hello-world:latest Hello from Docker! This message shows that your installation appears to be working correctly. #####################################
```

## 3.1 Setup permissions for docker-compose





```
sudo usermod -a -G docker $USER sudo reboot sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose sudo chmod +x /usr/local/bin/docker-compose sudo rm /usr/bin/docker-compose sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose sudo systemctl restart docker
```

## 3.2 Install NVidia-docker2





```
distribution=$(. /etc/os-release;echo $ID$VERSION_ID) \      && curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \      && curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.list | \            sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \            sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list sudo apt update && sudo apt upgrade -y sudo apt-get install -y nvidia-docker2 sudo systemctl restart docker # validate that nvidia docker is running sudo docker run --gpus all nvidia/cuda:11.0-base nvidia-smi # Note: this last can change according your os version. # For UBUNTU 20 or 22, it is: sudo docker run --rm --gpus all nvidia/cuda:11.0.3-base-ubuntu20.04 nvidia-smi
```

Expected output of `sudo docker run --gpus all nvidia/cuda:11.0-base nvidia-smi`





```
##################################### # expected output Unable to find image 'nvidia/cuda:11.0-base' locally 11.0-base: Pulling from nvidia/cuda 54ee1f796a1e: Pull complete  f7bfea53ad12: Pull complete  46d371e02073: Pull complete  b66c17bbf772: Pull complete  3642f1a6dfb3: Pull complete  e5ce55b8b4b9: Pull complete  155bc0332b0a: Pull complete  Digest: sha256:774ca3d612de15213102c2dbbba55df44dc5cf9870ca2be6c6e9c627fa63d67a Status: Downloaded newer image for nvidia/cuda:11.0-base Fri May 13 17:37:15 2022        +-----------------------------------------------------------------------------+ | NVIDIA-SMI 510.68.02    Driver Version: 510.68.02    CUDA Version: 11.6     | |-------------------------------+----------------------+----------------------+ | GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC | | Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. | |                               |                      |               MIG M. | |===============================+======================+======================| |   0  NVIDIA GeForce ...  Off  | 00000000:01:00.0  On |                  N/A | | N/A   46C    P8    14W /  N/A |     54MiB /  8192MiB |     16%      Default | |                               |                      |                  N/A | +-------------------------------+----------------------+----------------------+                                                                                +-----------------------------------------------------------------------------+ | Processes:                                                                  | |  GPU   GI   CI        PID   Type   Process name                  GPU Memory | |        ID   ID                                                   Usage      | |=============================================================================| +-----------------------------------------------------------------------------+ #####################################
```

Validate NVidia-docker2 works in docker-compose





```
# Validate that nvidia runtime works in docker-compose # For Ubuntu 18 echo "services:  test:    image: nvidia/cuda:11.0.3-base-ubuntu20.04    command: nvidia-smi    runtime: nvidia" > docker-compose.yml     # For Ubuntu 20 or higher echo "services:  test:    image: nvidia/cuda:11.0.3-base-ubuntu20.04    command: nvidia-smi    runtime: nvidia" > docker-compose.yml docker-compose up
```

Expected output of `docker-compose up`





```
##################################### # expected output Creating nicoloco_test_1 ... done Attaching to nicoloco_test_1 test_1  | Fri May 13 19:02:01 2022        test_1  | +-----------------------------------------------------------------------------+ test_1  | | NVIDIA-SMI 510.68.02    Driver Version: 510.68.02    CUDA Version: 11.6     | test_1  | |-------------------------------+----------------------+----------------------+ test_1  | | GPU  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC | test_1  | | Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | GPU-Util  Compute M. | test_1  | |                               |                      |               MIG M. | test_1  | |===============================+======================+======================| test_1  | |   0  NVIDIA GeForce ...  Off  | 00000000:01:00.0 Off |                  N/A | test_1  | | N/A   46C    P0    28W /  N/A |      5MiB /  8192MiB |      0%      Default | test_1  | |                               |                      |                  N/A | test_1  | +-------------------------------+----------------------+----------------------+ test_1  |                                                                                 test_1  | +-----------------------------------------------------------------------------+ test_1  | | Processes:                                                                  | test_1  | |  GPU   GI   CI        PID   Type   Process name                  GPU Memory | test_1  | |        ID   ID                                                   Usage      | test_1  | |=============================================================================| test_1  | +-----------------------------------------------------------------------------+ #####################################
```

# 4 Install Vimba, Teamviewer, FFMPEG, Git, wmctrl and Ansible 





```
# Download vimba software from https://www.alliedvision.com/en/products/vimba-sdk/#c1497 curl -L https://downloads.alliedvision.com/Vimba64_v6.0_Linux.tgz > Vimba64_v6.0_Linux.tgz tar -xzf Vimba64_v6.0_Linux.tgz -C ~/ sudo Vimba_6_0/VimbaUSBTL/Install.sh sudo reboot # TEAMVIEWER cd /tmp wget https://download.teamviewer.com/download/linux/signature/TeamViewer2017.asc sudo apt-key add TeamViewer2017.asc sudo sh -c 'echo "deb http://linux.teamviewer.com/deb stable main" >> /etc/apt/sources.list.d/teamviewer.list' sudo sh -c 'echo "deb http://linux.teamviewer.com/deb preview main" >> /etc/apt/sources.list.d/teamviewer.list' sudo apt update sudo apt install -y teamviewer ##################################################################################### # DEPRECATED                             # Install FFMPEG sudo apt install -y ffmpeg sudo apt update && sudo apt upgrade -y ##################################################################################### # Install net-tools sudo apt install -y net-tools # Install Git sudo apt install -y git # Install wmctrl sudo apt install -y wmctrl # Install Ansible sudo apt install -y ansible # Install jq sudo apt install -y jq # Install Python3 sudo add-apt-repository ppa:deadsnakes/ppa sudo apt-get update sudo apt-get install python3.8 
```

 

# 5 Enable xhost +





```
# Execute only once sudo echo "xhost +" >> /etc/profile # if it happens to give permission denied do the following: sudo gedit /etc/profile # add one new line at the end of the file: xhost + sudo reboot
```

# 6 Repositories setup





```
# We need to add credentials in first instalation. ## Set vault sudo git config --global credential.helper store ## Set vault in root folder sudo git config --global credential.helper "store --file /root/.git-credentials" # Repositories setup mkdir ~/alg-repositories cd ~/alg-repositories # input username and password for github # Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user sudo git clone https://username:password@git.jolibrain.com/cartier/cartier_lg_web.git # Delete the last history record  history -d $(history | cut -d " " -f 3 | tail -n 2 | head -n 1) # input username and password for github when asked # Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user sudo git clone https://username:password@github.com/cartier-lab/megacomputer-scripts.git # Delete the last history record  history -d $(history | cut -d " " -f 3 | tail -n 2 | head -n 1) # input username and password for github when asked # Complete command with credentials in https://brooklynlab.atlassian.net/wiki/spaces/A/pages/16973825/Mega+computer+credentials#BTS-repository-user sudo git clone https://username:password@github.com/cartier-lab/mc-playbooks.git # Delete the last history record  history -d $(history | cut -d " " -f 3 | tail -n 2 | head -n 1)
```

# 7 Disable Automatic Updates via Command Line

Update preferences are stored in the `/etc/apt/apt.conf.d/20auto-upgrades` file. Open it with nano or your favorite text editor to make some changes to it.





```
sudo nano /etc/apt/apt.conf.d/20auto-upgrades
```

 To disable automatic updates completely, make sure all these directives are set to “0”. When done, save your changes and exit the file.





```
APT::Periodic::Update-Package-Lists "0"; APT::Periodic::Download-Upgradeable-Packages "0"; APT::Periodic::AutocleanInterval "0"; APT::Periodic::Unattended-Upgrade "0";
```

# 8 Putting everything together and running it





```
# Login to cartier docker registry # Complete command in 'Mega computer credentials' docker login https://cartierdocker.deepdetect.com # Setup configuration export  DATA_ROOT_PATH=/data/cartier/cartier_realtime_data \ COMPOSE_UID=$(id -u) \ COMPOSE_GUID=$(id -g) sudo chown -R $USER ~/alg-repositories cd ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime docker-compose up -d ## If you want generate id's use this script cd ~/alg-repositories/megacomputer-scripts docker pull ghcr.io/car-ww-dataoffice/alg-backend:latest docker run --rm --entrypoint env ghcr.io/car-ww-dataoffice/alg-backend:latest | tail -n +5 | head -n -2 >>  ~/alg-repositories/megacomputer-scripts/config/.env_api_backend sudo ~/alg-repositories/megacomputer-scripts/oauth_config_generator.sh sudo ~/alg-repositories/megacomputer-scripts/id_generator.sh sudo cp  ~/alg-repositories/megacomputer-scripts/config/nginx.conf  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config/nginx sudo cp  ~/alg-repositories/megacomputer-scripts/config/nginx-cloud.conf  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config/nginx sudo cp  ~/alg-repositories/megacomputer-scripts/config/.htpasswd  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config/nginx mkdir  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/docker_app sudo cp  ~/alg-repositories/megacomputer-scripts/config/docker_app/docker-compose.yml  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/docker_app sudo cp  ~/alg-repositories/megacomputer-scripts/config/docker-compose.yml  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/ sudo cp  ~/alg-repositories/megacomputer-scripts/config/docker-compose-cloud.yml  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/ sudo cp  ~/alg-repositories/megacomputer-scripts/config/.env_api_backend  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/ sudo cp  ~/alg-repositories/megacomputer-scripts/config/.env_elastic  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/ sudo cp  ~/alg-repositories/megacomputer-scripts/config/filebeat.yml  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config sudo cp  ~/alg-repositories/megacomputer-scripts/config/logstash.conf  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config sudo cp  ~/alg-repositories/megacomputer-scripts/config/metricbeat.yml  ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config sudo  ~/alg-repositories/megacomputer-scripts/peer.sh sudo ~/alg-repositories/megacomputer-scripts/setup_manager_install.sh cp -r ~/alg-repositories/megacomputer-scripts/config/wireguard ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/config # Login to ALG cartier BTS docker registry # user gitlab+deploy-token-prod-1 # Complete command in 'Mega computer credentials' sudo docker login ghcr.io cd ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime sudo docker-compose up -d cd ~/alg-repositories/cartier_lg_web/docker/standalone_setup_realtime/docker_app sudo docker-compose up -d # BTS ALG setup cd ~/alg-repositories/megacomputer-scripts sudo ./install.sh MEGA_COMPUTER 192.168.121.101 root root /data/cartier/cartier_realtime_data sudo reboot sudo apt autoremove -y # Initial data should be automatically seeded # Verify seed data curl http://localhost:3000/rings
```

Sample output from `curl http://localhost:3000/rings`





```
##################################### [{ "is_deleted": false, "id": 1, "ring_id": "CRH4386653_essaouira_real_time", "name": "ESSAOUIRA RING", "reference_number": "CRH4386653", "description": "WHITE GOLD 750/1000, FLUTED EMERALD AND CHALCEDONY BEADS, CALIBRATED AND PRINCESS-CUT SAPPHIRES, ONYX, TURQUOISE, ALL SET WITH 56 BRILLIANT-CUT DIAMONDS TOTALING 1.22 CARATS", "ring_image": "/data/rings/CRH4386653_essaouira_real_time/ef7f94f3-7ef6-459b-a4a6-b35634ecc0ff.png", "ring_image_hd": "/data/rings/CRH4386653_essaouira_real_time/d688e36f-8550-445b-8abe-d03096604bb9.png", "overlay_image_url": "/data/rings/CRH4386653_essaouira_real_time/a9b2fc71-4b12-410f-bfcd-fc8fdc8e4ccf.png", "cover_image_url": "/data/rings/CRH4386653_essaouira_real_time/a933b838-a3a5-4622-8698-98a79427a7c3.png", "category": "FJ", "catalog_url": "https://plaza.cartier.com/products/25383/informations", "updated": 1652901981905 }, { "is_deleted": false, "id": 2, "ring_id": "CRZH420009_emerald_real_time", "name": "EMERALD RING", "reference_number": "CRZH420009", "description": "PLATINUM, ONE 9.01 CARAT EMERALD-CUT EMERALD FROM COLOMBIA, TWO TAPERED BAGUETTE DIAMONDS TOTALING 0.64 CARATS D-E/VVS1", "ring_image": "/data/rings/CRZH420009_emerald_real_time/f82c5eee-cda0-4fbc-86dc-06ef5d3b4cda.png", "ring_image_hd": "/data/rings/CRZH420009_emerald_real_time/cc0febaf-6271-4c2a-a656-e3803ddf850b.png", "overlay_image_url": "/data/rings/CRZH420009_emerald_real_time/190bea0f-c046-40e1-a138-345197590018.png", "cover_image_url": "/data/rings/CRZH420009_emerald_real_time/ab00b3f0-3bf1-4928-9a45-55beb893b4ee.png", "category": "HJ", "catalog_url": "https://plaza.cartier.com/products/30528/informations", "updated": 1652901981927 }, { "is_deleted": false, "id": 3, "ring_id": "CRH4372653_sanyogita_real_time", "name": "SANYOGITA RING", "reference_number": "CRH4372653", "description": "PLATINUM, ONE 10.00-CARAT RUBY FROM MOZAMBIQUE, CARVED RUBIES, SAPPHIRES AND EMERALDS, ONYX, BRILLIANT-CUT DIAMONDS", "ring_image": "/data/rings/CRH4372653_sanyogita_real_time/75c53d90-c8a7-4cbb-992b-2e4daecaf4d1.png", "ring_image_hd": "/data/rings/CRH4372653_sanyogita_real_time/5e462792-b4c1-4d88-8485-dfd32a3b1152.png", "overlay_image_url": "/data/rings/CRH4372653_sanyogita_real_time/ca81ac1c-ed9f-4240-b841-d116c1bbaa40.png", "cover_image_url": "/data/rings/CRH4372653_sanyogita_real_time/d34e5e79-9b77-452b-b73b-d91b84c56fb5.png", "category": "HJ", "catalog_url": "https://plaza.cartier.com/products/21965/informations", "updated": 1652901981947 }] ##################################### 
```

# 9 Install Gstreamer + Plugins





```
# Install gstreamer + good, bad and ugly plugins sudo apt install libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-libav gstreamer1.0-tools gstreamer1.0-x gstreamer1.0-alsa gstreamer1.0-gl gstreamer1.0-gtk3 gstreamer1.0-qt5 gstreamer1.0-pulseaudio # Check installed version gst-launch-1.0 --gst-version
```