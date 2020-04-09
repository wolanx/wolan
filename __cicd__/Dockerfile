FROM node:12.16.1-alpine3.9 AS go-builder

WORKDIR /www

RUN npm config set registry https://registry.npm.taobao.org --global
RUN npm config set disturl https://npm.taobao.org/dist --global

COPY frontend/package.json .
COPY frontend/package-lock.json .
RUN ls -l
RUN npm i

ADD frontend .
RUN npm run build

FROM openjdk:8-alpine

WORKDIR /www

COPY --from=go-builder /www/dist /www/resources/static
COPY --from=go-builder /www/dist/index.html /www/resources/templates/site/index.html
COPY ./build/libs/wolan.jar /www/

RUN ls -l

EXPOSE 8080

CMD ["java", "-jar","wolan.jar"]
