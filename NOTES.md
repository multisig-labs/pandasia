# TODO

Use this to check our work on if rewards have been recvd
https://ava-labs-inc.metabaseapp.com/public/question/78f9de45-1f09-4b08-847b-086372bdcc4c

# Command Scratchpad

```bash
just install

# Build everything
just build

just create-dev-db

# Start Anvil in one terminal
just anvil

# Deploy contracts in another terminal
just deploy

# Start Pandasia API server in another terminal
JOB_PERIOD=10h SERVE_EMBEDDED=false bin/pandasia serve --db data/pandasia-dev.db --node-url $ETH_RPC_URL --pandasia-addr $PANDASIA_ADDR

# Generate a merkle tree at current height and store in DB
bin/pandasia generate-tree --db data/pandasia-dev.db

# Submit current root to contract
export CURRENT_ROOT=$(curl --silent $PANDASIA_URL/trees | jq -r '.[0].Root'); echo ${CURRENT_ROOT}
just cast-submit-root ${CURRENT_ROOT}


# Test addrs and a sig from wallet.avax.network
export PADDR="P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww"
export CADDR="0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"
export SIG="24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx"

curl --silent "localhost:8000/proof/${CURRENT_ROOT}?addr=${PADDR}&sig=${SIG}"
scripts/register.ts ${PADDR} ${SIG}
just cast-is-validator ${CADDR}

# Export the merkle tree json to stdout
sqlite3 data/pandasia-dev.db 'select tree from merkle_trees where id = 1' | jq

curl --silent localhost:8000/airdrops

just forge-script createAirdrop

echo "0x0000000000000000000000000000000000000001\n0x0000000000000000000000000000000000000002" | bin/pandasia generate-tree-stdin --db data/pandasia-dev.db --desc "test tree"


cast call $PANDASIA_ADDR "getAirdropIds(address)(uint64[])" 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4
cast call $PANDASIA_ADDR "getAirdrops(uint64,uint64)()" 0 10

cast call $PANDASIA_ADDR "canClaimAirdrop(address, uint64, bytes32[])(bool)" $TEST_ADDR 1 "[0x6578616d706c6500000000000000000000000000000000000000000000000000,0x6578616d706c6500000000000000000000000000000000000000000000000000]"

cast call $PANDASIA_ADDR "canClaimAirdrop(address, uint64, bytes32[])(bool)" $TEST_ADDR 1 "[]"

cast call $PANDASIA_ADDR "stakingContract()(address)"

cast send $PANDASIA_ADDR "setStakingContract(address)" 0x9946e68490D71Fe976951e360f295c4Cf8531D00 --from $OWNER --private-key $PRIVATE_KEY


just forge-script canClaim


```

```bash
Mnemonic: test test test test test test test test test test test test test test test test test test test test test test test blade

P-Addr: P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww
address bytes: 0x424328bf10cdaeeda6bb05a78cff90a0bea12c02
priv key cb58: PrivateKey-2GCgSHhQMwycn28YMovLxDeka4dr6NeD1LFmMY7Dcbg14Sx7uX
priv key hex: a6367274e753df164ee95497c42c9a1f879307cb538eb34bbe29dc29950ed64d
PrivKey: 507269766174654b65792d32474367534868514d7779636e3238594d6f764c7844656b61346472364e6544314c466d4d59374463626731345378377558
serialized compressed pub key bytes: 03f9e73672eb9865f4e8fefd3cc508121661c59482c2677cbef5426163c86ee0df
<format byte = 0x02/0x03><32-byte X coordinate>


C-Addr: 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4
priv key: 93b3701cf8eeb6f7d3b22211c691734f24816a02efa933f67f34d37053182577

sign msg: 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4
hash to sign:
signature: 24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx

{
  v: "0x00",
  r: "0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337",
  s: "0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a"
}
```

### Sample custom Root

```json
{
  "format": "standard-v1",
  "tree": [
    "0x2e1dea9890e94d280361af414f0696bc7ba251d4e52a5f786d15629d5185a89c",
    "0x9fec67521532e3df3ca2461c12c79c2b89e8f633311a6b525fb8488fbcd1d177",
    "0xb5d9d894133a730aa651ef62d26b0ffa846233c74177a591a4a896adfda97d22",
    "0x20b2f891eaf390d96349554ee7297a8a8972c13215d1aa9dd752f9c6822c1888",
    "0x1ab0c6948a275349ae45a06aad66a8bd65ac18074615d53676c09b67809099e0"
  ],
  "values": [
    {
      "value": ["0x0000000000000000000000000000000000000001"],
      "treeIndex": 2
    },
    {
      "value": ["0x0000000000000000000000000000000000000002"],
      "treeIndex": 4
    },
    {
      "value": ["0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02"],
      "treeIndex": 3
    }
  ],
  "leafEncoding": ["address"]
}
```

If we want to use all lowercase (un-checksummed) c-chain addrs we can do this:

```
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
// Given an address, convert to its string (lowercase) format, and hash a message like the avalanche wallet would do
function hashMessage(address addr) public pure returns (bytes32) {
	bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
	// len of an ascii addr is 42 bytes
	uint32 addrLen = 42;
	bytes memory addrStr = bytes(Strings.toHexString(uint160(addr), 20));
	return sha256(abi.encodePacked(header, addrLen, addrStr));
}
```
