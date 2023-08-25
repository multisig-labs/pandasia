pragma solidity 0.8.19;

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ERC20Mock} from "@openzeppelin/contracts/mocks/ERC20Mock.sol";
import {console2} from "forge-std/console2.sol";
import {Pandasia} from "../../contracts/pandasia.sol";

contract AirdropTest is Test {
	ERC20Mock public erc20;
	Pandasia public pandasia;
	address public validator;
	address public deployer;

	bytes32[] public emptyProof;

	function setUp() public {
		emptyProof = new bytes32[](0);

		bytes32 root = bytes32(0x1733170f5a465a52692730efa67c11a3c9b1208a5acbe833057fac165ce6947b);
		validator = address(0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4);
		// address paddy = address(0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02);
		bytes32[] memory proof = new bytes32[](1);
		proof[0] = bytes32(0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f);

		deployer = address(999);
		pandasia = new Pandasia();
		pandasia.setValidatorRoot(root);
		pandasia.transferOwnership(deployer);

		// Signature generated on wallet.avax.network
		uint8 v = 0;
		bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
		bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

		vm.prank(validator);
		pandasia.registerPChainAddr(v, r, s, proof);

		erc20 = new ERC20Mock();
	}

	function testNewAirdrop() public {
		uint256 perClaimAmt = 10 ether;
		uint256 totalFundingAmt = 20 ether;
		address owner = address(1);
		address nonValidator = address(2);
		vm.startPrank(owner);

		uint64 id = pandasia.newAirdrop(bytes32(0), true, address(erc20), perClaimAmt, uint32(block.timestamp + 1000));
		uint64[] memory ids = pandasia.getAirdropIds(owner);
		assertEq(ids[0], id, "getAirdrops");

		erc20.mint(owner, totalFundingAmt);
		erc20.approve(address(pandasia), totalFundingAmt);

		// Under fund it
		vm.expectRevert(Pandasia.InvalidAmount.selector);
		pandasia.fundAirdrop(id, 1 ether);

		// Fund it
		pandasia.fundAirdrop(id, totalFundingAmt);
		assertEq(erc20.balanceOf(owner), 0);

		vm.stopPrank();

		vm.prank(nonValidator);
		vm.expectRevert(Pandasia.AddressNotEligible.selector);
		pandasia.claimAirdrop(id, emptyProof);

		vm.startPrank(validator);
		pandasia.claimAirdrop(id, emptyProof);
		assertEq(erc20.balanceOf(validator), perClaimAmt);

		// Can only claim once
		vm.expectRevert(Pandasia.AddressNotEligible.selector);
		pandasia.claimAirdrop(id, emptyProof);
		vm.stopPrank();

		vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
		pandasia.withdrawFunding(id, 1 ether);

		vm.startPrank(owner);
		vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
		pandasia.withdrawFunding(id, totalFundingAmt);

		vm.expectRevert(Pandasia.InvalidWithdrawRequest.selector);
		pandasia.withdrawFunding(id, 1 ether);

		vm.warp(block.timestamp + 1001);
		pandasia.withdrawFunding(id, 1 ether);
		assertEq(erc20.balanceOf(owner), 1 ether);
		vm.stopPrank();

		vm.expectRevert("Ownable: caller is not the owner");
		pandasia.emergencyWithdraw(id, 1 ether);

		vm.prank(deployer);
		pandasia.emergencyWithdraw(id, 1 ether);
		assertEq(erc20.balanceOf(deployer), 1 ether);
	}

	//
	// HELPERS
	//

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
}
