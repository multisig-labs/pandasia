// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {Pandasia} from "../contracts/Pandasia.sol";
import {console2} from "forge-std/console2.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployContract is Script {
  function run() external {
    vm.startBroadcast();

    Pandasia pandasiaImpl2 = new Pandasia();
    ProxyAdmin liveAdmin = ProxyAdmin(0xc9Ec877642100f9dEA388D03E28c4d9Ac5F66c8D);
    Pandasia pandasia = Pandasia(0x4de93c302475Dc25d6AE12a4eB43Aca119C76FA1);
    liveAdmin.upgrade(ITransparentUpgradeableProxy(address(0x4de93c302475Dc25d6AE12a4eB43Aca119C76FA1)), address(pandasiaImpl2));
    vm.stopBroadcast();
  }
}
