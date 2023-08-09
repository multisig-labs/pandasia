https://github.com/0xcyphered/secp256k1-solidity/blob/main/contracts/SECP256K1.sol

```
func DigestAvaMsg(msg string) []byte {
	msgb := []byte(msg)
	l := uint32(len(msgb))
	lb := make([]byte, 4)
	binary.BigEndian.PutUint32(lb, l)
	prefix := []byte("\x1AAvalanche Signed Message:\n")

	buf := new(bytes.Buffer)
	buf.Write(prefix)
	buf.Write(lb)
	buf.Write(msgb)
	fullmsg := buf.Bytes()
	h := sha256.Sum256(fullmsg)
	return h[:]
}
```

// prefix size: 26 bytes
0x1a
// prefix: Avalanche Signed Message:\n
0x41 0x76 0x61 0x6c 0x61 0x6e 0x63 0x68 0x65 0x20 0x53 0x69 0x67 0x6e 0x65 0x64 0x20 0x4d 0x65 0x73 0x73 0x61 0x67 0x65 0x3a 0x0a
// msg size: 30 bytes
0x00 0x00 0x00 0x1e
// msg: Through consensus to the stars
54 68 72 6f 75 67 68 20 63 6f 6e 73 65 6e 73 75 73 20 74 6f 20 74 68 65 20 73 74 61 72 73

```
const signature = "0x...."

const r = signature.slice(0, 66);
const s = "0x" + signature.slice(66, 130);
const v = parseInt(signature.slice(130, 132), 16);

function verify(bytes32 _data, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
    bytes memory prefix = "\x19Ethereum Signed Message:\n32";
    bytes32 hash = keccak256(abi.encodePacked(prefix, _data));
    address signer = ecrecover(hash, _v, _r, _s);
    return signer;
}

```

addr 1 from my ledger
P-avax15ayepmzwddzewjljl2dwwq7j9kvh4kj3x6cjp5
has no rewards
signing msg 0x63682bdc5f875e9bf69e201550658492c9763f89
hash to sign: 68C88E730ECED13EE4A68EFF65D3D250BB7B0F27C1CB4C8E20C52514D45D9390
sig: YPKVD1F1FXpeTrBeu2gN6xgniiVXGJVtFJnqrPvKsmbtkXF313XiQcBAYWFWjB13wXuJTz2XB3X5CuC3qtopBFGEX1uTpE
