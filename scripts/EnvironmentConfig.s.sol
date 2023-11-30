// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

// Helper funcs for the deploy scripts
// Store deployed addresses in /deployed directory.
// Because vm.writeJson can only replace keys, not create them,
// we use a template json with all contracts listed out.

contract EnvironmentConfig is Script {
  error InvalidChain();
  error InvalidUser();
  error InvalidAddress();

  string public addressesJson;

  struct User {
    uint256 pk;
    address addr;
  }
  string[] public userNames = ["deployer"];
  mapping(string => User) public namedUsers;

  // Map keys derived from mnemonic to our list of userNames
  function loadUsers() public {
    string memory mnemonic = vm.envString("MNEMONIC");
    for (uint32 i; i < userNames.length; i++) {
      string memory name = userNames[i];
      uint256 pk = vm.deriveKey(mnemonic, i);
      address addr = vm.rememberKey(pk);
      namedUsers[name] = User(pk, addr);
    }
  }

  string[] public contractNames = ["PandasiaImpl", "PandasiaAdmin", "Pandasia", "Storage"];

  function verifyContracts() public {
    for (uint256 i; i < contractNames.length; i++) {
      address addr = getAddress(contractNames[i]);
      console2.log(contractNames[i], addr);
    }
  }

  function loadAddresses() public {
    // these are not global so that we can use with vm.startFork
    uint256 chainID = getChainId();
    string memory pathToDeploymentTemplate = string(abi.encodePacked(vm.projectRoot(), "/deployed/", vm.toString(chainID), "-addresses.tmpl.json"));
    string memory pathToDeploymentFile = string(abi.encodePacked(vm.projectRoot(), "/deployed/", vm.toString(chainID), "-addresses.json"));

    if (fileExists(pathToDeploymentFile)) {
      addressesJson = vm.readFile(pathToDeploymentFile);
    } else {
      // TODO auto-create a blank template for a chainid if none exists
      addressesJson = vm.readFile(pathToDeploymentTemplate);
      require(bytes(addressesJson).length != 0, "Address template not found!");
      vm.writeFile(pathToDeploymentFile, addressesJson);
    }
  }

  function getChainId() public view returns (uint256) {
    console2.log("Current chain id:", block.chainid);
    return block.chainid;
  }

  function saveAddress(string memory name, address addr) public {
    // these are not global so that we can use with vm.startFork
    uint256 chainID = getChainId();
    string memory pathToDeploymentFile = string(abi.encodePacked(vm.projectRoot(), "/deployed/", vm.toString(chainID), "-addresses.json"));

    // key is a jq-ish locator of where to store data in the JSON
    string memory key = string(abi.encodePacked(".", name));
    console2.log("attempting to access key", key);
    require(keyExists(addressesJson, key), "Must update existing key");
    vm.writeJson(vm.toString(addr), pathToDeploymentFile, key);
    loadAddresses();
  }

  function getAddress(string memory name) public returns (address) {
    // key is a jq-ish locator of where to store data in the JSON
    string memory key = string(abi.encodePacked(".", name));
    address addr = vm.parseJsonAddress(addressesJson, key);
    if (addr == address(0)) {
      console2.log("Unable to access key in json addresses", key);
      revert InvalidAddress();
    }
    vm.label(addr, name);
    return addr;
  }

  function getUser(string memory name) public returns (address) {
    address addr = namedUsers[name].addr;
    if (addr == address(0)) {
      revert InvalidUser();
    }
    vm.label(addr, name);
    return addr;
  }

  function isContractDeployed(string memory name) public returns (bool) {
    // key is a jq-ish locator of where to store data in the JSON
    string memory key = string(abi.encodePacked(".", name));
    address addr = address(0);
    console2.log("json addresses", addressesJson);

    // address addr = vm.parseJsonAddress(addressesJson, key);
    return addr != address(0) && addr.code.length > 0;
  }

  function keyExists(string memory json, string memory key) internal pure returns (bool) {
    return vm.parseJson(json, key).length > 0;
  }

  // Must be better way?
  function fileExists(string memory path) internal returns (bool) {
    try vm.fsMetadata(path) {
      return true;
    } catch Error(string memory) {
      return false;
    } catch (bytes memory) {
      // data is �E &No such file or directory (os error 2)
      return false;
    }
  }

  modifier onlyMainnet() {
    if (block.chainid != 43114) {
      console2.log("Script only allowed on Mainnet");
      revert InvalidChain();
    }
    _;
  }

  modifier onlyDev() {
    if (block.chainid == 43114) {
      console2.log("Script only allowed on development chains");
      revert InvalidChain();
    }
    _;
  }
}
