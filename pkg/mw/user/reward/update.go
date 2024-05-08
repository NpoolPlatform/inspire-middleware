package reward

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update user_rewards "
	if h.ActionCredits != nil {
		_sql += fmt.Sprintf("%vaction_credits = '%v'", set, *h.ActionCredits)
		set = ""
	}
	if h.CouponAmount != nil {
		_sql += fmt.Sprintf("%vcoupon_amount = '%v'", set, *h.CouponAmount)
		set = ""
	}
	if h.CouponCashableAmount != nil {
		_sql += fmt.Sprintf("%vcoupon_cashable_amount = '%v'", set, *h.CouponCashableAmount)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateUserReward(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetUserReward(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid userreward")
	}

	if err := handler.constructSQL(); err != nil {
		return err
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateUserReward(_ctx, tx)
	})
}
