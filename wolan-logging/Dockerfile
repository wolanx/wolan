FROM alpine

LABEL author=zx5435

ENV TZ utc-8

WORKDIR /usr/local/bin

COPY filebeat .

RUN ls -l
ENTRYPOINT ["filebeat"]

# docker build -f Dockerfile -t zx5435/wolan:logging .
# docker run --restart=unless-stopped --name wlog -d -p 20100:20100 zx5435/wolan:logging
