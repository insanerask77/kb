# Ansible Handlers

Permiten que cuando consiga correr una tarea lo notifique para lanzar otras cosas

**No** corren de forma **lineal** aunque las declares en orden.

*Ejemplo:*

- Instala Apache2
- Renicia el servidor
  - Solo si la tarea anterior ha sido exitosa
  - Si Apache ya estaba instalado, el handler no se ejecuta

```yaml
# Install Apache2 with Ansible
---
- hosts: middlepay.zent.cash
  become: yes
  tasks:
    - name: install apache2
      apt:
        name: apache2
        state: present
        update_cache: yes
      notify:
        - start apache2
  handlers:
    - name: start apache2
      service:
        name: apache2
        state: started
        enabled: yes
```