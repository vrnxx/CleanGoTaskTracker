FROM golang:1.24.1 AS builder
WORKDIR /app
COPY /tracker /app
COPY tracker/go.mod tracker/go.sum /modules/

RUN go mod download
RUN go env -w CGO_ENABLED=0

RUN go build -o main ./cmd/app/main.go

FROM ubuntu:24.04
RUN apt-get update && apt-get install -y curl
COPY --from=builder /app/configs/app/prod.toml /app/main /usr/local/bin/