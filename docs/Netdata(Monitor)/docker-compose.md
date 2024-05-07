```yaml
version: '3.3'
services:
    netdata:
        container_name: netdata
        ports:
            - '19999:19999'
        volumes:
            - 'netdataconfig:/etc/netdata'
            - 'netdatalib:/var/lib/netdata'
            - 'netdatacache:/var/cache/netdata'
            - '/etc/passwd:/host/etc/passwd:ro'
            - '/etc/group:/host/etc/group:ro'
            - '/proc:/host/proc:ro'
            - '/sys:/host/sys:ro'
            - '/etc/os-release:/host/etc/os-release:ro'
        restart: unless-stopped
        image: netdata/netdata
volumes:
	netdataconfig:
	netdatalib:
	netdatacache:
```

```yaml
version: '3.3'
services:
    netdata:
        container_name: netdata
        ports:
            - '19999:19999'
        volumes:
            - 'netdataconfig:/etc/netdata'
            - 'netdatalib:/var/lib/netdata'
            - 'netdatacache:/var/cache/netdata'
            - '/etc/passwd:/host/etc/passwd:ro'
            - '/etc/group:/host/etc/group:ro'
            - '/proc:/host/proc:ro'
            - '/sys:/host/sys:ro'
            - '/etc/os-release:/host/etc/os-release:ro'
        restart: unless-stopped
        image: netdata/netdata
volumes:
  netdataconfig:
  netdatalib:
  netdatacache:
```

```
docker run -d --name=netdata \
  --pid=host \
  --network=host \
  -v netdataconfig:/etc/netdata \
  -v netdatalib:/var/lib/netdata \
  -v netdatacache:/var/cache/netdata \
  -v /etc/passwd:/host/etc/passwd:ro \
  -v /etc/group:/host/etc/group:ro \
  -v /etc/localtime:/etc/localtime:ro \
  -v /proc:/host/proc:ro \
  -v /sys:/host/sys:ro \
  -v /etc/os-release:/host/etc/os-release:ro \
  -v /var/log:/host/var/log:ro \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --restart unless-stopped \
  --cap-add SYS_PTRACE \
  --cap-add SYS_ADMIN \
  --security-opt apparmor=unconfined \
  netdata/netdata
```

```
docker run -d --name=netdata \
  -p 19999:19999 \
  -v /etc/passwd:/host/etc/passwd:ro \
  -v /etc/group:/host/etc/group:ro \
  -v /proc:/host/proc:ro \
  -v /sys:/host/sys:ro \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  --cap-add SYS_PTRACE \
  --security-opt apparmor=unconfined \
  netdata/netdata
```


