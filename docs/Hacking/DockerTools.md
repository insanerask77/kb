#### Bug Bounty Toolkit

```
docker pull warch/social-engineering-toolkit
```

```
docker pull hackersploit/bugbountytoolkit

docker run -it hackersploit/bugbountytoolkit /bin/bash

docker run -it hackersploit/bugbountytoolkit /bin/bash

docker run -it hackersploit/bugbountytoolkit /usr/bin/zsh
```

#### OWASP Juice Shop (lab)

```
docker run --rm -p 3000:3000 bkimminich/juice-shop
```



#### kasmweb/kali

```
docker pull kasmweb/kali

sudo docker run --rm  -it --shm-size=512m -p 6901:6901 -e VNC_PW=password kasmweb/kali
```

The container is now accessible via a browser : `https://<IP>:6901`

- **User** : `kasm_user`
- **Password**: `password`

#### Seeker

`https://github.com/thewhiteh4t/seeker.git`

Metasploit

```
sudo docker run --rm -it --name=metasploit metasploitframework/metasploit-framework:latest
```

