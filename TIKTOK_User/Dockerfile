FROM golang:alpine AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
ADD go.mod .
COPY . .
RUN go build -o tiktok_user ./cmd/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/tiktok_user /build/tiktok_user

CMD ["./tiktok_user"]