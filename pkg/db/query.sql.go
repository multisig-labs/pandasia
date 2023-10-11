// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createMerkleTree = `-- name: CreateMerkleTree :exec
INSERT OR IGNORE INTO merkle_trees (
	height, tree_type, tree, description
) VALUES (
 ?, ?, ?, ?
) RETURNING id
`

type CreateMerkleTreeParams struct {
	Height      int64
	TreeType    string
	Tree        string
	Description string
}

func (q *Queries) CreateMerkleTree(ctx context.Context, arg CreateMerkleTreeParams) error {
	_, err := q.db.ExecContext(ctx, createMerkleTree,
		arg.Height,
		arg.TreeType,
		arg.Tree,
		arg.Description,
	)
	return err
}

const createTx = `-- name: CreateTx :exec
INSERT OR IGNORE INTO txs (
  id, height, block_id, type_id, unsigned_tx, signer_addr_p
) VALUES (
  ?, ?, ?, ?, ?, ?
)
`

type CreateTxParams struct {
	ID          string
	Height      int64
	BlockID     string
	TypeID      int64
	UnsignedTx  string
	SignerAddrP string
}

func (q *Queries) CreateTx(ctx context.Context, arg CreateTxParams) error {
	_, err := q.db.ExecContext(ctx, createTx,
		arg.ID,
		arg.Height,
		arg.BlockID,
		arg.TypeID,
		arg.UnsignedTx,
		arg.SignerAddrP,
	)
	return err
}

const findAddrsForMerkleTree = `-- name: FindAddrsForMerkleTree :many
SELECT DISTINCT rewards_addr
FROM txs
WHERE has_earned_reward = 1
AND type_id = ?
AND height <= ?
ORDER BY rewards_addr
`

type FindAddrsForMerkleTreeParams struct {
	TypeID int64
	Height int64
}

func (q *Queries) FindAddrsForMerkleTree(ctx context.Context, arg FindAddrsForMerkleTreeParams) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, findAddrsForMerkleTree, arg.TypeID, arg.Height)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var rewards_addr sql.NullString
		if err := rows.Scan(&rewards_addr); err != nil {
			return nil, err
		}
		items = append(items, rewards_addr)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findMerkleTreeByRoot = `-- name: FindMerkleTreeByRoot :one
SELECT id, height, tree_type, root, tree, description
FROM merkle_trees
WHERE root = ?
`

type FindMerkleTreeByRootRow struct {
	ID          int64
	Height      int64
	TreeType    string
	Root        string
	Tree        string
	Description string
}

func (q *Queries) FindMerkleTreeByRoot(ctx context.Context, root string) (FindMerkleTreeByRootRow, error) {
	row := q.db.QueryRowContext(ctx, findMerkleTreeByRoot, root)
	var i FindMerkleTreeByRootRow
	err := row.Scan(
		&i.ID,
		&i.Height,
		&i.TreeType,
		&i.Root,
		&i.Tree,
		&i.Description,
	)
	return i, err
}

const findMerkleTreeByType = `-- name: FindMerkleTreeByType :one
SELECT id, height, tree_type, root, tree, description
FROM merkle_trees
WHERE tree_type = ?
ORDER BY height
LIMIT 1
`

type FindMerkleTreeByTypeRow struct {
	ID          int64
	Height      int64
	TreeType    string
	Root        string
	Tree        string
	Description string
}

func (q *Queries) FindMerkleTreeByType(ctx context.Context, treeType string) (FindMerkleTreeByTypeRow, error) {
	row := q.db.QueryRowContext(ctx, findMerkleTreeByType, treeType)
	var i FindMerkleTreeByTypeRow
	err := row.Scan(
		&i.ID,
		&i.Height,
		&i.TreeType,
		&i.Root,
		&i.Tree,
		&i.Description,
	)
	return i, err
}

const listMerkleRoots = `-- name: ListMerkleRoots :many
SELECT id, height, tree_type, root, description
FROM merkle_trees
ORDER BY height DESC
`

type ListMerkleRootsRow struct {
	ID          int64
	Height      int64
	TreeType    string
	Root        string
	Description string
}

func (q *Queries) ListMerkleRoots(ctx context.Context) ([]ListMerkleRootsRow, error) {
	rows, err := q.db.QueryContext(ctx, listMerkleRoots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMerkleRootsRow
	for rows.Next() {
		var i ListMerkleRootsRow
		if err := rows.Scan(
			&i.ID,
			&i.Height,
			&i.TreeType,
			&i.Root,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const maxHeight = `-- name: MaxHeight :one
SELECT cast(COALESCE(max(height),0) as integer) as maxheight from txs
`

func (q *Queries) MaxHeight(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, maxHeight)
	var maxheight int64
	err := row.Scan(&maxheight)
	return maxheight, err
}
