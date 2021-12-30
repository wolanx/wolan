```shell
curl svc-loki.wolan:3100
curl -H "Content-Type: application/json" -XPOST "http://svc-loki.wolan:3100/loki/api/v1/push" -d \
'{"streams": [{ "stream": { "foo": "bar2" }, "values": [ [ "1640594506000000000", "fizzbuzz" ] ] }]}'
curl -H "Content-Type: application/json" -XPOST "http://192.168.2.238:3100/loki/api/v1/push" -d \
'{"streams": [{ "stream": { "foo.a": "bar2" }, "values": [ [ "1640594506000000000", "fizzbuzz" ] ] }]}'
```

output demo
https://github1s.com/raboof/beats-output-http/blob/HEAD/README.md

v6 demo
https://www.fullstory.com/blog/writing-a-filebeat-output-plugin/

GOOS=linux go build -trimpath -buildmode=plugin

# build

```shell
git clone https://github.com.cnpmjs.org/elastic/beats.git --depth 1
vi /www/beats/libbeat/outputs/console/loki.go
cd /www/beats/filebeat
CGO_ENABLED=0 make
./filebeat -c 2.yml -e
```

```json
{
  "streams": [
    {
      "stream": {
        "label": "value"
      },
      "values": [
        [
          "<unix epoch in nanoseconds>",
          "<log line>"
        ],
        [
          "<unix epoch in nanoseconds>",
          "<log line>"
        ]
      ]
    }
  ]
}
```

## label ecs

```json
{
  "agent.ephemeral_id": "7e017ea8-007e-43e0-9b4c-bb503b771755",
  "agent.id": "9a267d76-8b87-4159-9aaf-d535a2a6b7ab",
  "agent.name": "iZuf6eciw1ef4tvi84chqdZ",
  "agent.type": "filebeat",
  "agent.version": "8.1.0",
  "ecs.version": "8.0.0",
  "host.name": "iZuf6eciw1ef4tvi84chqdZ",
  "input.type": "log",
  "log.file.path": "/var/log/1.log",
  "log.offset": 4,
  "message": "qwe"
}
```

## label k8s all

```json
{
  "agent.ephemeral_id": "a59df637-40cb-47d3-9cd8-f36e11633fb1",
  "agent.id": "bc07c01c-8be6-4185-af6d-a1c873845bc2",
  "agent.name": "iZuf6j6rl141pu4p2c4hcwZ",
  "agent.type": "filebeat",
  "agent.version": "8.1.0",
  "container.id": "a1dc10a805215c0695023a3862a15a2dea302b0843fe3ee4449adf7abf5f9a6b",
  "container.image.name": "registry.cn-shanghai.aliyuncs.com/digital-web/gimc-rt:20211203-1148",
  "container.runtime": "docker",
  "ecs.version": "8.0.0",
  "host.name": "iZuf6j6rl141pu4p2c4hcwZ",
  "input.type": "container",
  "kubernetes.container.name": "main",
  "kubernetes.labels.controller-uid": "0b3da476-d0d6-4f97-a3be-d192603a9f05",
  "kubernetes.labels.job-name": "job-device-state-spec-flush-1640833560",
  "kubernetes.namespace": "ccm-perf",
  "kubernetes.namespace_uid": "2f32296b-af84-44eb-aace-0d75c5540312",
  "kubernetes.node.hostname": "cn-shanghai.10.10.18.161",
  "kubernetes.node.labels.alibabacloud_com/nodepool-id": "np0c8c25c63aac44c89d7495170fac70e3",
  "kubernetes.node.labels.beta_kubernetes_io/arch": "amd64",
  "kubernetes.node.labels.beta_kubernetes_io/instance-type": "ecs.g5.xlarge",
  "kubernetes.node.labels.beta_kubernetes_io/os": "linux",
  "kubernetes.node.labels.failure-domain_beta_kubernetes_io/region": "cn-shanghai",
  "kubernetes.node.labels.failure-domain_beta_kubernetes_io/zone": "cn-shanghai-g",
  "kubernetes.node.labels.kubernetes_io/arch": "amd64",
  "kubernetes.node.labels.kubernetes_io/hostname": "cn-shanghai.10.10.18.161",
  "kubernetes.node.labels.kubernetes_io/os": "linux",
  "kubernetes.node.labels.node_kubernetes_io/instance-type": "ecs.g5.xlarge",
  "kubernetes.node.labels.topology_diskplugin_csi_alibabacloud_com/zone": "cn-shanghai-g",
  "kubernetes.node.labels.topology_kubernetes_io/region": "cn-shanghai",
  "kubernetes.node.labels.topology_kubernetes_io/zone": "cn-shanghai-g",
  "kubernetes.node.name": "cn-shanghai.10.10.18.161",
  "kubernetes.node.uid": "84611052-1ca6-4f6c-b22b-d73a8cfc69ee",
  "kubernetes.pod.ip": "172.20.0.82",
  "kubernetes.pod.name": "job-device-state-spec-flush-1640833560-bdq25",
  "kubernetes.pod.uid": "7b86564f-3f3c-4302-90ed-86de938dfaca",
  "log.file.path": "/var/log/containers/job-device-state-spec-flush-1640833560-bdq25_ccm-perf_main-a1dc10a805215c0695023a3862a15a2dea302b0843fe3ee4449adf7abf5f9a6b.log",
  "log.offset": 21154,
  "stream": "stderr"
}
```

## label k8s useful

```json
{
  "container.image.name": "registry.cn-shanghai.aliyuncs.com/digital-web/gimc-rt:20211203-1148",
  "host.name": "iZuf6j6rl141pu4p2c4hcwZ",
  "kubernetes.container.name": "main",
  "kubernetes.namespace": "ccm-perf",
  "kubernetes.node.name": "cn-shanghai.10.10.18.161",
  "kubernetes.pod.ip": "172.20.0.82",
  "kubernetes.pod.name": "job-device-state-spec-flush-1640833560-bdq25",
  "stream": "stderr"
}
```
