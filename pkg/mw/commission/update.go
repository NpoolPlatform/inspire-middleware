package commission

import (
	"context"

	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
)

func (h *Handler) UpdateCommission(ctx context.Context) (*npool.Commission, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissioncrud.UpdateSet(
			tx.Commission.UpdateOneID(*h.ID),
			&commissioncrud.Req{
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
