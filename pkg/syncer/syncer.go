package syncer

import (
	"context"
	"fmt"

	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

// Will run as a cron and keep DB up to date with chain
func SyncPChain(ctx context.Context, queries *db.Queries, uri string) error {
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

	for batch := int64(0); batch < batches; batch++ {
		height := startHeight + (batch * batchSize)
		slog.Info("superfetch", "batch", batch, "height", height)

		blks, err := pchain.SuperFetchBlocks(uri, int(height), int(batchSize))
		if err != nil {
			return fmt.Errorf("pchain.FetchBlocks %w", err)
		}

		for _, b := range blks {
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

				if t.TypeId == pchain.RewardValidatorTxId {
					err := queries.MarkAsRewarded(ctx, t.EarnedRewardForTxId)
					if err != nil {
						slog.Error("queries.MarkAsRewarded", "height", b.Height, "tid", t.Id, "err", err)
					}
				}

			}
		}
	}

	return nil
}
