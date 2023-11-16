// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {Pandasia} from "../contracts/Pandasia.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployContract is Script {
  function run() external {
    vm.startBroadcast();

    ProxyAdmin proxyAdmin = new ProxyAdmin();
    Pandasia pandasiaImpl = new Pandasia();
    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(
      address(pandasiaImpl),
      address(proxyAdmin),
      abi.encodeWithSelector(pandasiaImpl.initialize.selector)
    );
    Pandasia pandasia = Pandasia(payable(pandasiaProxy));

    pandasia.setStorageContract(0x1cEa17F9dE4De28FeB6A102988E12D4B90DfF1a9);

    vm.stopBroadcast();
  }
}
