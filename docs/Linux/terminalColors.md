```
#!/bin/bash
# Definimos colores, primer plano
PNegro='\033[30m'
PRojo='\033[31m'
PVerde='\033[32m'
PNaranja='\033[33m'
PAzul='\033[34m'
PMagenta='\033[35m'
PCian='\033[36m'
PGris='\033[37m'
NC='\033[39m'
# Colores de fondo, video inverso o segundo plano
IGris='\033[100m'		# Gris oscuro
IRojo='\033[101m'		# Luz roja
IVerde='\033[102m'		# Verde claro
IAmarillo='\033[103m'		# Amarillo
IAzul='\033[104m'		# Azul claro
IMagenta='\033[105m'		# Morado claro
ICian='\033[106m'		# Verde azulado
IBlanco='\033[107m'		# Blanco
INC='\033[40m'
# -----------
printf "${PRojo}${IAmarillo}"
echo "Hola Caracola"
printf "${PNegro}${IVerde}"
echo "Cada Texto"
printf "${PAzul}${ICian}"
echo "En un color"
printf "${NC}${INC}"
echo "Y vuelta a la normalidad"
```

