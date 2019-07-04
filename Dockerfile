FROM golang:1.12.0-alpine AS builder

RUN apk add --no-cache git

WORKDIR /target

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

ENTRYPOINT ["/target/aerospike-oom"]
