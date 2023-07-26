package allocated

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	entcouponallocated "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponallocated"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	// cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CouponAllocatedSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) selectCoupon(stm *ent.CouponAllocatedQuery) *ent.CouponAllocatedSelect {
	return stm.Select(
		entcouponallocated.FieldID,
	)
}

func (h *queryHandler) queryCoupon(cli *ent.Client) {
	stm := cli.
		CouponAllocated.
		Query().
		Where(
			entcouponallocated.ID(*h.ID),
			entcouponallocated.DeletedAt(0),
		)
	h.stmSelect = h.selectCoupon(stm)
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcouponallocated.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponallocated.FieldID),
			t.C(entcouponallocated.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcouponallocated.FieldAppID), "app_id"),
			sql.As(t.C(entcouponallocated.FieldUserID), "user_id"),
			sql.As(t.C(entcouponallocated.FieldCouponID), "coupon_id"),
			sql.As(t.C(entcouponallocated.FieldStartAt), "start_at"),
			sql.As(t.C(entcouponallocated.FieldUsed), "used"),
			sql.As(t.C(entcouponallocated.FieldUsedAt), "used_at"),
			sql.As(t.C(entcouponallocated.FieldUsedByOrderID), "used_by_order_id"),
			sql.As(t.C(entcouponallocated.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcouponallocated.FieldUpdatedAt), "updated_at"),
			sql.As(t.C(entcouponallocated.FieldDenomination), "denomination"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponallocated.FieldCouponID),
			t.C(entcoupon.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldName), "coupon_name"),
			sql.As(t.C(entcoupon.FieldCirculation), "circulation"),
			sql.As(t.C(entcoupon.FieldDurationDays), "duration_days"),
			sql.As(t.C(entcoupon.FieldMessage), "coupon_message"),
			sql.As(t.C(entcoupon.FieldGoodID), "good_id"),
			sql.As(t.C(entcoupon.FieldThreshold), "threshold"),
			sql.As(t.C(entcoupon.FieldAllocated), "allocated"),
			sql.As(t.C(entcoupon.FieldCouponConstraint), "coupon_constraint"),
			sql.As(t.C(entcoupon.FieldRandom), "random"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoupon(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		if *info.GoodID == uuid.Nil.String() {
			info.GoodID = nil
		}
		if *info.UsedByOrderID == uuid.Nil.String() {
			info.UsedByOrderID = nil
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
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCoupon(cli)
		handler.queryJoin()
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
