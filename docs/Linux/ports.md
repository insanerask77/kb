

\# Used Ports 

```sh
sudo lsof -i -P -n | grep LISTEN sudo ss -tnlp
```

\## Monit incoming

\# tcpdump filter for HTTP GET 

```sh
sudo tcpdump -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420' 
```

\# tcpdump filter for HTTP POST  

```sh
sudo tcpdump -s 0 -A 'tcp dst port 80 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504f5354)'
```

\# tcpdump filter for POST PORT 

```sh
tcpdump -s 0 -A 'tcp dst port 3000 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504f5354)'
```

\# tcpdump filter for GET PORT 

```sh
tcpdump -s 0 -A 'tcp dst port 3003 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420)'
```

```sh
kill -9 $(sudo lsof -i -P -n | grep nginx)
```

```
sudo netstat -tupln
```

```
sudo lsof -i -P -n
```

```
python -m SimpleHTTPServer 8000 &> /dev/null & pid=$!

```

