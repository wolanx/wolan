wolan
======

我懒

## Intro
> 基于k8s原生微服务后，最重要的3模块`logging` `tracing` `metrics`，主流方案都略麻烦、较重。用最低的依赖、最少的步骤、最懒的方式，是这个项目唯一的宗旨。

## wolan-logging
k8s 部署 `filebeat`后，会采集 `/var/log/containers/*.log`，设置`annotations` `co.elastic.logs/enabled: "true"` `co.elastic.logs.main/enabled: "true"`可以指定需要采集的`pod/container`。`input`为固定模板详见x，`output`通常为es，这两步通常步骤较繁琐，也是要解决的第一个问题，

- 为什么不用`filebeat`，不支持 signal reload
- 为什么不用`logstash`，java太重

## wolan-tracing


## wolan-metrics
