## SSL Install Cert

##### Install

```bash
## Try 1
sudo snap install
## Try 2
sudo apt install snap
## Cert INstall
sudo snap install --classic certbot
## Cert register
sudo certbot certonly
```



##### Monitoring Calls

```bash
## Monit incoming 
# tcpdump filter for HTTP GET  
sudo tcpdump -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420'

# tcpdump filter for HTTP POST  
sudo tcpdump -s 0 -A 'tcp dst port 80 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504f5354)'

# tcpdump filter for POST PORT 
tcpdump -s 0 -A 'tcp dst port 3000 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504f5354)' 

# tcpdump filter for GET PORT 
tcpdump -s 0 -A 'tcp dst port 3003 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420)'
```

##### Kill Nginx Ports

```bash
kill -9 $(sudo lsof -i -P -n | grep nginx)
sudo killall nginx
sudo certbot renew --dry-run
```

#### Intall tcpdump and lsof

```bash
apt-get install tcpdump
apt-get install lsof
```

#### Trubleshoot

```bash
sudo killall nginx
sudo certbot -v -i nginx -d sub.example.com -d sub2.example.com certonly
sudo nginx -t
sudo nginx -s reload
```

