// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.19;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract StorageMock {
  // Mock this from GoGoPool
  // function getAddress(bytes32 key) external view returns (address) {
  //   return addressStorage[key];
  // }

  mapping(bytes32 => address) private mockStorage;

  function setAddress(bytes32 key, address addr) public {
    mockStorage[key] = addr;
  }

  function getAddress(bytes32 key) public view returns (address) {
    return mockStorage[key];
  }
}
