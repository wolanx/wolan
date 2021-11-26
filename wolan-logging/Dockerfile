FROM golang:latest AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /www

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o wolan-logging wolan-logging.go

FROM alpine:latest

LABEL author=zx5435

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates
ENV TZ utc-8

WORKDIR /www
COPY --from=builder /www/wolan-logging /www/wolan-logging

CMD ["./wolan-logging"]

# docker build -f Dockerfile -t zx5435/wolan:logging .
# docker run --restart=unless-stopped --name wlog -d -p 20100:20100 zx5435/wolan:logging
