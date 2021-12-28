curl svc-loki.wolan:3100
curl -H "Content-Type: application/json" -XPOST -s "http://svc-loki.wolan:3100/loki/api/v1/push" --data-raw \
'{"streams": [{ "stream": { "foo": "bar2" }, "values": [ [ "1640594506000000000", "fizzbuzz" ] ] }]}'

```json
{
  "streams": [
    {
      "stream": {
        "label": "value"
      },
      "values": [
          [ "<unix epoch in nanoseconds>", "<log line>" ],
          [ "<unix epoch in nanoseconds>", "<log line>" ]
      ]
    }
  ]
}
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


