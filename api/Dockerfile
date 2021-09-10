FROM golang:1.16

# go moduleを使用
ENV GO111MODULE=on

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/pressly/goose/v3/cmd/goose
