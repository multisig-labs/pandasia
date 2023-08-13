-- name: MaxHeight :one
SELECT cast(COALESCE(max(height),0) as integer) as maxheight from txs;

-- name: CreateTx :exec
INSERT OR IGNORE INTO txs (
  id, height, block_id, type_id, unsigned_tx, signer_addr_p
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: MarkAsRewarded :exec
UPDATE txs
SET has_earned_reward = 1
WHERE id = ?;

-- name: FindAddrsForMerkleTree :many
SELECT DISTINCT rewards_addr
FROM txs
WHERE has_earned_reward = 1
AND type_id = ?
AND height <= ?
ORDER BY rewards_addr;

-- name: CreateMerkleRootAndReturnId :one
INSERT INTO merkle_roots (
	height, type, root
) VALUES (
 ?, ?, ?
) RETURNING id;

-- name: CreateMerkleProof :exec
INSERT INTO merkle_proofs (
	merkle_root_id, paddy, data, proof
) VALUES (
 ?, ?, ?, ?
);
