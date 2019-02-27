
docker cp ./cmd/wolan-ingress/wolan-ingress wolan-ingress:/usr/bin/wolan-ingress
docker cp ./tpl/ingress/rc wolan-ingress:/go/src/github.com/zx5435/wolan/tpl/ingress/rc

wolan-ingress -s new -d www.test.com
wolan-ingress -env=prod -s=new -d zx5435.com

# use 
docker run -it -d --name wolan-ingress -p80:80 -p443:443 zx5435/wolan:ingress
docker exec -it wolan-ingress sh
