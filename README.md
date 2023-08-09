# Pandasia

One of the Charities, (Πανδαισία) "banquet for everyone"

# Validator Airdrop Research

# 💡Idea

Be “the place” where user’s can come to “claim” their validator node, tying their C-chain addr to a P-chain addr. Then build web UI for projects to come and airdrop tokens / NFTs to the “verified” validator community. Of course any Minipool operator will automatically get boosted rewards 😈

# Slurp

CLI to download and process the P-Chain into a SQLite DB

[https://github.com/multisig-labs/slurp](https://github.com/multisig-labs/slurp)

P-Chain DB as of 8-2023 (7.5G)

[](http://gogopool.s3.amazonaws.com/slurp-mainnet.db.7z)

## Technical Approach

1. User must go to old avax wallet and go to Advanced tab, and sign a message with their C-chain addr, using the P-chain addr that was a rewards addr for a validator at anytime in the past.

2. Now user comes to our new site, and pastes in the signature from step 1, then submits the tx and signs with Metamask

🥵 This is a confusing ask maybe for users, but not sure how else to do it

So the trick is we need to verify the P-chain sig in solidity. `ecrecover` in solidity returns an _address_ which is a C-chain address, but we need the public key itself so that we can derive the P-chain address.

[](https://github.com/0xcyphered/secp256k1-solidity/blob/main/contracts/SECP256K1.sol)

The next tricky part is to generate offline, a Merkle tree of all validator rewards addrs, and then post the root to our contract every day?. The contract can then verify that any given C-chain addr maps to a P-chain addr that is or is not in the Merkle tree.

[https://github.com/OpenZeppelin/merkle-tree](https://github.com/OpenZeppelin/merkle-tree)

[Utilities - OpenZeppelin Docs](https://docs.openzeppelin.com/contracts/4.x/api/utils#MerkleProof)

[https://github.com/bloq/sol-mass-payouts](https://github.com/bloq/sol-mass-payouts)
