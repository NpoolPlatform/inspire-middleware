package scope

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	"github.com/google/uuid"

	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	appgoodscopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappgoodscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodscope"
	entcouponscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
)

type verifyHandler struct {
	*Handler
}

func (h *verifyHandler) verifyWhitelist(ctx context.Context, tx *ent.Tx, req *appgoodscopecrud.Req) (bool, error) {
	_, err := tx.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*req.GoodID),
			entcouponscope.CouponID(*req.CouponID),
			entcouponscope.CouponScope(types.CouponScope_Whitelist.String()),
			entcouponscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	_, err = tx.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppID(*req.AppID),
			entappgoodscope.AppGoodID(*req.AppGoodID),
			entappgoodscope.CouponID(*req.CouponID),
			entappgoodscope.CouponScope(types.CouponScope_Whitelist.String()),
			entappgoodscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (h *verifyHandler) verifyBlacklist(ctx context.Context, tx *ent.Tx, req *appgoodscopecrud.Req) (bool, error) {
	info, err := tx.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*req.GoodID),
			entcouponscope.CouponID(*req.CouponID),
			entcouponscope.CouponScope(types.CouponScope_Blacklist.String()),
			entcouponscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return false, err
		}
	}
	if info != nil {
		return false, nil
	}

	info1, err := tx.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppID(*req.AppID),
			entappgoodscope.AppGoodID(*req.AppGoodID),
			entappgoodscope.CouponID(*req.CouponID),
			entappgoodscope.CouponScope(types.CouponScope_Blacklist.String()),
			entappgoodscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return false, err
		}
	}
	if info1 != nil {
		return false, nil
	}

	return true, nil
}

func (h *verifyHandler) getCoupons(ctx context.Context) (bool, error) {
	handler, err := coupon1.NewHandler(ctx)
	if err != nil {
		return false, err
	}
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		ids = append(ids, *req.CouponID)
	}

	handler.Conds = &couponcrud.Conds{
		AppID: &cruder.Cond{Op: cruder.EQ, Val: h.Reqs[0].AppID},
		IDs:   &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	coupons, _, err := handler.GetCoupons(ctx)
	if err != nil {
		return false, err
	}
	return len(coupons) == len(h.Reqs), nil
}

func (h *Handler) VerifyCouponScopes(ctx context.Context) (bool, error) {
	if len(h.Reqs) == 0 {
		return false, nil
	}
	handler := &verifyHandler{
		Handler: h,
	}
	if exist, err := handler.getCoupons(ctx); !exist || err != nil {
		return false, err
	}

	available := []bool{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			var err error
			var valid bool
			_fn := func() error {
				if *req.CouponScope == types.CouponScope_Whitelist {
					valid, err = handler.verifyWhitelist(ctx, tx, req)
					available = append(available, valid)
				}
				if *req.CouponScope == types.CouponScope_Blacklist {
					valid, err = handler.verifyBlacklist(ctx, tx, req)
					available = append(available, valid)
				}
				if *req.CouponScope == types.CouponScope_AllGood {
					available = append(available, true)
				}
				if err != nil {
					return err
				}
				return nil
			}
			if err := _fn(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	_available := true
	for _, val := range available {
		_available = _available && val
	}
	return _available, err
}
