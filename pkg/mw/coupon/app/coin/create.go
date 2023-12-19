package coin

import (
	"context"
	"fmt"

	couponcoincrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/coin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCouponCoin(ctx context.Context, cli *ent.Client) error {
	if _, err := couponcoincrud.CreateSet(
		cli.CouponCoin.Create(),
		&couponcoincrud.Req{
			EntID:      h.EntID,
			AppID:      h.AppID,
			CoinTypeID: h.CoinTypeID,
			CouponID:   h.CouponID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCouponCoin(ctx context.Context) (*npool.CouponCoin, error) {
	handler := &createHandler{
		Handler: h,
	}

	h.Conds = &couponcoincrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CouponID:   &cruder.Cond{Op: cruder.EQ, Val: *h.CouponID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
	}
	exist, err := h.ExistCouponCoinConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coupon coin already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createCouponCoin(ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCouponCoin(ctx)
}
