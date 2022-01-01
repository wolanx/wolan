default:
	cat Makefile

build:
	CGO_ENABLED=0 go build -o filebeat
	docker build -f Dockerfile -t zx5435/wolan:ff .
	docker push zx5435/wolan:ff
