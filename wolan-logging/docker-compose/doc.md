
https://grafana.com/docs/loki/latest/clients/docker-driver/

docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

docker plugin disable loki --force
docker plugin upgrade loki grafana/loki-docker-driver:latest --grant-all-permissions
docker plugin enable loki
systemctl restart docker

docker plugin disable loki --force
docker plugin rm loki


