# How to Move Home Directory to New Partition or Disk in Linux

[Aaron Kili](https://www.tecmint.com/author/aaronkili/)July 4, 2017 Categories[Linux Tutorials](https://www.tecmint.com/category/linux-tutorials/) [68 Comments](https://www.tecmint.com/move-home-directory-to-new-partition-disk-in-linux/#comments)

On any Linux system, one of the directories that will surely grow in size has to be the `/home` directory. This is because system accounts (users) directories will reside in **/home** except root account – here users will continuously store documents and other files.

Another important directory with the same behavior is `/var`, it contains log files whose size will gradually increase as the system continues to run such as log files, web files, print spool files etc.

When these directories fill up, this can cause critical problems on the root file system resulting into system boot failure or some other related issues. However, sometimes you can only notice this after installing your system and configuring all directories on the root file system/partition.

**Suggested Read:** [Linux Directory Structure and Important Files Paths Explained](https://www.tecmint.com/linux-directory-structure-and-important-files-paths-explained/)

In this guide, we will show how to move the home directory into a dedicated partition possibly on a new storage disk in Linux.

### Installing and Partitioning a New Hard Disk in Linux

Before we proceed any further, we’ll briefly explain how to add a new hard disk to an existing Linux server.

**Note**: If you already have a partition ready for the operation, move to the section which explains the steps for moving `/home` directory in a partition of its own below.

We’ll assume you have attached the new disk to the system. On a hard disk, the number of partitions to be created as well as the partition table is normally determined by disk label type and the first few bytes of space will define the **MBR** (**Master Boot Record**) which stores the partition table as well as the boot loader (for bootable disks).

Although there are many label types, Linux only accepts two: **MSDOS MBR** (516 bytes in size) or **GPT** (**GUID Partition Table**) **MBR**.

Let’s also assume that the new new hard disk (**/dev/sdb** of size **270 GB** used for the purpose of this guide, you probably need a bigger capacity on a server for large user base.

First you need to [set the disk label type](https://www.tecmint.com/change-modify-linux-disk-partition-label-names/) using [fdisk](https://www.tecmint.com/fdisk-commands-to-manage-linux-disk-partitions/) or [parted](https://www.tecmint.com/parted-command-to-create-resize-rescue-linux-disk-partitions/); we have used **GPT** label name in this example.

```
# parted /dev/sdb mklabel gpt
```

**Note**: [fdisk](https://www.tecmint.com/fdisk-commands-to-manage-linux-disk-partitions/) only supports MSDOS MBR for now and [parted](https://www.tecmint.com/parted-command-to-create-resize-rescue-linux-disk-partitions/) supports both labels.

Now create the first partition (**/dev/sdb1**) with size **106GB**. We have reserved **1024MB** of space for the MBR.

```
# parted -a cylinder /dev/sdb mkpart primary 1074MB 107GB
```

Explaining the command above:

- **a** – option to specify the partition alignment.
- **mkpart** – sub command to create the partition.
- **primary** – sets partition type as primary on the hard disk (other values are logical or extended).
- **1074MB** – beginning of partition.
- **107GB** – end of partition.

Now check the free space on the disk as follows.

```
# parted /dev/sdb print free
```

We will create another partition (**/dev/sdb2**) with size **154GB**.

```
# parted -a cylinder /dev/sdb mkpart primary 115GB 268GB
```

Next, let’s set the filesystem type on each partition.

```
# mkfs.ext4 /dev/sdb1
# mkfs.xfs /dev/sdb2
```

To view all storage devices attached on the system, type.

```
# parted -l
```

[![List New Storage Device](https://www.tecmint.com/wp-content/uploads/2017/07/List-New-Storage-Device.png)](https://www.tecmint.com/wp-content/uploads/2017/07/List-New-Storage-Device.png)List New Storage Device

#### Moving Home Directory into a Dedicated Partition

Now we have added the new disk and created the necessary partition; it’s now time to move the **home** folder into one of the partitions. To use a fileysystem, it has to be mounted to the root filesystem at a mount point: the target directory such as **/home**.

First list the filesystem usage using [df command](https://www.tecmint.com/how-to-check-disk-space-in-linux/) on the system.

```
# df -l
```

[![Linux Filesystem Usage](https://www.tecmint.com/wp-content/uploads/2017/07/Linux-Filesystem-Usage.png)](https://www.tecmint.com/wp-content/uploads/2017/07/Linux-Filesystem-Usage.png)Linux Filesystem Usage

We will start by creating a new directory **/srv/home** where we can mount **/dev/sdb1** for the time being.

```
# mkdir -p /srv/home
# mount /dev/sdb1 /srv/home 
```

Then move the content of **/home** into **/srv/home** (so they will be practically stored in **/dev/sdb1**) using [rsync command](https://www.tecmint.com/rsync-local-remote-file-synchronization-commands/) or [cp command](https://www.tecmint.com/advanced-copy-command-shows-progress-bar-while-copying-files/).

```
# rsync -av /home/* /srv/home/
OR
# cp -aR /home/* /srv/home/
```

After that, we will find the [difference between the two directories](https://www.tecmint.com/compare-find-difference-between-two-directories-in-linux/) using the [diff tool](https://www.tecmint.com/best-linux-file-diff-tools-comparison/), if all is well, continue to the next step.

```
# diff -r /home /srv/home
```

Afterwards, delete all the old content in the **/home** as follows.

```
# rm -rf /home/*
or
# rm -rvf /home/*
```

Next unmount **/srv/home**.

```
# umount /srv/home
```

Finally, we have to mount the filesystem **/dev/sdb1** to **/home** for the mean time.

```
# mount /dev/sdb1 /home
# ls -l /home
```

The above changes will last only for the current boot, add the line below in the **/etc/fstab** to make the changes permanent.

Use following command to get the partition **UUID**.

```
# blkid /dev/sdb1

/dev/sdb1: UUID="e087e709-20f9-42a4-a4dc-d74544c490a6" TYPE="ext4" PARTLABEL="primary" PARTUUID="52d77e5c-0b20-4a68-ada4-881851b2ca99"
```

Once you know the partition **UUID**, open **/etc/fstab** file add following line.

```
UUID=e087e709-20f9-42a4-a4dc-d74544c490a6   /home   ext4   defaults   0   2
```

Explaining the field in the line above:

- **UUID** – specifies the block device, you can alternatively use the device file **/dev/sdb1**.
- **/home** – this is the mount point.
- **etx4** – describes the filesystem type on the device/partition.
- **defaults** – mount options, (here this value means rw, suid, dev, exec, auto, nouser, and async).
- **0** – used by dump tool, 0 meaning don’t dump if filesystem is not present.
- **2** – used by fsck tool for discovering filesystem check order, this value means check this device after root filesystem.

Save the file and reboot the system.

You can run following command to see that **/home** directory has been successfully moved into a dedicated partition.

```
# df -hl
```

[![Check Filesystem Usage on Linux](https://www.tecmint.com/wp-content/uploads/2017/07/Check-Filesystem-Usage-on-Linux.png)](https://www.tecmint.com/wp-content/uploads/2017/07/Check-Filesystem-Usage-on-Linux.png)Check Filesystem Usage on Linux

That’s It for now! To understand more about Linux file-system, read through these guides relating to filesystem management on Linux.

1. [How to Delete User Accounts with Home Directory in Linux](https://www.tecmint.com/delete-remove-a-user-account-with-home-directory-in-linux/)
2. [What is Ext2, Ext3 & Ext4 and How to Create and Convert Linux File Systems](https://www.tecmint.com/what-is-ext2-ext3-ext4-and-how-to-create-and-convert-linux-file-systems/)
3. [7 Ways to Determine the File System Type in Linux (Ext2, Ext3 or Ext4)](https://www.tecmint.com/find-linux-filesystem-type/)
4. [How to Mount Remote Linux Filesystem or Directory Using SSHFS Over SSH](https://www.tecmint.com/sshfs-mount-remote-linux-filesystem-directory-using-ssh/)

In this guide, we explained you how to move the **/home** directory into a dedicated partition in Linux. You can share any thoughts concerning this article via the comment form below.