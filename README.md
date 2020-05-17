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

task
 - [x] load 2020-05-17
 - [ ] config


git
 - [x] clone 2020-05-17
 - [ ] pull update


pipeline in Docker
 - [ ] load setting from
   - [ ] self
   - [ ] code/xx
 - [ ] job


k8s
 - [x] load ns pod
 - [x] apply


ingress
 - [ ] https


other
 - [ ] image register
