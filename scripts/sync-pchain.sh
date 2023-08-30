#!/bin/sh
# syncs pandasia with the P Chain

set -e

if [ $# -ne 1 ]; then
    echo "Usage: $0 <interval>"
    exit 1
fi

INTERVAL=$1

# sleep for a minute to let the server come online
sleep 60

while true; do
		echo "Syncing with P Chain..."
		/app/bin/pandasia sync-pchain --node-url $ETH_RPC_URL --db /data/pandasia.db
		echo "Done. Sleeping..."
		sleep $INTERVAL
done
