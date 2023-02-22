FROM golang:1.19-alpine3.17 AS builder

WORKDIR $GOPATH/src/tiktok-simple/TIKTOK_Gateway

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE=on

ADD go.mod .
COPY . .
RUN go mod download
RUN go build -o tiktok_gateway ./cmd/main.go
ENTRYPOINT  ["./tiktok_gateway"]