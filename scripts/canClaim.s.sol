// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {Pandasia} from "../contracts/Pandasia.sol";
import {ERC20PresetMinterPauser} from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol";

contract CanClaim is Script {
  bytes32[] public proof;

  function run() external {
    address pandasiaAddr = vm.envAddress("PANDASIA_ADDR");
    Pandasia pandasia = Pandasia(pandasiaAddr);

    address tester = vm.envAddress("TEST_ADDR");

    uint64[] memory ids = pandasia.getAirdropIds(0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4);

    proof.push(0x20bbea9e3ef756ae5bd741e1747488cd45bcbb4190ca0fcb6f6e1f8f4085dadc);
    proof.push(0x7bc7dc6bf7422ff4e29869c118dee787180350debf1321d9f388cd97b310d0b4);
    proof.push(0xe2be3db3f680ad7b8bed5efcd5b47c8abfd716e9608aeb9265be1be5c9f2d2f8);
    proof.push(0xea66d96e5940135479c3ca2224dfdd1761a6dc7b238cf7d8bf389b7f28af74e4);
    proof.push(0x342aeef56cf4f0239ab891ea00dd1bec637eb85d20dd95473a80e88b31506491);
    proof.push(0x912c0f94b90f5fe366b4c0b6b2f9e5428a4388e9ca14be69c98c4c4ee88e6cd9);

    bytes32 root = 0x5575f4c36b81aaa8dc2e6e460d66c6a450187cdd1d220869cbf190e515719cd5;

    bool verified = pandasia.verify(root, tester, proof);

    bool canClaim = pandasia.canClaimAirdrop(tester, 2, proof);
    console.log("verified?", verified);
    console.log("canClaim?", canClaim);
  }
}
