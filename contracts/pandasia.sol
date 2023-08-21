// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.17;

import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {AddressChecksumUtils} from "./AddressChecksumUtils.sol";
import "./SECP256K1.sol";

contract Pandasia is Ownable {
	error PAddrNotInValidatorMerkleTree();

	bytes32 public merkleRoot;
	mapping(address => address) private c2p;

	function setRoot(bytes32 root) public onlyOwner {
		merkleRoot = root;
	}

	// Given an address, convert to its checksummed string (mixedcase) format, and hash a message like the avalanche wallet would do
	function hashChecksummedMessage(address addr) public pure returns (bytes32) {
		bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
		// len of an ascii addr is 42 bytes
		uint32 addrLen = 42;
		string memory addrStr = AddressChecksumUtils.getChecksum(addr);
		return sha256(abi.encodePacked(header, addrLen, "0x", addrStr));
	}

	// Sign C-chain address in mixedcase hex format with P-chain addr on wallet.avax.network
	function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] memory proof) public {
		bytes32 msgHash = hashChecksummedMessage(msg.sender);
		(uint256 x, uint256 y) = SECP256K1.recover(uint256(msgHash), v, uint256(r), uint256(s));
		address paddy = pubKeyBytesToAvaAddressBytes(x, y);
		if (verify(merkleRoot, paddy, proof)) {
			// We could just store a flag as well, preserving privacy a little bit?
			// To really preserve priv we need a way to salt the message they sign
			c2p[msg.sender] = paddy;
		} else {
			revert PAddrNotInValidatorMerkleTree();
		}
	}

	function isRegisteredValidator(address addr) public view returns (bool) {
		return c2p[addr] != address(0);
	}

	function pubKeyBytesToAvaAddressBytes(uint256 x, uint256 y) public pure returns (address) {
		// Calc format prefix for compressed pub key
		bytes1 format = (y % 2 == 0) ? bytes1(0x02) : bytes1(0x03);
		bytes memory pubKey = abi.encodePacked(format, x);
		bytes32 pubKeySha = sha256(pubKey);
		return address(ripemd160(abi.encodePacked(pubKeySha)));
	}

	function verify(bytes32 root, address account, bytes32[] memory proof) public pure returns (bool) {
		return MerkleProof.verify(proof, root, _leaf(account));
	}

	function _leaf(address account) internal pure returns (bytes32) {
		return keccak256(bytes.concat(keccak256(abi.encode(account))));
	}
}
