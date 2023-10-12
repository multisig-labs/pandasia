pragma solidity 0.8.19;

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ERC20Mock} from "@openzeppelin/contracts/mocks/ERC20Mock.sol";
import {console2} from "forge-std/console2.sol";
import {Pandasia} from "../../contracts/Pandasia.sol";
import {StakingMock} from "../../contracts/StakingMock.sol";

contract AirdropTest is Test {
  ERC20Mock public erc20;
  Pandasia public pandasia;
  StakingMock public stakingContract;

  address public deployer; // Owner of Pandasia contract
  address public airdropOwner;
  address public validator; // eligible via validatorRoot
  address public minipoolOperator; // eligible via staking contract
  address public otherUser; // eligible via otherProof
  address public notEligibleUser; // not eligible in any way

  bytes32[] public emptyProof;
  bytes32[] public validatorProof;
  bytes32[] public otherProof;
  bytes32 public validatorRoot;
  bytes32 public otherRoot;

  function setUp() public {
    deployer = getActor("deployer");
    airdropOwner = getActor("airdropOwner");
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

    stakingContract = new StakingMock();
    stakingContract.setLastRewardsCycleCompleted(minipoolOperator, 1);

    pandasia = new Pandasia();
    pandasia.setMerkleRoot(validatorRoot);
    pandasia.setStakingContract(address(stakingContract));
    pandasia.transferOwnership(deployer);

    // Signature generated on wallet.avax.network
    uint8 v = 0;
    bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
    bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

    vm.prank(validator);
    pandasia.registerPChainAddr(v, r, s, validatorProof);

    erc20 = new ERC20Mock();
  }

  function testNewAirdropNoAdditionalMerkeRoot() public {
    bool isEligible;
    uint256 perClaimAmt = 10 ether;
    uint256 totalFundingAmt = 20 ether;

    vm.startPrank(airdropOwner);

    uint64 id = pandasia.newAirdrop(bytes32(0), false, address(erc20), perClaimAmt, uint64(block.timestamp), uint64(block.timestamp + 1000));
    uint64[] memory ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids[0], id, "getAirdrops");

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);

    // Fund it
    pandasia.fundAirdrop(id, totalFundingAmt);
    assertEq(erc20.balanceOf(airdropOwner), 0);

    vm.stopPrank();

    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(notEligibleUser, id, emptyProof);

    vm.prank(notEligibleUser);
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.claimAirdrop(id, emptyProof);

    isEligible = pandasia.canClaimAirdrop(validator, id, emptyProof);
    assertTrue(isEligible);

    vm.startPrank(validator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(validator), perClaimAmt);

    // Can only claim once
    vm.expectRevert(Pandasia.AddressAlreadyClaimed.selector);
    pandasia.claimAirdrop(id, emptyProof);
    vm.stopPrank();

    vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
    pandasia.withdrawFunding(id, 1 ether);

    vm.startPrank(airdropOwner);
    vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
    pandasia.withdrawFunding(id, totalFundingAmt);

    vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
    pandasia.withdrawFunding(id, 1 ether);

    vm.warp(block.timestamp + 1001);
    pandasia.withdrawFunding(id, 1 ether);
    assertEq(erc20.balanceOf(airdropOwner), 1 ether);
    vm.stopPrank();

    vm.expectRevert("Ownable: caller is not the owner");
    pandasia.emergencyWithdraw(id, 1 ether);

    vm.prank(deployer);
    pandasia.emergencyWithdraw(id, 1 ether);
    assertEq(erc20.balanceOf(deployer), 1 ether);
  }

  function testNewAirdropWithAdditionalMerkeRoot() public {
    bool isEligible;
    uint256 perClaimAmt = 10 ether;
    uint256 totalFundingAmt = 50 ether;

    vm.startPrank(airdropOwner);

    uint64 id = pandasia.newAirdrop(otherRoot, false, address(erc20), perClaimAmt, uint64(block.timestamp), uint64(block.timestamp + 1000));
    uint64[] memory ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids[0], id, "getAirdrops");

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);

    // Fund it
    pandasia.fundAirdrop(id, totalFundingAmt);
    assertEq(erc20.balanceOf(airdropOwner), 0);

    vm.stopPrank();

    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.canClaimAirdrop(notEligibleUser, id, emptyProof);

    vm.prank(notEligibleUser);
    vm.expectRevert(Pandasia.AddressNotEligible.selector);
    pandasia.claimAirdrop(id, emptyProof);

    isEligible = pandasia.canClaimAirdrop(validator, id, emptyProof);
    assertTrue(isEligible);

    vm.prank(validator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(validator), perClaimAmt);

    // Claim via otherRoot
    vm.prank(otherUser);
    pandasia.claimAirdrop(id, otherProof);
    assertEq(erc20.balanceOf(otherUser), perClaimAmt);

    // Claim via staking contract
    vm.prank(minipoolOperator);
    pandasia.claimAirdrop(id, emptyProof);
    assertEq(erc20.balanceOf(minipoolOperator), perClaimAmt);
  }

  function testGetAirdropWithoutCreating() public {
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(0);
    Pandasia.Airdrop memory expected = Pandasia.Airdrop(0, address(0), address(0), 0, bytes32(0), 0, 0, 0, false);

    assertEq(airdrop.balance, expected.balance);
    assertEq(airdrop.claimAmount, expected.claimAmount);
    assertEq(airdrop.erc20, expected.erc20);
    assertEq(airdrop.startsAt, expected.startsAt);
    assertEq(airdrop.expiresAt, expected.expiresAt);
    assertEq(airdrop.onlyRegistered, expected.onlyRegistered);
    assertEq(airdrop.owner, expected.owner);
    assertEq(airdrop.root, expected.root);
  }

  function testGetAirdrop() public {
    uint256 perClaimAmt = 10 ether;
    uint256 totalFundingAmt = 50 ether;
    uint64 startsAt = uint64(block.timestamp);
    uint64 expiresAt = uint64(block.timestamp + 1000);

    vm.startPrank(airdropOwner);

    uint64 id = pandasia.newAirdrop(otherRoot, false, address(erc20), perClaimAmt, uint32(block.timestamp), uint64(block.timestamp + 1000));
    uint64[] memory ids = pandasia.getAirdropIds(airdropOwner);
    assertEq(ids[0], id, "getAirdrops");

    erc20.mint(airdropOwner, totalFundingAmt);
    erc20.approve(address(pandasia), totalFundingAmt);

    // Fund it
    pandasia.fundAirdrop(id, totalFundingAmt);
    assertEq(erc20.balanceOf(airdropOwner), 0);

    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(ids[0]);
    Pandasia.Airdrop memory expected = Pandasia.Airdrop(
      0,
      airdropOwner,
      address(erc20),
      totalFundingAmt,
      otherRoot,
      perClaimAmt,
      startsAt,
      expiresAt,
      false
    );

    assertEq(airdrop.balance, expected.balance);
    assertEq(airdrop.claimAmount, expected.claimAmount);
    assertEq(airdrop.erc20, expected.erc20);
    assertEq(airdrop.expiresAt, expected.expiresAt);
    assertEq(airdrop.onlyRegistered, expected.onlyRegistered);
    assertEq(airdrop.owner, expected.owner);
    assertEq(airdrop.root, expected.root);
  }

  //
  // HELPERS
  //

  uint160 private actorCounter = 0;

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
    console2.log(airdrop.balance);
    console2.log(airdrop.claimAmount);
    console2.log(airdrop.erc20);
    console2.log(airdrop.expiresAt);
    console2.log(airdrop.onlyRegistered);
    console2.log(airdrop.owner);
    console2.logBytes32(airdrop.root);
  }
}
