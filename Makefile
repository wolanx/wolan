default:
	cat Makefile

build:
	go install github.com/zx5435/wolan/cmd/wolan-server
	go install github.com/zx5435/wolan/cmd/wolan-client
