```
version: "3"
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    #ports:
      # These ports are in format <host-port>:<container-port>
      #- '80:80' # Public HTTP Port
      #- '443:443' # Public HTTPS Port
      #- '81:81' # Admin Web Port
      # Add any other Stream port you want to expose
      # - '21:21' # FTP
    environment:
      DB_MYSQL_HOST: "localhost"
      DB_MYSQL_PORT: 3306
      DB_MYSQL_USER: "npm"
      DB_MYSQL_PASSWORD: "npm"
      DB_MYSQL_NAME: "npm"
      # Uncomment this if IPv6 is not enabled on your host
      # DISABLE_IPV6: 'true'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
      - ./_hsts.conf:/app/templates/_hsts.conf:ro
    depends_on:
      - db
    network_mode: host
    #networks:
    #  - caddy

  db:
    image: 'jc21/mariadb-aria:latest'
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: 'npm'
      MYSQL_DATABASE: 'npm'
      MYSQL_USER: 'npm'
      MYSQL_PASSWORD: 'npm'
    volumes:
      - ./data/mysql:/var/lib/mysql
    network_mode: host
    #networks:
    #  - caddy
#networks:
#  caddy:
#    external: true
```

`_hsts.conf`

```

{% if certificate and certificate_id > 0 -%}
{% if ssl_forced == 1 or ssl_forced == true %}
{% if hsts_enabled == 1 or hsts_enabled == true %}
  # HSTS (ngx_http_headers_module is required) (63072000 seconds = 2 years)
  add_header Strict-Transport-Security "max-age=63072000;{% if hsts_subdomains == 1 or hsts_subdomains == true -%} includeSubDomains;{% endif %} preload" always;
  add_header Referrer-Policy strict-origin-when-cross-origin; 
  add_header X-Content-Type-Options nosniff;
  add_header X-XSS-Protection "1; mode=block";
  add_header X-Frame-Options SAMEORIGIN;
  add_header Content-Security-Policy upgrade-insecure-requests;
  add_header Permissions-Policy interest-cohort=();
  add_header Expect-CT 'enforce; max-age=604800';
  more_set_headers 'Server: Proxy';
  more_clear_headers 'X-Powered-By';
{% endif %}
{% endif %}
{% endif %}

```

```
{% if certificate and certificate_id > 0 -%}
{% if ssl_forced == 1 or ssl_forced == true %}
{% if hsts_enabled == 1 or hsts_enabled == true %}
  # HSTS (ngx_http_headers_module is required) (63072000 seconds = 2 years)
  add_header Strict-Transport-Security "max-age=15552000;{% if hsts_subdomains == 1 or hsts_subdomains == true -%} includeSubDomains;{% endif %} preload" always;
  add_header Referrer-Policy strict-origin-when-cross-origin;
  add_header X-Content-Type-Options nosniff;
  add_header Content-Security-Policy upgrade-insecure-requests;
  add_header Permissions-Policy interest-cohort=();
  add_header Expect-CT 'enforce; max-age=604800';
  more_set_headers 'Server: Proxy';
  more_clear_headers 'X-Powered-By';
{% endif %}
{% endif %}
{% endif %}
```



```
# /etc/nginx/nginx.conf
server_tokens off;
proxy_pass_header 'lukaku';
etag off;
add_header Last-Modified '';
if_modified_since off;
sendfile        on;
max_ranges 0;

# /etc/nginx/conf.d/default.conf
add_header Referrer-Policy strict-origin-when-cross-origin; 
add_header X-Content-Type-Options nosniff;
add_header X-XSS-Protection "1; mode=block";
add_header X-Frame-Options SAMEORIGIN;
add_header Content-Security-Policy upgrade-insecure-requests;
add_header Permissions-Policy interest-cohort=();
add_header Expect-CT 'enforce; max-age=604800';
```

```
So I made so changes :

First here is my _hsts.conf file :

{% if certificate and certificate_id > 0 -%}
{% if ssl_forced == 1 or ssl_forced == true %}
{% if hsts_enabled == 1 or hsts_enabled == true %}
add_header Strict-Transport-Security $hdr_strict_transport_security;
add_header Referrer-Policy $hdr_referrer_policy;
add_header X-Content-Type-Options $hdr_x_content_type_options;
add_header X-XSS-Protection $hdr_x_xss_protection;
add_header X-Frame-Options $hdr_x_frame_options;
add_header Content-Security-Policy $hdr_content_security_policy;
add_header Permissions-Policy $hdr_permissions_policy;
add_header Expect-CT $hdr_expect_ct;
more_set_headers 'Server: Proxy';
more_clear_headers 'X-Powered-By';
{% endif %}
{% endif %}
{% endif %}

Then I created (or modified) in /data/nginx/custom a http_top.conf file :

map $upstream_http_strict_transport_security $hdr_strict_transport_security {
'' "max-age=63072000; includeSubDomains; preload";
}

map $upstream_http_referrer_policy $hdr_referrer_policy {
'' "strict-origin-when-cross-origin";
}

map $upstream_http_x_content_type_options $hdr_x_content_type_options {
'' "nosniff";
}

map $upstream_http_x_xss_protection $hdr_x_xss_protection {
'' "1; mode=block";
}

map $upstream_http_x_frame_options $hdr_x_frame_options {
'' "SAMEORIGIN";
}

map $upstream_http_content_security_policy $hdr_content_security_policy {
'' "upgrade-insecure-requests";
}

map $upstream_http_permissions_policy $hdr_permissions_policy {
'' "interest-cohort=()";
}

map $upstream_http_expect_ct $hdr_expect_ct {
'' "enforce, max-age=604800";
}

Than restarting nginx and disabling SSL and saving in NPM, and enabling it again.
With this configuration nginx only adds those header if they are not already send by the underling application.

The only thing is if you want to disable on of those header (X-frame for me) on a specified application, you have to put the entire configuration for this application in the advanced tap with "location/{ ....}
```

