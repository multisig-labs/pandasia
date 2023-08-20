package syncer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

// Sync the DB with the current state of the P-chain
func SyncPChain(ctx context.Context, queries *db.Queries, uri string, progressFn func(tot int, n int)) error {
	if progressFn == nil {
		progressFn = func(tot int, n int) {}
	}
	batchSize := int64(1000)
	keepTypeIds := []int64{pchain.RewardValidatorTxId, pchain.AddValidatorTxId, pchain.AddDelegatorTxId}

	startHeight, err := queries.MaxHeight(ctx)
	if err != nil {
		return fmt.Errorf("queries.MaxHeight %w", err)
	}

	maxHeight := pchain.MaxHeight(uri)
	numBlksToFetch := (maxHeight - startHeight)

	batches := numBlksToFetch / batchSize
	if numBlksToFetch%batchSize != 0 {
		batches++
	}

	progressFn(int(numBlksToFetch), 0)

	for batch := int64(0); batch < batches; batch++ {
		height := startHeight + (batch * batchSize)
		// slog.Info("superfetch", "batch", batch, "height", height, "remaining", maxHeight-height)

		blks, err := pchain.SuperFetchBlocks(uri, int(height), int(batchSize))
		if err != nil {
			return fmt.Errorf("pchain.FetchBlocks %w", err)
		}

		for _, b := range blks {
			progressFn(int(numBlksToFetch), 1)
			for _, t := range b.Txs {
				// To save space we only keep the json for txs we are interested in
				unsignedTx := "{}"
				if slices.Contains(keepTypeIds, t.TypeId) {
					unsignedTx = t.UnsignedTx
				}

				dbTx := db.CreateTxParams{
					ID:          t.Id,
					Height:      int64(b.Height),
					BlockID:     b.Id,
					TypeID:      t.TypeId,
					UnsignedTx:  unsignedTx,
					SignerAddrP: t.RecoveredAddrP,
				}
				err := queries.CreateTx(ctx, dbTx)
				if err != nil {
					slog.Error("queries.CreateTx", "height", b.Height, "err", err)
				}
			}
		}
	}

	return nil
}

// Scan all txs and mark the validator txs that have earned rewards
func UpdateRewards(ctx context.Context, dbFile *sql.DB) error {
	// sqlc doesnt like this query so just run it manually
	sql := `
		UPDATE txs
		SET has_earned_reward = 1
		FROM txs AS t2
		WHERE txs.type_id = 12
		AND txs.id = t2.rewards_for_id;
		`
	_, err := dbFile.Exec(sql)
	if err != nil {
		return err
	}

	// num, err := r.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	return nil
}
