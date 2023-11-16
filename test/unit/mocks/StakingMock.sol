// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.19;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract StakingMock {
  // Mock this from GoGoPool
  // function getLastRewardsCycleCompleted(address stakerAddr) public view returns (uint256) {
  // 	int256 stakerIndex = getIndexOf(stakerAddr);
  // 	return getUint(keccak256(abi.encodePacked("staker.item", stakerIndex, ".lastRewardsCycleCompleted")));
  // }

  mapping(address => uint256) private mockData;

  function setLastRewardsCycleCompleted(address stakerAddr, uint256 num) public {
    mockData[stakerAddr] = num;
  }

  function getLastRewardsCycleCompleted(address stakerAddr) public view returns (uint256) {
    return mockData[stakerAddr];
  }
}
