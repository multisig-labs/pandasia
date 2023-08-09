// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.17;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import "./SECP256K1.sol";

contract Pandasia {
	mapping(address => address) p2c;

	// Given an address, convert to its string (lowercase) format, and hash a message like the avalanche wallet would do
	function hashMessage(address addr) public pure returns (bytes32) {
		bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
		// len of an ascii addr is 42 bytes
		uint32 addrLen = 42;
		bytes memory addrStr = bytes(Strings.toHexString(uint160(addr), 20));
		return sha256(abi.encodePacked(header, addrLen, addrStr));
	}

	function registerPChainAddr(bytes32 pChainSig) external {
		address hash = msg.sender;
		//pubKey = recoverPubKey(hash, pChainSig)
	}

	function recoverPersonalSignAddress(bytes32 message, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
		bytes memory publicKey = recoverPersonalSignPublicKey(message, v, r, s);
		return address(uint160(uint256(keccak256(publicKey))));
	}

	// TODO make this work
	function recoverPersonalSignPublicKey(bytes32 message, uint8 v, bytes32 r, bytes32 s) public pure returns (bytes memory) {
		// string memory header = "\x1AAvalanche Signed Message:\n";
		bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
		bytes32 _message = keccak256(abi.encodePacked(header, message));
		(uint256 x, uint256 y) = SECP256K1.recover(uint256(_message), v - 27, uint256(r), uint256(s));
		return abi.encodePacked(x, y);
	}
}
