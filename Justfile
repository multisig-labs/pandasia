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
	forge install

# Delete artifacts
clean:
	forge clean

# Build Go and Solidity
build:
	sqlc generate
	go build -ldflags "{{LDFLAGS}}" -o bin/pandasia main.go
	forge build

test:
	forge test

deploy: (_ping ETH_RPC_URL)
	#!/bin/bash
	forge script --broadcast --slow --ffi --fork-url=${ETH_RPC_URL} --private-key=${PRIVATE_KEY} scripts/deploy.s.sol
	addr=$(cat broadcast/deploy.s.sol/31337/run-latest.json | jq -r ".transactions[2].contractAddress")
	echo "Pandasia deployed to $addr"
	sed -i '' "s/^PANDASIA_ADDR=.*/PANDASIA_ADDR=${addr}/" .env

cast-submit-root root: (_ping ETH_RPC_URL)
	cast send --private-key=${PRIVATE_KEY} ${PANDASIA_ADDR} "setRoot(bytes32)" {{root}}

anvil:
	anvil --port 9650 --mnemonic "${MNEMONIC}"

# Delete and recreate a dev sqlite db
create-db:
	rm -f data/pandasia-dev.db*
	cat schema.sql | sqlite3 data/pandasia-dev.db

# Check if there is an http(s) server listening on [url]
_ping url:
	@if ! curl -k --silent --connect-timeout 2 {{url}} >/dev/null 2>&1; then echo 'No server at {{url}}!' && exit 1; fi
