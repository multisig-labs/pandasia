// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.19;

import {Airdrop} from "./Airdrop.sol";
import {AddressChecksumUtils} from "./AddressChecksumUtils.sol";
import {SECP256K1} from "./SECP256K1.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Clones} from "@openzeppelin/contracts/proxy/Clones.sol";

interface Staking {
  function getLastRewardsCycleCompleted(address stakerAddr) external view returns (uint256);
}

contract Pandasia is Ownable, Initializable {
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

  address public airdropImplementation;

  // Storage is sorted for slot optimization
  // _owner address comes from Ownable Slot 0
  uint64 public airdropCount; // counter for AirdropIds (max 18,446,744,073,709,551,615 LFGG)

  mapping(uint64 => address) public airdrops;
  mapping(address => uint64[]) public airdropIds; // index of owners => airdropIds

  bytes32 public merkleRoot; // Merkle root defining all verified validator P-chain addresses
  mapping(address => address) public c2p; // c-chain addr => verified p-chain addr
  mapping(address => address) public p2c; // verified p-chain addr => c-chain addr

  uint32 public feePct;
  address public stakingContract;

  function initialize() public initializer {
    airdropImplementation = address(new Airdrop());
  }

  constructor() {
    _disableInitializers();
  }

  /**************************************************************************************************************************************/
  /*** Airdrop Functions                                                                                                              ***/
  /**************************************************************************************************************************************/

  function newAirdrop(
    bytes32 root,
    bool onlyRegistered,
    address erc20,
    uint256 claimAmount,
    uint32 expires
  ) external returns (uint64) {
    if (erc20 == address(0)) {
      revert InvalidAddress();
    }

    if (claimAmount == 0) {
      revert InvalidAmount();
    }

    if (expires < block.timestamp) {
      revert AirdropExpired();
    }

    address clone = Clones.clone(airdropImplementation);

    Airdrop(clone).initialize(
      _msgSender(),
      root,
      onlyRegistered,
      erc20,
      claimAmount,
      expires,
      feePct,
      stakingContract,
      address(this)
    );

    airdropCount++;
    airdrops[airdropCount] = clone;
    airdropIds[msg.sender].push(airdropCount);

    // TODO Emit event for airdrop creation
    return airdropCount;
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
    return
      proof.length > 0 &&
      account != address(0) &&
      root != bytes32(0) &&
      MerkleProof.verify(proof, root, _leaf(account));
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

  /********************************************************************************************************************/
  /*** Owner Functions                                                                                             ****/
  /********************************************************************************************************************/

  function setMerkleRoot(bytes32 root) external onlyOwner {
    merkleRoot = root;
  }

  function setFeePct(uint32 _feePct) external onlyOwner {
    feePct = _feePct;
  }

  function setStakingContract(address _stakingContract) external onlyOwner {
    stakingContract = _stakingContract;
  }

  /********************************************************************************************************************/
  /*** Read Functions                                                                                               ***/
  /********************************************************************************************************************/

  function isRegisteredValidator(address addr) public view returns (bool) {
    return c2p[addr] != address(0);
  }

  function isMinipoolOperator(address addr) public view returns (bool) {
    // TODO verify this is going to work
    return Staking(stakingContract).getLastRewardsCycleCompleted(addr) > 0;
  }

  function getAirdrops(uint64 offset, uint64 limit) external view returns (address[] memory pageOfAirdrops) {
    uint64 max = offset + limit;
    if (max > airdropCount || limit == 0) {
      max = airdropCount;
    }
    pageOfAirdrops = new address[](max - offset);
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
}
