

# How to Add Swap Space on Ubuntu 20.04

Swap is a space on a disk that is used when the amount of physical RAM memory is full. When a Linux system runs out of RAM, inactive pages are moved from the RAM to the swap space.

Swap space can take the form of either a dedicated swap partition or a swap file. Typically, when running Ubuntu on a virtual machine, a swap partition is not present, and the only option is to create a swap file.

This tutorial explains how to add a swap file on Ubuntu 20.04.

## Before You Begin

Swap should not be seen as a replacement to physical memory. Since swap space is a section of the hard drive, it has a slower access time than physical memory. If your system constantly runs out of memory, you should add more RAM.

Generally, the size of the swap file depends on how much RAM your system has:

- Systems with **less than 2 GB RAM** - 2 times the amount of RAM.
- Systems with **2 to 8 GB RAM** - the same size as the amount of RAM.
- Systems with **more than 8 GB RAM** - at least 4 GB of Swap.

Only root or user with [sudo privileges](https://linuxize.com/post/how-to-create-a-sudo-user-on-ubuntu/) can activate the swap file.

## Creating a Swap File

In this example, we will create `2 GB` swap file. If you want to add more swap, replace `2G` with the size of the swap space you need.

Complete the steps below to add swap space on Ubuntu 20.04:

1. First, create a file that will be used as swap:

   ```
   sudo fallocate -l 2G /swapfile
   ```

   If the `fallocate` utility is not present on your system, or you get an error message saying `fallocate failed: Operation not supported`, use the following command to create the swap file:

   ```
   sudo dd if=/dev/zero of=/swapfile bs=1024 count=2097152
   ```

2. Set the file [permissions](https://linuxize.com/post/chmod-command-in-linux/) to `600` to prevent regular users to write and read the file:

   ```
   sudo chmod 600 /swapfile
   ```

3. Create a Linux swap area on the file:

   ```
   sudo mkswap /swapfile
   ```

   ```output
   Setting up swapspace version 1, size = 2 GiB (2147479552 bytes)
   no label, UUID=fde7d2c8-06ea-400a-9027-fd731d8ab4c8
   
   ```

4. Activate the swap file by running the following command:

   ```
   sudo swapon /swapfile
   ```

   To make the change permanent open the `/etc/fstab` file:

   ```
   sudo nano /etc/fstab
   ```

   and paste the following line:

   /etc/fstab

   ```ini
   /swapfile swap swap defaults 0 0
   ```

   

5. Verify that the swap is active by using either the `swapon` or the [`free`](https://linuxize.com/post/free-command-in-linux/) command, as shown below:

   ```
   sudo swapon --show
   ```

   ```output
   NAME      TYPE      SIZE  USED PRIO
   /swapfile file        2G    0B   -1
   ```

   ```
   sudo free -h
   ```

   ```output
                 total        used        free      shared  buff/cache   available
   Mem:          981Mi        97Mi        68Mi       0.0Ki       814Mi       735Mi
   Swap:         2.0Gi        10Mi       1.9Gi
   ```

## Adjusting the Swappiness Value

Swappiness is a Linux kernel property that defines how often the system will use the swap space. It can have a value between 0 and 100. A low value will make the kernel to try to avoid swapping whenever possible, while a higher value will make the kernel to use the swap space more aggressively.

On Ubuntu, the default swappiness value is set to `60`. You can check the current value by typing the following command:

```
cat /proc/sys/vm/swappiness
60

```

While the swappiness value of `60` is OK for most Linux systems, for production servers, you may need to set a lower value.

For example, to set the swappiness value to `10`, run:

```
sudo sysctl vm.swappiness=10
```

To make this parameter persistent across reboots, append the following line to the `/etc/sysctl.conf` file:

/etc/sysctl.conf

```ini
vm.swappiness=10
```



The optimal swappiness value depends on your system workload and how the memory is being used. You should adjust this parameter in small increments to find an optimal value.

## Removing a Swap File

To deactivate and delete the swap file, follow these steps:

1. First, deactivate the swap space:

   ```
   sudo swapoff -v /swapfile
   ```

2. Next, remove the swap file entry `/swapfile swap swap defaults 0 0` from the `/etc/fstab` file.

3. Finally, remove the actual swapfile file using the [`rm`](https://linuxize.com/post/rm-command-in-linux/) command:

   ```
   sudo rm /swapfile
   ```

## Conclusion

We have shown you how to create a swap file and activate and configure swap space on your Ubuntu 20.04 system.

If you hit a problem or have feedback, leave a comment below.