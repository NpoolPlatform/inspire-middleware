package allocated

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	devicecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coin/allocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcoinallocated "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coinallocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CoinAllocatedSelect
	stmCount  *ent.CoinAllocatedSelect
	infos     []*npool.CoinAllocated
	total     uint32
}

func (h *queryHandler) selectCoinAllocated(stm *ent.CoinAllocatedQuery) {
	h.stmSelect = stm.Select(entcoinallocated.FieldID)
}

func (h *queryHandler) queryCoinAllocated(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CoinAllocated.Query().Where(entcoinallocated.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcoinallocated.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcoinallocated.EntID(*h.EntID))
	}
	h.selectCoinAllocated(stm)
	return nil
}

func (h *queryHandler) queryCoinAllocateds(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.CoinAllocated.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectCoinAllocated(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entcoinallocated.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcoinallocated.FieldEntID),
			t1.C(entcoinallocated.FieldEntID),
		).
		AppendSelect(
			t1.C(entcoinallocated.FieldEntID),
			t1.C(entcoinallocated.FieldAppID),
			t1.C(entcoinallocated.FieldUserID),
			t1.C(entcoinallocated.FieldCoinConfigID),
			t1.C(entcoinallocated.FieldCoinTypeID),
			t1.C(entcoinallocated.FieldValue),
			t1.C(entcoinallocated.FieldExtra),
			t1.C(entcoinallocated.FieldCreatedAt),
			t1.C(entcoinallocated.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Value)
		if err != nil {
			info.Value = decimal.NewFromInt(0).String()
		} else {
			info.Value = amount.String()
		}
	}
}

func (h *Handler) GetCoinAllocated(ctx context.Context) (*npool.CoinAllocated, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinAllocated(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
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

func (h *Handler) GetCoinAllocateds(ctx context.Context) ([]*npool.CoinAllocated, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinAllocateds(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
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
