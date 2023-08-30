# syntax=docker/dockerfile:1
FROM golang:1.21 AS builder

# install dependencies
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# copy and install libs
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go /src/main.go
COPY pkg/ /src/pkg/
COPY sqlc.yaml /src/sqlc.yaml
COPY public/ /src/public/

ARG CGO_CFLAGS="-O -D__BLST_PORTABLE__"
ARG CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"

RUN sqlc generate
RUN go build -o bin/pandasia main.go

ENV SERVE_EMBEDDED=false
ENV JOB_PERIOD=10h
ENV ETH_RPC_URL=https://api.avax-test.network/ext/bc/C/rpc
ENV PANDASIA_ADDR=none


CMD /src/bin/pandasia serve --db /data/pandasia-dev.db --node-url ${ETH_RPC_URL} --pandasia-addr ${PANDASIA_ADDR}
