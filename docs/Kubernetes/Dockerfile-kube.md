```dockerfile
# BUILD 
FROM node:12.16-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
ENV NODE_PATH=src/
ENV PORT=443
COPY package.json ./
COPY package-lock.json ./
RUN npm install -qy
# RUN npm install react-scripts@3.4.1 -g -qy
COPY . ./
# Change env to production
COPY ./src/config/env.prod.js ./src/config/env.js
ENV ENVIRONMENT=production
RUN npm run build

# NGINX
FROM nginx:stable-alpine

# Permissions
ENV USR=nginx
ENV PRODUCT_DIR="/var/www"
ARG NGINX_RUN_DIR="/var/run"
ENV NGINX_RUN_DIR $NGINX_RUN_DIR
ARG NGINX_CACHE_DIR="/var/cache/nginx"
ENV NGINX_CACHE_DIR $NGINX_CACHE_DIR
ARG NGINX_LOG_DIR="/var/log/nginx"
ENV NGINX_LOG_DIR $NGINX_LOG_DIR

# Clear Nginx
RUN rm -rf /var/www/*
RUN rm -rf /etc/nginx/conf.d/*

# Copy files
COPY --from=build /app/build $PRODUCT_DIR
COPY nginx.conf /etc/nginx/nginx.conf

RUN chgrp -R 0 $NGINX_RUN_DIR && chmod -R g=u $NGINX_RUN_DIR
RUN chgrp -R 0 $NGINX_CACHE_DIR && chmod -R g=u $NGINX_CACHE_DIR
RUN chgrp -R 0 $NGINX_LOG_DIR && chmod -R g=u $NGINX_LOG_DIR

RUN touch /var/run/nginx.pid && chmod -R g=u /var/run/nginx.pid


## users are not allowed to listen on priviliged ports
RUN touch /etc/nginx/conf.d/default.conf
RUN sed -i.bak 's/listen\(.*\)80;/listen 8080;/' /etc/nginx/conf.d/default.conf

# Expose Port
EXPOSE 8080

# comment user directive as master process is run as user in OpenShift anyhow
RUN sed -i.bak 's/^user/#user/' /etc/nginx/nginx.conf

# Get Datetime to Test if image is updated
RUN date >> /tmp/date.txt

# Using the User
RUN addgroup nginx root
USER nginx

# TEST Nginx
RUN nginx -t

# Start Nginx
CMD ["/bin/sh", "-c", "nginx -g \"daemon off;\""]
```

