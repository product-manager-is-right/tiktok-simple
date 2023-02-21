FROM golang:1.19-alpine3.17 AS builder

WORKDIR /build
ENV GOPROXY https://goproxy.cn
ENV GOPATH=/build
ENV GO111MODULE=on

ADD go.mod .
COPY . .
RUN go mod download
RUN go build -o tiktok_video ./cmd/main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/tiktok_video /build/tiktok_video
CMD ["./tiktok_video"]