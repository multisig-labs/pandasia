CREATE TABLE txs (
  id text PRIMARY KEY,
	height integer NOT NULL,
  block_id text NOT NULL,
  type_id integer NOT NULL,
  unsigned_tx text NOT NULL, -- JSON of the unsignedTx key
	signer_addr_p text NOT NULL DEFAULT '',
  memo text GENERATED ALWAYS AS (cast(unhex(substr(json_extract(unsigned_tx, '$.memo'),3)) AS text)) STORED,
  node_id text GENERATED ALWAYS AS (json_extract(unsigned_tx, '$.validator.nodeID')) STORED,
  validator_start_ts integer GENERATED ALWAYS AS (json_extract(unsigned_tx, '$.validator.start')) STORED,
  validator_end_ts integer GENERATED ALWAYS AS (json_extract(unsigned_tx, '$.validator.end')) STORED,
  validator_weight integer GENERATED ALWAYS AS (json_extract(unsigned_tx, '$.validator.weight')) STORED,
  rewards_addr text GENERATED ALWAYS AS (json_extract(unsigned_tx, '$.rewardsOwner.addresses[0]')) STORED,
	has_earned_reward integer,
  FOREIGN KEY(type_id) REFERENCES types(id)
) STRICT;

CREATE INDEX txs_height ON txs(height);
CREATE INDEX txs_block_id ON txs(block_id);
CREATE INDEX txs_type_id ON txs(type_id);
CREATE INDEX txs_node_id ON txs(node_id);
CREATE INDEX txs_rewards_addr ON txs(rewards_addr);
CREATE INDEX txs_signer_addr_p ON txs(signer_addr_p);
CREATE INDEX txs_has_earned_reward ON txs(has_earned_reward);

CREATE TABLE merkle_roots (
	id integer PRIMARY KEY,
	height integer NOT NULL,
	type text NOT NULL, -- validator or delegator
	root text NOT NULL
);

CREATE TABLE merkle_proofs (
	id integer PRIMARY KEY,
	merkle_root_id integer NOT NULL,
	paddy text NOT NULL, -- P-avax1blahblah
	data text NOT NULL, -- the 20 bytes of the address in hex
	proof text NOT NULL, -- hex proof data
	FOREIGN KEY(merkle_root_id) REFERENCES merkle_roots(id)
);


CREATE TABLE types (
  id integer PRIMARY KEY,
  name text NOT NULL
) STRICT;

INSERT INTO types (id, name) VALUES (0,  "ApricotProposalBlock");
INSERT INTO types (id, name) VALUES (1,  "ApricotAbortBlock");
INSERT INTO types (id, name) VALUES (2,  "ApricotCommitBlock");
INSERT INTO types (id, name) VALUES (3,  "ApricotStandardBlock");
INSERT INTO types (id, name) VALUES (4,  "ApricotAtomicBlock");
INSERT INTO types (id, name) VALUES (5,  "secp256k1fx.TransferInput");
INSERT INTO types (id, name) VALUES (6,  "secp256k1fx.MintOutput");
INSERT INTO types (id, name) VALUES (7,  "secp256k1fx.TransferOutput");
INSERT INTO types (id, name) VALUES (8,  "secp256k1fx.MintOperation");
INSERT INTO types (id, name) VALUES (9,  "secp256k1fx.Credential");
INSERT INTO types (id, name) VALUES (10, "secp256k1fx.Input");
INSERT INTO types (id, name) VALUES (11, "secp256k1fx.OutputOwners");
INSERT INTO types (id, name) VALUES (12, "AddValidatorTx");
INSERT INTO types (id, name) VALUES (13, "AddSubnetValidatorTx");
INSERT INTO types (id, name) VALUES (14, "AddDelegatorTx");
INSERT INTO types (id, name) VALUES (15, "CreateChainTx");
INSERT INTO types (id, name) VALUES (16, "CreateSubnetTx");
INSERT INTO types (id, name) VALUES (17, "ImportTx");
INSERT INTO types (id, name) VALUES (18, "ExportTx");
INSERT INTO types (id, name) VALUES (19, "AdvanceTimeTx");
INSERT INTO types (id, name) VALUES (20, "RewardValidatorTx");
INSERT INTO types (id, name) VALUES (21, "stakeable.LockIn");
INSERT INTO types (id, name) VALUES (22, "stakeable.LockOut");
INSERT INTO types (id, name) VALUES (23, "RemoveSubnetValidatorTx");
INSERT INTO types (id, name) VALUES (24, "TransformSubnetTx");
INSERT INTO types (id, name) VALUES (25, "AddPermissionlessValidatorTx");
INSERT INTO types (id, name) VALUES (26, "AddPermissionlessDelegatorTx");
INSERT INTO types (id, name) VALUES (27, "EmptyProofOfPossession");
INSERT INTO types (id, name) VALUES (28, "BLSProofOfPossession  ");
INSERT INTO types (id, name) VALUES (29, "BanffProposalBlock");
INSERT INTO types (id, name) VALUES (30, "BanffAbortBlock");
INSERT INTO types (id, name) VALUES (31, "BanffCommitBlock");
INSERT INTO types (id, name) VALUES (32, "BanffStandardBlock");
