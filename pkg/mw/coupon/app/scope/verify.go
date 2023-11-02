package scope

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappgoodscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodscope"
	entcouponscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
)

type verifyHandler struct {
	*Handler
}

func (h *verifyHandler) verifyWhitelist(ctx context.Context) (bool, error) {
    err := db.WithClient(ctx, fn func(_ctx context.Context, cli *ent.Client) error {
	_, err := cli.
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
		return err
	}
    })

	_, err = tx.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppGoodID(*h.AppGoodID),
			entappgoodscope.CouponID(*h.CouponID),
			entappgoodscope.AppID(*h.AppID),
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

func (h *verifyHandler) verifyBlacklist(ctx context.Context, tx *ent.Tx) error {
	if *h.CouponScope != types.CouponScope_Blacklist {
		return nil
	}
	_, err := tx.
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
		return err
	}

	_, err = tx.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.AppGoodID(*h.AppGoodID),
			entappgoodscope.CouponID(*h.CouponID),
			entappgoodscope.AppID(*h.AppID),
			entappgoodscope.CouponScope(types.CouponScope_Blacklist.String()),
			entappgoodscope.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) VerifyCouponScope(ctx context.Context) (bool, error) {
	handler := &verifyHandler{
		Handler: h,
	}
	case types.CouponScope_Whitelist:
		return handler.verifyWhitelist(ctx)
	case types.CouponScope_Blacklist:
		return handler.verifyBlacklist(ctx)
	}
	return false, nil
}
