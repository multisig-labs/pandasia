#!/usr/bin/env deno run --allow-env --allow-net
import { createWalletClient, http } from "https://esm.sh/viem@1.6.4";
import { mnemonicToAccount } from "https://esm.sh/viem@1.6.4/accounts";
import { chains } from "http://localhost:8000/js/chains.js";
const abi = await fetch("http://localhost:8000/js/abi.json").then((res) =>
	res.json()
);

const pandasiaAddr = Deno.env.get("PANDASIA_ADDR");
const pandasiaUrl = Deno.env.get("PANDASIA_URL");
const mnemonic = Deno.env.get("MNEMONIC");
const account = mnemonicToAccount(mnemonic, { addressIndex: 0 });

const chain = chains["anvil"];

const client = createWalletClient({
	account,
	chain,
	transport: http(chain.rpcUrls.public.http[0]),
});

// TODO

const { request } = await client.writeContract({
	address: pandasiaAddr,
	abi: abi.abi,
	functionName: "newAirdrop",
	args: [],
	account,
});
