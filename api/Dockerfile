FROM golang:1.17-alpine

RUN apk update
RUN apk add --no-cache git curl make gcc g++
RUN go install github.com/cosmtrek/air@v1
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest


# go moduleを使用
ENV GO111MODULE=on

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /main
