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
	error AirdropExpired();
	error AirdropOutOfFunds();
	error AddressNotEligible();
	error AddressAlreadyClaimed();
	error InvalidWithdrawRequest();

	// Storage is sorted for slot optimization
	// _owner address comes from Ownable Slot 0
	uint64 public airdropCount; // counter for AirdropIds (max 18,446,744,073,709,551,615 LFGG)
	uint32 public feePct; // 10_000 = 100% fee charged on funding an airdrop
	mapping(uint64 => Airdrop) public airdrops;
	mapping(uint64 => mapping(address => bool)) public claimed;
	mapping(address => uint64[]) public airdropIds; // index of owners => airdropIds

	bytes32 public validatorRoot; // Merkle root defining all verified validator P-chain addresses
	mapping(address => address) public c2p; // maps c-chain addr => verified p-chain addr
	mapping(address => address) public p2c; // maps verified p-chain addr => c-chain addr

	struct Airdrop {
		bytes32 root; // optional merkle root that applies for this airdrop
		uint256 balance; // current balance of asset
		uint256 amount; // amount claimable by each address
		address owner; // account that contributed the funds
		address erc20; // claimable asset
		uint32 expires; // time that airdop expires and no further claims can be made
		bool union; // if union=t, then an addr in root OR validatorRoot is eligble, else addr must be in root AND validatorRoot
	}

	// Pandasia DAO will update this root (daily?) with a new tree of all validators
	function setValidatorRoot(bytes32 root) external onlyOwner {
		validatorRoot = root;
	}

	function setFee(uint32 fee) external onlyOwner {
		feePct = fee;
	}

	function newAirdrop(bytes32 root, bool union, address erc20, uint256 amount, uint32 expires) external returns (uint64) {
		if (erc20 == address(0)) {
			revert InvalidAddress();
		}
		if (amount == 0) {
			revert InvalidAmount();
		}
		if (expires < block.timestamp) {
			revert AirdropExpired();
		}
		airdropCount++;
		Airdrop storage airdrop = airdrops[airdropCount];
		airdrop.root = root;
		airdrop.union = union;
		// Do we care? Can they set any owner on creation?
		airdrop.owner = msg.sender;
		airdrop.erc20 = erc20;
		airdrop.amount = amount;
		airdrop.expires = expires;

		airdropIds[msg.sender].push(airdropCount);
		return airdropCount;
	}

	function fundAirdrop(uint64 airdropId, uint256 amount) external {
		Airdrop storage airdrop = airdrops[airdropId];
		IERC20 token = IERC20(airdrop.erc20);
		uint256 balance = token.balanceOf(msg.sender);
		if (amount == 0 || balance < amount) {
			revert InvalidAmount();
		}

		uint256 feeAmt = (amount * feePct) / 10_000;
		uint256 fundAmt = amount - feeAmt;
		airdrop.balance = airdrop.balance + fundAmt;
		token.safeTransferFrom(msg.sender, address(this), fundAmt);
	}

	function canClaimAirdrop(address addr, uint64 airdropId, bytes32[] memory proof) public view returns (bool) {
		if (claimed[airdropId][addr]) {
			revert AddressAlreadyClaimed();
		}

		Airdrop memory airdrop = airdrops[airdropId];

		if (block.timestamp > airdrop.expires) {
			revert AirdropExpired();
		}

		if (airdrop.balance < airdrop.amount) {
			revert AirdropOutOfFunds();
		}

		bool isInAirdropRoot = verify(airdrop.root, addr, proof);
		bool isValidator = isRegisteredValidator(addr) || isMinipoolOperator(addr);
		bool isEligible;

		if (airdrop.union) {
			if (isValidator || isInAirdropRoot) {
				isEligible = true;
			}
		} else {
			if (isValidator && isInAirdropRoot) {
				isEligible = true;
			}
		}

		if (!isEligible) {
			revert AddressNotEligible();
		}

		return true;
	}

	// TODO measure gas costs of this, do we need to optimize?
	function claimAirdrop(uint64 airdropId, bytes32[] memory proof) external {
		Airdrop memory airdrop = airdrops[airdropId];
		if (canClaimAirdrop(msg.sender, airdropId, proof)) {
			claimed[airdropId][msg.sender] = true;
			airdrop.balance = airdrop.balance - airdrop.amount;
			IERC20(airdrop.erc20).safeTransfer(msg.sender, airdrop.amount);
		}
	}

	function hasClaimed(uint64 airdropId, address addr) public view returns (bool) {
		return claimed[airdropId][addr] == true;
	}

	function isRegisteredValidator(address addr) public view returns (bool) {
		return c2p[addr] != address(0);
	}

	function isMinipoolOperator(address addr) public view returns (bool) {
		// Can use Staking.sol getLastRewardsCycleCompleted > 0 I think?
		return false;
	}

	function getAirdropIds(address owner) public view returns (uint64[] memory) {
		return airdropIds[owner];
	}

	function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) external {
		Airdrop memory airdrop = airdrops[airdropId];
		if (airdrop.owner != msg.sender || airdrop.balance < withdrawAmt || block.timestamp < airdrop.expires) {
			revert InvalidWithdrawRequest();
		}
		airdrop.balance = airdrop.balance - withdrawAmt;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
	}

	function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) external onlyOwner {
		Airdrop memory airdrop = airdrops[airdropId];
		if (airdrop.balance < withdrawAmt) {
			revert InvalidWithdrawRequest();
		}
		airdrop.balance = airdrop.balance - withdrawAmt;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
	}

	function withdrawFees(uint64 airdropId) external onlyOwner {
		Airdrop memory airdrop = airdrops[airdropId];
		uint256 fees = IERC20(airdrop.erc20).balanceOf(address(this)) - airdrop.balance;
		IERC20(airdrop.erc20).safeTransfer(msg.sender, fees);
	}

	function getAirdrops(uint64 offset, uint64 limit) external view returns (Airdrop[] memory pageOfAirdrops) {
		uint64 max = offset + limit;
		if (max > airdropCount || limit == 0) {
			max = airdropCount;
		}
		pageOfAirdrops = new Airdrop[](max - offset);
		uint64 total = 0;
		for (uint64 i = offset; i < max; i++) {
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

	// Sign C-chain address in mixedcase hex format with P-chain addr on wallet.avax.network
	function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] memory proof) external {
		bytes32 msgHash = hashChecksummedMessage(msg.sender);
		(uint256 x, uint256 y) = SECP256K1.recover(uint256(msgHash), v, uint256(r), uint256(s));
		address paddy = pubKeyBytesToAvaAddressBytes(x, y);
		if (p2c[paddy] != address(0)) {
			revert PAddrAlreadyRegistered();
		}
		if (verify(validatorRoot, paddy, proof)) {
			c2p[msg.sender] = paddy;
			p2c[paddy] = msg.sender;
		} else {
			revert PAddrNotInValidatorMerkleTree();
		}
	}

	// Given an address, convert to its checksummed string (mixedcase) format, and hash a message like the avalanche wallet would do
	function hashChecksummedMessage(address addr) public pure returns (bytes32) {
		bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
		// len of an ascii addr is 42 bytes
		uint32 addrLen = 42;
		string memory addrStr = AddressChecksumUtils.getChecksum(addr);
		return sha256(abi.encodePacked(header, addrLen, "0x", addrStr));
	}

	function pubKeyBytesToAvaAddressBytes(uint256 x, uint256 y) public pure returns (address) {
		// Calc format prefix for compressed pub key
		bytes1 format = (y % 2 == 0) ? bytes1(0x02) : bytes1(0x03);
		bytes memory pubKey = abi.encodePacked(format, x);
		bytes32 pubKeySha = sha256(pubKey);
		return address(ripemd160(abi.encodePacked(pubKeySha)));
	}

	function verify(bytes32 root, address account, bytes32[] memory proof) public pure returns (bool) {
		return proof.length > 0 && account != address(0) && root != bytes32(0) && MerkleProof.verify(proof, root, _leaf(account));
	}

	function _leaf(address account) internal pure returns (bytes32) {
		return keccak256(bytes.concat(keccak256(abi.encode(account))));
	}
}
