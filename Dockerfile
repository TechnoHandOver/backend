FROM golang:latest AS build

MAINTAINER Nikita Lobaev

RUN mkdir /go/src/backend

COPY . /go/src/backend

WORKDIR /go/src/backend/app/main

RUN go build -o backend .

WORKDIR /go/src/backend

CMD ./app/main/backend configFileName=config.json
