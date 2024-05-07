# Redimensionar discos duros de Servidores Cloud en Linux sin reinicios

Como comentábamos en nuestro post anterior, en el que utilizábamos la aplicación Gparted para que Linux interpretara el redimensionamiento de discos duros realizado desde Cloudbuilder Next, si queremos prescindir de reiniciar el servidor para que su nueva capacidad esté operativa, **necesitamos emplear algo de código**. En este artículo, explicamos cómo hacerlo con las distribuciones Ubuntu y CentOS y te recordamos el proceso en Windows para redimensionar los discos duros de los Servidores Cloud.



Antes de seguir avanzando y meternos con el código, te recordamos que el primer paso es [modificar la capacidad del almacenamiento del Servidor](https://www.arsys.es/soporte/cloud-y-servidores/editar-servidor) desde el Panel de Control de Cloudbuilder Next.

A continuación, explicamos cómo conseguir que el sistema operativo Linux interprete el nuevo tamaño asignado al disco duro sin reiniciar los [servidores](https://www.arsys.es/servidores/vps). Esta opción es más compleja ya que hay que ejecutar una serie de comandos, conectados por SSH al servidor en cuestión, así que vamos a describir todo este proceso por medio de una serie de comandos en las distribuciones Ubuntu y CentOS. Sin embargo, es importante darse cuenta que en los casos particulares de cada máquina puedan cambiar nombres de discos y volúmenes.

**Tabla de contenidos**

[1 Redimensionar discos duros en Ubuntu](https://www.arsys.es/blog/redimensionar-discos-duros-linux-no-reinicio#Redimensionar_discos_duros_en_Ubuntu)

[2 Redimensionar particiones en CentOS](https://www.arsys.es/blog/redimensionar-discos-duros-linux-no-reinicio#Redimensionar_particiones_en_CentOS)

## Redimensionar discos duros en Ubuntu

Vamos a comenzar viendo la secuencia de comandos de consola que nos permitirá reorganizar las particiones en una distribución Ubuntu:

Comprobamos el identificador del disco usando el comando:

```
ls /sys/class/scsi_disk/
```

Será algo como *2:0:0:0, 3:0:0:0, 4:0:0:0*. La salida la podemos ver en la siguiente imagen.

![Redimensionar discos duros en Ubuntu](https://www.arsys.es/blog/file/uploads/2019/09/4-idisco.png)

Hacemos un escaneo del disco, usando ese identificador:

```
echo 1 > /sys/class/scsi_device/4:0:0:0/device/rescan
```

En el comando anterior, ten cuenta que probablemente necesitarás cambiar el identificador *4:0:0:0* por el que te haya aparecido al hacer *ls /sys/class/scsi_disk/*.

Ahora vamos a hacer un análisis para comprobar discos, particiones y tamaños. Esto se consigue con *fdisk*.

```
fdisk -l
```

![fdisk](https://www.arsys.es/blog/file/uploads/2019/09/4fdisk-1.png)Deberíamos observar que el tamaño de los discos ya se han actualizado y qué particiones tenemos creadas en este momento, junto con sus dimensiones.

Ahora vamos a crear una partición con el espacio no particionado que se ha generado en el disco como consecuencia de la redimensión. Para ello, lanzamos el comando:

```
fdisk /dev/sda
```

Este comando nos abre una interfaz de menú capaz de recibir varias instrucciones para administrar las particiones. Tenemos que ir introduciendo uno a uno las distintas opciones de menú por medio de su inicial. Serán las siguientes:

- ***n #*** Nueva partición.
- ***p #*** Primaria.
- ***3 #*** Número de la partición.
- ***Enter #\*** Primer sector, por defecto el primero libre.
- ***Enter #*** Último sector, por defecto el último del disco.
- ***w #** E*scribir cambios.

![administrar las particiones](https://www.arsys.es/blog/file/uploads/2019/09/4-crear-particion.png)

Una vez creada esta partición, tenemos que marcarla como tipo Linux LVM.

Tenemos que volver a lanzar el menú anterior con el comando:

```
fdisk /dev/sda
```

Y vamos usando las siguientes opciones de menú:

- ***t #***Cambiar tipo de partición.
- ***3 #*** Número de partición.
- ***8e #*** Tipo LVM.
- ***w #*** Escribir cambios.

![opciones de menú](https://www.arsys.es/blog/file/uploads/2019/09/4linux-lvm.png)

El siguiente paso consiste en escanear las particiones de disco para detectar los cambios.

```
apt update -yy
apt install -yy parted
partprobe /dev/sda
```

![escanear las particiones de disco para detectar los cambios](https://www.arsys.es/blog/file/uploads/2019/09/4detectar-cambios-particiones.png)
Ahora usaremos LVM, primero creando el volumen físico.

```
pvcreate /dev/sda3
```

Luego añadimos este volumen al grupo de volúmenes, en este caso se llama *vg00*.

```
vgextend vg00 /dev/sda3
```

Redimensionamos el volumen lógico, *vg00-lv01* en este caso, para ocupar todo el espacio libre.

```
lvextend -l +100%FREE /dev/mapper/vg00-lv01
```

Acabamos redimensionando el sistema de ficheros del volumen lógico.

```
resize2fs /dev/mapper/vg00-lv01
```

Esta secuencia de comandos se puede ver resumida en la siguiente imagen:![redimensionando el sistema de ficheros del volumen lógico](https://www.arsys.es/blog/file/uploads/2019/09/4secuencia-final.png)Solo nos queda comprobar el estado de las particiones del disco, para ver cómo ha quedado nuestra configuración. Podemos hacerlo mediante el comando que ya conocemos:

```
fdisk -l
```

![comprobar el estado de las particiones del disco](https://www.arsys.es/blog/file/uploads/2019/09/4estadofinal.png)

## Redimensionar particiones en CentOS

Ahora vamos a repetir toda esta secuencia de comandos sobre una distribución CentOS. El proceso es muy similar al que hemos señalado para Ubuntu, con algunas pequeñas modificaciones. Lo más importante es fijarse en la salida de los comandos para saber cuáles son los discos que tenemos actualmente, sus nombres, particiones, etc., de modo que podamos personalizar las instrucciones a nuestro caso particular.

Como el proceso es prácticamente el mismo que para Ubuntu, vamos a resumir los pasos haciendo mayor hincapié en las diferencias que nos podemos encontrar entre uno y otro sistema.

El primer paso es comprobar el identificador del disco.

```
ls /sys/class/scsi_disk/
```

![comprobar el identificador del disco](https://www.arsys.es/blog/file/uploads/2019/09/5-identificadordiscos.png)
Como puedes ver en la imagen, en nuestro caso es *2:0:0:0*.

Hacemos un escaneo del disco, usando ese identificador:

```
echo 1 > /sys/class/scsi_device/2:0:0:0/device/rescan
```

Ejecutamos el análisis de particiones, del mismo modo que en Ubuntu, con *fdisk*.

```
fdisk -l
```

![análisis de particiones](https://www.arsys.es/blog/file/uploads/2019/09/5-analis-fdisk.png)
En la parte de arriba vemos que los discos ya tienen el tamaño que hemos asignado en el panel de control de Cloud Builder. Sin embargo, la partición principal todavía tiene el tamaño antiguo.

Creamos una partición con el espacio no particionado:

```
fdisk /dev/sda
```

Todo el proceso de creación de estas particiones es el mismo que el explicado más arriba para Ubuntu. En resumen, los comandos que se necesita ir indicando son los siguientes:

- ***n #*** Nueva partición.
- ***p #\*** Primaria.
- ***3 #*** Número de la partición.
- ***Enter #*** Primer sector, por defecto el primero libre.
- ***Enter #*** Último sector, por defecto el último del disco.
- ***w #*** Escribir cambios.

Ahora debemos lanzar de nuevo este proceso para marcar la partición como tipo Linux LVM.

```
fdisk /dev/sda
```

Introducimos las siguientes opciones de menú:

- ***t #*** Cambiar tipo de partición.
- ***3 #*** Número de partición.
- ***8e #*** Tipo LVM.
- ***w #*** Escribir cambios.

Para el paso de escanear las particiones de disco, con objetivo de detectar los cambios, CentOS ya viene con el programa Parted instalado, así que simplemente tenemos que hacer lo siguiente:

```
partprobe /dev/sda
```

Recuerda que, **si este comando no está disponible**, probablemente lo tengas que instalar usando Yum: *yum install parted*.

Ahora usaremos LVM, primero creando el volumen físico.

```
pvcreate /dev/sda3
```

Ahora añadimos este volumen al grupo de volúmenes. Para este caso te tienes que fijar especialmente en el nombre del volumen, que pudimos ver con *fdisk -l*. En nuestra máquina se llama *centos*, tal y como aparece marcado en la siguiente imagen.![añadimos este volumen al grupo de volúmenes](https://www.arsys.es/blog/file/uploads/2019/09/5-marca-centos.png)Así que escribimos:

```
vgextend centos /dev/sda3
```

![añadimos este volumen al grupo de volúmenes](https://www.arsys.es/blog/file/uploads/2019/09/5-centos-extended.png)
Estamos terminando. Pero antes debemos redimensionar el volumen lógico, *centos-root* en este caso, para ocupar todo el espacio libre.

```
lvextend -l +100%FREE /dev/mapper/centos-root
```

![redimensionar el volumen lógico](https://www.arsys.es/blog/file/uploads/2019/09/5-redimension-logico.png)

Redimensionamos el volumen lógico, *centos-root* en este caso, para ocupar el espacio libre.

```
xfs_growfs /dev/mapper/centos-root
```

Terminado el proceso, podemos comprobar el estado de las particiones del disco, usando de nuevo el comando *fdisk*.

```
fdisk -l
```

El resultado que encontraremos será similar al de la imagen siguiente, donde se puede comprobar que la partición principal tiene el tamaño del disco, una vez agregado el nuevo espacio asignado.

![comprobar que la partición principal tiene el tamaño del disco](https://www.arsys.es/blog/file/uploads/2019/09/5-fdisk-final.png)