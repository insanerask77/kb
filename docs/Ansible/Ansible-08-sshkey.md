# [Specifying ssh key in ansible playbook file](https://stackoverflow.com/questions/44734179/specifying-ssh-key-in-ansible-playbook-file)

Ansible playbook can specify the key used for ssh connection using `--key-file` on the command line.

```
ansible-playbook -i hosts playbook.yml --key-file "~/.ssh/mykey.pem"
```

Is it possible to specify the location of this key in playbook file instead of using `--key-file` on command line?

Because I want to write the location of this key into a `var.yaml` file, which will be read by ansible playbook with `vars_files:`.

The followings are parts of my configuration:

vars.yml file

```
key1: ~/.ssh/mykey1.pem
key2: ~/.ssh/mykey2.pem
```

playbook.yml file

```
---

- hosts: myHost
  remote_user: ubuntu
  key_file: {{ key1 }}  # This is not a valid syntax in ansible. Does there exist this kind of directive which allows me to specify the ssh key used for this connection?
  vars_files:
    - vars.yml
  tasks:
    - name: Echo a hello message
      command: echo hello
```

I've tried adding `ansible_ssh_private_key_file` under `vars`. But it doesn't work on my machine.

```
vars_files:
  - vars.yml
vars:
  ansible_ssh_private_key_file: "{{ key1 }}"
tasks:
  - name: Echo a hello message
    command: echo hello
```

If I run `ansible-playbook` with the `playbook.yml` above. I got the following error:

```
TASK [Gathering Facts] ******************************************************************************************************************************
Using module file /usr/local/lib/python2.7/site-packages/ansible/modules/system/setup.py
<192.168.5.100> ESTABLISH SSH CONNECTION FOR USER: ubuntu
<192.168.5.100> SSH: EXEC ssh -C -o ControlMaster=auto -o ControlPersist=60s -o StrictHostKeyChecking=no -o KbdInteractiveAuthentication=no -o PreferredAuthentications=gssapi-with-mic,gssapi-keyex,hostbased,publickey -o PasswordAuthentication=no -o User=ubuntu -o ConnectTimeout=10 -o ControlPath=/Users/myName/.ansible/cp/2d18691789 192.168.5.100 '/bin/sh -c '"'"'echo ~ && sleep 0'"'"''
<192.168.5.100> (255, '', 'Permission denied (publickey).\r\n')
fatal: [192.168.5.100]: UNREACHABLE! => {
    "changed": false,
    "msg": "Failed to connect to the host via ssh: Permission denied (publickey).\r\n",
    "unreachable": true
}
    to retry, use: --limit @/Users/myName/playbook.retry
```



- in the inventory file:

  ```
  myHost ansible_ssh_private_key_file=~/.ssh/mykey1.pem
  myOtherHost ansible_ssh_private_key_file=~/.ssh/mykey2.pem
  ```

- in the `host_vars`:

  ```
  # host_vars/myHost.yml
  ansible_ssh_private_key_file: ~/.ssh/mykey1.pem
  
  # host_vars/myOtherHost.yml
  ansible_ssh_private_key_file: ~/.ssh/mykey2.pem
  ```

- in a `group_vars` file if you use the same key for a group of hosts

- in the `vars` section of an entry in a play:

  ```yaml
  - hosts: myHost
     remote_user: ubuntu
     vars_files:
       - vars.yml
     vars:
       ansible_ssh_private_key_file: "{{ key1 }}"
     tasks:
       - name: Echo a hello message
         command: echo hello
  ```

- in *setting a fact* in a play entry (task):

  ```yaml
  - name: 'you name it'
     ansible.builtin.set_fact:
       ansible_ssh_private_key_file: "{{ key1 }}"
  ```



You can use the ansible.cfg file, it should look like this (There are other parameters which you might want to include):

```
[defaults]
inventory = <PATH TO INVENTORY FILE>
remote_user = <YOUR USER>
private_key_file =  <PATH TO KEY_FILE>
```

```
ansible-playbook -i hosts playbook.yml --key-file "~/.ssh/mykey.pem"
```

