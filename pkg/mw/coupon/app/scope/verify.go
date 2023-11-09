package scope

import (
	"context"
	"fmt"

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

func (h *verifyHandler) verifyWhitelist(ctx context.Context, cli *ent.Client, req *appgoodscopecrud.Req) error {
	if *req.CouponScope != types.CouponScope_Whitelist {
		return nil
	}
	_, err := cli.
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
		return err
	}

	_, err = cli.
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
		return err
	}
	return nil
}

func (h *verifyHandler) verifyBlacklist(ctx context.Context, cli *ent.Client, req *appgoodscopecrud.Req) error {
	if *req.CouponScope != types.CouponScope_Blacklist {
		return nil
	}
	info, err := cli.
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
			return err
		}
	}
	if info != nil {
		return fmt.Errorf("couponid in blacklist(good)")
	}

	info1, err := cli.
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
			return err
		}
	}
	if info1 != nil {
		return fmt.Errorf("couponid in blacklist(appgood)")
	}

	return nil
}

func (h *verifyHandler) checkCoupons(ctx context.Context) error {
	handler, err := coupon1.NewHandler(ctx)
	if err != nil {
		return err
	}
	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		ids = append(ids, *req.CouponID)
	}

	handler.Conds = &couponcrud.Conds{
		AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.Reqs[0].AppID},
		IDs:   &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	coupons, _, err := handler.GetCoupons(ctx)
	if err != nil {
		return err
	}

	if len(coupons) != len(h.Reqs) {
		return fmt.Errorf("invalid couponid")
	}
	return nil
}

func (h *Handler) VerifyCouponScopes(ctx context.Context) error {
	if len(h.Reqs) == 0 {
		return fmt.Errorf("invalid infos")
	}
	handler := &verifyHandler{
		Handler: h,
	}
	if err := handler.checkCoupons(ctx); err != nil {
		return err
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			if err := handler.verifyWhitelist(ctx, cli, req); err != nil {
				return err
			}
			if err := handler.verifyBlacklist(ctx, cli, req); err != nil {
				return err
			}
			if *req.CouponScope == types.CouponScope_AllGood {
				continue
			}
		}
		return nil
	})
	return err
}
