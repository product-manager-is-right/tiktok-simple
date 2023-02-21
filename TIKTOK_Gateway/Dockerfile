FROM golang:alpine AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
ADD go.mod .
COPY . .
RUN go build -o tiktok_gateway ./cmd/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/tiktok_gateway /build/tiktok_gateway

CMD ["./tiktok_gateway"]