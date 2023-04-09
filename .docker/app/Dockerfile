# gRPC Greeting
FROM golang:1.20.0-alpine3.17 AS base
RUN apk update && apk add git

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