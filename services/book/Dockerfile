FROM golang:1.17.5-bullseye AS builder

WORKDIR /go/src/github.com/taroooth/go-microservices-example

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

RUN go get -u google.golang.org/grpc

COPY . .
RUN go build -o /usr/bin/book ./services/book

CMD ["/usr/bin/book"]
