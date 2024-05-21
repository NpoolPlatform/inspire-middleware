package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) ExistCashControlConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm, err := cashcontrolcrud.SetQueryConds(cli.CashControl.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
