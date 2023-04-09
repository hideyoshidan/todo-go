# gRPC Greeting
FROM golang:1.20.0-alpine3.17 AS base
RUN apk update && apk add git

# 現状分ける必要性はないが、将来的にeach containerにoriginalな何かを加えるかもしれないので
# マルチステージビルドを定義しておく。
# 開発が完成したら、go buildしたマルチビルドも作りたい
###################################
# gPRC greeting
###################################
FROM base AS greeting-server
WORKDIR /go/src/app

COPY ../../src /go/src/app
RUN go mod tidy

###################################
# gPRC todo
###################################
FROM base AS todo-server
WORKDIR /go/src/app

COPY ../../src /go/src/app
RUN go mod tidy

###################################
# gPRC appmixer
###################################
FROM base AS appmixer-server
WORKDIR /go/src/app

COPY ../../src /go/src/app
RUN go mod tidy

###################################
# gPRC user
###################################
FROM base AS user-server
WORKDIR /go/src/app

COPY ../../src /go/src/app
RUN go mod tidy

