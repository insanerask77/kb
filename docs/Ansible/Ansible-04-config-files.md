# Ansible - Configuration Files

Changes can be made and used in a configuration file which will be searched for in the following order:

> - `ANSIBLE_CONFIG` (environment variable if set)
> - `ansible.cfg` (in the current directory)
> - `~/.ansible.cfg` (in the home directory)
> - `/etc/ansible/ansible.cfg`

#### ENV VARS

```yaml
ANSIBLE_CONFIG
```



#### Config Files

Create basic config file

```shell
ansible-config init --disabled > ansible.cfg
```

You can also have a more complete file that includes existing plugins:

```shell
ansible-config init --disabled -t all > ansible.cfg
```

- `ansible.cfg`
  - Current Directory
  - **or** `/etc/ansible/ansible.cfg`
- `.ansible.cfg`
  - Home Directory



#### Examples

```sh
[defaults]
# (string) Sets the login user for the target machines
# When blank it uses the connection plugin's default, normally the user currently executing Ansible.
remote_user=user
```

Now we can launch the ansible command without `-u user`

`ansible -i hosts.txt coin-gateway.com -m ping`

or

`ANSIBLE_CONFIG=ansible2.cfg ansible -i hosts.txt coin-gateway.com -m ping`



### WSL Issues

> [WARNING]: Ansible is being run in a world writable directory

The problem is related to the permissions given to the files by default, as in fact the latest versions these permissions (777) as a security problem and ignore the file.

To modify the permissions of the **WSL** by creating the file `/etc/wsl.conf`

```shell
[Automount]
enabled = true
mountFsTab = false
root = /mnt/
options = "metadata, umask = 22, fmask = 11"

[network]
generateHosts = true
generateResolvConf = true
```

**Restart WSL**

```powershell
wsl --shutdown
```

**Change folder permissions:**

```shell
su myuser
sudo chmod -R o-w ANSIBLE/
```

















