package scope

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappgoodscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodscope"
	entcouponscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"
)

type verifyHandler struct {
	*Handler
}

func (h *verifyHandler) verifyWhitelist(ctx context.Context, tx *ent.Tx) (bool, error) {
	_, err := tx.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*h.GoodID),
			entcouponscope.CouponID(*h.CouponID),
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
			entappgoodscope.AppID(*h.AppID),
			entappgoodscope.AppGoodID(*h.AppGoodID),
			entappgoodscope.CouponID(*h.CouponID),
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

func (h *verifyHandler) verifyBlacklist(ctx context.Context, tx *ent.Tx) (bool, error) {
	info, err := tx.
		CouponScope.
		Query().
		Where(
			entcouponscope.GoodID(*h.GoodID),
			entcouponscope.CouponID(*h.CouponID),
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
			entappgoodscope.AppID(*h.AppID),
			entappgoodscope.AppGoodID(*h.AppGoodID),
			entappgoodscope.CouponID(*h.CouponID),
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

func (h *Handler) VerifyCouponScope(ctx context.Context) (bool, error) {
	handler := &verifyHandler{
		Handler: h,
	}
	valid := false
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if *h.CouponScope == types.CouponScope_Whitelist {
			valid1, err := handler.verifyWhitelist(ctx, tx)
			valid = valid1
			return err
		}
		if *h.CouponScope == types.CouponScope_Blacklist {
			valid2, err := handler.verifyBlacklist(ctx, tx)
			valid = valid2
			return err
		}
		return nil
	})
	return valid, err
}
