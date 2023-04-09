# gRPC Greeting
FROM golang:1.20.0-alpine3.17 AS base
RUN apk update && apk add git

# Current status there is no need to separate images with docker multi stage building.
# But there is possibility that adds some original thing to each image in the future.
# Hence defines these images as multi building.
# 
# Want to make multi building with built go file.
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

