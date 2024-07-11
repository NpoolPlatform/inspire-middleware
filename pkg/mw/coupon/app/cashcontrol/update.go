package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/cashcontrol"
)

func (h *Handler) UpdateCashControl(ctx context.Context) (*npool.CashControl, error) {
	info, err := h.GetCashControl(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid cashcontrol")
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := cashcontrolcrud.UpdateSet(
			cli.CashControl.UpdateOneID(*h.ID),
			&cashcontrolcrud.Req{
				Value: h.Value,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCashControl(ctx)
}
