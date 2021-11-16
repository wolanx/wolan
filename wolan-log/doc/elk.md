# doc
https://www.elastic.co/guide/en/beats/filebeat/current/configuration-autodiscover-hints.html#_co_elastic_logsenabled


# config
kubectl config set-context --current --namespace=$(basename $PWD)
k delete configmap filebeat-config
k create configmap filebeat-config --from-file=config/


# pod annotations
co.elastic.logs/enabled: "true"
co.elastic.logs.main/enabled: "true"



```yaml
#https://www.elastic.co/guide/en/beats/filebeat/current/console-output.html
output.console:
  pretty: true
```

```yaml
output.file:
  path: "/tmp/filebeat"
  filename: filebeat
  #rotate_every_kb: 10000
  #number_of_files: 7
  #permissions: 0600
```
