package commission

import (
	"context"
	"fmt"

	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
)

func (h *Handler) UpdateCommission(ctx context.Context) (*npool.Commission, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissioncrud.UpdateSet(
			tx.Commission.UpdateOneID(*h.ID),
			&commissioncrud.Req{
				ID:              h.ID,
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				Threshold:       h.Threshold,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCommission(ctx)
}
