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

RUN sqlc generate
RUN go build -o bin/pandasia main.go

ENV SERVE_EMBEDDED=false
ENV JOB_PERIOD=1h

CMD ["/src/bin/pandasia", "serve", "--db", "data/pandasia-dev.db", "--node-url", "$ETH_RPC_URL", "--pandasia-addr", "$PANDASIA_ADDR"]
