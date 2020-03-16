docker cp ./cmd/wolan-ingress/wolan-ingress wolan-ingress:/usr/bin/wolan-ingress
docker cp ./tpl/ingress/rc wolan-ingress:/go/src/github.com/zx5435/wolan/tpl/ingress/rc

# use
## 1 
docker run -it -d --name wolan-ingress -p443:443 -p80:80 zx5435/wolan:ingress
docker run -it -d --name wolan-ingress -p443:443 -p80:80 -v "$PWD":/etc/nginx/conf.d zx5435/wolan:ingress
docker run -it -d --name wolan-ingress -p443:443 -p80:80 -v "$PWD":/usr/share/nginx/html zx5435/wolan:ingress
docker run -it -d --name wolan-ingress --net=host -p443:443 -p80:80 zx5435/wolan:ingress

## 2
docker exec -it wolan-ingress sh

## 3 
wolan-ingress -s new -d www.test.com
wolan-ingress -env=prod -s=new -d zx5435.com
