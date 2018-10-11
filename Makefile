default:
	cat Makefile

asdad:
	go install github.com/zx5435/wolan/cmd/wolan-server
	go install github.com/zx5435/wolan/cmd/wolan-client

build:
	docker build -f __cicd__/Dockerfile.rt -t wolan.rt .
	docker build -f __cicd__/Dockerfile -t wolan .
