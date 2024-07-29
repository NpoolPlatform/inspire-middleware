package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) updateUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateUserReward(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetUserReward(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid userreward")
	}

	sql, err := handler.constructUpdateSQL()
	if err != nil {
		return wlog.WrapError(err)
	}
	handler.sql = sql
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateUserReward(_ctx, tx)
	})
}
