# Netplan static IP on Ubuntu configuration

13 May 2023 by [Lubos Rendek](https://linuxconfig.org/author/lubos)

In this tutorial, we will discuss a netplan static IP configuration on Ubuntu Linux. Netplan allows for straightforward network IP address configuration using human-readable data-serialization language YAML. The article will also discuss a default Netplan network settings and the location of the Netplan configuration file.



You have two options when configuring the IP address on your Ubuntu system, and that is either a static IP address or DHCP. A static IP address allows you to manually select your IP address by configuring it on the Linux system, whereas DHCP relies on the router or DHCP server to lease you an IP address – either a reserved one or the next available one that is currently free, depending on the setup.

**In this tutorial you will learn how to:**

- Use netplan to set static IP on Ubuntu Server
- Configure netplan to set static IP on Ubuntu Server

------

![Netplan static ip on Ubuntu configuration](https://linuxconfig.org/wp-content/uploads/2018/01/00-set-configure-static-ip-address-ubuntu-18.04-bionic-linux.png)Netplan static ip on Ubuntu configuration

## Software Requirements and Conventions Used

| Category    | Requirements, Conventions or Software Version Used           |
| :---------- | :----------------------------------------------------------- |
| System      | Any [version of Ubuntu](https://linuxconfig.org/how-to-check-ubuntu-version) Linux system |
| Software    | Netplan.io                                                   |
| Other       | Privileged access to your Linux system as root or via the `sudo` command. |
| Conventions | **#** – requires given [linux commands](https://linuxconfig.org/linux-commands) to be executed with root privileges either directly as a root user or by use of `sudo` command **$** – requires given [linux commands](https://linuxconfig.org/linux-commands) to be executed as a regular non-privileged user |

## Configure static IP address using Netplan



------

------

Netplan network configuration had been first introduced starting from Ubuntu 18.04, hence Netplan is available to all new Ubuntu from this version and higher. Let’s get started with some basic understating on how netplan works on Ubuntu.



Netplan allows network configuration via both: **networkd** daemon or **NetworkManager**. **networkd** daemon is mainly used for server configuration, whereas **NetworkManager** is used by GUI users. To switch between both you need to specify `renderer` explicitly via netplan configuration file.

**NOTE**
If no `renderer` is specified within the netplan’s configuration file, then the default handler for the network configuration on this particular device will be **networkd** daemon.

The netplan configuration file location is set to `/etc/netplan/` directory. Other possible locations are `/lib/netplan/` and `/run/netplan/`. Depending on your Ubuntu installation the actual Netplan configuration file can take one of the following three forms:

- 01-netcfg.yaml
- 01-network-manager-all.yaml
- 50-cloud-init.yaml

In case you cannot find your configuration file, you may attempt to generate the new netplan config by executing the below command:

```
$ sudo netplan generate
```

**DID YOU KNOW YOU CAN CONFIGURE STATIC IP USING DHCP SERVER?**
Most likely your current Ubuntu system uses DHCP server to configure its networking settings. Hence, the configuration of your IP address is dynamic. In many scenarios, simply configuring your router or local DHCP server is a preferred way to set a static address to any host regardless of the operating system in use. Check your router manual and assign the static IP address to your host based on its [MAC address](https://linuxconfig.org/how-to-display-my-internal-ip-address-on-ubuntu-18-04-bionic-beaver-linux#h6-1-command-line) using the DHCP service.

## Netplan static ip step by step instructions

#### Ubuntu Server

1. To configure a static IP address on your Ubuntu server you need to find and modify a relevant netplan network configuration file. See the above section for all possible Netplan configuration file locations and forms.For example you might find there a default netplan configuration file called

    

   ```
   01-netcfg.yaml
   ```

    

   with a following content instructing the

    

   ```
   networkd
   ```

    

   deamon to configure your network interface via DHCP:

   ```
   # This file describes the network interfaces available on your system
   # For more information, see netplan(5).
   network:
     version: 2
     renderer: networkd
     ethernets:
       enp0s3:
         dhcp4: yes
   ```

2. To set your network interface

    

   ```
   enp0s3
   ```

    

   to static IP address

    

   ```
   192.168.1.222
   ```

    

   with gateway

    

   ```
   192.168.1.1
   ```

    

   and DNS server as

    

   ```
   8.8.8.8
   ```

    

   and

    

   ```
   8.8.4.4
   ```

    

   replace the above configuration with the one below.

   **WARNING**
   You must adhere to a correct code indent for each line of the block. In other words, the number of spaces before each configuration stanza matters. Otherwise you may end up with an error message similar to:**
   Invalid YAML at //etc/netplan/01-netcfg.yaml line 7 column 6: did not find expected key**

   ```
   # This file describes the network interfaces available on your system
   # For more information, see netplan(5).
   network:
     version: 2
     renderer: networkd
     ethernets:
       enp0s3:
        dhcp4: no
        addresses: [192.168.1.222/24]
        gateway4: 192.168.1.1
        nameservers:
          addresses: [8.8.8.8,8.8.4.4]
   ```

3. Once ready apply the new Netplan configuration changes with the following commands:

   ```
   $ sudo netplan apply
   ```

   In case you run into some issues execute:

   ```
   $ sudo netplan --debug apply
   ```