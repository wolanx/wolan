FROM alpine:3.8

WORKDIR /app

COPY ./cmd/wolan-server/wolan-server /app/cmd/wolan-server/wolan-server
COPY ./frontend/build /app/frontend/build

RUN ls -l

EXPOSE 23456

CMD ["./cmd/wolan-server/wolan-server"]
