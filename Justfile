# Justfiles are better Makefiles (Don't @ me)
# Install the `just` command from here https://github.com/casey/just
# or `cargo install just` or `brew install just`
# https://cheatography.com/linux-china/cheat-sheets/justfile/

# Build vars for versioning the binary
VERSION := `grep "const Version " pkg/version/version.go | sed -E 's/.*"(.+)"$$/\1/'`
GIT_COMMIT := `git rev-parse HEAD`
BUILD_DATE := `date '+%Y-%m-%d'`
VERSION_PATH := "github.com/multisig-labs/pandasia/pkg/version"
LDFLAGS := "-X " + VERSION_PATH + ".BuildDate=" + BUILD_DATE + " -X " + VERSION_PATH + ".Version=" + VERSION + " -X " + VERSION_PATH + ".GitCommit=" + GIT_COMMIT
DOCKER_IMAGE_NAME := "ghcr.io/multisig-labs/pandasia"
DOCKER_IMAGE_TAG := "latest"
PANDASIA_ADDR := env_var("PANDASIA_ADDR")
export ETH_RPC_URL := env_var_or_default("ETH_RPC_URL", "http://127.0.0.1:9650")
export MNEMONIC := env_var_or_default("MNEMONIC", "test test test test test test test test test test test junk")
# First key from MNEMONIC
export PRIVATE_KEY := env_var_or_default("PRIVATE_KEY", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")

# Autoload a .env if one exists
set dotenv-load

# Print out some help
default:
	@just --list --unsorted

# Install dependencies
install:
	brew install sqlc sqlite3
	forge install
	# echo 'export PATH="/opt/homebrew/opt/sqlite/bin:$PATH"' >> ~/.zshrc
	echo "Make sure sqlite3 from homebrew is in your path"

# Delete artifacts
clean:
	forge clean

# Build Go and Solidity
build:
	sqlc generate
	go build -ldflags "{{LDFLAGS}}" -o bin/pandasia main.go
	forge build

# Run forge unit tests
test contract="." test="." *flags="":
	forge test --match-contract {{contract}} --match-test {{test}} {{flags}}

deploy: (_ping ETH_RPC_URL)
	#!/bin/bash
	forge script --broadcast --slow --ffi --fork-url=${ETH_RPC_URL} --private-key=${PRIVATE_KEY} scripts/deploy.s.sol
	chain_id=$(cast chain-id)
	addr=$(cat broadcast/deploy.s.sol/$chain_id/run-latest.json | jq -r ".transactions[2].contractAddress")
	echo "Pandasia deployed to $addr"
	sed -i '' "s/^PANDASIA_ADDR=.*/PANDASIA_ADDR=${addr}/" .env
	rm -f public/js/abi.json
	cat artifacts-forge/pandasia.sol/Pandasia.json | jq "{abi: .abi}" > public/js/abi.json

# Execute a Forge script
forge-script cmd:
	#!/usr/bin/env bash
	fn={{cmd}}
	forge script --broadcast --slow --ffi --fork-url=${ETH_RPC_URL} --private-key=${PRIVATE_KEY} scripts/${fn%.*.*}.s.sol

cast-submit-root root: (_ping ETH_RPC_URL)
	cast send --private-key=${PRIVATE_KEY} ${PANDASIA_ADDR} "setMerkleRoot(bytes32)" {{root}}

cast-is-validator caddr: (_ping ETH_RPC_URL)
	cast call ${PANDASIA_ADDR} "isRegisteredValidator(address)" {{caddr}}

# TODO create a P Chain testing table

sync: (_ping ETH_RPC_URL)
	bin/pandasia sync-pchain --node-url=${ETH_RPC_URL} --db data/pandasia-fuji.db

anvil:
	anvil --port 9650 --mnemonic "${MNEMONIC}"

anvil-fork:
	anvil --port 9650 --mnemonic "${MNEMONIC}" --fork-url https://nd-058-850-167.p2pify.com/4e4706b8fc3a3bb4a5559c84671a1cf4/ext/bc/C/rpc

# Delete and recreate a dev sqlite db
create-dev-db:
	mkdir -p data
	rm -f data/pandasia-dev.db*
	cat pkg/db/schema.sql | sqlite3 data/pandasia-dev.db
	unzip -p pkg/db/txs-sample.sql.zip | sqlite3 data/pandasia-dev.db

# Generate Go code interface for /contracts
codegen:
	#!/bin/bash
	CORETH=0.11.9
	forge build
	THISDIR=$PWD
	echo "Generating GO code with Coreth v${CORETH}"
	cd $GOPATH/pkg/mod/github.com/ava-labs/coreth@v${CORETH}
	cat $THISDIR/artifacts-forge/contracts/Pandasia.sol/Pandasia.json | jq '.abi' | go run cmd/abigen/main.go --abi - --pkg pandasia --out $THISDIR/pkg/contracts/pandasia/pandasia.go

# Check if there is an http(s) server listening on [url]
_ping url:
	@if ! curl -k --silent --connect-timeout 2 {{url}} >/dev/null 2>&1; then echo 'No server at {{url}}!' && exit 1; fi

serve:
	JOB_PERIOD=10h SERVE_EMBEDDED=false bin/pandasia serve --db data/pandasia-dev.db --node-url $ETH_RPC_URL --pandasia-addr $PANDASIA_ADDR

keys:
	ggt utils mnemonic-keys "${MNEMONIC}"

build-docker:
	docker build --platform linux/amd64 --build-arg LDFLAGS="{{LDFLAGS}}" -t {{DOCKER_IMAGE_NAME}}:{{DOCKER_IMAGE_TAG}} .

run-docker:
	docker run --platform linux/amd64 --name "pandasia" -p 8000:8000 -v $(pwd)/data:/data -e ETH_RPC_URL={{ETH_RPC_URL}} -e PANDASIA_ADDR={{PANDASIA_ADDR}}  -e PRIVATE_KEY={{PRIVATE_KEY}} {{DOCKER_IMAGE_NAME}}:{{DOCKER_IMAGE_TAG}}

run-docker-it:
	docker run --platform linux/amd64 --name "pandasia" -it -p 8000:8000 -v $(pwd)/data:/data -e ETH_RPC_URL={{ETH_RPC_URL}} -e PANDASIA_ADDR={{PANDASIA_ADDR}} -e PRIVATE_KEY={{PRIVATE_KEY}} {{DOCKER_IMAGE_NAME}}:{{DOCKER_IMAGE_TAG}} /bin/bash

rm-docker:
	docker rm pandasia

kill-docker:
	docker kill pandasia

chain-id:
	cast chain-id

decoded-errors:
	#!/usr/bin/env bash
	join() { local d=$1 s=$2; shift 2 && printf %s "$s${@/#/$d}"; }
	shopt -s globstar # so /**/ works
	errors=$(cat artifacts-forge/**/*.json | jq -r '.abi[]? | select(.type == "error") | .name' | sort | uniq)
	sigsArray=()
	for x in $errors;	do
		sigsArray+=("\"$(cast sig "${x}()")\":\"${x}()\"")
	done
	sigs=$(join ',' ${sigsArray[*]})
	echo "{${sigs}}" | jq
