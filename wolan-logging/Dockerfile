FROM alpine

LABEL author=github.com/zx5435
ENV TZ utc-8

WORKDIR /usr/share/filebeat

COPY filebeat .

ENTRYPOINT ["./filebeat"]

# docker build -f Dockerfile -t zx5435/wolan:logging .
# docker run --restart=unless-stopped --name wlog -d -p 20100:20100 zx5435/wolan:logging
