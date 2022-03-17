FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod .
RUN go mod download
RUN go get github.com/gin-gonic/gin@v1.7.7
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/spf13/viper
