FROM golang:1.17.1-alpine3.14

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get github.com/gin-gonic/gin \
    github.com/go-sql-driver/mysql \
    github.com/jinzhu/gorm