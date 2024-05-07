# KALI  WSL



```console
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
```

- Restart
- Open PowerShell as administrator and run:

```console
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```

- Restart

- Download and install the WSL2 Linux Kernel from here: https://aka.ms/wsl2kernel

- Open PowerShell as administrator and run: `wsl --set-default-version 2`

- Install Kali Linux from the Microsoft Store

  Note: to upgrade an existing WSL1 kali-linux installation, type: `wsl --set-version kali-linux 2`

- Run Kali and finish the initial setup

### Install Win-KeX

- Install win-kex via:

```console
kali@kali:~$ sudo apt update
kali@kali:~$
kali@kali:~$ sudo apt install -y kali-win-kex
```



----- *<u>**NO ES ORO TODO**</u>*



## Run Win-KeX

#### Win-KeX supports three modes:

- Window Mode:[![img](resources/kali-wsl/win-kex.png)](https://www.kali.org/docs/wsl/win-kex/win-kex.png)

  To start Win-KeX in Window mode with sound support, run

  `kex --win -s`

  Refer to the [Win-KeX Win usage documentation](https://www.kali.org/docs/wsl/win-kex-win/) for further information.

- Enhanced Session Mode:[![img](resources/kali-wsl/win-kex-2.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-2.png)

- To start Win-KeX in Enhanced Session Mode with sound support and arm workaround, run

  `kex --esm --ip -s`

  Refer to the [Win-KeX ESM usage documentation](https://www.kali.org/docs/wsl/win-kex-esm/) for further information.

- Seamless mode:[![img](resources/kali-wsl/win-kex-sl.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-sl.png)

  To start Win-KeX in Seamless mode with sound support, run

  `kex --sl -s`

  Refer to the [Win-KeX SL usage documentation](https://www.kali.org/docs/wsl/win-kex-sl/) for further information.

## Optional Steps:

- If you have the space, why not install “Kali with the lot”?: `sudo apt install -y kali-linux-large`

[![img](resources/kali-wsl/win-kex-thelot.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-thelot.png)

- Create a [Windows Terminal](https://www.microsoft.com/en-us/p/windows-terminal/9n0dx20hk701) Shortcut:

[![img](resources/kali-wsl/win-kex-wt1.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-wt1.png)

Choose amongst these options:

**Basic Win-KeX in window mode with sound:**

```plaintext
{
      "guid": "{55ca431a-3a87-5fb3-83cd-11ececc031d2}",
      "hidden": false,
      "name": "Win-KeX",
      "commandline": "wsl -d kali-linux kex --wtstart -s",
},
```

**Advanced Win-KeX in window mode with sound - Kali icon and start in kali home directory:**

Copy the kali-menu.png icon across to your windows picture directory and add the icon and start directory to your WT config:

```plaintext
{
        "guid": "{55ca431a-3a87-5fb3-83cd-11ececc031d2}",
        "hidden": false,
        "icon": "file:///c:/users/<windows user>/pictures/icons/kali-menu.png",
        "name": "Win-KeX",
        "commandline": "wsl -d kali-linux kex --wtstart -s",
        "startingDirectory" : "//wsl$/kali-linux/home/<kali user>"
},
```

**Basic Win-KeX in seamless mode with sound:**

```plaintext
{
      "guid": "{55ca431a-3a87-5fb3-83cd-11ececc031d2}",
      "hidden": false,
      "name": "Win-KeX",
      "commandline": "wsl -d kali-linux kex --sl --wtstart -s",
},
```

**Advanced Win-KeX in seamless mode with sound - Kali icon and start in kali home directory:**

Copy the kali-menu.png icon across to your windows picture directory and add the icon and start directory to your WT config:

```plaintext
{
        "guid": "{55ca431a-3a87-5fb3-83cd-11ececc031d2}",
        "hidden": false,
        "icon": "file:///c:/users/<windows user>/pictures/icons/kali-menu.png",
        "name": "Win-KeX",
        "commandline": "wsl -d kali-linux kex --sl --wtstart -s",
        "startingDirectory" : "//wsl$/kali-linux/home/<kali user>"
},
```

**Basic Win-KeX in ESM mode with sound:**

```plaintext
{
      "guid": "{55ca431a-3a87-5fb3-83cd-11ecedc031d2}",
      "hidden": false,
      "name": "Win-KeX",
      "commandline": "wsl -d kali-linux kex --esm --wtstart -s",
},
```

**Advanced Win-KeX in ESM mode with sound - Kali icon and start in kali home directory:**

Copy the kali-menu.png icon across to your windows picture directory and add the icon and start directory to your WT config:

```plaintext
{
        "guid": "{55ca431a-3a87-5fb3-83cd-11ecedd031d2}",
        "hidden": false,
        "icon": "file:///c:/users/<windows user>/pictures/icons/kali-menu.png",
        "name": "Win-KeX",
        "commandline": "wsl -d kali-linux kex --esm --wtstart -s",
        "startingDirectory" : "//wsl$/kali-linux/home/<kali user>"
},
```

[![img](resources/kali-wsl/win-kex-wt1.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-wt1.png)

[![img](resources/kali-wsl/win-kex-wt2.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-wt2.png)

[![img](resources/kali-wsl/win-kex-full.png)](https://www.kali.org/docs/wsl/win-kex/win-kex-full.png)

### Help

For more information, ask for help via:

```
kex --help
```

or consult the manpage via:

```
man kex
```

[![img](resources/kali-wsl/manpage.png)](https://www.kali.org/docs/wsl/win-kex/manpage.png)

or join us in the [Kali Forums](https://forums.kali.org/)

## Troubleshoots

```
sudo mount -o remount,rw /tmp/.X11-unix
sudo apt remove -y kali-win-kex && sudo apt install -y kali-win-kex
```

```
These steps solved it for me:

sudo su (some commands require root)
cd /tmp
ls -a
Delete all .XX-lock files you can find (just do rm .X1-lock, rm .X2-lock ...)
rm -r /tmp/.X11-unix
run vncserver
run kex
Enjoy
```

```
After switching to the Windows App version of WSL with
wsl.exe --update
I also had this issue.

This fixed the issue:

sudo apt remove -y kali-win-kex && sudo apt install -y kali-win-kex
sudo apt-get update
sudo apt-get upgrade
sudo su
umount /tmp/X11-unix
rm -r /tmp/.X11-unix
rm /tmp/.X1-lock
exit
sudo rm /home/YOURUSERNAME/.Xauthority
vncserver
kex

Not all steps are necessary..

Somehow

kex --win -s
fixed the audio so I have audio now. So input that if you have problems with the audio and WSL. Windows is asking you for permission for audiopulse. Confirm and then you would have the audio..
```

