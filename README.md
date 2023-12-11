# Pandasia.io

One of the Charities, (Πανδαισία) "banquet for everyone"

![Pandasia](docs/pandasia.jpg)

# Problem

<<<<<<< HEAD
Be “the place” where user’s can come to “claim” their validator node, cryptographically tying their C-chain addr to a P-chain addr. Then build web UI for projects to come and airdrop tokens / NFTs to the “verified” validator community. Of course any Minipool operator will automatically be a "verified" validator.
=======
The primary way to bootstrap your novel blockchain is through an airdrop to as wide and decentralized set of node
operators as possible. But because of the avalanche network design, this is very tricky to do. Validators exist on
the P-chain, but airdrops, wallets, and smart contracts exist on the C-Chain. There is no easy way to link a validators
profile on the P-chain to their C-Chain wallets. BEcause of that there isn't actually an easy way for subnets
to airdrop tokens to validator nodes which is one key blocker to helping subnets grow through decentralization.
Introducing Pandasia.

> > > > > > > 47727a6 (Update README)

# Idea

Pandasia is the tool to help projects airdrop to AVAX validators by allowing validators to verify ownership of
their node.

![](docs/pandasia-ui.png)

## Technical Approach

### Registration

Users register with Pandasia by signing a message with their P-Chain rewards address. That message should be
the C-Chain address they want to use on Pandasia.

To verify a validator, we have them sign a message on wallet.avax.network with their P-Chain address, then construct
that same message in our smart contracts. If the signatures match they're allowed to register.

### User creates signed message

When you sign a message with [wallet.avax.network](https://wallet.avax.network), it prepends your message with additional information before hashing. For example we want the user to sign a message with their C-chain address, i.e. `0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4`. So the actual messgage the wallet constructs for the user to sign is:

`\x1AAvalanche Signed Message:\n\x0000002A0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4`

where `\x1A` is 26, the length of `Avalanche Signed Message:\n` and `\x0000002A` is 42, the length of the text string `0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4`. The wallet then takes the `sha256` hash of those bytes, and that is the message that gets signed.

### Reconstructing the signed message

Using signature parts and msg.sender, we create the P-Chain address that will be registered. If the provided signature
and our created message don't mesh, we won't be able to construct a valid P-Chain address.

We do this in two parts:

1. We want to verify that the user signed their correct C-chain address, and not somebody elses. (After all its
   whatever they type in the message box so they can of course lie). We do this by constructing the message ourselves, in
   the Solidity contract, and using `msg.sender` as the C-chain address.

2. Then apply the user-provided signature against **our** message. Since the user cannot forge `msg.sender` we are sure that the P-chain key did in fact sign a message containing `msg.sender`.

Decomposing the signature is a little tricky, because the `ecrecover` precompile will take a signature, a message hash, and return an _ethereum_ address that signed it. But in our case a P-chain key signed the message, so `ecrecover` will give us the wrong address because the P-Chain address is derived differently than eth addresses.
So we use a Solidity library `SECP256K1` to do this instead of the `ecrecover` precompile.

The C-chain address the user types in the message box MUST be of the mixed-case, checksummed variety. (Remember that an
Ethereum address has a clever built-in checksum that uses the case of the various letters as the checksum.) If they use
an all lowercase address (i.e. `0x0961ca10d49b9b8e371aa0bcf77fe5730b18f2e4`) then our system wont work, as we are
expecting mixed case.

### Tracking Validators

So now we have a method for cryptographically linking a C-chain address to a P-chain address, the final piece to the
puzzle is determining if a specific P-Chain address was actually running a validator at one point. We cannot determine
this from Solidity, since the C-chain cannot query the P-chain for information.

The validator data must be collected off-chain, so we have built a Go program that slurps in the entire P-chain into a
SQLite DB, and tags addresses that have been used as a validator rewards address at any time in the past. The program
will also periodically create a giant Merkle Tree with all of these addresses, and post the merkle root to the Pandasia
contract. It will also provide an API so that a user can obtain the necessary merkle proof for their address, to submit
to the contract, which can verify their address and proof, against the merkle root.

With all those pieces in place, a user can now "register" with Pandasia and their C-chain address will be tagged as a verified validator. Projects can use Pandasia to distribute tokens to this group as airdrops to build community and reward those who are most heavily invested in the success of the Avalanche blockchain.

## Airdrops

Airdrops! Airdrops! Airdrops!

Projects create airdrops, depositing tokens to be distributed to registered validators!
The start and end time can be configued, as well as the distribution amount.

![](docs/airdrop.png)

# What's coming

Currently you need to contact MultisigLabs to create an airdrop, but we plan to open it up to the public soon.

We're also working on getting delegators involved in the airdrop action!
