# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./api ./api
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY .golangci.yml .
RUN go build -o service ./cmd/main.go

# Run stage
FROM alpine:3.18
RUN addgroup -S koubonigrp && adduser -S kouboni -G koubonigrp
WORKDIR /root/
COPY --from=builder /app/service .
USER kouboni
EXPOSE 8080
ENTRYPOINT ["./service"]
