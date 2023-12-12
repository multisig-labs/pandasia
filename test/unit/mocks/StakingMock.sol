// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.19;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract StakingMock {
  // Mock these methods from GoGoPool
  // function getLastRewardsCycleCompleted(address stakerAddr) public view returns (uint256) {
  // 	int256 stakerIndex = getIndexOf(stakerAddr);
  // 	return getUint(keccak256(abi.encodePacked("staker.item", stakerIndex, ".lastRewardsCycleCompleted")));
  // }

  // function getAVAXValidatingHighWater(address stakerAddr) public view returns (uint256) {
  // 	int256 stakerIndex = getIndexOf(stakerAddr);
  // 	return getUint(keccak256(abi.encodePacked("staker.item", stakerIndex, ".avaxValidatingHighWater")));
  // }

  mapping(address => uint256) private mockCycle;
  mapping(address => uint256) private mockHighWater;

  function setLastRewardsCycleCompleted(address stakerAddr, uint256 num) public {
    mockCycle[stakerAddr] = num;
  }

  function getLastRewardsCycleCompleted(address stakerAddr) public view returns (uint256) {
    return mockCycle[stakerAddr];
  }

  function setAVAXValidatingHighWater(address stakerAddr, uint256 num) public {
    mockHighWater[stakerAddr] = num;
  }

  function getAVAXValidatingHighWater(address stakerAddr) public view returns (uint256) {
    return mockHighWater[stakerAddr];
  }
}
