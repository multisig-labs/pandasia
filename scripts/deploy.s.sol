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

    pandasia.setStakingContract(0x9946e68490D71Fe976951e360f295c4Cf8531D00);

    vm.stopBroadcast();
  }
}
