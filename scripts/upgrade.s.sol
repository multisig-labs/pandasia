// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {EnvironmentConfig} from "./EnvironmentConfig.s.sol";

import {Pandasia} from "../contracts/Pandasia.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";

contract UpgradeContract is Script, EnvironmentConfig {
  function run() external {
    loadAddresses();
    loadUsers();

    address deployer = getUser("deployer");
    require(deployer.balance > 0.5 ether, "Insufficient funds to upgrade");

    vm.startBroadcast(deployer);

    Pandasia pandasiaImplV2 = new Pandasia();
    saveAddress("PandasiaImplV2", address(pandasiaImpl));

    ProxyAdmin proxyAdmin = ProxyAdmin(getAddress("PandasiaAdmin"));
    address pandasiaProxy = getAddress("Pandasia");

    proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(pandasiaProxy), address(pandasiaImplV2), new bytes(0));

    vm.stopBroadcast();
  }
}
