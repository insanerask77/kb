## Desinstala aplicaciones y juegos que ya no utilices

**La más obvia y sencilla de todas**, y por la que menos nos preocupamos en la mayoría de las ocasiones, es **ir eliminando aquellos programas o juegos que ya no utilizamos**. Quizás tenías planificado utilizarlos más adelante o los mantenías por simple nostalgia, pero ocupan un valioso espacio en tu disco duro que puedes aprovechar.

No hay motivo para mantener al mismo tiempo varios navegadores en el equipo (Chromium, Opera, Firefox, …), varios gestores de correo electrónico (Thunderbird, Claws, Evolution, …) o un sinfín de programas que realizan una función similar pero de los que sólo empleamos unos pocos. Y lo mismo sucede con los juegos. Deshazte de aquellos que no uses y recuperarás más espacio en tu unidad del que esperabas. Para ello, emplea tan solo el siguiente comando:

```
sudo apt remove paquete1 paquete2 paquete3
```

Y observa los resultados con:

```
df -h
```

Si además quiere **eliminar aquellos paquetes o dependencias que ya no se necesitan dentro del sistema**, puedes utilizar el siguiente comando:

```
sudo apt autoremove
```

## Comprime tus datos

Aunque es importante tener siempre disponibles nuestros datos, quizas quieras ahorrar algo de espacio comprimiendo aquellos ficheros que llevas tiempo sin utilizar. Seguirán estando igualmente accesibles dentro del sistema, aunque no de una forma tan directa, y a cambio ganarás algo de espacio de almacenamiento. Ya que **el periodo que determinemos para comprimir puede variar**, os dejamos un ejemplo para archivos cuyo valor sea mayor de 30 días (parámetro -mtime), que podéis modificar a vuestro gusto:

```
find . -type f -name "*" -mtime +30 -print -exec gzip {} \;
```

## Limpia la caché del APT

Quizás no habías caído en ello, pero la aplicación ***apt\* guarda mucha información en forma de caché** respecto a las actualizaciones de cada paquete que se encuentra instalado dentro del sistema. Sal de dudas y **consulta en tu sistema cuánto espacio se está desperdiciando** en tu equipo con el siguiente comando:

```
du -sh /var/cache/apt/archives
```

Si eres de los usuarios que te gusta probar aplicaciones y te pasas el día instalando, reconfigurando y desinstalando programas, puedes **deshacerte de toda aquella información inservible** que se almacena en la caché de *apt* con el siguiente comando:

```
sudo apt clean
```

Con esta función se eliminarán de Ubuntu todos los paquetes almacenados en la caché de *apt* sin importar su antigüedad. Sin embargo, **si dispones de una conexión de Internet lenta**, deberías considerar qué factor te beneficia más, el espacio de tu disco duro o el tiempo de descarga.

## Actualiza frecuentemente tu sistema

Aunque pudiera sonar confuso, en muchas ocasiones **las actualizaciones de paquetes logran optimizar los recursos de espacio** y ocupar menor tamaño dentro del equipo. Por ello, consultad a menudo las actualizaciones de paquetes y no dudéis en emplear el comando *upgrade* de vuestro *apt-get*.

## Utilizad un limpiador del sistema

Existen, como ya supondréis, **programas de terceros** que permiten de una forma más o menos eficaz **realizar una limpieza general de todo vuestro sistema**. Uno de ellos es [BleachBit](http://www.bleachbit.org/), y dada su especialidad puede llevar a cabo una tarea general de limpieza en pocos minutos.

Soporta **hasta 70 de las aplicaciones más conocidas** del entorno Linux (navegadores, gestores de correo electrónico, historial de bash, etc.) y es capaz de **eliminar archivos duplicados del sistema o aquellos con una antigüedad que indiquemos**, por lo que puede ser una alternativa muy a tener en cuenta. No obstante, sed cautos al utilizar este tipo de herramientas, pues perdemos gran parte del control sobre qué hacen y pueden arruinar nuestro sistema o información si no la manejamo con cuidado.

![bleachbit](https://ubunlog.com/wp-content/uploads/2016/08/bleachbit.jpg)

## Elimina los ficheros de kernel que no utilices

Por último y algo más alejada del ámbito convencional está la **eliminación de aquellos ficheros de \*kernel\* que no empleemos** en el sistema. La hemos reservado para el final por ser la más extrema de todas pero, si estáis seguros de que no empleáis ningún otro kernel dentro del sistema, para qué almacenar sus ficheros. Eliminadlos con este comando y liberad algunos megas de vuestro equipo:

```
sudo apt autoremove --purge
```

```
sudo apt-get -s clean
```

```
du -h --max-depth=1
```

```
find . -type f -size +100000k -exec ls -lh {} \;
```

