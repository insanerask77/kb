# Metricbeats in docker

```yaml
# docker-compose.yml this docker-compose needs be executed by sudo.
↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑
#############################-Rest of the code-#########################
# Metricbeats
  metricbeat:
    container_name: metricbeats_docker
    #image: eeacms/metricbeat
    image:  docker.elastic.co/beats/metricbeat:7.16.1
    # https://github.com/docker/swarmkit/issues/1951
    hostname: "MC0-metricbeat"  # ← This name will appear in kibana to diferenciate other hosts
    user: root
    #network_mode: "host"
    #configs:
    #  - source: mb_config
    #    target: /usr/share/metricbeat/metricbeat.yml
    volumes:
      - ./config/metricbeat.yml:/usr/share/metricbeat/metricbeat.yml
      - /proc:/hostfs/proc:ro
      - /sys/fs/cgroup:/hostfs/sys/fs/cgroup:ro
      - /:/hostfs:ro
      - /var/run/docker.sock:/var/run/docker.sock
      - metricbeat:/usr/share/metricbeat/data
    env_file:
      - .env_elastic
    command: ["--strict.perms=false", "-system.hostfs=/hostfs", "-e"]
    deploy:
      mode: global
#############################-Rest of the code-#########################      
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
```

##### Config File

```yaml
#./config/metricbeat.yml
metricbeat.config.modules:
  path: ${path.config}/modules.d/*.yml
  reload.enabled: false
#============================== System =========================================
metricbeat.modules:
- module: system
  period: 10s
  metricsets:
    - cpu
    - load
    - memory
    - network
    - process
    - process_summary
    - core
    - diskio
    - socket
  process.include_top_n:
    by_cpu: 5
    by_memory: 5
  enabled: true
#============================== Docker =========================================
- module: docker
  hosts: ["unix:///var/run/docker.sock"]
  period: 10s
  enabled: true
  metricsets:
    - container
    - cpu
    - diskio
    - healthcheck
    - info
    - memory
    - network

processors:
- add_cloud_metadata: ~
- add_host_metadata:
    netinfo.enabled: true
- add_docker_metadata:
    host: "unix:///var/run/docker.sock"
#============================== Elastic =========================================
output.elasticsearch:
  hosts: ["${ELASTICSEARCH_HOST}"]
  username: ${ELASTICSEARCH_USERNAME}
  password: ${ELASTICSEARCH_PASSWORD}

#============================== Kibana =========================================
setup.kibana:
  host: "${KIBANA_URL}"
  username: ${ELASTICSEARCH_USERNAME}
  password: ${ELASTICSEARCH_PASSWORD}
```

