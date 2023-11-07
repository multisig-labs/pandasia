// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import {Pandasia} from "../contracts/Pandasia.sol";
import {ERC20PresetMinterPauser} from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol";

contract CreateAirdrop is Script {
  function run() external {
    address pandasiaAddr = vm.envAddress("PANDASIA_ADDR");
    Pandasia pandasia = Pandasia(pandasiaAddr);

    uint256 deployerPk = vm.envUint("PRIVATE_KEY");
    address deployer = vm.rememberKey(deployerPk);
    vm.startBroadcast(deployerPk);

    ERC20PresetMinterPauser erc20 = new ERC20PresetMinterPauser("Meme from Heaven", "MEME");
    uint64 startsAt = uint64(block.timestamp + 1 days);
    uint64 expiresAt = uint64(block.timestamp + 30 days);
    console.log("addresseerc20", address(erc20));

    uint64 id = pandasia.newAirdrop(
      bytes32(0x5575f4c36b81aaa8dc2e6e460d66c6a450187cdd1d220869cbf190e515719cd5),
      true,
      address(erc20),
      1 ether,
      startsAt,
      expiresAt
    );
    console.log("AIRDROP ID", id);
    // Fund it
    erc20.mint(deployer, 100 ether);
    erc20.approve(address(pandasia), 100 ether);
    pandasia.fundAirdrop(id, 100 ether);

    vm.stopBroadcast();
  }
}
