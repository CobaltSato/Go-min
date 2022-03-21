FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod .

EXPOSE 9090
