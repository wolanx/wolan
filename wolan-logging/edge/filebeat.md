```yaml
#filebeat.inputs:
#- type: container
#  paths:
#    - /var/log/containers/*.log
#  processors:
#    - add_kubernetes_metadata:
#        in_cluster: true
#        include_annotations:
#          - kubectl.kubernetes.io/default-logs-container
#        host: ${NODE_NAME}
#        matchers:
#        - logs_path:
#            logs_path: "/var/log/containers/"


filebeat.autodiscover:
  providers:
    - type: kubernetes
      node: ${NODE_NAME}
      hints.enabled: true
      hints.default_config:
        enabled: false
        type: container
        paths:
          - /var/log/containers/*${data.kubernetes.container.id}.log
#      templates:
#        - condition:
#            equals:
#              kubernetes.namespace: ccm-prod
#          config:
#            - type: container
#              paths:
#                - /var/log/containers/*-${data.kubernetes.container.id}.log
#              exclude_lines: [ "^\\s+[\\-`('.|_]" ]

#    - type: kubernetes
#      node: ${NODE_NAME}
#      hints.enabled: true
#      hints.default_config:
#        type: container
#        paths:
#          - /var/log/containers/*${data.kubernetes.container.id}.log


processors:
  - drop_fields:
      fields: [ "host", "tags", "ecs", "log", "prospector", "agent", "input", "beat", "offset", "kubernetes.node.labels" ]
      ignore_missing: true


#processors:
#  - add_cloud_metadata:
#  - add_host_metadata:
#
#cloud.id: ${ELASTIC_CLOUD_ID}
#cloud.auth: ${ELASTIC_CLOUD_AUTH}

output.elasticsearch:
  hosts: [ '${ELASTICSEARCH_HOST:elasticsearch}:${ELASTICSEARCH_PORT:9200}' ]
  username: ${ELASTICSEARCH_USERNAME}
  password: ${ELASTICSEARCH_PASSWORD}
#setup.kibana:
#  host: kibana:5601

```

```yaml
filebeat.autodiscover:
  providers:
    - type: kubernetes
      node: ${NODE_NAME}
      hints.enabled: true
      hints.default_config:
        enabled: false
        type: container
        paths:
          - /var/log/containers/*${data.kubernetes.container.id}.log
output.console:
  pretty: true
```

```yaml
filebeat.autodiscover:
  providers:
    - type: kubernetes
      templates:
        - condition:
            equals:
              kubernetes.namespace: ccm-perf
          config:
            - type: container
              paths:
                - /var/log/containers/*-${data.kubernetes.container.id}.log
              exclude_lines: [ "^\\s+[\\-`('.|_]" ]  # drop asciiart lines
processors:
  - drop_fields:
      fields: [ "host", "tags", "ecs", "log", "prospector", "agent", "input", "beat", "offset", "kubernetes.node.labels" ]
      ignore_missing: true
output.console:
  pretty: true
```
