default:
	cat Makefile

build: build-fe build-be

build-be:
	docker run -it --rm -v "$$GOPATH/src":/go/src -w /go/src/github.com/zx5435/wolan/cmd/wolan-server golang:1.10.2 \
        go build -v -ldflags "-linkmode external -extldflags -static -w" -o wolan-server
	docker build -f __cicd__/Dockerfile -t zx5435/wolan .

build-fe:
	cd frontend && npm run build

restart: down up

up:
	cd __work__ && docker run -it -d --name wolan -p 4321:23456 \
	    -v "$$PWD":/go/src/github.com/zx5435/wolan/__work__ zx5435/wolan

down:
	docker stop wolan
	docker rm wolan
