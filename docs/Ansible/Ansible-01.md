# Ansible

- Funciona por módulos
- Galaxy es un repositorio de Ansible
- Funciona con Python
- Funciona en sistemas basados en Unix
- Se puede utilizar con Windows Server
- Sólo tienes que instalar el software en tu ordenador, sin necesidad de configurar servidores intermedios
- Ansible es Software Libre GPL3 



## Instalar

##### Habilitar conexion por claves SSH

*Nos vendrá bien para que no pida el pass continuamente*

**Se puede usar un Cert sin contraseña**

```shell
# Crear la SSH Key	
ssh-keygen

# Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
#Your identification has been saved in /root/.ssh/id_rsa
#Your public key has been saved in /root/.ssh/id_rsa.pub
#The key fingerprint is:
#SHA256:JMg6QcgI..... root@....
The key's randomart image is:
+---[RSA 3072]----+
|B+A=o            |
|OB+Ao+           |
       ...
|+ ..             |
| E ..            |
+----[SHA256]-----+
```

Hay que subir el Public Key `/root/.ssh/id_rsa.pub` al servidor

```shell
ssh-copy-id -i /root/.ssh/id_rsa.pub user@coin-gateway.com
# Type remote user password
```



```shell
ssh coin-gateway.com
eval $(ssh-agent)
ssh add

# Ver la versión de Linux
cat /etc/lsb-release
```

##### Install Ansible

```shell
sudo add-apt-repository ppa:ansible/ansible


python -m pip -V

python -m pip install --user ansible
python -m pip install --user argcomplete
```



## Tareas de Ansible

Palabros chungos:

- **Invertory**
  - Donde listar cada uno de los Hosts a aprovisionar
  - Configurar el Inventory Global editando el archivo hosts.
    - Podemos usar grupos poniendo etiqueta  entre una corchetes 

```shell
nano /etc/ansible/hosts

# Ejemplo
domains.com
8.8.8.8

# GOUPS en corchetes
[webserver]
nodo1.coin-gateway.com
nodo2.coin-gateway.com
127.0.0.1

[dbservers]
db1.coin-gateway.com
db2.coin-gateway.com
127.0.0.2

```

## Tasks

De manera declarativa, no ponemos comandos. Le decimos como debería quedar el sistema cuando termine el aprovisionamiento.

Cuando hagamos playbooks, donde le decimos que ha de tener instalado y que datos necesita.

Si le decimos que instale Nginx y no está instalado, lo instala.



#### Comandos

Comandos elemntales, para comprobar que todo va correctamente

```shell
# Verifica Host conectado
# A maquinas del inventario
ansible user@coin-gateway.com -m ping
# Con tu propia lista de hosts
ansible user@coin-gateway.com -m ping -i hosts.txt
```

