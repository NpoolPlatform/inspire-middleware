package cashcontrol

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"
)

func (h *Handler) DeleteCashControl(ctx context.Context) (*npool.Scope, error) {
	info, err := h.GetCashControl(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := cashcontrolcrud.UpdateSet(
			cli.CashControl.UpdateOneID(*h.ID),
			&cashcontrolcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
