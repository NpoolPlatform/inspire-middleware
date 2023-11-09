package coupon

import (
	"context"
	"fmt"

	couponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

func (h *queryHandler) queryCoupons(ctx context.Context, cli *ent.Client) error {
	stm, err := couponcrud.SetQueryConds(cli.Coupon.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.stmSelect = stm.Select()
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])

		if info.UserID != nil && *info.UserID == uuid.Nil.String() {
			info.UserID = nil
		}
		amount, err := decimal.NewFromString(info.Denomination)
		if err != nil {
			info.Denomination = decimal.NewFromInt(0).String()
		} else {
			info.Denomination = amount.String()
		}
		amount, err = decimal.NewFromString(info.Circulation)
		if err != nil {
			info.Circulation = decimal.NewFromInt(0).String()
		} else {
			info.Circulation = amount.String()
		}
		amount, err = decimal.NewFromString(info.Allocated)
		if err != nil {
			info.Allocated = decimal.NewFromInt(0).String()
		} else {
			info.Allocated = amount.String()
		}
		switch info.CouponConstraint {
		case types.CouponConstraint_PaymentThreshold:
		case types.CouponConstraint_GoodThreshold:
		default:
			info.Threshold = nil
		}
		if info.Threshold != nil {
			amount, err = decimal.NewFromString(*info.Threshold)
			if err != nil {
				_amount := decimal.NewFromInt(0).String()
				info.Threshold = &_amount
			} else {
				_amount := amount.String()
				info.Threshold = &_amount
			}
		}
	}
}

func (h *Handler) GetCoupon(ctx context.Context) (*npool.Coupon, error) {
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

func (h *Handler) GetCoupons(ctx context.Context) ([]*npool.Coupon, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoupons(_ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
