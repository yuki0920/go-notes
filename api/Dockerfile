FROM golang:1.18-alpine

RUN apk update
RUN apk add --no-cache git curl make gcc g++
RUN go install github.com/cosmtrek/air@v1
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/golang/mock/mockgen@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/xo/xo@latest
ENV GO111MODULE=on

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /main
