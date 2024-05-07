# Ansible Playbooks



**Usamos YAML**

editamos el fichero sites.yml

```yaml
---
- hosts: user@coin-gateway.com
  # become: yes
  # remote_user: root
  tasks:
  - name: Install dependencies
    apt: name={{ item }} state=present
    become: yes
  - name: saluda
    shell: echo hola
```

Con `ansible-playbook site.yml -K ` 

```
ansible-playbook -i hosts playbook.yml --key-file "~/.ssh/mykey.pem"
```

### Servicios

https://docs.ansible.com/ansible/2.9/modules/service_module.html#service-manage-services

- Controla los servicios en hosts remotos.
- Para Windows targets, usa [win_service](https://docs.ansible.com/ansible/2.9/modules/win_service_module.html#win-service-module)



##### [Examples](https://docs.ansible.com/ansible/2.9/modules/service_module.html#id5)

```yaml
---
- hosts: user@coin-gateway.com
  # become: yes
  tasks:
  - name: dependencies
    #apt: name={{ item }} state=present
    apt: name=nano state=present
    become: yes
  - name: saluda
    shell: echo hola
  - name: servicio
    service: name=nginx state=stopped
    become: yes
```

```yaml
- name: Start service httpd, if not started
  service:
    name: httpd
    state: started

- name: Stop service httpd, if started
  service:
    name: httpd
    state: stopped

- name: Restart service httpd, in all cases
  service:
    name: httpd
    state: restarted

- name: Reload service httpd, in all cases
  service:
    name: httpd
    state: reloaded

- name: Enable service httpd, and not touch the state
  service:
    name: httpd
    enabled: yes

- name: Start service foo, based on running process /usr/bin/foo
  service:
    name: foo
    pattern: /usr/bin/foo
    state: started

- name: Restart network service for interface eth0
  service:
    name: network
    state: restarted
    args: eth0
```

```yaml
- hosts: 192.168.0.90
  # become: yes
  # remote_user: root
  tasks:
    - name: Install aptitude
      apt:
        name: aptitude
        state: latest
        update_cache: true

    - name: Install required system packages
      apt:
        pkg:
          - apt-transport-https
          - ca-certificates
          - curl
          - software-properties-common
          - python3-pip
          - virtualenv
          - python3-setuptools
        state: latest
        update_cache: true

    - name: Add Docker GPG apt Key
      apt_key:
        url: https://download.docker.com/linux/ubuntu/gpg
        state: present

    - name: Add Docker Repository
      apt_repository:
        repo: deb https://download.docker.com/linux/ubuntu focal stable
        state: present

    - name: Update apt and install docker-ce
      apt:
        name: docker-ce
        state: latest
        update_cache: true

    - name: Install Docker Module for Python
      pip:
        name: docker
```

```yaml
---
- name: Do something that requires a reboot when it results in a change.
  ...
  register: task_result

- name: Reboot immediately if there was a change.
  shell: "sleep 5 && reboot"
  async: 1
  poll: 0
  when: task_result is changed

- name: Wait for the reboot to complete if there was a change.
  wait_for_connection:
    connect_timeout: 20
    sleep: 5
    delay: 5
    timeout: 300
  when: task_result is changed
```
