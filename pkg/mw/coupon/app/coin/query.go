package coin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	couponcoincrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coupon/app/coin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	entcouponcoin "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponcoin"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CouponCoinSelect
	stmSelect *ent.CouponCoinSelect
	infos     []*npool.CouponCoin
	total     uint32
}

func (h *queryHandler) selectCouponCoin(stm *ent.CouponCoinQuery) *ent.CouponCoinSelect {
	return stm.Select(
		entcouponcoin.FieldID,
	)
}

func (h *queryHandler) queryCouponCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CouponCoin.Query().Where(entcouponcoin.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcouponcoin.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcouponcoin.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCouponCoin(stm)
	return nil
}

func (h *queryHandler) queryCouponCoins(cli *ent.Client) (*ent.CouponCoinSelect, error) {
	stm, err := couponcoincrud.SetQueryConds(cli.CouponCoin.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCouponCoin(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcouponcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponcoin.FieldID),
			t.C(entcouponcoin.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcouponcoin.FieldEntID), "ent_id"),
			sql.As(t.C(entcouponcoin.FieldAppID), "app_id"),
			sql.As(t.C(entcouponcoin.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(entcouponcoin.FieldCouponID), "coupon_id"),
			sql.As(t.C(entcouponcoin.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcouponcoin.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinCoupon(s *sql.Selector) {
	t := sql.Table(entcoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entcouponcoin.FieldCouponID),
			t.C(entcoupon.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoupon.FieldName), "coupon_name"),
			sql.As(t.C(entcoupon.FieldDenomination), "coupon_denomination"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoupon(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinCoupon(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		denomination, err := decimal.NewFromString(info.CouponDenomination)
		if err != nil {
			info.CouponDenomination = decimal.NewFromInt(0).String()
		} else {
			info.CouponDenomination = denomination.String()
		}
	}
}

func (h *Handler) GetCouponCoin(ctx context.Context) (*npool.CouponCoin, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CouponCoin{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCouponCoin(cli); err != nil {
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

func (h *Handler) GetCouponCoins(ctx context.Context) ([]*npool.CouponCoin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CouponCoin{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCouponCoins(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCouponCoins(cli)
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
