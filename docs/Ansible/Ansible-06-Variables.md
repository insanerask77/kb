```yaml
---
- hosts: middlepay.zent.cash
  become: true
  vars:
    container_count: "1"
    default_container_name: "docker"
    default_container_image: "ubuntu"
    default_container_command: "sleep 1"
  tasks:
  - name: Pull default Docker image
    community.docker.docker_image:
      name: "{{ default_container_image }}"
      source: pull
  - name: Create default containers
    community.docker.docker_container:
      name: "{{ default_container_name }}{{ item }}"
      image: "{{ default_container_image }}"
      command: "{{ default_container_command }}"
      state: present
    with_sequence: count={{ container_count }}
```