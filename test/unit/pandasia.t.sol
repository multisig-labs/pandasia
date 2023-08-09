pragma solidity 0.8.17;

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {console2} from "forge-std/console2.sol";
import {Pandasia} from "../../contracts/pandasia.sol";

contract PandasiaTest is Test {
	function setUp() public {}

	function testSig() public {
		string memory caddrStr = "0x63682bDC5f875e9bF69E201550658492C9763F89";
		// address caddr = address(0x63682bDC5f875e9bF69E201550658492C9763F89);
		// string memory header = "\x1AAvalanche Signed Message:\n";
		bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
		console2.logBytes(header);
		// len of an ascii addr is 42 bytes
		uint32 addrLen = 42;
		bytes32 message = sha256(abi.encodePacked(header, addrLen, caddrStr));
		console2.logBytes32(message);
	}

	function testMessageHash() public {
		// address caddr = address(0x63682bDC5f875e9bF69E201550658492C9763F89);
		// string memory s = Strings.toHexString(uint160(caddr), 20);
		// console2.log(s);
		// bytes memory addrStr = bytes(Strings.toHexString(uint160(caddr), 20));
		// console2.logBytes(addrStr);

		Pandasia pandasia = new Pandasia();
		bytes32 h = pandasia.hashMessage(address(0x63682bDC5f875e9bF69E201550658492C9763F89));
		bytes32 hashFromWallet = bytes32(0x68C88E730ECED13EE4A68EFF65D3D250BB7B0F27C1CB4C8E20C52514D45D9390);
		// console2.logBytes32(hashFromWallet);
		assertEq(h, hashFromWallet);
	}
}
