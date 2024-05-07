## hecking the Linux version in the terminal

Whether you’re using Linux privately or professionally, it’s always important to know which Linux version and distribution you’re working with. That way you’ll know which package manager you’ll need for downloading new tools and updates, and which Linux forum you should turn to when you have questions or experience problems.

If you’re looking for details about your Linux version, there are two words which will be of particular significance:

1. The version number of the **distribution**
2. The version of the **Linux kernel**

To find out these two values, you’ll need to use [Linux commands](https://www.ionos.com/digitalguide/server/configuration/linux-commands-an-overview-of-terminal-commands/). In general, when working in Linux, user input is entered into so-called “shells”, which are interfaces between systems and users. Shells run using a graphic terminal that processes the commands in the relevant programming language. This will serve as your starting point in checking your Linux version.

### Step 1: Distribution version number

Open the Linux terminal with the keys [[Ctrl\]](https://www.ionos.com/digitalguide/websites/web-development/ctrl-key/) + [[Alt\]](https://www.ionos.com/digitalguide/websites/web-development/alt-key/) + [T] or by using the search function. Type the following command into the terminal and then press enter:

 

```mixed
cat /etc/*release
```

The asterisk in the code ensures that the command will apply to all distributions and shows you the installed version. The data that you see now may look a bit messy, with some lines appearing twice or several ending in “release”. The most important line here is “**PRETTY_NAME=**”, which contains the name of the distribution and version number that you’re currently using.

Another command that works on all distributions without the need for a special tool is the following:



```mixed
cat /etc/os-release
```

If you only need the name and version number of your current distribution, the following command will suffice:



```mixed
lsb_release -d
```

### CPU Stats

```
uname -a

lscpu
```

