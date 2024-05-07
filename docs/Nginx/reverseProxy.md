```
server {
listen 443 ssl;

server_name preprod.myluna.tk;
# Edit the above _YOUR-DOMAIN_ to your domain name

ssl_certificate /etc/letsencrypt/live/preprod..myluna.tk/fullchain.pem;
# If you use Lets Encrypt, you should just need to change the domain.
# Otherwise, change this to the path to full path to your domains public certificate file.

ssl_certificate_key /etc/letsencrypt/live/preprod.myluna.tk/privkey.pem;
# If you use Let's Encrypt, you should just need to change the domain.
# Otherwise, change this to the direct path to your domains private key certificate file.

ssl_session_cache builtin:1000 shared:SSL:10m;
# Defining option to share SSL Connection with Passed Proxy

ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
# Defining used protocol versions.

ssl_ciphers HIGH:!aNULL:!eNULL:!EXPORT:!CAMELLIA:!DES:!MD5:!PSK:!RC4;
# Defining ciphers to use.

ssl_prefer_server_ciphers on;
# Enabling ciphers

access_log /var/log/nginx/access.log;
# Log Location. the Nginx User must have R/W permissions. Usually by ownership.

location / {
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_pass http://localhost:3000;
#proxy_pass unix:/path/to/php7.3.sock # This is an example of how to define a unix socket.
proxy_read_timeout 90;
}

}
```

