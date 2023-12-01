#!/usr/bin/env deno run --allow-env --allow-net

// Example of how to use JS lib to get a proof from pandasia API server and submit to contract

import { createWalletClient, http } from "https://esm.sh/viem@1.6.4";
import { mnemonicToAccount } from "https://esm.sh/viem@1.6.4/accounts";
const abi = await fetch("http://localhost:8000/js/abi.json").then((res) =>
  res.json(),
);

const fork = {
  id: 43114,
  name: "anvil",
  network: "anvil",
  nativeCurrency: {
    decimals: 18,
    name: "AVAX",
    symbol: "AVAX",
  },
  rpcUrls: {
    public: { http: [Deno.env.get("ETH_RPC_URL")] },
    default: { http: [Deno.env.get("ETH_RPC_URL")] },
  },
  blockExplorers: {
    default: { name: "Blockscout", url: "https://todo.com" },
  },
  testnet: true,
};

// pass in addr P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww on command line
const addrToRegister = Deno.args[0];
// pass in sig from wallet.avax.network 24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx
const sigToVerify = Deno.args[1];

const pandasiaAddr = Deno.env.get("PANDASIA_ADDR");
const pandasiaUrl = Deno.env.get("PANDASIA_URL");
const mnemonic = Deno.env.get("MNEMONIC");
const account = mnemonicToAccount(mnemonic, { addressIndex: 0 });

const chain = fork;

const client = createWalletClient({
  account,
  chain,
  transport: http(chain.rpcUrls.public.http[0]),
});

console.log(`Register ${addrToRegister} with sig ${sigToVerify}`);
let resp = await fetch(`${pandasiaUrl}/trees`);
const trees = await resp.json();
const root = trees[0].Root;

resp = await fetch(
  `${pandasiaUrl}/proof/${root}?addr=${addrToRegister}&sig=${sigToVerify}`,
);
const proof = await resp.json();
console.log(proof);

const { request } = await client.writeContract({
  address: pandasiaAddr,
  abi: abi.abi,
  functionName: "registerPChainAddr",
  args: [proof.SigV, proof.SigR, proof.SigS, proof.Proof],
  account,
});
