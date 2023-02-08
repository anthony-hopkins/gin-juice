FROM golang:alpine3.17

RUN apk update && apk upgrade --no-cache
COPY ./<APP> /<APP>
WORKDIR /<APP>
