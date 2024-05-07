# Elastalert

```yacas
docker-compose.yml
elastalert
    ├── elastalert.yaml
    └── rules
        ├── slack_error.yaml
        └── slack_info.yaml
```

```yaml
#docker-compose.yml
  elastalert:
      container_name: elastalert
      restart: always
      volumes:
        - './elastalert/elastalert.yaml:/opt/elastalert/config.yaml'
        - './elastalert/rules:/opt/elastalert/rules'
      image: jertel/elastalert2
      ports:
        - 3030:3030
```

```yaml
#elastalert.yaml
rules_folder: /opt/elastalert/rules

run_every:
  seconds: 10

buffer_time:
  minutes: 15

es_host: elasticsearch
es_port: 9200
es_username: alg_kibana
es_password: cartier419

writeback_index: elastalert_status

alert_time_limit:
  days: 2
```

```yaml
#slack_error.yaml
name: info error
type: frequency
index: logs*
num_events: 3
timeframe:
  hours: 1
filter:
- terms:
  #  host.name: "b6f798466e9c"
    log.file.path.keyword: ["/data/logs/be/error.log", "/data/error.log"]
alert:
- "slack"
slack:
slack_webhook_url: 'https://hooks.slack.com/services/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'
slack_emoji_override: ":warning:"
slack_msg_color: "danger"
```

```yaml
#slack_info.yml
name: info warning
type: frequency
index: logs*
num_events: 3
timeframe:
  hours: 1
filter:
- terms:
    #host.name: "b6f798466e9c"
    log.file.path.keyword: ["/data/logs/be/info.log", "/data/info.log"]
alert:
- "slack"
slack:
slack_webhook_url: 'https://hooks.slack.com/services/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'
slack_emoji_override: ":warning:"
slack_msg_color: "warning"
```

