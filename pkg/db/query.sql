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
WHERE type_id = ?
AND height <= ?
AND validator_start_ts < strftime('%s','now')
AND validator_end_ts > strftime('%s','now')
ORDER BY rewards_addr;

-- name: CreateMerkleTree :exec
INSERT OR IGNORE INTO merkle_trees (
	height, tree_type, tree, description
) VALUES (
 ?, ?, ?, ?
) RETURNING id;

-- name: FindPchainAddr :one
SELECT count(*)
FROM txs
WHERE rewards_addr = ?
AND type_id = ?
AND height <= ?
AND validator_start_ts < strftime('%s','now')
AND validator_end_ts > strftime('%s','now');

-- name: FindMerkleTreeByType :one
SELECT id, height, tree_type, root, tree, description
FROM merkle_trees
WHERE tree_type = ?
ORDER BY height
LIMIT 1;

-- name: FindMerkleTreeByRoot :one
SELECT id, height, tree_type, root, tree, description
FROM merkle_trees
WHERE root = ?;

-- name: ListMerkleRoots :many
SELECT id, height, tree_type, root, description
FROM merkle_trees
ORDER BY height DESC;
