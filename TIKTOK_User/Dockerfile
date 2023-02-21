FROM golang:1.19-alpine3.17 AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
ADD go.mod .
COPY . .
RUN go build -o TIKTOK_User ./cmd/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/tiktok_user /build/tiktok_user

CMD ["./tiktok_user"]