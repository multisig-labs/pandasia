import { toHex } from "https://esm.sh/viem@1.6.4";
import { utils as ethersUtils } from "https://esm.sh/ethers@5.7.2";

async function sha256(message) {
	const buffer = await window.crypto.subtle.digest("SHA-256", message.buffer);
	return new Uint8Array(buffer);
}

async function cb58Encode(message) {
	const payload = ethersUtils.arrayify(message);
	const checksum = await sha256(payload);
	const buffer = new Uint8Array(payload.length + 4);
	buffer.set(payload);
	buffer.set(checksum.slice(-4), payload.length);
	return ethersUtils.base58.encode(new Uint8Array(buffer));
}

async function cb58Decode(message) {
	const buffer = ethersUtils.base58.decode(message);
	const payload = buffer.slice(0, -4);
	const checksum = buffer.slice(-4);
	const newChecksum = (await sha256(payload)).slice(-4);

	if (
		(checksum[0] ^ newChecksum[0]) |
		(checksum[1] ^ newChecksum[1]) |
		(checksum[2] ^ newChecksum[2]) |
		(checksum[3] ^ newChecksum[3])
	)
		throw new Error("Invalid checksum");
	return payload;
}

async function decodeSig(sig) {
	const sigBytes = await cb58Decode(sig);
	const v = toHex(sigBytes.slice(64));
	const r = toHex(sigBytes.slice(0, 32));
	const s = toHex(sigBytes.slice(32, 64));
	return { v, r, s };
}

export { decodeSig, cb58Decode, cb58Encode, sha256 };
