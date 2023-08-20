-- name: MaxHeight :one
SELECT cast(COALESCE(max(height),0) as integer) as maxheight from txs;

-- name: CreateTx :exec
INSERT OR IGNORE INTO txs (
  id, height, block_id, type_id, unsigned_tx, signer_addr_p
) VALUES (
  ?, ?, ?, ?, ?, ?
);

-- name: FindAddrsForMerkleTree :many
SELECT DISTINCT rewards_addr
FROM txs
WHERE has_earned_reward = 1
AND type_id = ?
AND height <= ?
ORDER BY rewards_addr;

-- name: CreateMerkleTreeAndReturnId :one
INSERT INTO merkle_trees (
	height, tree_type, tree
) VALUES (
 ?, ?, ?
) RETURNING id;

-- name: FindMerkleTreeByType :one
SELECT id, height, tree_type, root, tree
FROM merkle_trees
WHERE tree_type = ?
ORDER BY height
LIMIT 1;

-- name: FindMerkleTreeByRoot :one
SELECT id, height, tree_type, root, tree
FROM merkle_trees
WHERE root = ?;
