package cashcontrol

import (
	"context"

	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
)

func (h *Handler) UpdateCashControl(ctx context.Context) (*npool.CashControl, error) {
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := cashcontrolcrud.UpdateSet(
			cli.CashControl.UpdateOneID(*h.ID),
			&cashcontrolcrud.Req{
				Value: h.Value,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCashControl(ctx)
}
