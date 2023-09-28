// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.19;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {Pandasia} from "./pandasia.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract Airdrop is Initializable {
  using SafeERC20 for IERC20;
  address owner; // account that contributed the funds
  address erc20; // claimable asset
  uint256 balance; // current balance of asset in the airdrop
  bytes32 root; // optional merkle root for this airdrop
  uint256 claimAmount; // claimAmount claimable by each address
  uint32 expires; // time that airdop expires and no further claims can be made
  bool onlyRegistered; // if onlyRegistered=true than addr must be a registered validator AND merkleRoot,

  uint32 public feePct; // 10_000 = 100% fee charged on funding an airdrop

  mapping(address => bool) public claimed;

  address public stakingContract;
  address public pandasiaAddress;

  error AddressNotEligible();
  error AddressAlreadyClaimed();
  error AirdropExpired();
  error AirdropOutOfFunds();
  error InvalidAddress();
  error InvalidAmount();
  error InvalidWithdrawRequest();
  error PAddrAlreadyRegistered();
  error PAddrNotInValidatorMerkleTree();

  function initialize(
    address _owner,
    bytes32 _root,
    bool _onlyRegistered,
    address _erc20,
    uint256 _claimAmount,
    uint32 _expires,
    uint32 _feePct,
    address _stakingContract,
    address _pandasia
  ) public initializer {
    owner = _owner;
    root = _root;
    onlyRegistered = _onlyRegistered;
    erc20 = _erc20;
    claimAmount = _claimAmount;
    expires = _expires;
    feePct = _feePct;

    stakingContract = _stakingContract;
    pandasiaAddress = _pandasia;
  }

  constructor() {
    _disableInitializers();
  }

  /**************************************************************************************************************************************/
  /*** Airdrop Functions                                                                                                              ***/
  /**************************************************************************************************************************************/

  function fundAirdrop(uint256 amount) external {
    IERC20 token = IERC20(erc20);

    uint256 tokenBalance = token.balanceOf(msg.sender);
    if (amount == 0 || tokenBalance < amount) {
      revert InvalidAmount();
    }

    uint256 feeAmt = (amount * feePct) / 10_000;
    uint256 fundAmt = amount - feeAmt;
    balance = balance + fundAmt;
    token.safeTransferFrom(msg.sender, address(this), fundAmt);
  }

  function withdrawFunding(uint256 withdrawAmt) external {
    if (owner != msg.sender || balance < withdrawAmt || block.timestamp < expires) {
      revert InvalidWithdrawRequest();
    }
    balance = balance - withdrawAmt;
    IERC20(erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  // do we want to be able to withdraw funding like this?
  function emergencyWithdraw(uint256 withdrawAmt) external {
    if (balance < withdrawAmt) {
      revert InvalidWithdrawRequest();
    }
    balance = balance - withdrawAmt;
    IERC20(erc20).safeTransfer(msg.sender, withdrawAmt);
  }

  /**************************************************************************************************************************************/
  /*** Claimant Functions                                                                                                             ***/
  /**************************************************************************************************************************************/

  function claimAirdrop(bytes32[] memory proof) external {
    if (canClaimAirdrop(msg.sender, proof)) {
      claimed[msg.sender] = true;
      balance = balance - claimAmount;
      IERC20(erc20).safeTransfer(msg.sender, claimAmount);
    }

    // emit event
  }

  function canClaimAirdrop(address cChainAddr, bytes32[] memory proof) public view returns (bool) {
    if (claimed[cChainAddr]) {
      revert AddressAlreadyClaimed();
    }

    if (block.timestamp > expires) {
      revert AirdropExpired();
    }

    if (balance < claimAmount) {
      revert AirdropOutOfFunds();
    }

    bool isInAirdropRoot = Pandasia(pandasiaAddress).verify(root, cChainAddr, proof);

    bool isKnownValidator = Pandasia(pandasiaAddress).isRegisteredValidator(cChainAddr) ||
      Pandasia(pandasiaAddress).isMinipoolOperator(cChainAddr);
    bool isEligible;
    if (onlyRegistered) {
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

  function hasClaimed(address addr) public view returns (bool) {
    return claimed[addr] == true;
  }

  // owner functions
  function withdrawFees() external {
    uint256 fees = IERC20(erc20).balanceOf(address(this)) - balance;
    IERC20(erc20).safeTransfer(msg.sender, fees);
  }
}
