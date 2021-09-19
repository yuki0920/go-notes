FROM golang:1.16-alpine

RUN apk update
RUN apk add --no-cache git curl make gcc g++ \
  && go get -u github.com/cosmtrek/air \
  && go get -u github.com/pressly/goose/v3/cmd/goose

# go moduleを使用
ENV GO111MODULE=on

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /main
