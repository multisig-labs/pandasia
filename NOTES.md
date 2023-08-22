# TODO

Use this to check our work on if rewards have been recvd
https://ava-labs-inc.metabaseapp.com/public/question/78f9de45-1f09-4b08-847b-086372bdcc4c

# Command Scratchpad

```bash
export PADDR="P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww"
export CADDR="0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"
export SIG="24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx"

just anvil
JOB_PERIOD=10h SERVE_EMBEDDED=false bin/pandasia serve --db data/pandasia-dev.db --node-url http://100.83.243.106:9650

just deploy

export CURRENT_ROOT=$(curl --silent localhost:8000/trees | jq -r '.[0].Root'); echo ${CURRENT_ROOT}
just cast-submit-root ${CURRENT_ROOT}

curl --silent "localhost:8000/proof/${CURRENT_ROOT}?addr=${PADDR}&sig=${SIG}"
scripts/register.ts ${PADDR} ${SIG}
just cast-is-validator ${CADDR}

sqlite3 data/pandasia-dev.db '.once out.json' 'select tree from merkle_trees'

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
