package coupon

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

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
	stmCount  *ent.CouponSelect
	stmSelect *ent.CouponSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) selectCoupon(stm *ent.CouponQuery) *ent.CouponSelect {
	return stm.Select(
		entcoupon.FieldID,
	)
}

func (h *queryHandler) queryCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Coupon.Query().Where(entcoupon.DeletedAt(0))
	if h.ID != nil {
		h.stmSelect.Where(entcoupon.ID(*h.ID))
	}
	if h.EntID != nil {
		h.stmSelect.Where(entcoupon.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCoupon(stm)
	return nil
}

func (h *queryHandler) queryCoupons(cli *ent.Client) (*ent.CouponSelect, error) {
	stm, err := couponcrud.SetQueryConds(cli.Coupon.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCoupon(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcoupon.FieldID),
			t.C(entcoupon.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldEntID), "ent_id"),
			sql.As(t.C(entcoupon.FieldAppID), "app_id"),
			sql.As(t.C(entcoupon.FieldName), "name"),
			sql.As(t.C(entcoupon.FieldMessage), "message"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
			sql.As(t.C(entcoupon.FieldDenomination), "denomination"),
			sql.As(t.C(entcoupon.FieldCirculation), "circulation"),
			sql.As(t.C(entcoupon.FieldDurationDays), "duration_days"),
			sql.As(t.C(entcoupon.FieldCouponScope), "coupon_scope"),
			sql.As(t.C(entcoupon.FieldStartAt), "start_at"),
			sql.As(t.C(entcoupon.FieldIssuedBy), "issued_by"),
			sql.As(t.C(entcoupon.FieldAllocated), "allocated"),
			sql.As(t.C(entcoupon.FieldCouponConstraint), "coupon_constraint"),
			sql.As(t.C(entcoupon.FieldRandom), "random"),
			sql.As(t.C(entcoupon.FieldUserID), "user_id"),
			sql.As(t.C(entcoupon.FieldThreshold), "threshold"),
			sql.As(t.C(entcoupon.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcoupon.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return err
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
		if err := handler.queryCoupon(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
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

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCoupons(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCoupons(cli)
		if err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
