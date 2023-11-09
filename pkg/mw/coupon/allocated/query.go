package allocated

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/allocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	entcouponallocated "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponallocated"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CouponAllocatedSelect
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

func (h *queryHandler) queryCoupons(cli *ent.Client) (*ent.CouponAllocatedSelect, error) {
	stm, err := allocatedcrud.SetQueryConds(cli.CouponAllocated.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCoupon(stm), nil
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
			sql.As(t.C(entcouponallocated.FieldCouponScope), "coupon_scope"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) error {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponallocated.FieldCouponID),
			t.C(entcoupon.FieldID),
		)

	if h.Conds != nil && h.Conds.CouponType != nil {
		couponType, ok := h.Conds.CouponType.Val.(types.CouponType)
		if !ok {
			return fmt.Errorf("invalid coupontype")
		}
		s.Where(
			sql.EQ(t.C(entcoupon.FieldCouponType), couponType.String()),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entcoupon.FieldName), "coupon_name"),
		sql.As(t.C(entcoupon.FieldCirculation), "circulation"),
		sql.As(t.C(entcoupon.FieldDurationDays), "duration_days"),
		sql.As(t.C(entcoupon.FieldMessage), "coupon_message"),
		sql.As(t.C(entcoupon.FieldThreshold), "threshold"),
		sql.As(t.C(entcoupon.FieldAllocated), "allocated"),
		sql.As(t.C(entcoupon.FieldCouponConstraint), "coupon_constraint"),
		sql.As(t.C(entcoupon.FieldRandom), "random"),
		sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinCoupon(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinCoupon(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EndAt = info.StartAt + info.DurationDays*timedef.SecondsPerDay
		info.Expired = uint32(time.Now().Unix()) > info.EndAt
		info.Valid = uint32(time.Now().Unix()) >= info.StartAt && uint32(time.Now().Unix()) <= info.EndAt
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponConstraint = types.CouponConstraint(types.CouponConstraint_value[info.CouponConstraintStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])
		if info.CouponName == "" {
			continue
		}
		if info.GoodID != nil && *info.GoodID == uuid.Nil.String() {
			info.GoodID = nil
		}
		if info.UsedByOrderID != nil && *info.UsedByOrderID == uuid.Nil.String() {
			info.UsedByOrderID = nil
		}
		if info.AppGoodID != nil && *info.AppGoodID == uuid.Nil.String() {
			info.AppGoodID = nil
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
