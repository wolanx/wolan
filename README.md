wolan
======
cicdcm

# start

```sh
mkdir wolan
cd wolan
docker run -it -d --name wolan -p 4321:23456 \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    -v "$PWD":/app/__work__ \
    zx5435/wolan
open http://localhost:4321
```

# 简单的自动部署工具，将具有

- [ ] manage
    - [x] task run
- [ ] ci
- [ ] cd
- [ ] gateway
- [ ] auto ssl
