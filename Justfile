# Justfiles are better Makefiles (Don't @ me)
# Install the `just` command from here https://github.com/casey/just
# or `cargo install just` or `brew install just`
# https://cheatography.com/linux-china/cheat-sheets/justfile/

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

#
build:
	forge build

