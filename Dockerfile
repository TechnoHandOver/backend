FROM golang:latest

MAINTAINER Nikita Lobaev

RUN mkdir /go/src/backend

COPY . /go/src/backend

WORKDIR /go/src/backend/app/main

RUN go build -o backend .

WORKDIR /go/src/backend
