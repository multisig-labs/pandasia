// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
)

type MerkleTree struct {
	ID         int64
	Height     int64
	TreeType   string
	Tree       string
	Root       string
	AddrsCount int64
}

type Tx struct {
	ID               string
	Height           int64
	BlockID          string
	TypeID           int64
	UnsignedTx       string
	SignerAddrP      string
	Memo             sql.NullString
	NodeID           sql.NullString
	ValidatorStartTs sql.NullInt64
	ValidatorEndTs   sql.NullInt64
	ValidatorWeight  sql.NullInt64
	RewardsAddr      sql.NullString
	RewardsForID     sql.NullString
	HasEarnedReward  sql.NullInt64
}

type Type struct {
	ID   int64
	Name string
}
