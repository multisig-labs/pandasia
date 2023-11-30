// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {EnvironmentConfig} from "./EnvironmentConfig.s.sol";

import {Pandasia} from "../contracts/Pandasia.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployContract is Script, EnvironmentConfig {
  function run() external {
    loadAddresses();
    loadUsers();

    address deployer = getUser("deployer");
    require(deployer.balance > 1 ether, "Insufficient funds to deploy");

    vm.startBroadcast(deployer);

    if (isContractDeployed("Pandasia")) {
      console2.log("Pandasia exists, skipping...");
    } else {
      ProxyAdmin proxyAdmin = new ProxyAdmin();
      saveAddress("PandasiaAdmin", address(proxyAdmin));

      Pandasia pandasiaImpl = new Pandasia();
      saveAddress("PandasiaImpl", address(pandasiaImpl));

      TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(
        address(pandasiaImpl),
        address(proxyAdmin),
        abi.encodeWithSelector(pandasiaImpl.initialize.selector)
      );
      Pandasia pandasia = Pandasia(payable(pandasiaProxy));
      saveAddress("Pandasia", address(pandasia));

      console2.log("Setting storage address to", getAddress("Storage"));
      pandasia.setStorageContract(getAddress("Storage"));
    }

    vm.stopBroadcast();
  }
}
