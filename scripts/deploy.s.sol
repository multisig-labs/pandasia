// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {Pandasia} from "../contracts/Pandasia.sol";

import {console2} from "forge-std/console2.sol";
import {Script} from "forge-std/Script.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployContract is Script {
  function run() external {
    vm.startBroadcast();

    ERC20 ppj = new ERC20("PapaJohn", "PPJ", 1 ether, 0xd8fF2A2E79417046A41f26A7AFFd38fbe87CF10c);
    ERC20 marcos = new ERC20("Marcos", "MC", 1 ether, 0xd8fF2A2E79417046A41f26A7AFFd38fbe87CF10c);
    console2.log("ppj", address(ppj));

    // ppj._mint(0xd8fF2A2E79417046A41f26A7AFFd38fbe87CF10c, 1 ether);

    bytes memory data = "";

    ProxyAdmin proxyAdmin = new ProxyAdmin();
    Pandasia pandasiaImpl = new Pandasia();
    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(
      address(pandasiaImpl),
      address(proxyAdmin),
      data
    );
    Pandasia pandasia = Pandasia(payable(pandasiaProxy));

    console2.log("proxy admin            ", address(proxyAdmin));
    console2.log("pandasia implementation", address(pandasiaImpl));
    console2.log("pandasia proxy         ", address(pandasiaProxy));
    console2.log("pandasia deployed to   ", address(pandasia));

    vm.stopBroadcast();
  }
}
