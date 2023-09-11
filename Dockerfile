# syntax=docker/dockerfile:1
# build go code
FROM golang:1.21 AS builder

# install dependencies
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN GOBIN=/src/bin/ go install github.com/DarthSim/overmind/v2@latest

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

# build cast
FROM ghcr.io/foundry-rs/foundry:nightly AS foundry

# minimal runtime container
FROM debian:bookworm-slim AS runtime

WORKDIR /app

ENV SERVE_EMBEDDED=false
ENV JOB_PERIOD=10h
ENV ETH_RPC_URL=https://api.avax-test.network/ext/bc/C/rpc
ENV PANDASIA_ADDR=none
ENV PRIVATE_KEY=none

RUN apt-get update && apt-get install -y ca-certificates jq curl tmux git && rm -rf /var/lib/apt/lists/*
RUN mkdir /data
RUN mkdir -p /app/bin
# copy scripts and Procfile
COPY scripts/ /app/scripts/
COPY Procfile /app/Procfile
# copy go binaries to /app/bin
COPY --from=builder /src/bin/pandasia /app/bin/pandasia
COPY --from=builder /src/bin/overmind /app/bin/overmind
# copy cast command to /app/bin
COPY --from=foundry /usr/local/bin/cast /app/bin/cast

CMD /app/bin/overmind start -f /app/Procfile
