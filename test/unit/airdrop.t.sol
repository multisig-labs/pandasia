pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ERC20Mock} from "@openzeppelin/contracts/mocks/token/ERC20Mock.sol";
import {console2} from "forge-std/console2.sol";
import {Pandasia} from "../../contracts/Pandasia.sol";
import {StakingMock} from "./mocks/StakingMock.sol";
import {StorageMock} from "./mocks/StorageMock.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AirdropTest is Test {
  ERC20Mock public erc20;
  Pandasia public pandasia;
  StorageMock public storageContract;
  StakingMock public stakingContract;

  address public deployer; // Owner of Pandasia contract
  address public airdropOwner; // Owner of individual airdrops
  address public validator; // eligible via validatorRoot
  address public minipoolOperator; // eligible via staking contract
  address public otherUser; // eligible via otherProof
  address public notEligibleUser; // not eligible in any way

  bytes32[] public emptyProof;
  bytes32[] public validatorProof;
  bytes32[] public otherProof;
  bytes32 public validatorRoot;
  bytes32 public otherRoot;
  uint64 public blockHeight;

  function setUp() public {
    Pandasia pandasiaImpl = new Pandasia();

    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(address(pandasiaImpl), address(this), bytes(""));
    pandasia = Pandasia(payable(pandasiaProxy));
    pandasia.initialize();

    pandasia.grantRole(pandasia.ROOT_UPDATER(), address(this));
    pandasia.grantRole(pandasia.AIRDROP_ADMIN(), address(this));

    deployer = getActor("deployer");
    airdropOwner = getActor("airdropOwner");
    pandasia.grantRole(pandasia.AIRDROP_ADMIN(), address(airdropOwner));

    minipoolOperator = getActor("minipoolOperator");
    otherUser = getActor("otherUser", 0x0000000000000000000000000000000000000001);
    validator = getActor("validator", 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4);

    emptyProof = new bytes32[](0);

    otherRoot = bytes32(0x2e1dea9890e94d280361af414f0696bc7ba251d4e52a5f786d15629d5185a89c);
    otherProof = new bytes32[](1); // for addr 0x0000000000000000000000000000000000000001
    otherProof[0] = bytes32(0x9fec67521532e3df3ca2461c12c79c2b89e8f633311a6b525fb8488fbcd1d177);

    validatorRoot = bytes32(0x1733170f5a465a52692730efa67c11a3c9b1208a5acbe833057fac165ce6947b);
    // address paddy = address(0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02);
    validatorProof = new bytes32[](1);
    validatorProof[0] = bytes32(0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f);

    blockHeight = uint64(1211);

    stakingContract = new StakingMock();
    stakingContract.setLastRewardsCycleCompleted(minipoolOperator, 1);

    storageContract = new StorageMock();
    storageContract.setAddress(keccak256(abi.encodePacked("contract.address", "Staking")), address(stakingContract));

    pandasia.setMerkleRoot(validatorRoot, blockHeight);
    pandasia.setStorageContract(address(storageContract));

    pandasia.transferOwnership(deployer);

    // Signature generated on wallet.avax.network
    uint8 v = 0;
    bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
    bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

    vm.prank(validator);
    pandasia.registerPChainAddr(v, r, s, validatorProof);

    erc20 = new ERC20Mock();
  }

  /**************************************************************************************************************************************/
  /*** New Airdrop Tests                                                                                                              ***/
  /**************************************************************************************************************************************/

  function testNewAirdropNoAdditionalMerkeRoot() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    // Not eligible user can't claim
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(notEligibleUser, id, emptyProof);

    // Other user can't claim with either proof
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(otherUser, id, emptyProof);
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(otherUser, id, otherProof);

    // Validator can claim
    assertTrue(pandasia.canClaimAirdrop(validator, id, emptyProof));

    // minipool operator is eligible
    assertTrue(pandasia.canClaimAirdrop(minipoolOperator, id, emptyProof));
  }

  function testNewAirdropWithCustomMerkleRoot() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    // Pass a custom root to the airdrop, in this case validatorRoot
    uint64 id = pandasia.newAirdrop(validatorRoot, address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    // Not eligible users can't claim
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(notEligibleUser, id, emptyProof);

    // Other user without proof can't claim
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(otherUser, id, emptyProof);

    // Other user is also not registered and can't claim
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(otherUser, id, otherProof);

    // Minipool operator is not in merkle root and can't claim
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(minipoolOperator, id, emptyProof);

    // Validator is registered and in merkle root, thus they can claim
    assertTrue(pandasia.canClaimAirdrop(validator, id, validatorProof));
  }

  function testNewAirdropInvalidERC20() public {
    address validErc20 = address(erc20);
    address invalidErc20 = address(0);

    vm.expectRevert(Pandasia.InvalidAddress.selector);
    pandasia.newAirdrop(bytes32(0), invalidErc20, 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    uint64 id = pandasia.newAirdrop(bytes32(0), validErc20, 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);
    assertEq(airdrop.erc20, validErc20);
  }

  function testNewAirdropInvalidAmount() public {
    uint256 validAmount = 10 ether;
    uint256 invalidAmount = 0 ether;

    vm.expectRevert(Pandasia.InvalidAmount.selector);
    pandasia.newAirdrop(bytes32(0), address(erc20), invalidAmount, uint32(block.timestamp), uint64(block.timestamp + 1000));

    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), validAmount, uint32(block.timestamp), uint64(block.timestamp + 1000));
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);
    assertEq(airdrop.claimAmount, validAmount);
  }

  function testNewAirdropExpired() public {
    uint64 validExpiresAt = uint64(block.timestamp + 1000);
    uint64 invalidExpiresAt = uint64(block.timestamp - 1);

    vm.expectRevert(Pandasia.AirdropExpired.selector);
    pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), invalidExpiresAt);

    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), validExpiresAt);
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);
    assertEq(airdrop.expiresAt, validExpiresAt);
  }

  function testNewAirdropZeroCustomRoot() public {
    uint256 totalFundingAmt = 50 ether;

    // Verify that neither 0x0 or 0 triggers custom root check
    assertEq(bytes32(0), bytes32(0x0));

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    // Validator is registered and in merkle root, thus they can claim
    assertTrue(pandasia.canClaimAirdrop(validator, id, validatorProof));

    vm.startPrank(airdropOwner);
    id = pandasia.newAirdrop(bytes32(0x0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    // Validator is registered and in merkle root, thus they can claim
    assertTrue(pandasia.canClaimAirdrop(validator, id, validatorProof));
  }

  function testNewAirdropNotAdmin() public {
    address user1 = address(0x01);

    vm.startPrank(user1);
    bytes4 selector = bytes4(keccak256("AccessControlUnauthorizedAccount(address,bytes32)"));
    vm.expectRevert(abi.encodeWithSelector(selector, user1, pandasia.AIRDROP_ADMIN()));
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));
    vm.stopPrank();

    pandasia.grantRole(pandasia.AIRDROP_ADMIN(), user1);

    vm.prank(user1);
    id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));
  }

  function testDeleteAirdrop() public {
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);

    assertEq(airdrop.erc20, address(erc20));

    pandasia.deleteAirdrop(id);

    airdrop = pandasia.getAirdrop(id);
    assertEq(airdrop.erc20, address(0));
  }

  function testDeleteAirdropNotAdmin() public {
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);

    assertEq(airdrop.erc20, address(erc20));

    address user1 = address(0x01);

    vm.startPrank(user1);
    bytes4 selector = bytes4(keccak256("AccessControlUnauthorizedAccount(address,bytes32)"));
    vm.expectRevert(abi.encodeWithSelector(selector, user1, pandasia.AIRDROP_ADMIN()));
    pandasia.deleteAirdrop(id);
    vm.stopPrank();

    pandasia.grantRole(pandasia.AIRDROP_ADMIN(), user1);
    vm.prank(user1);
    pandasia.deleteAirdrop(id);

    airdrop = pandasia.getAirdrop(id);
    assertEq(airdrop.erc20, address(0));
  }

  function testDeleteAirdropWithFunding() public {
    uint256 perClaimAmt = 10 ether;
    uint64 id = createFundedAirdrop(perClaimAmt);

    vm.expectRevert(Pandasia.AirdropStillHasFunding.selector);
    pandasia.deleteAirdrop(id);

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);

    vm.prank(deployer);
    pandasia.emergencyWithdraw(id);

    pandasia.deleteAirdrop(id);
    airdrop = pandasia.getAirdrop(id);

    assertEq(airdrop.claimAmount, 0);
  }

  /**************************************************************************************************************************************/
  /*** Can Claim Tests                                                                                                                ***/
  /**************************************************************************************************************************************/

  function testCanClaimAirdropNotStarted() public {
    uint256 totalFundingAmt = 50 ether;
    uint256 futureTime = 100;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp + futureTime), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    vm.expectRevert(Pandasia.AirdropNotStarted.selector);
    pandasia.canClaimAirdrop(validator, id, emptyProof);

    vm.warp(block.timestamp + futureTime);
    assertTrue(pandasia.canClaimAirdrop(validator, id, emptyProof));
  }

  function testCanClaimAirdropExpired() public {
    uint256 totalFundingAmt = 50 ether;
    uint256 endTime = 100;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + endTime));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    vm.warp(block.timestamp + endTime + 1);

    vm.expectRevert(Pandasia.AirdropExpired.selector);
    pandasia.canClaimAirdrop(validator, id, emptyProof);
  }

  function testCanClaimAirdropOutOfFunds() public {
    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    vm.expectRevert(Pandasia.AirdropOutOfFunds.selector);
    pandasia.canClaimAirdrop(validator, id, emptyProof);

    uint256 totalFundingAmt = 50 ether;
    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    pandasia.canClaimAirdrop(validator, id, emptyProof);
  }

  function testCanOnlyClaimOnce() public {
    uint256 perClaimAmt = 10 ether;
    uint64 id = createFundedAirdrop(perClaimAmt);

    uint256 preBalance = erc20.balanceOf(validator);
    vm.prank(validator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(validator), preBalance + perClaimAmt);

    vm.prank(validator);
    vm.expectRevert(Pandasia.AddressAlreadyClaimed.selector);
    pandasia.claimAirdrop(id, emptyProof);
  }

  function testCanClaimAndClaimMatch() public {
    uint256 perClaimAmt = 10 ether;
    uint64 id = createFundedAirdrop(perClaimAmt);

    assertTrue(pandasia.canClaimAirdrop(validator, id, emptyProof));

    uint256 preBalance = erc20.balanceOf(validator);
    vm.prank(validator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(validator), preBalance + perClaimAmt);

    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(notEligibleUser, id, emptyProof);

    vm.prank(notEligibleUser);
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.claimAirdrop(id, emptyProof);
  }

  function testHasClaimed() public {
    uint256 perClaimAmt = 10 ether;
    uint64 id = createFundedAirdrop(perClaimAmt);

    bool hasClaimed = pandasia.hasClaimed(id, validator);
    assertFalse(hasClaimed);

    uint256 preBalance = erc20.balanceOf(validator);
    vm.prank(validator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(validator), preBalance + perClaimAmt);

    hasClaimed = pandasia.hasClaimed(id, validator);
    assertTrue(hasClaimed);
  }

  /**************************************************************************************************************************************/
  /*** Funding Airdrop Tests                                                                                                                ***/
  /**************************************************************************************************************************************/

  function testFundAirdropInvalidAmount() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);

    vm.expectRevert(Pandasia.InvalidAmount.selector);
    pandasia.fundAirdrop(id, totalFundingAmt + 10);

    vm.expectRevert(Pandasia.ZeroAmount.selector);
    pandasia.fundAirdrop(id, 0);

    pandasia.fundAirdrop(id, totalFundingAmt);
    assertEq(erc20.balanceOf(address(pandasia)), totalFundingAmt);
    vm.stopPrank();
  }

  function testFundAirdropFee() public {
    uint256 totalFundingAmt = 50 ether;
    uint32 feePct = 1000; // 10%

    vm.prank(deployer);
    pandasia.setFee(feePct);

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);

    uint256 feeAmount = (totalFundingAmt * feePct) / 10_000;
    uint256 fundAmount = totalFundingAmt - feeAmount;

    assertEq(airdrop.balance, fundAmount);
    assertEq(erc20.balanceOf(address(pandasia)), totalFundingAmt);

    uint256 previousBalance = erc20.balanceOf(deployer);
    vm.prank(deployer);
    pandasia.withdrawFees(id);

    assertEq(erc20.balanceOf(deployer), previousBalance + feeAmount);
  }

  function testFundAirdrop() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);
    vm.stopPrank();

    assertEq(airdrop.balance, totalFundingAmt);
    assertEq(erc20.balanceOf(address(pandasia)), totalFundingAmt);
  }

  /**************************************************************************************************************************************/
  /*** Withdraw Airdrop Tests                                                                                                              ***/
  /**************************************************************************************************************************************/

  function testWithdrawNotOwnerFunding() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);

    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);

    vm.warp(block.timestamp + 1001);
    vm.stopPrank();

    vm.expectRevert(Pandasia.NotOwner.selector);
    pandasia.withdrawFunding(id, 1 ether);

    vm.prank(airdropOwner);
    pandasia.withdrawFunding(id, 1 ether);
    assertEq(erc20.balanceOf(airdropOwner), 1 ether);
  }

  function testWithdrawInvalidAmount() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);

    vm.warp(block.timestamp + 1001);

    vm.expectRevert(Pandasia.InvalidAmount.selector);
    pandasia.withdrawFunding(id, totalFundingAmt + 10 ether);

    pandasia.withdrawFunding(id, totalFundingAmt);
    assertEq(erc20.balanceOf(airdropOwner), totalFundingAmt);
    vm.stopPrank();
  }

  function testWithdrawFundingAirdropStillActive() public {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(bytes32(0), address(erc20), 10 ether, uint64(block.timestamp), uint64(block.timestamp + 1000));

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);

    vm.expectRevert(Pandasia.AirdropStillActive.selector);
    pandasia.withdrawFunding(id, 1 ether);

    vm.warp(block.timestamp + 1001);

    pandasia.withdrawFunding(id, 1 ether);
    assertEq(erc20.balanceOf(airdropOwner), 1 ether);
    vm.stopPrank();
  }

  function testEmergencyWithdrawFundingUnauthorized() public {
    uint64 id = createFundedAirdrop(10 ether);

    bytes4 selector = bytes4(keccak256("OwnableUnauthorizedAccount(address)"));
    vm.expectRevert(abi.encodeWithSelector(selector, address(this)));
    pandasia.emergencyWithdraw(id);

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(id);
    uint256 airdropBalance = airdrop.balance;

    vm.prank(deployer);
    pandasia.emergencyWithdraw(id);

    airdrop = pandasia.getAirdrop(id);

    assertEq(erc20.balanceOf(deployer), airdropBalance);
    assertEq(airdrop.balance, 0);
  }

  /**************************************************************************************************************************************/
  /*** Get Airdrop Tests                                                                                                              ***/
  /**************************************************************************************************************************************/

  function testGetAirdropNotExistent() public {
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(0);
    Pandasia.Airdrop memory expected = Pandasia.Airdrop(0, address(0), address(0), 0, bytes32(0), 0, 0, 0);

    assertEq(airdrop.balance, expected.balance);
    assertEq(airdrop.claimAmount, expected.claimAmount);
    assertEq(airdrop.erc20, expected.erc20);
    assertEq(airdrop.startsAt, expected.startsAt);
    assertEq(airdrop.expiresAt, expected.expiresAt);
    assertEq(airdrop.owner, expected.owner);
    assertEq(airdrop.customRoot, expected.customRoot);
  }

  function testGetAirdrop() public {
    uint256 perClaimAmt = 10 ether;
    uint256 totalFundingAmt = 50 ether;
    uint64 startsAt = uint64(block.timestamp);
    uint64 expiresAt = uint64(block.timestamp + 1000);

    vm.startPrank(airdropOwner);

    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), perClaimAmt, uint32(block.timestamp), uint64(block.timestamp + 1000));
    uint64[] memory ids = pandasia.getAirdropIds(airdropOwner);

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    assertEq(ids[0], id);

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(ids[0]);
    Pandasia.Airdrop memory expected = Pandasia.Airdrop(
      id,
      airdropOwner,
      address(erc20),
      totalFundingAmt,
      otherRoot,
      perClaimAmt,
      startsAt,
      expiresAt
    );

    assertEq(airdrop.id, expected.id);
    assertEq(airdrop.balance, expected.balance);
    assertEq(airdrop.claimAmount, expected.claimAmount);
    assertEq(airdrop.erc20, expected.erc20);
    assertEq(airdrop.startsAt, expected.startsAt);
    assertEq(airdrop.expiresAt, expected.expiresAt);
    assertEq(airdrop.owner, expected.owner);
    assertEq(airdrop.customRoot, expected.customRoot);
  }

  function testGetAirdropIds() public {
    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    uint64[] memory ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids.length, 1);

    pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));
    uint64 id3 = pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));
    vm.stopPrank();

    ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids.length, 3);
    assertEq(ids[0], id);
    assertEq(ids[2], id3);

    address otherOwner = address(0x45);
    pandasia.grantRole(pandasia.AIRDROP_ADMIN(), otherOwner);

    vm.startPrank(otherOwner);
    pandasia.newAirdrop(otherRoot, address(erc20), 10 ether, uint32(block.timestamp), uint64(block.timestamp + 1000));

    ids = pandasia.getAirdropIds(otherOwner);
    assertEq(ids.length, 1);

    ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids.length, 3);
  }

  function testGetAirdrops() public {
    uint256 perClaimAmt = 10 ether;
    uint256 totalFundingAmt = 50 ether;
    uint64 startsAt = uint64(block.timestamp);
    uint64 expiresAt = uint64(block.timestamp + 1000);

    // create an airdrop
    vm.startPrank(airdropOwner);
    uint64 id = pandasia.newAirdrop(otherRoot, address(erc20), perClaimAmt, startsAt, expiresAt);
    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();

    Pandasia.Airdrop memory expected1 = Pandasia.Airdrop(
      0,
      airdropOwner,
      address(erc20),
      totalFundingAmt,
      otherRoot,
      perClaimAmt,
      startsAt,
      expiresAt
    );

    // second airdrop
    vm.startPrank(airdropOwner);
    uint64 id2 = pandasia.newAirdrop(validatorRoot, address(erc20), perClaimAmt, startsAt, expiresAt);
    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id2, totalFundingAmt);
    vm.stopPrank();

    Pandasia.Airdrop memory expected2 = Pandasia.Airdrop(
      id2,
      airdropOwner,
      address(erc20),
      totalFundingAmt,
      validatorRoot,
      perClaimAmt,
      startsAt,
      expiresAt
    );

    // Test
    Pandasia.Airdrop[] memory airdrops = pandasia.getAirdrops(0, 1);
    assertEq(airdrops.length, 1);

    Pandasia.Airdrop memory actual1 = airdrops[0];

    assertEq(actual1.id, expected1.id);
    assertEq(actual1.balance, expected1.balance);
    assertEq(actual1.claimAmount, expected1.claimAmount);
    assertEq(actual1.erc20, expected1.erc20);
    assertEq(actual1.startsAt, expected1.startsAt);
    assertEq(actual1.expiresAt, expected1.expiresAt);
    assertEq(actual1.owner, expected1.owner);
    assertEq(actual1.customRoot, expected1.customRoot);

    airdrops = pandasia.getAirdrops(1, 2);
    assertEq(airdrops.length, 1);

    Pandasia.Airdrop memory actual2 = airdrops[0];

    assertEq(actual2.id, expected2.id);
    assertEq(actual2.balance, expected2.balance);
    assertEq(actual2.claimAmount, expected2.claimAmount);
    assertEq(actual2.erc20, expected2.erc20);
    assertEq(actual2.startsAt, expected2.startsAt);
    assertEq(actual2.expiresAt, expected2.expiresAt);
    assertEq(actual2.owner, expected2.owner);
    assertEq(actual2.customRoot, expected2.customRoot);

    airdrops = pandasia.getAirdrops(0, 2);
    assertEq(airdrops.length, 2);
  }

  /**************************************************************************************************************************************/
  /*** Helpers                                                                                                                        ***/
  /**************************************************************************************************************************************/

  uint160 private actorCounter = 0;

  // create generic funded airdrop
  function createFundedAirdrop(uint256 perClaimAmt) internal returns (uint64 id) {
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);
    id = pandasia.newAirdrop(bytes32(0), address(erc20), perClaimAmt, uint64(block.timestamp), uint64(block.timestamp + 1000));

    // fund it
    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);
    pandasia.fundAirdrop(id, totalFundingAmt);
    vm.stopPrank();
  }

  function getActor(string memory name) internal returns (address) {
    actorCounter++;
    address addr = address(uint160(0x50000 + actorCounter));
    vm.label(addr, name);
    return addr;
  }

  function getActor(string memory name, address addr) internal returns (address) {
    vm.label(addr, name);
    return addr;
  }

  string private checkpointLabel;
  uint256 private checkpointGasLeft;

  function startMeasuringGas(string memory label) internal virtual {
    checkpointLabel = label;
    checkpointGasLeft = gasleft();
  }

  function stopMeasuringGas() internal virtual {
    uint256 checkpointGasLeft2 = gasleft();

    string memory label = checkpointLabel;

    emit log_named_uint(string(abi.encodePacked(label, " Gas")), checkpointGasLeft - checkpointGasLeft2);
  }

  function logAirdrop(Pandasia.Airdrop memory airdrop) internal virtual {
    console2.log(airdrop.id);
    console2.log(airdrop.balance);
    console2.log(airdrop.claimAmount);
    console2.log(airdrop.erc20);
    console2.log(airdrop.startsAt);
    console2.log(airdrop.expiresAt);
    console2.log(airdrop.owner);
    console2.logBytes32(airdrop.customRoot);
  }
}
