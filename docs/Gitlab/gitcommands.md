```sh
git config --global user.name "Rafael Madolell Vargas"
git config --global user.email "rafaelm@bluetrailsoft.com"
```

##### Create a new repository

```sh
git clone git@gitlab.com:rafaelm2/canal-denuncias.git
cd canal-denuncias
git switch -c main
touch README.md
git add README.md
git commit -m "add README"
git push -u origin main
```

##### Push an existing folder

```sh
cd existing_folder
git init --initial-branch=main
git remote add origin git@gitlab.com:rafaelm2/canal-denuncias.git
git add .
git commit -m "Initial commit"
git push -u origin main
```

##### Push an existing Git repository

```sh
cd existing_repo
git remote rename origin old-origin
git remote add origin git@gitlab.com:rafaelm2/canal-denuncias.git
git push -u origin --all
git push -u origin --tags
```

**Steps** 

```sh
ssh-keygen -b 2048 -t rsa -C "email"
```

```sh
git clone "ssh token" #clonar repo
git status # estado del git
git add . # añadir archivos
git commit -m "Firts commit" # comparar archivos añadidos
git push origin <name> # subir archivos
git pull origin <name> # descargar archivos
git branch <name> # crear rama
git checkout <name> # escoger rama
Todos los dias
git checkout main 
git pull origin main
git checkout test
git merge main

git branch -D <name> # brorrar rama

```

