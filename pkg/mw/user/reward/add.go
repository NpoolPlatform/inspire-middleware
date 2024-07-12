package reward

import (
	"context"

	userrewardcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/user/reward"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type addHandler struct {
	*Handler
	sql  string
	info *npool.UserReward
}

func (h *addHandler) addUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *addHandler) getUserCoinReward(ctx context.Context) error {
	h.Conds = &userrewardcrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	h.Limit = int32(1)
	infos, _, err := h.GetUserRewards(ctx)
	if err != nil {
		return err
	}
	if len(infos) == 0 {
		return nil
	}
	h.info = infos[0]
	id := infos[0].ID
	h.ID = &id
	return nil
}

func (h *addHandler) calculateReward() error {
	if h.info == nil {
		if h.ActionCredits == nil {
			value := decimal.NewFromInt32(0)
			h.ActionCredits = &value
		}
		if h.CouponAmount == nil {
			value := decimal.NewFromInt32(0)
			h.CouponAmount = &value
		}
		if h.CouponCashableAmount == nil {
			value := decimal.NewFromInt32(0)
			h.CouponCashableAmount = &value
		}
		if h.EntID == nil {
			id := uuid.New()
			h.EntID = &id
		}
		sql := h.constructCreateSQL()
		h.sql = sql
		return nil
	}

	if h.ActionCredits != nil {
		credits, err := decimal.NewFromString(h.info.ActionCredits)
		if err != nil {
			return err
		}
		newCredits := h.ActionCredits.Add(credits)
		h.ActionCredits = &newCredits
	}
	if h.CouponAmount != nil {
		couponAmount, err := decimal.NewFromString(h.info.CouponAmount)
		if err != nil {
			return err
		}
		newCouponAmount := h.CouponAmount.Add(couponAmount)
		h.CouponAmount = &newCouponAmount
	}
	if h.CouponCashableAmount != nil {
		couponCashableAmount, err := decimal.NewFromString(h.info.CouponCashableAmount)
		if err != nil {
			return err
		}
		newCouponCashableAmount := h.CouponCashableAmount.Add(couponCashableAmount)
		h.CouponCashableAmount = &newCouponCashableAmount
	}
	sql, err := h.constructUpdateSQL()
	if err != nil {
		return err
	}
	h.sql = sql
	return nil
}

func (h *Handler) AddUserReward(ctx context.Context) error {
	handler := &addHandler{
		Handler: h,
	}

	if err := handler.getUserCoinReward(ctx); err != nil {
		return err
	}

	if err := handler.calculateReward(); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.addUserReward(_ctx, tx)
	})
}
