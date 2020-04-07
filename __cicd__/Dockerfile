FROM openjdk:8-alpine

WORKDIR /app

COPY ./build/libs/wolan.jar /app/

RUN ls -l

EXPOSE 8080

CMD ["java", "-jar","wolan.jar"]
