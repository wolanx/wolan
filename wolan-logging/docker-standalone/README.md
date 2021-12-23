# docker standalone

[官方文档](https://grafana.com/docs/loki/latest/clients/docker-driver/) 主要提供两种部署方式，

- 方式A：直接部署 promtail，缺点：只获取 log，没有 label
- 方式B：使用 docker plugin 部署
  - 每个容器 run 时 --log-driver=loki
  - **推荐** 修改 daemon 的默认 log driver           

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
