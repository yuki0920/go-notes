FROM golang:1.16-alpine

RUN apk update
RUN apk add --no-cache git curl gcc g++

# go moduleを使用
ENV GO111MODULE=on

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /main

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/pressly/goose/v3/cmd/goose
