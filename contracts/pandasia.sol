// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.19;

import {AddressChecksumUtils} from "./AddressChecksumUtils.sol";
import "./SECP256K1.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

interface Staking {
  function getLastRewardsCycleCompleted(address stakerAddr) external view returns (uint256);
}

interface Storage {
  function getAddress(bytes32 key) external view returns (address);
}

contract Pandasia is OwnableUpgradeable {
  using SafeERC20 for IERC20;

  error AddressNotEligible();
  error AddressAlreadyClaimed();
  error AirdropExpired();
  error AirdropNotStarted();
  error AirdropOutOfFunds();
  error AirdropStillActive();
  error InvalidAddress();
  error InvalidAmount();
  error InvalidWithdrawRequest();
  error NotOwner();
  error PAddrAlreadyRegistered();
  error PAddrNotInMerkleTree();
  error ZeroAmount();

  event AirdropCreated(uint64 indexed id);
  event AirdropClaimed(uint64 indexed id, address indexed claimant);

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

  address public storageContract;

  struct Airdrop {
    uint64 id;
    address owner; // account that contributed the funds
    address erc20; // claimable asset
    uint256 balance; // current balance of asset in the airdrop
    bytes32 customRoot; // optional merkle root for this airdrop
    uint256 claimAmount; // claimAmount claimable by each address
    uint64 startsAt; // time that airdop starts and claims can be made
    uint64 expiresAt; // time that airdop expires and no further claims can be made
  }

  constructor() {
    _disableInitializers();
  }

  function initialize() public initializer {
    __Ownable_init(msg.sender);
  }

  /**************************************************************************************************************************************/
  /*** Airdrop Functions                                                                                                              ***/
  /**************************************************************************************************************************************/

  function newAirdrop(bytes32 customRoot, address erc20, uint256 claimAmount, uint64 startsAt, uint64 expiresAt) external returns (uint64) {
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
    airdrop.customRoot = customRoot;
    airdrop.startsAt = startsAt;
    airdrop.expiresAt = expiresAt;

    airdropIds[msg.sender].push(currentAirdropId);

    emit AirdropCreated(currentAirdropId);

    return currentAirdropId;
  }

  function fundAirdrop(uint64 airdropId, uint256 fundAmount) external {
    Airdrop storage airdrop = airdrops[airdropId];
    IERC20 token = IERC20(airdrop.erc20);

    uint256 balance = token.balanceOf(msg.sender);

    if (fundAmount == 0) {
      revert ZeroAmount();
    }

    if (balance < fundAmount) {
      revert InvalidAmount();
    }

    uint256 feeAmt = (fundAmount * feePct) / 10_000;
    uint256 fundAmt = fundAmount - feeAmt;
    airdrop.balance = airdrop.balance + fundAmt;
    token.safeTransferFrom(msg.sender, address(this), fundAmt + feeAmt);
  }

  function withdrawFunding(uint64 airdropId, uint256 withdrawAmt) external {
    Airdrop memory airdrop = airdrops[airdropId];
    if (airdrop.owner != msg.sender) {
      revert NotOwner();
    }

    if (airdrop.balance < withdrawAmt) {
      revert InvalidAmount();
    }

    if (block.timestamp < airdrop.expiresAt) {
      revert AirdropStillActive();
    }
    airdrop.balance = airdrop.balance - withdrawAmt;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  /**************************************************************************************************************************************/
  /*** Claimant Functions                                                                                                             ***/
  /**************************************************************************************************************************************/

  function claimAirdrop(uint64 airdropId, bytes32[] memory proof) external {
    Airdrop storage airdrop = airdrops[airdropId];

    if (!canClaimAirdrop(msg.sender, airdropId, proof)) {
      return;
    }

    claimed[airdropId][msg.sender] = true;
    airdrop.balance = airdrop.balance - airdrop.claimAmount;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, airdrop.claimAmount);
    emit AirdropClaimed(airdropId, msg.sender);
  }

  function canClaimAirdrop(address cChainAddr, uint64 airdropId, bytes32[] memory proof) public view returns (bool) {
    if (claimed[airdropId][cChainAddr]) {
      revert AddressAlreadyClaimed();
    }

    Airdrop memory airdrop = airdrops[airdropId];

    if (block.timestamp < airdrop.startsAt) {
      revert AirdropNotStarted();
    }

    if (block.timestamp > airdrop.expiresAt) {
      revert AirdropExpired();
    }

    if (airdrop.balance < airdrop.claimAmount) {
      revert AirdropOutOfFunds();
    }

    address addr = c2p[cChainAddr];
    if (addr == address(0)) {
      addr = cChainAddr;
    }

    bool isKnownValidator = isRegisteredValidator(cChainAddr) || isMinipoolOperator(cChainAddr);

    bool isEligible;
    if (airdrop.customRoot != bytes32(0x0)) {
      isEligible = isKnownValidator && verify(airdrop.customRoot, addr, proof);
    } else {
      isEligible = isKnownValidator;
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
    bytes32 key = keccak256(abi.encodePacked("contract.address", "Staking"));
    address stakingContract = Storage(storageContract).getAddress(key);
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
      revert PAddrNotInMerkleTree();
    }
  }

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

  function emergencyWithdraw(uint64 airdropId, uint256 withdrawAmt) external onlyOwner {
    Airdrop memory airdrop = airdrops[airdropId];
    if (airdrop.balance < withdrawAmt) {
      revert InvalidWithdrawRequest();
    }
    airdrop.balance = airdrop.balance - withdrawAmt;
    IERC20(airdrop.erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  function setMerkleRoot(bytes32 root) external onlyOwner {
    merkleRoot = root;
  }

  function setFee(uint32 fee) external onlyOwner {
    feePct = fee;
  }

  function setStorageContract(address addr) external onlyOwner {
    storageContract = addr;
  }

  /**************************************************************************************************************************************/
  /*** View Functions                                                                                                                 ***/
  /**************************************************************************************************************************************/

  function getAirdropIds(address owner) public view returns (uint64[] memory) {
    return airdropIds[owner];
  }

  function getAirdrop(uint64 airdropId) external view returns (Airdrop memory) {
    return airdrops[airdropId];
  }

  function getAirdrops(uint64 offset, uint64 limit) external view returns (Airdrop[] memory pageOfAirdrops) {
    uint64 max = offset + limit;
    if (max > airdropCount || limit == 0) {
      max = airdropCount;
    }
    pageOfAirdrops = new Airdrop[](max - offset);
    uint64 total = 0;
    for (uint64 i = offset; i < max; i++) {
      Airdrop memory airdrop = airdrops[i];
      pageOfAirdrops[total] = airdrop;
      total++;
    }
    // Dirty hack to cut unused elements off end of return value (from RP)
    // solhint-disable-next-line no-inline-assembly
    assembly {
      mstore(pageOfAirdrops, total)
    }
  }
}
