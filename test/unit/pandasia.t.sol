pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {console2} from "forge-std/console2.sol";
import {Pandasia} from "../../contracts/pandasia.sol";
import {SECP256K1} from "../../contracts/SECP256K1.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";

// Test Data
// Mnemonic: test test test test test test test test test test test test test test test test test test test test test test test blade
// Eth Addr: 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4
// Ava: P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww
// P-Addr Bytes: 0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02

// P-avax1gfpj... signing 0x0000000000000000000000000000000000000001
// sig: MczkY1ar24JgNxmxcRA9KkQJSVo4rWorg2XwDtnAzakMw4PLnLwFqrEyTYL7goJjLX3Gnim4UoXhjcCdEPAHrxpXh9PpNR
// {"r":"0x23b5b54651e48c075395b537775219548920f453332fdc586d4d0c8fadfb6072","s":"0x155ffa9f0c72d2dc3e60890a641ba28346271e9fb51ec0710e69095038cee1d6","v":"0x01"}

// P-avax1gfpj... signing "0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4" gives
// sig: 24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx
// {"r":"0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337","s":"0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a","v":"0x00"}

// Merkle Tree
// {"format":"standard-v1","tree":["0xfcd7a701a861392c67cef2baaaf08063a1214f4ba4c3948c45b8e2008d28a35e","0xbfc74daa8eab55692b857896491b60b6d44cf47bd8629cd66ae2aca38f6fbb37","0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f"],"values":[{"value":["0x1111111111111111111111111111111111111111"],"treeIndex":2},{"value":["0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"],"treeIndex":1}],"leafEncoding":["address"]}
// Root: 0xfcd7a701a861392c67cef2baaaf08063a1214f4ba4c3948c45b8e2008d28a35e
// Proof for 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4: [0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f]

contract PandasiaTest is Test {
  Pandasia public pandasia;

  address public cAddress = address(0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4);
  address public pAddressBytes = address(0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02);

  function setUp() public {
    ProxyAdmin proxyAdmin = new ProxyAdmin(address(this));
    Pandasia pandasiaImpl = new Pandasia();

    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(address(pandasiaImpl), address(proxyAdmin), bytes(""));
    pandasia = Pandasia(payable(pandasiaProxy));
    pandasia.initialize();
  }

  /**************************************************************************************************************************************/
  /*** Registration Tests                                                                                                             ***/
  /**************************************************************************************************************************************/

  function testRegisterPChainAddr() public {
    bytes32 root = bytes32(0x1733170f5a465a52692730efa67c11a3c9b1208a5acbe833057fac165ce6947b);
    bytes32[] memory proof = new bytes32[](1);
    proof[0] = bytes32(0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f);

    pandasia.setMerkleRoot(root);
    assertFalse(pandasia.isRegisteredValidator(cAddress));

    // Signature generated on wallet.avax.network
    uint8 v = 0;
    bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
    bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

    vm.expectRevert(Pandasia.PAddrNotInMerkleTree.selector);
    pandasia.registerPChainAddr(v, r, s, proof);
    assertFalse(pandasia.isRegisteredValidator(cAddress));

    startMeasuringGas("registerPChainAddr");
    vm.prank(cAddress);
    pandasia.registerPChainAddr(v, r, s, proof);
    stopMeasuringGas();
    assertTrue(pandasia.isRegisteredValidator(cAddress));

    // Try to register a different cAddress with the same pAddress
    address hacker = address(1);
    v = 1;
    r = bytes32(0x23b5b54651e48c075395b537775219548920f453332fdc586d4d0c8fadfb6072);
    s = bytes32(0x155ffa9f0c72d2dc3e60890a641ba28346271e9fb51ec0710e69095038cee1d6);
    vm.prank(hacker);
    vm.expectRevert(Pandasia.PAddrAlreadyRegistered.selector);
    pandasia.registerPChainAddr(v, r, s, proof);
  }

  function testUnregisterPChainAddr() public {
    bytes32 root = bytes32(0x1733170f5a465a52692730efa67c11a3c9b1208a5acbe833057fac165ce6947b);
    bytes32[] memory proof = new bytes32[](1);
    proof[0] = bytes32(0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f);

    pandasia.setMerkleRoot(root);
    assertFalse(pandasia.isRegisteredValidator(cAddress));

    // Signature generated on wallet.avax.network
    uint8 v = 0;
    bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
    bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

    vm.prank(cAddress);
    pandasia.unregisterPChainAddr();
    assertFalse(pandasia.isRegisteredValidator(cAddress));

    vm.prank(cAddress);
    pandasia.registerPChainAddr(v, r, s, proof);
    assertTrue(pandasia.isRegisteredValidator(cAddress));

    vm.prank(cAddress);
    pandasia.unregisterPChainAddr();
    assertFalse(pandasia.isRegisteredValidator(cAddress));
  }

  /**************************************************************************************************************************************/
  /*** Merkle Tests                                                                                                                   ***/
  /**************************************************************************************************************************************/

  // Test against known-good message using wallet.avax.network
  function testMessageHashAlgo() public {
    string memory caddrStr = "0x63682bdc5f875e9bf69e201550658492c9763f89";
    bytes memory header = bytes("\x1AAvalanche Signed Message:\n");
    // len of an ascii addr is 42 bytes
    uint32 addrLen = 42;
    bytes32 message = sha256(abi.encodePacked(header, addrLen, caddrStr));
    // Signing "0x63682bdc5f875e9bf69e201550658492c9763f89" on https://wallet.avax.network/ gives 0x68c88e730eced13ee4a68eff65d3d250bb7b0f27c1cb4c8e20c52514d45d9390
    assertEq(message, 0x68c88e730eced13ee4a68eff65d3d250bb7b0f27c1cb4c8e20c52514d45d9390);
  }

  // Ensure pandasia is hashing messages the same as the avax wallet
  function testMessageHash() public {
    bytes32 msgHash = pandasia.hashChecksummedMessage(cAddress);
    // Use gogotools ggt utils msgdigest 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4 to get hashFromWallet
    bytes32 hashFromWallet = 0x1627404f56d262d498dd02e4fd880f38fafd6ed220dc9a3c3c9e75fe9846dc30;
    assertEq(msgHash, hashFromWallet);
  }

  function testRecoverPubKey() public {
    // signer: P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww
    // bech32 decode gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww => 0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02
    // actualPchainAddrBytes is what we need to get from recovering a msg + signature
    // using that we can then check the Merkle tree to see if that addr was a validator

    bytes32 msgHash = pandasia.hashChecksummedMessage(cAddress);
    // Known-good sig from wallet.avax.network
    // signer: P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww msg: 0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4
    // sig: 24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx
    // cb58 decode of sig gives (r, s, v) 6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337 39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a 00
    uint8 v = 0;
    uint256 r = uint256(bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337));
    uint256 s = uint256(bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a));
    (uint256 x, uint256 y) = SECP256K1.recover(uint256(msgHash), v, uint256(r), uint256(s));

    address pAddress = pandasia.pubKeyBytesToAvaAddressBytes(x, y);
    assertEq(pAddress, pAddressBytes);
  }

  function testProof() public {
    bytes32 root = bytes32(0xfcd7a701a861392c67cef2baaaf08063a1214f4ba4c3948c45b8e2008d28a35e);
    bytes32[] memory proof = new bytes32[](1);
    proof[0] = bytes32(0xa7409058568815d08a7ad3c7d4fd44cf1dec90c620cb31e55ad24c654f7ba34f);
    bool ok = pandasia.verify(root, cAddress, proof);
    assertTrue(ok);
  }

  // Merkle tree generated from P-chain snapshot
  function testLargeProof() public {
    bytes32 root = bytes32(0x1e2e8273c778f47235885ca5ab7db6645b768acc25c307b3069e6dc403b09551);
    bytes32[] memory proof = new bytes32[](18);
    proof[0] = bytes32(0x2fac3b7f6bf64f871937debc74f5086362a5da55855b0af732f80c9c0e547a17);
    proof[1] = bytes32(0xd970444f7cc592752242e22ecb13d739dc1734c7b00ba05aa86c34d831b3401b);
    proof[2] = bytes32(0xf62d36803f79fd136af598f33d9c26e551a8582e0374455ae5e6711bcaf09598);
    proof[3] = bytes32(0xb6eab16254380ca1c1189266b6de43de1422d9496ccd29de6e214bdb67f58d03);
    proof[4] = bytes32(0x653898d0812b7eb54797d1b056e5f51c085e8e652a4000b0ed1e094f4158fe31);
    proof[5] = bytes32(0xe3e95697bbf724316cfefff313dcf1efd335b15c9757af110fbf3c250a721b70);
    proof[6] = bytes32(0x61fb7aa6878508148d2b1fe2e548a49e330d0eb9ebfaa86a0e51bc6482398fc1);
    proof[7] = bytes32(0x16e09fec546547fb7438719cf49b47471096d9e8e1bed30d1e9a0c20722e9a07);
    proof[8] = bytes32(0x15a5a577ff7099aed0ace7301d0bc93653565be6396c2c5a7c3435e827cc0199);
    proof[9] = bytes32(0x9c00615deccec18f4c4d27c776d6d91bcd9c2c6d20c0c9e367786f8d705bda02);
    proof[10] = bytes32(0xed2381b58c990a90b5dbc183663b44d32d096b80f426d8a29b529f36ccc428f9);
    proof[11] = bytes32(0xb712d4310d1e6ff92259fab89bbdc99b2cc420f6c82113d572625ca12d4ab9b4);
    proof[12] = bytes32(0xb2dc703fcbeb0dc5eb85f6205f422d43f9dc472aa4a197f1eb4fe9a76a0dbd11);
    proof[13] = bytes32(0x3721827324639c3be37f7f8e220a4ffeb30754d6887a92e1585ebe2ed6b3a096);
    proof[14] = bytes32(0x3f799f2cde9ef1d50e6d0a3ca2a1a3ada14801396c9988b2156730be15a6d85c);
    proof[15] = bytes32(0xfd306cba9e1e94a29c9397fc33956b2af4424d0c1fd8dbfeb2b495ddb2deb48a);
    proof[16] = bytes32(0x8650d014c99057e6c2e56e99fc2abe8875141a600a0a34c9d30ddd4d7468a5b5);
    proof[17] = bytes32(0x8789ac606f7545a65049120d57b3b06b839ba6cf0a341a8d6edc0d7922d281e6);
    address account = address(0x7bDF8B86561d98d77e5BFc4B0eD20b7beB8fCdb6);
    startMeasuringGas("verify large proof");
    bool ok = pandasia.verify(root, account, proof);
    stopMeasuringGas();
    assertTrue(ok);
  }

  function testRecoverMessage() public {
    // Signature generated on wallet.avax.network
    uint8 v = 0;
    bytes32 r = bytes32(0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337);
    bytes32 s = bytes32(0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a);

    // not using our c-chain addr the recovered address will be wrong
    assertFalse(pandasia.recoverMessage(v, r, s) == pAddressBytes);

    vm.prank(cAddress);
    assertEq(pandasia.recoverMessage(v, r, s), pAddressBytes);
  }

  /**************************************************************************************************************************************/
  /*** Upgrade Tests                                                                                                                  ***/
  /**************************************************************************************************************************************/

  function testUpgrades() public {
    Pandasia pandasiaImpl = new Pandasia();
    TransparentUpgradeableProxy pandasiaProxy = new TransparentUpgradeableProxy(address(pandasiaImpl), address(this), bytes(""));
    Pandasia v1 = Pandasia(payable(pandasiaProxy));
    v1.initialize();

    bytes32 adminSlot = vm.load(address(v1), ERC1967Utils.ADMIN_SLOT);
    ProxyAdmin proxyAdmin = ProxyAdmin(address(uint160(uint256(adminSlot))));

    address storageContract = address(0x01);
    v1.setStorageContract(storageContract);

    // Can't initialize the implementation contract
    vm.expectRevert(Initializable.InvalidInitialization.selector);
    pandasiaImpl.initialize();

    // Can't re-initialize the same proxy contract
    vm.expectRevert(Initializable.InvalidInitialization.selector);
    v1.initialize();

    // ownership checks
    assertEq(address(this), proxyAdmin.owner());

    Pandasia v2Implementation = new Pandasia();

    vm.prank(address(0x01));
    bytes4 selector = bytes4(keccak256("OwnableUnauthorizedAccount(address)"));
    vm.expectRevert(abi.encodeWithSelector(selector, address(0x01)));
    proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(pandasiaProxy)), address(v2Implementation), new bytes(0));

    proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(v1)), address(v2Implementation), new bytes(0));

    assertEq(v1.storageContract(), storageContract);
  }

  /**************************************************************************************************************************************/
  /*** Helpers                                                                                                                        ***/
  /**************************************************************************************************************************************/

  string private checkpointLabel;
  uint256 private checkpointGasLeft;

  function startMeasuringGas(string memory label) internal virtual {
    checkpointLabel = label;
    checkpointGasLeft = gasleft();
  }

  function stopMeasuringGas() internal virtual {
    uint256 checkpointGasLeft2 = gasleft();

    string memory label = checkpointLabel;

    emit log_named_uint(string(abi.encodePacked(label, " Gas")), checkpointGasLeft - checkpointGasLeft2);
  }
}
