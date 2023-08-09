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
signing msg 0x63682bDC5f875e9bF69E201550658492C9763F89
hash to sign: 22BC946873FAEDBD2201CD7F0740244B555CC2B60C018F3E74A63D1BBBC32336
sig: 3TGqznbmnJ4tSTgCkfy6V2V5LjkwmkDrNqHovUbsqyAkYMp7QN5BPGyUeZzvYHk7FGqfGVzdbyGnvk4V2PTeoHQmVpmBzMi
