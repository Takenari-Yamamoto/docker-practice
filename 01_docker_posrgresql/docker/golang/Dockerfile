FROM golang:1.18-alpine

RUN apk --no-cache add bash tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk add --no-cache gcc musl-dev
