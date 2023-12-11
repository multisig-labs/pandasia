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
    TREES=$(curl --silent localhost:8000/trees)
		CURRENT_ROOT=$(echo $TREES | jq -r '.[0].Root')
    CURRENT_HEIGHT=$(echo $TREES | jq -r '.[0].Height')
		echo "Posting root to contract..."
		if ! /app/bin/cast send --private-key=$PRIVATE_KEY $PANDASIA_ADDR "setMerkleRoot(bytes32,uint64)" $CURRENT_ROOT $CURRENT_HEIGHT; then
		    echo "Error posting root to contract to $ETH_RPC_URL. Will try again after sleeping..."
		fi
		echo "Done. Sleeping..."
		sleep $INTERVAL
done
