# build a tiny docker image
FROM alpine:3.20.0

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]