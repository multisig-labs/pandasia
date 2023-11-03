#!/bin/sh
# posts trees to the c chain contract

set -e

if [ $# -ne 1 ]; then
    echo "Usage: $0 <interval>"
    exit 1
fi

INTERVAL=$1
ETH_RPC_URL=$ETH_RPC_URL/ext/bc/C/rpc

# sleep for a minute to let the server come online
sleep 60

while true; do
		echo "Getting current root..."
		CURRENT_ROOT=$(curl --silent localhost:8000/trees | jq -r '.[0].Root')
		echo "Posting root to contract..."
		if ! /app/bin/cast send --private-key=$PRIVATE_KEY $PANDASIA_ADDR "setValidatorRoot(bytes32)" $CURRENT_ROOT; then
		    echo "Error posting root to contract. Will try again after sleeping..."
		fi
		echo "Done. Sleeping..."
		sleep $INTERVAL
done
