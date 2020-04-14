wolan
======

我懒

# start

demo http://zx5435.com:8080

```sh
mkdir wolan
cd wolan
docker run -it -d --name wolan -p 4321:23456 \
    -v "/var/run/docker.sock:/var/run/docker.sock" \
    -v "$PWD":/app/__work__ \
    zx5435/wolan
open http://localhost:4321
```

# include

- [x] page
    - [x] task run
- [x] ingress
    - [ ] auto https
- [ ] ci
- [ ] cd
- [ ] gateway
- [ ] auto ssl
