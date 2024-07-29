package reward

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	userrewardcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/user/reward"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
	"github.com/shopspring/decimal"
)

type subHandler struct {
	*Handler
	sql  string
	info *npool.UserReward
}

func (h *subHandler) updateUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if _, err := rc.RowsAffected(); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

//nolint:dupl
func (h *subHandler) getUserCoinReward(ctx context.Context) error {
	h.Conds = &userrewardcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	h.Limit = int32(1)
	infos, _, err := h.GetUserRewards(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil
	}
	h.info = infos[0]
	id := infos[0].ID
	h.ID = &id
	return nil
}

func (h *subHandler) calculateReward() error {
	if h.ActionCredits != nil {
		credits, err := decimal.NewFromString(h.info.ActionCredits)
		if err != nil {
			return wlog.WrapError(err)
		}
		newCredits := h.ActionCredits.Sub(credits)
		h.ActionCredits = &newCredits
		if h.ActionCredits.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid actioncredits")
		}
	}
	if h.CouponAmount != nil {
		couponAmount, err := decimal.NewFromString(h.info.CouponAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		newCouponAmount := h.CouponAmount.Sub(couponAmount)
		h.CouponAmount = &newCouponAmount
		if h.CouponAmount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid couponamount")
		}
	}
	if h.CouponCashableAmount != nil {
		couponCashableAmount, err := decimal.NewFromString(h.info.CouponCashableAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		newCouponCashableAmount := h.CouponCashableAmount.Sub(couponCashableAmount)
		h.CouponCashableAmount = &newCouponCashableAmount
		if h.CouponCashableAmount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid couponcashableamount")
		}
	}
	sql, err := h.constructUpdateSQL()
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sql = sql
	return nil
}

func (h *Handler) SubUserReward(ctx context.Context) error {
	handler := &subHandler{
		Handler: h,
	}

	err := handler.getUserCoinReward(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if handler.info == nil {
		return nil
	}
	if err := handler.calculateReward(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateUserReward(_ctx, tx)
	})
}
