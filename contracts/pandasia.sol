// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.19;

import {AddressChecksumUtils} from "./AddressChecksumUtils.sol";
import "./SECP256K1.sol";
import {console2} from "forge-std/console2.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

interface Staking {
  function getLastRewardsCycleCompleted(address stakerAddr) external view returns (uint256);
}

// TODO Make this contract a TransparentUpgradeableProxy so we can upgrade without losing state

contract Pandasia is Ownable {
  using SafeERC20 for IERC20;

  error AddressNotEligible();
  error AddressAlreadyClaimed();
  error AirdropExpired();
  error AirdropOutOfFunds();
  error InvalidAddress();
  error InvalidAmount();
  error InvalidWithdrawRequest();
  error PAddrAlreadyRegistered();
  error PAddrNotInValidatorMerkleTree();

  // Storage is sorted for slot optimization
  // _owner address comes from Ownable Slot 0
  uint64 public airdropCount; // counter for AirdropIds (max 18,446,744,073,709,551,615 LFGG)
  uint32 public feePct; // 10_000 = 100% fee charged on funding an airdrop

  mapping(uint64 => Airdrop) public airdrops;
  mapping(uint64 => mapping(address => bool)) public claimed; // airdropIds => users claim status
  mapping(address => uint64[]) public airdropIds; // index of owners => airdropIds

  bytes32 public merkleRoot; // Merkle root defining all verified validator P-chain addresses
  mapping(address => address) public c2p; // c-chain addr => verified p-chain addr
  mapping(address => address) public p2c; // verified p-chain addr => c-chain addr

  // TODO Switch this to gogopool Storage addr, and then get staking contract addr from storage
  address public stakingContract;

  struct Airdrop {
    uint64 id;
    address owner; // account that contributed the funds
    address erc20; // claimable asset
    uint256 balance; // current balance of asset in the airdrop
    bytes32 root; // optional merkle root for this airdrop
    uint256 claimAmount; // claimAmount claimable by each address
    uint32 expiresAt; // time that airdop expires and no further claims can be made
    bool onlyRegistered; // if onlyRegistered=true than addr must be in root AND merkleRoot, else an addr in root OR (previously seen valdiator in pandasia or googpool) is eligble
  }

  /**************************************************************************************************************************************/
  /*** Airdrop Functions                                                                                                              ***/
  /**************************************************************************************************************************************/

  function newAirdrop(bytes32 root, bool onlyRegistered, address erc20, uint256 claimAmount, uint32 expiresAt) external returns (uint64) {
    if (erc20 == address(0)) {
      revert InvalidAddress();
    }

    if (claimAmount == 0) {
      revert InvalidAmount();
    }

    if (expiresAt < block.timestamp) {
      revert AirdropExpired();
    }

    uint64 currentAirdropId = airdropCount;
    airdropCount++;

    Airdrop storage airdrop = airdrops[currentAirdropId];

    // Do we care? Can they set any owner on creation?
    airdrop.id = currentAirdropId;
    airdrop.owner = msg.sender;
    airdrop.erc20 = erc20;
    airdrop.claimAmount = claimAmount;
    airdrop.root = root;
    airdrop.expiresAt = expiresAt;
    airdrop.onlyRegistered = onlyRegistered;

    airdropIds[msg.sender].push(currentAirdropId);

    // TODO Emit event for airdrop creation
    return currentAirdropId;
  }

  function fundAirdrop(uint64 airdropId, uint256 claimAmount) external {
    Airdrop storage airdrop = airdrops[airdropId];
    IERC20 token = IERC20(airdrop.erc20);

    uint256 balance = token.balanceOf(msg.sender);
    if (claimAmount == 0 || balance < claimAmount) {
      revert InvalidAmount();
    }

    uint256 feeAmt = (claimAmount * feePct) / 10_000;
    uint256 fundAmt = claimAmount - feeAmt;
    airdrop.balance = airdrop.balance + fundAmt;
    token.safeTransferFrom(msg.sender, address(this), fundAmt);
  }

  function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) external {
    Airdrop memory airdrop = airdrops[airdropId];
    if (airdrop.owner != msg.sender || airdrop.balance < withdrawAmt || block.timestamp < airdrop.expiresAt) {
      revert InvalidWithdrawRequest();
    }
    airdrop.balance = airdrop.balance - withdrawAmt;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  // do we want to be able to withdraw funding like this?
  function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) external onlyOwner {
    Airdrop memory airdrop = airdrops[airdropId];
    if (airdrop.balance < withdrawAmt) {
      revert InvalidWithdrawRequest();
    }
    airdrop.balance = airdrop.balance - withdrawAmt;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  function getAirdropIds(address owner) public view returns (uint64[] memory) {
    return airdropIds[owner];
  }

  function getAirdrop(uint64 airdropId) external view returns (Airdrop memory) {
    return airdrops[airdropId];
  }

  /**************************************************************************************************************************************/
  /*** Claimant Functions                                                                                                             ***/
  /**************************************************************************************************************************************/

  function claimAirdrop(uint64 airdropId, bytes32[] memory proof) external {
    Airdrop memory airdrop = airdrops[airdropId];
    if (canClaimAirdrop(msg.sender, airdropId, proof)) {
      claimed[airdropId][msg.sender] = true;
      airdrop.balance = airdrop.balance - airdrop.claimAmount;
      IERC20(airdrop.erc20).safeTransfer(msg.sender, airdrop.claimAmount);
    }

    // emit some event
  }

  function canClaimAirdrop(address cChainAddr, uint64 airdropId, bytes32[] memory proof) public view returns (bool) {
    if (claimed[airdropId][cChainAddr]) {
      revert AddressAlreadyClaimed();
    }

    Airdrop memory airdrop = airdrops[airdropId];

    if (block.timestamp > airdrop.expiresAt) {
      revert AirdropExpired();
    }

    if (airdrop.balance < airdrop.claimAmount) {
      revert AirdropOutOfFunds();
    }

    bool isInAirdropRoot = verify(airdrop.root, cChainAddr, proof);

    // this should be isKnownValidator
    bool isKnownValidator = isRegisteredValidator(cChainAddr) || isMinipoolOperator(cChainAddr);
    bool isEligible;
    if (airdrop.onlyRegistered) {
      if (isKnownValidator && isInAirdropRoot) {
        isEligible = true;
      }
    } else {
      if (isKnownValidator || isInAirdropRoot) {
        isEligible = true;
      }
    }

    if (!isEligible) {
      revert AddressNotEligible();
    }

    // I wonder if this should return true or false rather than that last revert
    return true;
  }

  function hasClaimed(uint64 airdropId, address addr) public view returns (bool) {
    return claimed[airdropId][addr] == true;
  }

  function isRegisteredValidator(address addr) public view returns (bool) {
    return c2p[addr] != address(0);
  }

  function isMinipoolOperator(address addr) public view returns (bool) {
    // TODO verify this is going to work
    return Staking(stakingContract).getLastRewardsCycleCompleted(addr) > 0;
  }

  /**************************************************************************************************************************************/
  /*** Merkle Tree Functions                                                                                                          ***/
  /**************************************************************************************************************************************/

  // Sign C-chain address in mixedcase hex format with P-chain addr on wallet.avax.network
  function registerPChainAddr(uint8 v, bytes32 r, bytes32 s, bytes32[] memory proof) external {
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

  // TODO: Test function, remove before going to production
  function unregisterPChainAddr() external {
    address paddr = c2p[msg.sender];
    delete c2p[msg.sender];
    delete p2c[paddr];
  }

  // Uses about 35K gas for a large proof
  function verify(bytes32 root, address account, bytes32[] memory proof) public pure returns (bool) {
    // I want to short circuit this root check somewhere else.
    // it's weird to me that the negative case is a 0 and onyl checked
    // when verification happens.
    return proof.length > 0 && account != address(0) && root != bytes32(0) && MerkleProof.verify(proof, root, _leaf(account));
  }

  function _leaf(address account) internal pure returns (bytes32) {
    return keccak256(bytes.concat(keccak256(abi.encode(account))));
  }

  // TODO: Test function, remove before going to production
  function recoverMessage(uint8 v, bytes32 r, bytes32 s) external view returns (address) {
    bytes32 msgHash = hashChecksummedMessage(msg.sender);
    (uint256 x, uint256 y) = SECP256K1.recover(uint256(msgHash), v, uint256(r), uint256(s));
    address paddy = pubKeyBytesToAvaAddressBytes(x, y);
    return paddy;
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

  /**************************************************************************************************************************************/
  /*** Owner Functions                                                                                                                ***/
  /**************************************************************************************************************************************/

  function withdrawFees(uint64 airdropId) external onlyOwner {
    Airdrop memory airdrop = airdrops[airdropId];
    uint256 fees = IERC20(airdrop.erc20).balanceOf(address(this)) - airdrop.balance;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, fees);
  }

  function getAirdrops(uint64 offset, uint64 limit) external returns (Airdrop[] memory pageOfAirdrops) {
    uint64 max = offset + limit;
    if (max > airdropCount || limit == 0) {
      max = airdropCount;
    }
    pageOfAirdrops = new Airdrop[](max - offset);
    uint64 total = 0;
    for (uint64 i = offset; i < max; i++) {
      Airdrop memory airdrop = airdrops[i];
      logAirdrop(airdrop);

      logAirdrop(airdrops[i]);

      pageOfAirdrops[total] = airdrop;
      total++;
    }
    // Dirty hack to cut unused elements off end of return value (from RP)
    // solhint-disable-next-line no-inline-assembly
    assembly {
      mstore(pageOfAirdrops, total)
    }
  }

  function logAirdrop(Pandasia.Airdrop memory airdrop) internal virtual {
    console2.log(airdrop.balance);
    console2.log(airdrop.claimAmount);
    console2.log(airdrop.erc20);
    console2.log(airdrop.expiresAt);
    console2.log(airdrop.onlyRegistered);
    console2.log(airdrop.owner);
    console2.logBytes32(airdrop.root);
  }

  function setMerkleRoot(bytes32 root) external onlyOwner {
    merkleRoot = root;
  }

  function setFee(uint32 fee) external onlyOwner {
    feePct = fee;
  }

  function setStakingContract(address addr) external onlyOwner {
    stakingContract = addr;
  }
}
