// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.19;

import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {AddressChecksumUtils} from "./AddressChecksumUtils.sol";
import "./SECP256K1.sol";

contract Pandasia is Ownable {
	using SafeERC20 for IERC20;

	error PAddrNotInValidatorMerkleTree();
	error PAddrAlreadyRegistered();
	error InvalidAddress();
	error InvalidAmount();
	error AirdropOutOfFunds();
	error AddressNotEligible();
	error InvalidWithdrawRequest();

	struct Airdrop {
		address owner; // account that contributed the funds
		address erc20; // claimable asset
		uint256 balance; // current balance of asset
		uint256 amount; // amount claimable by each validator
		uint256 expires; // time that airdop expires and owner can re-claim any left over funds
		bytes32 memo; // notes about the airdrop
	}

	mapping(uint256 => Airdrop) public airdrops;
	mapping(uint256 => mapping(address => bool)) public claimed;
	mapping(address => uint256[]) public airdropIds; // owners => airdropIds
	uint256 public airdropCount;

	bytes32 public merkleRoot;
	mapping(address => address) public c2p;
	mapping(address => address) public p2c;

	function setRoot(bytes32 root) external onlyOwner {
		merkleRoot = root;
	}

	function newAirdrop(address erc20, uint256 amount, uint256 expires, bytes32 memo) external returns (uint256) {
		if (erc20 == address(0)) {
			revert InvalidAddress();
		}
		if (amount == 0) {
			revert InvalidAmount();
		}
		airdropCount++;
		Airdrop storage airdrop = airdrops[airdropCount];
		// Do we care? Can they set any owner on creation?
		airdrop.owner = msg.sender;
		airdrop.erc20 = erc20;
		airdrop.amount = amount;
		airdrop.expires = expires;
		airdrop.memo = memo;

		airdropIds[msg.sender].push(airdropCount);
		return airdropCount;
	}

	function fundAirdrop(uint256 airdropId, uint256 amount) external {
		Airdrop storage airdrop = airdrops[airdropId];
		IERC20 token = IERC20(airdrop.erc20);
		uint256 balance = token.balanceOf(msg.sender);
		if (amount == 0 || balance < amount || amount < airdrop.amount) {
			revert InvalidAmount();
		}

		token.safeTransferFrom(msg.sender, address(this), amount);
		airdrop.balance = airdrop.balance + amount;
	}

	function fundAirdropWithPermit(uint256 airdropId, uint256 amount, address funder, uint256 deadline, uint8 v, bytes32 r, bytes32 s) external {
		Airdrop storage airdrop = airdrops[airdropId];
		IERC20 token = IERC20(airdrop.erc20);
		uint256 balance = token.balanceOf(msg.sender);
		if (balance < amount || amount == 0) {
			revert InvalidAmount();
		}

		IERC20Permit(airdrop.erc20).permit(funder, address(this), amount, deadline, v, r, s);
		token.safeTransferFrom(msg.sender, address(this), amount);
		airdrop.balance = airdrop.balance + amount;
	}

	function claimAirdrop(uint256 airdropId) external {
		if (c2p[msg.sender] == address(0) || claimed[airdropId][msg.sender]) {
			revert AddressNotEligible();
		}

		Airdrop storage airdrop = airdrops[airdropId];
		if (airdrop.balance < airdrop.amount) {
			revert AirdropOutOfFunds();
		}
		claimed[airdropId][msg.sender] = true;
		airdrop.balance = airdrop.balance - airdrop.amount;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, airdrop.amount);
	}

	// TODO Safes? Wat do?
	function claimAirdropWithPermit(uint256 airdropId) external {}

	function getAirdropIds(address owner) public view returns (uint256[] memory) {
		return airdropIds[owner];
	}

	function withdrawFunding(uint256 airdropId, uint256 withdrawAmt) external {
		Airdrop storage airdrop = airdrops[airdropId];
		if (airdrop.owner != msg.sender || airdrop.balance < withdrawAmt || block.timestamp < airdrop.expires) {
			revert InvalidWithdrawRequest();
		}
		airdrop.balance = airdrop.balance - withdrawAmt;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
	}

	function emergencyWithdraw(uint256 airdropId, uint256 withdrawAmt) external onlyOwner {
		Airdrop storage airdrop = airdrops[airdropId];
		if (airdrop.balance < withdrawAmt) {
			revert InvalidWithdrawRequest();
		}
		airdrop.balance = airdrop.balance - withdrawAmt;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
	}

	function getAirdrops(uint256 offset, uint256 limit) external view returns (Airdrop[] memory pageOfAirdrops) {
		uint256 max = offset + limit;
		if (max > airdropCount || limit == 0) {
			max = airdropCount;
		}
		pageOfAirdrops = new Airdrop[](max - offset);
		uint256 total = 0;
		for (uint256 i = offset; i < max; i++) {
			pageOfAirdrops[total] = airdrops[i];
			total++;
		}
		// Dirty hack to cut unused elements off end of return value (from RP)
		// solhint-disable-next-line no-inline-assembly
		assembly {
			mstore(pageOfAirdrops, total)
		}
	}

	/* Merkle Tree Functions */

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
		if (p2c[paddy] != address(0)) {
			revert PAddrAlreadyRegistered();
		}
		if (verify(merkleRoot, paddy, proof)) {
			c2p[msg.sender] = paddy;
			p2c[paddy] = msg.sender;
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
