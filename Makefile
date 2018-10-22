default:
	cat Makefile

build: build-fe build-be build-pkg

build-be:
	docker run -it --rm \
	    -v "$$GOPATH/src":/go/src \
	    -w /go/src/github.com/zx5435/wolan/cmd/wolan-server \
	    golang:1.10.2 \
        go build -v -ldflags "-linkmode external -extldflags -static -w" -o wolan-server

build-fe:
	cd frontend && npm run build

build-pkg:
	docker build -f __cicd__/Dockerfile -t zx5435/wolan .

ingress-build:
	docker run -it --rm \
	    -v "$$GOPATH/src":/go/src \
	    -w /go/src/github.com/zx5435/wolan/cmd/wolan-ingress \
	    golang:1.10.2 \
        go build -v -ldflags "-linkmode external -extldflags -static -w" -o wolan-ingress

ingress-pkg:
	docker build -f __cicd__/Dockerfile.ingress -t zx5435/wolan:ingress .

ingress-test:
	docker run -it -d --name wolan-ingress -p2323:80 zx5435/wolan:ingress

restart: down up

up:
	cd __work__ && docker run -it -d --name wolan -p 4321:23456 \
	    -v "$$PWD":/app/__work__ \
	    -v "/var/run/docker.sock:/var/run/docker.sock" \
	    zx5435/wolan

down:
	docker stop wolan
	docker rm wolan
