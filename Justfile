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
	brew install sqlc
	forge install

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
	addr=$(cat broadcast/deploy.s.sol/31337/run-latest.json | jq -r ".transactions[2].contractAddress")
	echo "Pandasia deployed to $addr"
	sed -i '' "s/^PANDASIA_ADDR=.*/PANDASIA_ADDR=${addr}/" .env
	rm -f public/js/abi.json
	cat artifacts-forge/contracts/Pandasia.sol/Pandasia.json | jq "{abi: .abi}" > public/js/abi.json

# Execute a Forge script
forge-script cmd:
	#!/usr/bin/env bash
	fn={{cmd}}
	forge script --broadcast --slow --ffi --fork-url=${ETH_RPC_URL} --private-key=${PRIVATE_KEY} scripts/${fn%.*.*}.s.sol

cast-submit-root root: (_ping ETH_RPC_URL)
	cast send --private-key=${PRIVATE_KEY} ${PANDASIA_ADDR} "setValidatorRoot(bytes32)" {{root}}

cast-is-validator caddr: (_ping ETH_RPC_URL)
	cast call ${PANDASIA_ADDR} "isRegisteredValidator(address)" {{caddr}}

anvil:
	anvil --port 9650 --mnemonic "${MNEMONIC}"

# Delete and recreate a dev sqlite db
create-dev-db:
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
