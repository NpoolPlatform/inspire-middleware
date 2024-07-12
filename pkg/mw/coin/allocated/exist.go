package allocated

import (
	"context"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coin/allocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) ExistCoinAllocatedConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := allocatedcrud.SetQueryConds(
			cli.CoinAllocated.Query(),
			h.Conds,
		)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
