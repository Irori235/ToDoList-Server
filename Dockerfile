# ベースとなるDockerイメージ指定
FROM golang:1.17-alpine as server-build

RUN mkdir ./ToDoList-server
WORKDIR  /go/src/ToDoList-Server
COPY . .
RUN apk upgrade --update && \
    apk --no-cache add git