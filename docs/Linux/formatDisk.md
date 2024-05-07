## Formatting Disk Partition in Linux

There are three ways to format disk partitions using the **`mkfs`** command, depending on the file system type:

- ext4
- FAT32
- NTFS

The general syntax for formatting disk partitions in Linux is:

```
mkfs [options] [-t type fs-options] device [size]
```

```
fdisk -l | grep '^Disk'
```



### Formatting Disk Partition with ext4 File System

\1. Format a disk partition with the ext4 file system using the following command:

```
sudo mkfs -t ext4 /dev/sdb1
```

\2. Next, verify the file system change using the command:

```
lsblk -f
```

The terminal prints out a list of block devices.

\3. Locate the preferred partition and confirm that it uses the ext4 file system.

![The process of formatting disk partition using the ext4 file system in Linux.](https://phoenixnap.com/kb/wp-content/uploads/2021/04/formatting-ext4-partition.png)

### Formatting Disk Partition with FAT32 File System

\1. To format a disk with a FAT32 file system, use:

```
sudo mkfs -t vfat /dev/sdb1
```

\2. Again, run the **`lsblk`** command to verify the file system change and locate the preferred partition from the list.

```
lsblk -f
```

The expected output is:

![Formatting disk partition using FAT32 file system in Linux.](https://phoenixnap.com/kb/wp-content/uploads/2021/04/formatting-fat32-partition.png)

### Formatting Disk Partition with NTFS File System

\1. Run the **`mkfs`** command and specify the NTFS file system to format a disk:

```
sudo mkfs -t ntfs /dev/sdb1
```

The terminal prints a confirmation message when the formatting process completes.

\2. Next, verify the file system change using:

```
lsblk -f
```

\3. Locate the preferred partition and confirm that it uses the NFTS file system.

![Formatting a partition with NTFS file system in Linux.](https://phoenixnap.com/kb/wp-content/uploads/2021/04/formatting-ntfs-partition.png)

## Mounting the Disk Partition in Linux

Before using the disk, create a mount point and mount the partition to it. A mount point is a directory used to access data stored in disks.

\1. Create a mount point by entering:

```
sudo mkdir -p [mountpoint]
```

\2. After that, mount the partition by using the following command:

```
sudo mount -t auto /dev/sdb1 [mountpoint]
```

**Note:** Replace **`[mountpoint]`** with the preferred mount point (example:**`/usr/media`**).

There is no output if the process completes successfully.

![The process of creating mount point and mounting.](https://phoenixnap.com/kb/wp-content/uploads/2021/04/creating-mounting-point-and-mounting-partition.png)

\3. Verify if the partition is mounted using the following command:

```
lsblk -f
```

The expected output is:

![Verifying formatting and mounting partition process. ](https://phoenixnap.com/kb/wp-content/uploads/2021/04/verifying-formatting-and-mounting.png)

## Understanding the Linux File System

Choosing the right file system before formatting a storage disk is crucial. Each type of file system has different file size limitations or different operating system compatibility.

The most commonly used file systems are:

- FAT32
- NTFS
- ext4

Their main features and differences are:

| **File System** | **Supported File Size** | **Compatibility**                                      | **Ideal Usage**                             |
| --------------- | ----------------------- | ------------------------------------------------------ | ------------------------------------------- |
| **FAT32**       | up to 4 GB              | Windows, Mac, Linux                                    | For maximum compatibility                   |
| **NTFS**        | 16 EiB – 1 KB           | Windows, Mac (read-only), most Linux distributions     | For internal drives and Windows system file |
| **Ext4**        | 16 GiB – 16 TiB         | Windows, Mac, Linux (requires extra drivers to access) | For files larger than 4 GB                  |