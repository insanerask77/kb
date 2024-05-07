# DockerLabs

### PivotingGame (Beta)

![image-20240419151014005](C:\Users\rafab\AppData\Roaming\Typora\typora-user-images\image-20240419151014005.png)

Una vez desplegado el entorno vemos que tenemos varias IP's hemos comprobado que de estas IP's solo tenemos acceso a las 2 que estan marcadas en rojo. `10.10.10.2 y 20.20.20.2`

Comenzaremos con un escaneo con nmap de todos sus puertos:

```
sudo nmap -p- --open -sS --min-rate 5000 -vvv -Pn 10.10.10.2 -oG allPorts 
```

Observamos que los puertos abiertos son el `80 y 443`

Vamos a proceder a escanear los recursos mas en profuncidad con la finalidad de obtener mas informacion de los servicios expuestos en estos puertos .

```
 nmap -p80,443 -sCV 10.10.10.2
```

Vemos que nos reporta que es un apache 2.4.59 la cual no es una version vulnerable pero todo va a depender de como este configurado el entorno.

Observamos el puerto 80 y vemos la pagina de presentación de apache. Pero si observamos el puerto 443 o https vemos una pagina web.

![image-20240419151827938](C:\Users\rafab\AppData\Roaming\Typora\typora-user-images\image-20240419151827938.png)

Veo partes interesantes en la web pero voy a enumerar tambien posibles directorios con gobuster.

```
gobuster dir -u https://10.10.10.2/ -w /usr/share/wordlists/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt -x html,php,sh,py,pdf -k
```

Obtenemos el directorio assets

![image-20240419152648389](C:\Users\rafab\AppData\Roaming\Typora\typora-user-images\image-20240419152648389.png)

