# Guides - Setting Up and Securing a Compute Instance

Updated August 9, 2023, by [Linode](https://www.linode.com/docs/authors/linode/)

Linux virtual machines equipped with a tailored set of resources designed to run any cloud-based workload.

After you have successfully created a Compute Instance, there are a few initial configuration steps you should perform within your new Linux system. This includes updating your system, setting the timezone, configuring a custom hostname, adding a limited user, hardening SSH to prevent unauthorized access, and configuring a firewall. These steps ensure your instance is up to date, secure, and ready for use.

Note

While this guide is optional, it walks you through best practices and covers important steps to secure your server. It is recommended that you follow these instructions when deploying a new Compute Instance. Some guides within our library assume that you have performed these steps, such as setting your hostname and updating your software.

1. [View your Instance in the Cloud Manager](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#view-your-instance-in-the-cloud-manager)
2. [Connect to the Instance](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#connect-to-the-instance)
3. [Perform System Updates](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#perform-system-updates)
4. [Set the Timezone](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#set-the-timezone)
5. [Configure a Custom Hostname](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#configure-a-custom-hostname)
6. [Add a Limited User Account](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#add-a-limited-user-account)
7. [Harden SSH Access](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#harden-ssh-access)
8. [Configure a Firewall](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#configure-a-firewall)
9. [Common Lockout Recovery Steps](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#common-lockout-recovery-steps)

## Before You Begin

If you haven’t done so already, review the following guides to learn more about using Linode and Compute Instances.

- [Getting Started with Linode](https://www.linode.com/docs/products/platform/get-started/)
- [Creating a Compute Instance](https://www.linode.com/docs/products/compute/compute-instances/guides/create/)
- [Linode Beginner’s Guide](https://www.linode.com/docs/products/compute/compute-instances/faqs/)

## View your Instance in the Cloud Manager

Log in to the [Cloud Manager](https://cloud.linode.com/), click the **Linodes** link in the left menu, and select your Compute Instance from the list. This opens the details page for that instance, which allows you to view key information and further configure it to meet your needs.



![Details page in Cloud Manager](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/create-instance-details_hu135105cff8031a37f082b923896969b0_49418_1388x0_resize_q71_bgfafafc_catmullrom_3.jpg)



## Connect to the Instance

Once the Compute Instance has been created and has finished booting up, you can connect to it. Connecting to your instance is usually done through the SSH (Secure Shell) protocol, though you can use the [Lish Console](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/) to bypass SSH and connect directly to your instance. The Lish Console can be accessed through a web browser (Weblish) or via SSH on the command line.

- **Weblish (via the Cloud Manager):** Click the **Launch LISH Console** link at the top right corner of the Compute Instance’s detail page. See [Using the Lish Console > Through a Browser](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/#through-the-cloud-manager-weblish).

- **SSH:** Copy the command from the *SSH Access* field under the **Access** section on the Compute Instance’s detail page (see screenshot above) and paste it into your local computer’s terminal. The command should look similar to the following, only with the IP address of your newly created instance.

  ```bash
  ssh root@192.0.2.17
  ```

  

  - **Windows:** Windows 10 and 11 users can connect to their Compute Instance using the [Command Prompt (or PowerShell)](https://www.linode.com/docs/guides/connect-to-server-over-ssh-on-windows/#command-prompt-or-powershell---windows-10-or-11) application, provided their system is fully updated. For users of Windows 8 and earlier, [Secure Shell on Chrome](https://www.linode.com/docs/guides/connect-to-server-over-ssh-on-chrome/), [PuTTY](https://www.linode.com/docs/guides/connect-to-server-over-ssh-using-putty/), or many other third party tools can be used instead. See [Connecting to a Remote Server Over SSH on Windows](https://www.linode.com/docs/guides/connect-to-server-over-ssh-on-windows/).
  - **macOS:** The *Terminal* application is pre-installed on macOS. See [Connecting to a Remote Server Over SSH on a Mac](https://www.linode.com/docs/guides/connect-to-server-over-ssh-on-mac/).
  - **Linux:** You can use a terminal window, regardless of desktop environment or window manager. See [Connecting to a Remote Server Over SSH on Linux](https://www.linode.com/docs/guides/connect-to-server-over-ssh-on-linux/)

- **Lish (via SSH):** Copy the command from the *LISH Console via SSH* field under the **Access** section on the Compute Instance’s detail page (see screenshot above) and paste it into your local computer’s terminal. The command should look similar to the one below, only with your username, data center, and Compute Instance label. Review [Using the Lish Console > Through SSH](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/#through-ssh-using-a-terminal) for more instructions.

  ```bash
  ssh -t user@lish-newark.linode.com example-instance
  ```

  

## Perform System Updates

Updating your system frequently is the single biggest security precaution you can take for any operating system. Software updates range from critical vulnerability patches to minor bug fixes and many software vulnerabilities are actually patched by the time they become public. Updating also provides you with the latest software versions available for your distribution.

Ubuntu, Debian, Kali LinuxCentOS/RHEL 8+, FedoraCentOS 7openSUSEAlpineArchGentooSlackware

```bash
apt update && apt upgrade
```



Note

When updating some packages, you may be prompted to use updated configuration files. If prompted, it is typically safer to keep the locally installed version.

Note

Linode’s Kali Linux distribution image is a [minimum installation](https://www.kali.org/docs/troubleshooting/common-minimum-setup/). You will likely want to install individual [tools](https://www.kali.org/tools/) or [metapackages](https://www.kali.org/tools/kali-meta/), such as the [kali-linux-headless](https://www.kali.org/tools/kali-meta/#kali-linux-headless) metapackage.

## Set the Timezone

All new Compute Instances are set to UTC time by default. However, you may prefer to use the time zone which you live in so log file timestamps are relative to your local time.

Most DistributionsUbuntu, Debian, KaliAlpineGentooopenSUSESlackware

*This includes CentOS Stream 8 (and newer), CentOS 7 (and newer), other RHEL derivatives (including AlmaLinux 8 and Rocky Linux 8), Fedora, and Arch. These instructions also work for most Ubuntu, Debian, and openSUSE distributions, though other methods may be preferred in those cases.*

1. Use `timedatectl` to output a list of available timezones.

   ```bash
   timedatectl list-timezones
   ```

   

2. Use the arrow keys, `Page Up`, and `Page Down` to navigate through the list. Copy or make note of your desired time zone and press **q** to exit the list.

3. Set the time zone using the command below, replacing *America/New_York* with your preferred time zone.

   ```bash
   timedatectl set-timezone 'America/New_York'
   ```

   

### Check the Time

Use the `date` command to view the current date and time according to your server.

```
root@localhost:~# date
Thu Feb 16 12:17:52 EST 2018
```

## Configure a Custom Hostname

A hostname is used to identify your Compute Instance using an easy-to-remember name. It can be descriptive and structured (detailing what the system is used for) or a generic word or phrase. Here are some examples of hostnames:

- **Descriptive and/or Structured:** `web`, `staging`, `blog`, or something more structured like `[purpose]-[number]-[environment]` (ex: `web-01-prod`).
- **Generic/Series:** Such as the name of a fruit (`apple`, `watermelon`), a planet (`mercury`, `venus`), or animal (`leopard`, `sloth`).

This hostname can be used as part of a FQDN (fully qualified domain name) for the system (ex: `web-01-prod.example.com`).

Most distributionsAlpineGentooSlackware

*This includes Ubuntu 16.04 (and newer), CentOS Stream 8 (and newer), CentOS 7 (and newer), other RHEL derivatives (including AlmaLinux 8 and Rocky Linux 8), Debian 8 (and newer), Fedora, openSUSE, Kali Linux, and Arch.*

Replace `example-hostname` with one of your choice.

```bash
hostnamectl set-hostname example-hostname
```



After you’ve made the changes above, you may need to log out and log back in again to see the terminal prompt change from `localhost` to your new hostname. The command `hostname` should also show it correctly. See our guide on using the [hosts file](https://www.linode.com/docs/guides/using-your-systems-hosts-file/) if you want to configure a fully qualified domain name.

### Update Your System’s `hosts` File

The `hosts` file creates static associations between IP addresses and hostnames or domains which the system prioritizes before DNS for name resolution.

1. Open the `hosts` file in a text editor, such as [Nano](https://www.linode.com/docs/guides/use-nano-to-edit-files-in-linux/).

   ```bash
   nano /etc/hosts
   ```

   

2. Add a line for your Compute Instance’s public IP address. You can associate this address with your instance’s **Fully Qualified Domain Name** (FQDN) if you have one, and with the local hostname you set in the steps above. In the example below, `203.0.113.10` is the public IP address, `example-hostname` is the local hostname, and `example-hostname.example.com` is the FQDN.

   - File: /etc/hosts

     `1 2 ``127.0.0.1 localhost.localdomain localhost 203.0.113.10 example-hostname.example.com example-hostname`

3. Add a line for your Compute Instance’s IPv6 address. Applications requiring IPv6 will not work without this entry:

   - File: /etc/hosts

     `1 2 3 ``127.0.0.1 localhost.localdomain localhost 203.0.113.10 example-hostname.example.com example-hostname 2600:3c01::a123:b456:c789:d012 example-hostname.example.com example-hostname`

The value you assign as your system’s FQDN should have an “A” record in DNS pointing to your Compute Instance’s IPv4 address. For IPv6, you should also set up a DNS “AAAA” record pointing to your instance’s IPv6 address.

See our guide to [Adding DNS Records](https://www.linode.com/docs/products/networking/dns-manager/) for more information on configuring DNS. For more information about the `hosts` file, see [Using your System’s hosts File](https://www.linode.com/docs/guides/using-your-systems-hosts-file/)

## Add a Limited User Account

Up to this point, you have accessed your Compute Instance as the `root` user, which has unlimited privileges and can execute *any* command–even one that could accidentally disrupt your server. We recommend creating a limited user account and using that at all times. Administrative tasks will be done using `sudo` to temporarily elevate your limited user’s privileges so you can administer your server. Later, when you want to restrict sudo access for users, see [Linux Users and Groups](https://www.linode.com/docs/guides/linux-users-and-groups/#understanding-sudo).

Note

Not all Linux distributions include `sudo` on the system by default, but all the images provided by Linode have sudo in their package repositories. If you get the output `sudo: command not found`, install sudo before continuing.

Ubuntu, Debian, Kali LinuxCentOS/RHEL, Fedora

1. Create the user, replacing `example_user` with your desired username. You’ll then be asked to assign the user a password:

   ```bash
   adduser example_user
   ```

   

2. Add the user to the `sudo` group so you’ll have administrative privileges:

   ```bash
   adduser example_user sudo
   ```

   

### Log in as the New User

1. After creating your limited user, disconnect from your Compute Instance:

   ```bash
   exit
   ```

   

2. Log back in as your new user. Replace `example_user` with your username, and the example IP address with your instance’s IP address:

   ```bash
   ssh example_user@192.0.2.17
   ```

   

Now you can administer your Compute Instance from your new user account instead of `root`. Nearly all superuser commands can be executed with `sudo` (example: `sudo iptables -L -nv`) and those commands will be logged to `/var/log/auth.log`.

## Harden SSH Access

By default, password authentication is used to connect to your Compute Instance via SSH. A cryptographic key-pair is more secure because a private key takes the place of a password, which is generally much more difficult to decrypt by brute-force. In this section we’ll create a key-pair and configure your system to not accept passwords for SSH logins.

### Create and Upload Your SSH Key

To protect your user account with public key authentication, you first need to create an SSH key pair and upload the public key to your server.

1. Locate your existing SSH public key or, if you don’t yet have one, create a new SSH key pair.

   - **If you have an existing SSH key,** find the public key on your local machine. SSH keys are typically stored in a hidden `.ssh` directory within the user’s home directory:

     - **Linux:** `/home/username/.ssh/`
     - **macOS:** `/Users/username/.ssh/`
     - **Windows:** `C:\Users\Username\.ssh\`

     Since SSH keys are generated as a private and public key pair, there should be two files for each SSH key. They have similar file names, with the public key using a `.pub` extension and the private key using no extension. While SSH keys can have custom file names, many people generate them using their default names. These default file names start with `id_` followed by the type of key, such as `id_rsa`, `id_ed25519`, and `id_ecdsa`. See example private and public key file names below:

     - **Private key:** `id_ed25519`
     - **Public key:** `id_ed25519.pub`

   - **If you do not yet have an SSH key pair,** generate one now. We recommend using the Ed25519 algorithm with a secure passphrase. The command below works for Linux, macOS, and most fully updated Windows 10 and 11 machines. Replace `user@domain.tld` with your own email address or whatever custom comment string you wish to use. This helps with differentiate SSH keys and identify the owner.

     ```bash
     ssh-keygen -t ed25519 -C "user@domain.tld"
     ```

     

     When prompted for the filename, you can press Enter to use the defaults. When prompted for the optional passphrase, we recommend using a string similar to a strong password (with a mix of letters, numbers, and symbols).

     For more detailed instructions, on creating an SSH key, review the [Generate an SSH Key Pair](https://www.linode.com/docs/guides/use-public-key-authentication-with-ssh/#generate-an-ssh-key-pair) guide. Users of Windows 7 and earlier should review the [PuTTY](https://www.linode.com/docs/guides/use-public-key-authentication-with-ssh/#public-key-authentication-with-putty-on-windows) section.

2. Upload the public key to your Compute Instance. Replace `example_user` with the name of the user you plan to administer the server as and `192.0.2.17` with your instance’s IP address.

   LinuxmacOSWindows 10 or 11Earlier Windows Versions

   From your local computer:

   ```bash
   ssh-copy-id example_user@192.0.2.17
   ```

   

3. Finally, you’ll want to set permissions for the public key directory and the key file itself. On your Compute Instance, run the following command:

   ```bash
   sudo chmod -R 700 ~/.ssh && chmod 600 ~/.ssh/authorized_keys
   ```

   

   This provides an extra layer of security by preventing other users from accessing the public key directory as well as the file itself. For more information on how this works, see our guide on [how to modify file permissions](https://www.linode.com/docs/guides/modify-file-permissions-with-chmod/).

4. Now exit and log back in to your Compute Instance. In most cases, the first authentication method attempted will be public key authentication. If you’ve successfully uploaded a public key for your user, you should be logged in without entering your user’s password (though you will need to enter the passphrase for the SSH key).

This should trigger If you specified a passphrase for your private key, you’ll need to enter it.

### SSH Daemon Options

Lastly, edit the SSH configuration file to disallow root login and disable password authentication over SSH.

1. Open the SSH configuration file on your Compute Instance using a Linux text editor, such as nano or vim:

   ```bash
   sudo nano /etc/ssh/sshd_config
   ```

   

2. **Disallow root logins over SSH.** This requires all SSH connections be by non-root users. Once a limited user account is connected, administrative privileges are accessible either by using `sudo` or changing to a root shell using `su -`.

   - File: /etc/ssh/sshd_config

     `1 2 3 ``# Authentication: ... PermitRootLogin no`

3. **Disable SSH password authentication.** This requires all users connecting via SSH to use key authentication. Depending on the Linux distribution, the line `PasswordAuthentication` may need to be added, or uncommented by removing the leading `#`.

   - File: /etc/ssh/sshd_config

     `1 2 ``# Change to no to disable tunnelled clear text passwords PasswordAuthentication no`

   Note

   You may want to leave password authentication enabled if you connect to your Compute Instance from many different computers. This will allow you to authenticate with a password instead of generating and uploading a key-pair for every device.

4. **Listen on only one internet protocol.** The SSH daemon listens for incoming connections over both IPv4 and IPv6 by default. Unless you need to SSH into your Compute Instance using both protocols, disable whichever you do not need. *This does not disable the protocol system-wide, it is only for the SSH daemon.* Depending on the Linux distribution, the line `AddressFamily` may need to be added, or uncommented by removing the leading `#`

   Use the option:

   - `AddressFamily inet` to listen only on IPv4.
   - `AddressFamily inet6` to listen only on IPv6.

   - File: /etc/ssh/sshd_config

     `1 2 ``# Port 22 AddressFamily inet`

5. Restart the SSH service to load the new configuration.

   - If you’re using a Linux distribution which uses systemd (CentOS 7, Debian 8, Fedora, Ubuntu 15.10+)

     ```bash
     sudo systemctl restart sshd
     ```

     

   - If your init system is SystemV or Upstart (CentOS 6, Debian 7, Ubuntu 14.04):

     ```bash
     sudo service sshd restart
     ```

     

### Use Fail2Ban for SSH Login Protection

[*Fail2Ban*](http://www.fail2ban.org/wiki/index.php/Main_Page) is an application that bans IP addresses from logging into your server after too many failed login attempts. Since legitimate logins usually take no more than three tries to succeed (and with SSH keys, no more than one), a server being spammed with unsuccessful logins indicates attempted malicious access.

Fail2Ban can monitor a variety of protocols including SSH, HTTP, and SMTP. By default, Fail2Ban monitors SSH only, and is a helpful security deterrent for any server since the SSH daemon is usually configured to run constantly and listen for connections from any remote IP address.

For complete instructions on installing and configuring Fail2Ban, see our guide: [A Tutorial for Using Fail2ban to Secure Your Server](https://www.linode.com/docs/guides/using-fail2ban-to-secure-your-server-a-tutorial/).

## Configure a Firewall

Note

Linode’s free [Cloud Firewall](https://www.linode.com/products/cloud-firewall/) service can be used to replace or supplement internal firewall configuration. For more information on Cloud Firewalls, see our [Getting Started with Cloud Firewalls](https://www.linode.com/docs/products/networking/cloud-firewall/get-started/) guide. For help with solving general firewall issues, see the [Troubleshooting Firewalls](https://www.linode.com/docs/products/compute/compute-instances/guides/troubleshooting-firewall-issues/) guide.

Using a *firewall* to block unwanted inbound traffic to your Compute Instance provides a highly effective security layer. By being very specific about the traffic you allow in, you can prevent intrusions and network mapping. A best practice is to allow only the traffic you need, and deny everything else. See our documentation on some of the most common firewall applications:

- [nftables](https://www.linode.com/docs/guides/how-to-use-nftables/) or its predecessor, [iptables](https://www.linode.com/docs/guides/control-network-traffic-with-iptables/), is the controller for netfilter, the Linux kernel’s packet filtering framework. One of these utilities is included in most Linux distributions by default.
- [firewalld](https://www.linode.com/docs/guides/introduction-to-firewalld-on-centos/) is a firewall management tool that serves as a frontend to nftables or iptables. It is preinstalled on the RHEL family of distributions (and others), including CentOS, AlmaLinux, Rocky Linux, Fedora, and openSUSE Leap.
- [UFW](https://www.linode.com/docs/guides/configure-firewall-with-ufw/) is another firewall management tool that operates as a frontend to nftables or iptables. It is used by default on Ubuntu and is also available on other Debian-based distributions.

## Common Lockout Recovery Steps

If for whatever reason you find yourself locked out of your Compute Instance after putting your security controls into place, there are still a number of ways that you can regain access to your instance.

- Access your Compute Instance through our out-of-band [Lish console](https://www.linode.com/docs/products/compute/compute-instances/guides/lish/) to regain access to the internals of your system without relying on SSH.

- If you need to re-enable password authentication and/or root login over ssh to your instance, you can do this by reversing the following sections of this file to reflect these changes

  - File: /etc/ssh/sshd_config

    `1 2 3 4 5 ``# Authentication: ... PermitRootLogin yes ... PasswordAuthentication yes`

  From there, you just need to restart SSH.

  If you’re using a Linux distribution which uses systemd (CentOS 7, Debian 8, Fedora, Ubuntu 15.10+)

  ```bash
  sudo systemctl restart sshd
  ```

  

  If your init system is SystemV or Upstart (CentOS 6, Debian 7, Ubuntu 14.04):

  ```bash
  sudo service sshd restart
  ```

  

- If you need to remove your public key from your Compute Instance, you can enter the following command:

  ```bash
  rm ~/.ssh/authorized_keys
  ```

  

  You can then replace your key by re-following the [Create an Authentication Key-pair](https://www.linode.com/docs/products/compute/compute-instances/guides/set-up-and-secure/#create-an-authentication-key-pair) section of this guide.