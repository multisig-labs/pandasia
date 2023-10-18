// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import "forge-std/Script.sol";
import {ECoin} from "../contracts/ECoin.sol";

contract DeployECoin is Script {
  function run() external {
    vm.startBroadcast();

    new ECoin();

    vm.stopBroadcast();
  }
}
