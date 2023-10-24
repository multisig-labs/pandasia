// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {Pandasia} from "../contracts/Pandasia.sol";

contract DeployContract is Script {
  function run() external {
    vm.startBroadcast();

    new Pandasia();

    vm.stopBroadcast();
  }
}
