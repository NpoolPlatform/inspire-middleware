package commission

import (
	"context"
	"time"

	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) DeleteCommission(ctx context.Context) error {
	info, err := h.GetCommission(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissioncrud.UpdateSet(
			tx.Commission.UpdateOneID(*h.ID),
			&commissioncrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
