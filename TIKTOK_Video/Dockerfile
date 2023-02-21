FROM golang:alpine AS builder

WORKDIR /build

ENV GOPROXY https://goproxy.cn
ADD go.mod .
COPY . .
RUN go build -o tiktok_video ./cmd/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/tiktok_video /build/tiktok_video

CMD ["./tiktok_video"]