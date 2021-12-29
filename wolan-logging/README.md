curl svc-loki.wolan:3100 curl -H "Content-Type: application/json" -XPOST
-s "http://svc-loki.wolan:3100/loki/api/v1/push" --data-raw \
'{"streams": [{ "stream": { "foo": "bar2" }, "values": [ [ "1640594506000000000", "fizzbuzz" ] ] }]}'


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

```json
{
  "@timestamp": "2021-12-29T07:21:26.782Z",
  "@metadata": {
    "beat": "",
    "type": "_doc",
    "version": "8.1.0"
  },
  "log": {
    "offset": 169,
    "file": {
      "path": "/var/log/1.log"
    }
  },
  "message": "asd",
  "input": {
    "type": "log"
  },
  "ecs": {
    "version": "8.0.0"
  },
  "host": {
    "name": "iZuf6eciw1ef4tvi84chqdZ"
  },
  "agent": {
    "version": "8.1.0",
    "ephemeral_id": "63a40da7-500d-4045-a11d-6b75bd795c45",
    "id": "01f03c5c-eb20-4176-8d02-e0c1a223794a",
    "name": "iZuf6eciw1ef4tvi84chqdZ",
    "type": "filebeat"
  }
}
```

```json
{
  "Timestamp": "2021-12-29T17:05:21.137550585+08:00",
  "Meta": null,
  "Fields": {
    "agent": {
      "ephemeral_id": "4b0326a5-d85d-41b3-8567-eb046d3d2871",
      "id": "1fe8141a-94db-4dfc-abbf-52b2a29ab93e",
      "name": "iZuf6eciw1ef4tvi84chqdZ",
      "type": "filebeat",
      "version": "8.1.0"
    },
    "ecs": {
      "version": "8.0.0"
    },
    "host": {
      "name": "iZuf6eciw1ef4tvi84chqdZ"
    },
    "input": {
      "type": "log"
    },
    "log": {
      "file": {
        "path": "/var/log/1.log"
      },
      "offset": 173
    },
    "message": "asd"
  },
  "Private": {
    "id": "native::926802-64769",
    "prev_id": "",
    "source": "/var/log/1.log",
    "offset": 177,
    "timestamp": "2021-12-29T17:05:21.136902467+08:00",
    "ttl": -1,
    "type": "log",
    "meta": null,
    "FileStateOS": {
      "inode": 926802,
      "device": 64769
    },
    "identifier_name": "native"
  },
  "TimeSeries": false
}
```

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