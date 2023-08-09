// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.17;

import "./SECP256K1.sol";

contract Pandasia {
	mapping(address => address) p2c;

	function registerPChainAddr(bytes32 pChainSig) external {
		address hash = msg.sender;
		//pubKey = recoverPubKey(hash, pChainSig)
	}

	function recoverPersonalSignAddress(bytes32 message, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
		bytes memory publicKey = recoverPersonalSignPublicKey(message, v, r, s);
		return address(uint160(uint256(keccak256(publicKey))));
	}

	function recoverPersonalSignPublicKey(bytes32 message, uint8 v, bytes32 r, bytes32 s) public pure returns (bytes memory) {
		string memory header = "\x1AAvalanche Signed Message:\n32";
		bytes32 _message = keccak256(abi.encodePacked(header, message));
		(uint256 x, uint256 y) = SECP256K1.recover(uint256(_message), v - 27, uint256(r), uint256(s));
		return abi.encodePacked(x, y);
	}

	function addressToString(address _address) public pure returns (string memory) {
		bytes32 _bytes = bytes32(uint256(_address));
		bytes memory HEX = "0123456789abcdef";
		bytes memory _string = new bytes(42);
		_string[0] = "0";
		_string[1] = "x";
		for (uint i = 0; i < 20; i++) {
			_string[2 + i * 2] = HEX[uint8(_bytes[i + 12] >> 4)];
			_string[3 + i * 2] = HEX[uint8(_bytes[i + 12] & 0x0f)];
		}
		return string(_string);
	}

	// Strings.toHexString(uint160(address), 20)
	function toHexString(uint256 value, uint256 length) internal pure returns (string memory) {
		uint256 localValue = value;
		bytes memory buffer = new bytes(2 * length + 2);
		buffer[0] = "0";
		buffer[1] = "x";
		for (uint256 i = 2 * length + 1; i > 1; --i) {
			buffer[i] = _HEX_DIGITS[localValue & 0xf];
			localValue >>= 4;
		}
		if (localValue != 0) {
			revert StringsInsufficientHexLength(value, length);
		}
		return string(buffer);
	}
}
