### Bind 80 port in pm2 

```
sudo setcap CAP_NET_BIND_SERVICE=+eip `which node`
```

