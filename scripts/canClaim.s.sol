// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {Pandasia} from "../contracts/Pandasia.sol";

contract CanClaim is Script {
  bytes32[] public proof = [
    bytes32(0x20b897affce761596e11062ca10c605c466fcc119a577230b5081a9b48dff0fc),
    bytes32(0x961d49e58acd55d827554615d2ea7e69665af62b509c5d1e7bd0f00966f5dcac),
    bytes32(0x24a15ee8e4297f8d709bdc45be2e0fd9f341183303a25b09c869c2164f9d6846),
    bytes32(0x774464ce2d430d2a7894f90d5fb2898f70a7f03482c6b12869082406a69e323a),
    bytes32(0xda2fe588f888ea5cfd231ba97e95880185426cfd038511e67b40840bfe790772),
    bytes32(0x636157f82de22f2592b4ff88d50f470d7b255f4930e612fef3eb4cb00b98067b),
    bytes32(0x649b1ae11aeb489dd9fae059945044139cba157809ca32fa3ba0d1e0666e1691),
    bytes32(0xd10e55e6ba7cf2be10d81fd0b2720b8fc4c0d52c08c2d8e3011faeb21cf7f049),
    bytes32(0x3f458034777a48d48f3add4f214e708509c890621dc270b9ca4dd5305580debc),
    bytes32(0xb5e92c60df716c6cc2f63013645f13bbe0b14180cd5a58ca29bb88e4cb9b4726),
    bytes32(0x5e026c98737c0b293a04f351ed66a65455d84e4cb8a52145c5f14f74e7e8918c),
    bytes32(0xa4936775dbd0bb4482e9fb5b3e7d9f737f9c3371522b18007b35a2aca1ece41b),
    bytes32(0x62f069ba79e64ad42e306297ba26f641c33bf3b151e32d8d85804c4c444c946b),
    bytes32(0xc551c06d7d46daa08603c941da06f4c80eb3ec28ab569c52b023735ae9518623)
  ];

  function run() external {
    address pandasiaAddr = vm.envAddress("PANDASIA_ADDR");
    Pandasia pandasia = Pandasia(pandasiaAddr);

    address tester = address(0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02);

    // uint64[] memory ids = pandasia.getAirdropIds(0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4);

    // proof.push(0x20bbea9e3ef756ae5bd741e1747488cd45bcbb4190ca0fcb6f6e1f8f4085dadc);
    // proof.push(0x7bc7dc6bf7422ff4e29869c118dee787180350debf1321d9f388cd97b310d0b4);
    // proof.push(0xe2be3db3f680ad7b8bed5efcd5b47c8abfd716e9608aeb9265be1be5c9f2d2f8);
    // proof.push(0xea66d96e5940135479c3ca2224dfdd1761a6dc7b238cf7d8bf389b7f28af74e4);
    // proof.push(0x342aeef56cf4f0239ab891ea00dd1bec637eb85d20dd95473a80e88b31506491);
    // proof.push(0x912c0f94b90f5fe366b4c0b6b2f9e5428a4388e9ca14be69c98c4c4ee88e6cd9);

    bytes32 root = 0x6128219813313b6056381b56d5699edcda31e2afa6558aa9866e0f9021f7725d;

    bool verified = pandasia.verify(root, tester, proof);
    pandasia.registerPChainAddr(
      0,
      0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337,
      0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a,
      courseMap
    );

    bool canClaim = pandasia.canClaimAirdrop(tester, 2, proof);
    console.log("verified?", verified);
    console.log("canClaim?", canClaim);
  }
}
