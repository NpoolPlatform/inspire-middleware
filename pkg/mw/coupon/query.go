package coupon

import (
	"context"
	"fmt"

	// couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CouponSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) queryCoupon(cli *ent.Client) {
	h.stmSelect = cli.
		Coupon.
		Query().
		Where(
			entcoupon.ID(*h.ID),
			entcoupon.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		if *info.UserID == uuid.Nil.String() {
			info.UserID = nil
		}
		if *info.GoodID == uuid.Nil.String() {
			info.GoodID = nil
		}
		switch info.CouponConstraint {
		case types.CouponConstraint_PaymentThreshold:
		case types.CouponConstraint_GoodThreshold:
			info.Threshold = nil
		}
	}
}

func (h *Handler) GetCoupon(ctx context.Context) (*npool.Coupon, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCoupon(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}
