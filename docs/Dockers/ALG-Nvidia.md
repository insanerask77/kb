# 1 Install Nvidia Drivers

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y build-essential libglvnd-dev pkg-config

# Note: These gcc versions will change according your kernel version.
# For utuntu 22.04, it's gcc-11 g++-11 and not the suggested 7.

sudo apt -y install gcc-7 g++-7 gcc-8 g++-8 gcc-9 g++-9
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 7
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 7
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-8 8
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-8 8
sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-9 9
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-9 9

# nvidia driver needs gcc 7, run the command and choose gcc version 7
sudo update-alternatives --config gcc

# nvidia driver needs g++ 7, run the command and choose g++ version 7
sudo update-alternatives --config g++

# verify gcc version
gcc --version
# Sample output
gcc (Ubuntu 7.5.0-6ubuntu2) 7.5.0
Copyright (C) 2017 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

# verify g++ version
g++ --version
# Sample output
g++ (Ubuntu 7.5.0-6ubuntu2) 7.5.0
Copyright (C) 2017 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

# if add-apt repository takes too long, run this comand before hand
sudo sysctl net.ipv6.conf.all.disable_ipv6=1
sudo add-apt-repository ppa:graphics-drivers/ppa -y
sudo apt update && sudo apt upgrade -y
ubuntu-drivers devices
#####
== /sys/devices/pci0000:00/0000:00:03.1/0000:09:00.0 ==
modalias : pci:v000010DEd00002204sv00001462sd00003881bc03sc00i00
vendor   : NVIDIA Corporation
driver   : nvidia-driver-515 - third-party non-free recommended
driver   : nvidia-driver-470-server - distro non-free
driver   : nvidia-driver-510 - distro non-free
driver   : nvidia-driver-510-server - distro non-free
driver   : nvidia-driver-470 - distro non-free
driver   : xserver-xorg-video-nouveau - distro free builtin
#####
# pick up the latest one with format nvidia-driver-###
sudo apt install nvidia-driver-515 -y
```

If not, try it and reinatll NVIDIA drivers



```
sudo apt-get remove --purge '^nvidia-.*'
```

## 3.2 Install NVidia-docker2





```
distribution=$(. /etc/os-release;echo $ID$VERSION_ID) \
      && curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \
      && curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.list | \
            sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
            sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
sudo apt update && sudo apt upgrade -y
sudo apt-get install -y nvidia-docker2
sudo systemctl restart docker

# validate that nvidia docker is running
sudo docker run --gpus all nvidia/cuda:11.0-base nvidia-smi

# Note: this last can change according your os version.
# For UBUNTU 20 or 22, it is:
sudo docker run --rm --gpus all nvidia/cuda:11.0.3-base-ubuntu20.04 nvidia-smi
```