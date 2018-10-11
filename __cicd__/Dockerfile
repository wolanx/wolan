FROM wolan.rt AS go-builder

COPY . /go/src/github.com/zx5435/wolan

WORKDIR /go/src/github.com/zx5435/wolan/cmd/wolan-server

RUN go build -v -ldflags "-linkmode external -extldflags -static -w" -o wolan

FROM alpine:3.8

COPY . /go/src/github.com/zx5435/wolan

COPY --from=go-builder /go/src/github.com/zx5435/wolan/cmd/wolan-server/wolan-server /go/src/github.com/zx5435/wolan/

WORKDIR /go/src/github.com/zx5435/wolan

EXPOSE 23456

CMD ["./wolan-server"]
