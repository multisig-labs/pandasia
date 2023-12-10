// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {EnvironmentConfig} from "./EnvironmentConfig.s.sol";

import {Pandasia} from "../contracts/Pandasia.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";

contract DeployContract is Script, EnvironmentConfig {
  function run() external {
    loadAddresses();
    loadUsers();

    address deployer = getUser("deployer");
    require(deployer.balance > 1 ether, "Insufficient funds to deploy");

    vm.startBroadcast(deployer);

    Pandasia pandasiaImpl = new Pandasia();
    saveAddress("PandasiaImpl", address(pandasiaImpl));

    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(
      address(pandasiaImpl),
      deployer,
      abi.encodeWithSelector(pandasiaImpl.initialize.selector)
    );

    Pandasia pandasia = Pandasia(payable(pandasiaProxy));
    saveAddress("Pandasia", address(pandasia));

    // TransparentUpgradeableProxy makes it's own ProxyAdmin now
    // so we have to get the ProxyAdmin from it's storage slot
    bytes32 adminSlot = vm.load(address(pandasia), ERC1967Utils.ADMIN_SLOT);
    ProxyAdmin proxyAdmin = ProxyAdmin(address(uint160(uint256(adminSlot))));
    saveAddress("PandasiaAdmin", address(proxyAdmin));

    console2.log("Setting storage address to", getAddress("Storage"));
    pandasia.setStorageContract(getAddress("Storage"));

    vm.stopBroadcast();
  }
}
