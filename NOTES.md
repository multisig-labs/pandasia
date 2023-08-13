```
Mnemonic: test test test test test test test test test test test test test test test test test test test test test test test blade

P-Addr: P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww
address bytes: 0x424328bf10cdaeeda6bb05a78cff90a0bea12c02
priv key cb58: PrivateKey-2GCgSHhQMwycn28YMovLxDeka4dr6NeD1LFmMY7Dcbg14Sx7uX
priv key hex: a6367274e753df164ee95497c42c9a1f879307cb538eb34bbe29dc29950ed64d
PrivKey: 507269766174654b65792d32474367534868514d7779636e3238594d6f764c7844656b61346472364e6544314c466d4d59374463626731345378377558
serialized compressed pub key bytes: 03f9e73672eb9865f4e8fefd3cc508121661c59482c2677cbef5426163c86ee0df
<format byte = 0x02/0x03><32-byte X coordinate>


C-Addr: 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4 0x0961ca10d49b9b8e371aa0bcf77fe5730b18f2e4
priv key: 93b3701cf8eeb6f7d3b22211c691734f24816a02efa933f67f34d37053182577

sign msg: 0x0961ca10d49b9b8e371aa0bcf77fe5730b18f2e4
hash to sign:
signature: YezQscPPaK8pXRbLuy1abhCozJbLZnEkZnJcbv3c3x5YG2XQo1Cqb7HQ2GTPLj9uRr2NWrY2vD1XTggnmVczdETza21y2C
36d190a0b334483474936da7fb45912df86a894a265e93bd0ea6f2e96abe24f0 38d596998a3667b3f1f44416f92c2d0234e414db37f7473da42d1ab9ddab24f4 00 94d99d15
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
