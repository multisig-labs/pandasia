// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {Pandasia} from "../contracts/Pandasia.sol";
import {ERC20PresetMinterPauser} from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol";

contract GetAirdrop is Script {
  bytes32[] public proof;

  function run() external {
    address pandasiaAddr = vm.envAddress("PANDASIA_ADDR");
    Pandasia pandasia = Pandasia(pandasiaAddr);

    // Pandasia.Airdrop[] memory airdrops = pandasia.getAirdrops(0, 10);
    Pandasia.Airdrop memory airdrop = pandasia.getAirdrop(0);
    console.log();
    console.log();
    // console.log(airdrops[1].onlyRegistered);
    console.log(airdrop.onlyRegistered);
  }
}
