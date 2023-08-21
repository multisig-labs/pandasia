#!/usr/bin/env deno run --allow-env --allow-net
import {
	createWalletClient,
	http,
	WalletClient,
} from "https://esm.sh/viem@1.6.4";
import { Chain, toHex } from "https://esm.sh/viem@1.6.4";
import { mnemonicToAccount } from "https://esm.sh/viem@1.6.4/accounts";

const addrToVerify = Deno.args[0];

const ethRpcUrl = Deno.env.get("ETH_RPC_URL");
const pandasiaAddr = Deno.env.get("PANDASIA_ADDR");
const pandasiaUrl = Deno.env.get("PANDASIA_URL");
const mnemonic = Deno.env.get("MNEMONIC");
const account = mnemonicToAccount(mnemonic, { addressIndex: 0 });

const chain = {
	id: 31337,
	name: "anvil",
	network: "anvil",
	nativeCurrency: {
		decimals: 18,
		name: "AVAX",
		symbol: "AVAX",
	},
	rpcUrls: {
		public: { http: [ethRpcUrl] },
		default: { http: [ethRpcUrl] },
	},
	blockExplorers: {
		default: { name: "Blockscout", url: "https://todo.com" },
	},
	testnet: true,
} as const satisfies Chain;

const client = createWalletClient({
	account,
	chain,
	transport: http(chain.rpcUrls.public.http[0]),
});

const pandasiaAbi = [
	{
		inputs: [
			{
				internalType: "uint8",
				name: "v",
				type: "uint8",
			},
			{
				internalType: "bytes32",
				name: "r",
				type: "bytes32",
			},
			{
				internalType: "bytes32",
				name: "s",
				type: "bytes32",
			},
			{
				internalType: "bytes32[]",
				name: "proof",
				type: "bytes32[]",
			},
		],
		name: "registerPChainAddr",
		outputs: [],
		stateMutability: "nonpayable",
		type: "function",
	},
];

const { request } = await client.writeContract({
	address: pandasiaAddr,
	abi: pandasiaAbi,
	functionName: "registerPChainAddr",
	args: [
		"0x00",
		"0x6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337",
		"0x39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a",
		[
			"0x20ab9da0017c72517588d8ace595ebf2398049bcfbb266848e31fefe7bc8b142",
			"0x2b52bc313d3d9a86d4b41b6a795a8a6c61ab8074691a74c26df2cb2d4a3d4e12",
			"0x4443334965035b66270a2e6f8a753c9901726a800506685bc3df129d3e931c28",
			"0x394863a96a875773bf1ed12b0e66bce0ee99d79eb60f08170337f3092a8960f2",
			"0x3524b800264fb6439eae944b920793eafcb59dcc5b062c5907bee8cadcd8ec54",
			"0x25991bbbeb0f65fc62adddfa537a99cace78cc1762b697c5a10113f5d9589a36",
			"0x282fef08f12424db2405b994a551fd5b10dd9abd68cea5ab4b24713136bb5314",
			"0xca342ecfa04e605efb7b1daff7cf5fbc6694ad0017463d318c903465529928d7",
			"0x611e28c1cba9cb98dc0e29c55c0b61773243c86d7e2aad5ec6241c1c4ea5e8d3",
			"0x708659e9ebb82629b7eb4497f2f25a797f266d725ec69fd18ad4aae967808c12",
			"0xa23787ca0ec5cf330142718ba59c38a65812e9e4dc7622fa5afb7304e64f7c03",
			"0x6f258b7e118c298281291e3446a03f6202edeb5ffc605cac29fcb315c5badd14",
			"0xc3a21d4312fca63b56ab60c89c54037d2d5f28b7d113c3b2ceca5a676379ff42",
			"0x4d22deb9967ae8793f6f289519a0fc9d63b752acb91d11d92b3eb3cc24426b4e",
		],
	],
	account,
});

// console.log(`Verify ${addrToVerify}`);
// let resp = await fetch(`${pandasiaUrl}/trees`);
// const trees = await resp.json();
// const root = trees[0].Root;

// resp = await fetch(`${pandasiaUrl}/trees/${root}/${addrToVerify}`);
// const proof = await resp.json();
// console.log(proof);
