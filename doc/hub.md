


docker run -d -p 5000:5000 -p 443:443 -v ~/registry:/var/lib/registry registry:2.6.2

docker tag zx5435/cdemo-php:a 127.0.0.1:5000/cdemo-php:a
docker tag zx5435/cdemo-php:a 192.168.199.115/cdemo-php:a

docker push 127.0.0.1:5000/cdemo-php:a
docker push 192.168.199.115/cdemo-php:a


docker run -it --rm certbot/certbot certbot-auto --help all
docker run -it -v "$PWD":/mk -w /mk --rm certbot/certbot certonly --email 82547762@qq.com -d www.825407762.com -d 825407762.com -d hub.825407762.com
