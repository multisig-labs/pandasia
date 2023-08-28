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
	rewards_for_id text GENERATED ALWAYS AS (CASE WHEN type_id = 20 THEN json_extract(unsigned_tx, '$.txID') ELSE null END) STORED, -- This tx signifies that another txid should earn rewards
	has_earned_reward integer,
  FOREIGN KEY(type_id) REFERENCES types(id)
) STRICT;
-- Some Sql browsers dont show generated columns, so make a view as well
CREATE VIEW txs_v AS SELECT * FROM txs;

CREATE INDEX txs_height ON txs(height);
CREATE INDEX txs_block_id ON txs(block_id);
CREATE INDEX txs_type_id ON txs(type_id);
CREATE INDEX txs_node_id ON txs(node_id);
CREATE INDEX txs_rewards_addr ON txs(rewards_addr);
CREATE INDEX txs_signer_addr_p ON txs(signer_addr_p);
CREATE INDEX txs_rewards_for_id ON txs(rewards_for_id);
CREATE INDEX txs_has_earned_reward ON txs(has_earned_reward);


CREATE TABLE merkle_trees (
	id integer PRIMARY KEY,
	height integer NOT NULL, -- pchain height tree was generated at
	tree_type text NOT NULL, -- validator or custom
	description text NOT NULL DEFAULT '',
	tree text NOT NULL, -- json of tree in OpenZeppelin format
	root text NOT NULL GENERATED ALWAYS AS (json_extract(tree, '$.tree[0]')) STORED,
	addrs_count integer NOT NULL GENERATED ALWAYS AS (json_array_length(tree, '$.values')) STORED
);
-- Some Sql browsers dont show generated columns, so make a view as well
CREATE VIEW merkle_trees_v AS SELECT * FROM merkle_trees;

CREATE UNIQUE INDEX merkle_trees_height_tree_type ON merkle_trees(height,tree_type);
CREATE UNIQUE INDEX merkle_trees_root ON merkle_trees(root);

-- Only keep the last N large validator trees to save space
CREATE TRIGGER merkle_trees_insert AFTER INSERT ON merkle_trees
WHEN (SELECT count(*) FROM merkle_trees WHERE NEW.tree_type = "validator") > 5
BEGIN
	DELETE FROM merkle_trees
	WHERE tree_type = "validator"
		AND height = (SELECT min(height) FROM merkle_trees WHERE tree_type = "validator");
END;


CREATE TABLE types (
  id integer PRIMARY KEY,
  name text NOT NULL
) STRICT;

INSERT INTO types (id, name) VALUES (0,  'ApricotProposalBlock');
INSERT INTO types (id, name) VALUES (1,  'ApricotAbortBlock');
INSERT INTO types (id, name) VALUES (2,  'ApricotCommitBlock');
INSERT INTO types (id, name) VALUES (3,  'ApricotStandardBlock');
INSERT INTO types (id, name) VALUES (4,  'ApricotAtomicBlock');
INSERT INTO types (id, name) VALUES (5,  'secp256k1fx.TransferInput');
INSERT INTO types (id, name) VALUES (6,  'secp256k1fx.MintOutput');
INSERT INTO types (id, name) VALUES (7,  'secp256k1fx.TransferOutput');
INSERT INTO types (id, name) VALUES (8,  'secp256k1fx.MintOperation');
INSERT INTO types (id, name) VALUES (9,  'secp256k1fx.Credential');
INSERT INTO types (id, name) VALUES (10, 'secp256k1fx.Input');
INSERT INTO types (id, name) VALUES (11, 'secp256k1fx.OutputOwners');
INSERT INTO types (id, name) VALUES (12, 'AddValidatorTx');
INSERT INTO types (id, name) VALUES (13, 'AddSubnetValidatorTx');
INSERT INTO types (id, name) VALUES (14, 'AddDelegatorTx');
INSERT INTO types (id, name) VALUES (15, 'CreateChainTx');
INSERT INTO types (id, name) VALUES (16, 'CreateSubnetTx');
INSERT INTO types (id, name) VALUES (17, 'ImportTx');
INSERT INTO types (id, name) VALUES (18, 'ExportTx');
INSERT INTO types (id, name) VALUES (19, 'AdvanceTimeTx');
INSERT INTO types (id, name) VALUES (20, 'RewardValidatorTx');
INSERT INTO types (id, name) VALUES (21, 'stakeable.LockIn');
INSERT INTO types (id, name) VALUES (22, 'stakeable.LockOut');
INSERT INTO types (id, name) VALUES (23, 'RemoveSubnetValidatorTx');
INSERT INTO types (id, name) VALUES (24, 'TransformSubnetTx');
INSERT INTO types (id, name) VALUES (25, 'AddPermissionlessValidatorTx');
INSERT INTO types (id, name) VALUES (26, 'AddPermissionlessDelegatorTx');
INSERT INTO types (id, name) VALUES (27, 'EmptyProofOfPossession');
INSERT INTO types (id, name) VALUES (28, 'BLSProofOfPossession  ');
INSERT INTO types (id, name) VALUES (29, 'BanffProposalBlock');
INSERT INTO types (id, name) VALUES (30, 'BanffAbortBlock');
INSERT INTO types (id, name) VALUES (31, 'BanffCommitBlock');
INSERT INTO types (id, name) VALUES (32, 'BanffStandardBlock');
