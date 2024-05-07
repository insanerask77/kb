```
server {
        add_header 'Access-Control-Allow-Origin' '*' always;
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Allow-Headers' 'Content-Type,Accept';
        add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS, PUT, DELETE';
        #root /home/gitlab-runner/OLD/bts_internship_2019_fe_app/dist/bts-admin;
        root /home/gitlab-runner/UAT/nginx/;
        index index.html index.htm index.nginx-debian.html;
        server_name mybts-uat.bluetrail.software;

        location / {
                try_files $uri $uri/ /index.html;
        }

        location /api {
                proxy_pass http://localhost:3002;
        }

        location /swagger {
                proxy_pass http://localhost:3002;
        }

        location = /robots.txt {
                add_header Content-Type text/plain;
                return 200 "User-agent: *\nDisallow: /\n";
        }

        location /legacy {
                proxy_pass http://mybts-legacy.bluetrail.software;
        }




        listen 443 ssl;
        listen [::]:443 ssl;
        ssl_certificate_key /etc/letsencrypt/live/mybts-uat.bluetrail.software/privkey.pem;
        ssl_certificate /etc/letsencrypt/live/mybts-uat.bluetrail.software/fullchain.pem;
        # managed by Certbot
}

server {
    if ($host = mybts-uat.bluetrail.software) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


        listen 80;
        listen [::]:80;
        server_name mybts-uat.bluetrail.software;
    return 404; # managed by Certbot


}

```

