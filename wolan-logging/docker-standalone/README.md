# docker standalone

[官方文档](https://grafana.com/docs/loki/latest/clients/docker-driver/) 主要提供两种部署方式，

- 方式B：每个容器需要配置 --log-driver=loki
- 方式B+：使用 docker plugin 部署

## 0x01 docker setting

注意修改 loki-url ip

```shell
cat > /etc/docker/daemon.json << EOF
{
  "debug": true,
  "log-driver": "loki",
  "log-opts": {
    "loki-url": "http://192.168.2.238:3100/loki/api/v1/push"
  }
}
EOF
```

## 0x02 docker plugin install

```shell
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

#docker plugin disable loki --force
#docker plugin upgrade loki grafana/loki-docker-driver:latest --grant-all-permissions
#docker plugin enable loki

#docker plugin disable loki --force
#docker plugin rm loki
```

## 0x03 restart

```shell
systemctl restart docker
docker ps -a -q | xargs docker restart
```
