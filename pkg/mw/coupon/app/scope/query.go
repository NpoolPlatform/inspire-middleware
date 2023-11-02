package scope

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodscopecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/scope"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappgoodscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodscope"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	entcouponscope "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponscope"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.AppGoodScopeSelect
	stmSelect *ent.AppGoodScopeSelect
	infos     []*npool.Scope
	total     uint32
}

func (h *queryHandler) selectAppGoodScope(stm *ent.AppGoodScopeQuery) *ent.AppGoodScopeSelect {
	return stm.Select(
		entappgoodscope.FieldID,
	)
}

func (h *queryHandler) queryAppGoodScope(cli *ent.Client) {
	stm := cli.
		AppGoodScope.
		Query().
		Where(
			entappgoodscope.ID(*h.ID),
			entappgoodscope.DeletedAt(0),
		)
	h.stmSelect = h.selectAppGoodScope(stm)
}

func (h *queryHandler) queryAppGoodScopes(cli *ent.Client) (*ent.AppGoodScopeSelect, error) {
	stm, err := appgoodscopecrud.SetQueryConds(cli.AppGoodScope.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectAppGoodScope(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgoodscope.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodscope.FieldID),
			t.C(entappgoodscope.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappgoodscope.FieldAppID), "app_id"),
			sql.As(t.C(entappgoodscope.FieldScopeID), "scope_id"),
			sql.As(t.C(entappgoodscope.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entappgoodscope.FieldCouponID), "coupon_id"),
			sql.As(t.C(entappgoodscope.FieldCouponScope), "coupon_scope"),
			sql.As(t.C(entappgoodscope.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappgoodscope.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinCouponScope(s *sql.Selector) {
	t := sql.Table(entcouponscope.Table)
	s.LeftJoin(t).
		On(
			t.C(entcouponscope.FieldID),
			s.C(entappgoodscope.FieldScopeID),
		).
		AppendSelect(
			sql.As(t.C(entcouponscope.FieldGoodID), "good_id"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodscope.FieldCouponID),
			t.C(entcoupon.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldName), "coupon_name"),
			sql.As(t.C(entcoupon.FieldCouponType), "coupon_type"),
			sql.As(t.C(entcoupon.FieldDenomination), "coupon_denomination"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoupon(s)
		h.queryJoinCouponScope(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinCoupon(s)
		h.queryJoinCouponScope(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CouponType = types.CouponType(types.CouponType_value[info.CouponTypeStr])
		info.CouponScope = types.CouponScope(types.CouponScope_value[info.CouponScopeStr])
		denomination, err := decimal.NewFromString(info.CouponDenomination)
		if err != nil {
			info.CouponDenomination = decimal.NewFromInt(0).String()
		} else {
			info.CouponDenomination = denomination.String()
		}
	}
}

func (h *Handler) GetAppGoodScope(ctx context.Context) (*npool.Scope, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Scope{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppGoodScope(cli)
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

func (h *Handler) GetAppGoodScopes(ctx context.Context) ([]*npool.Scope, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Scope{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAppGoodScopes(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAppGoodScopes(cli)
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

func (h *Handler) GetAppGoodScopeOnly(ctx context.Context) (*npool.Scope, error) {
	h.Limit = 1
	infos, _, err := h.GetAppGoodScopes(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos[0], nil
}
