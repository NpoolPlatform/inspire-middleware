package commission

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.CommissionSelect
	stmSelect *ent.CommissionSelect
	infos     []*npool.Commission
	total     uint32
}

func (h *queryHandler) selectCommission(stm *ent.CommissionQuery) *ent.CommissionSelect {
	return stm.Select(
		entcommission.FieldID,
	)
}

func (h *queryHandler) queryCommission(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}

	stm := cli.Commission.Query().Where(entcommission.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcommission.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcommission.EntID(*h.EntID))
	}
	h.selectCommission(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcommission.Table)
	s.LeftJoin(t).
		On(
			s.C(entcommission.FieldID),
			t.C(entcommission.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcommission.FieldEntID), "ent_id"),
			sql.As(t.C(entcommission.FieldAppID), "app_id"),
			sql.As(t.C(entcommission.FieldUserID), "user_id"),
			sql.As(t.C(entcommission.FieldGoodID), "good_id"),
			sql.As(t.C(entcommission.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entcommission.FieldSettleType), "settle_type"),
			sql.As(t.C(entcommission.FieldSettleMode), "settle_mode"),
			sql.As(t.C(entcommission.FieldSettleInterval), "settle_interval"),
			sql.As(t.C(entcommission.FieldSettleAmountType), "settle_amount_type"),
			sql.As(t.C(entcommission.FieldAmountOrPercent), "amount_or_percent"),
			sql.As(t.C(entcommission.FieldThreshold), "threshold"),
			sql.As(t.C(entcommission.FieldStartAt), "start_at"),
			sql.As(t.C(entcommission.FieldEndAt), "end_at"),
			sql.As(t.C(entcommission.FieldOrderLimit), "order_limit"),
			sql.As(t.C(entcommission.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcommission.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryCommissions(cli *ent.Client) (*ent.CommissionSelect, error) {
	stm, err := commissioncrud.SetQueryConds(cli.Commission.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCommission(stm), nil
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
		info.SettleType = types.SettleType(types.SettleType_value[info.SettleTypeStr])
		info.SettleMode = types.SettleMode(types.SettleMode_value[info.SettleModeStr])
		info.SettleAmountType = types.SettleAmountType(types.SettleAmountType_value[info.SettleAmountTypeStr])
		info.SettleInterval = types.SettleInterval(types.SettleInterval_value[info.SettleIntervalStr])
		amount, err := decimal.NewFromString(info.AmountOrPercent)
		if err != nil {
			info.AmountOrPercent = decimal.NewFromInt(0).String()
		} else {
			info.AmountOrPercent = amount.String()
		}
		amount, err = decimal.NewFromString(info.Threshold)
		if err != nil {
			info.Threshold = decimal.NewFromInt(0).String()
		} else {
			info.Threshold = amount.String()
		}
	}
}

func (h *Handler) GetCommission(ctx context.Context) (*npool.Commission, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Commission{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCommission(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		return handler.scan(ctx)
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

func (h *Handler) GetCommissions(ctx context.Context) ([]*npool.Commission, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Commission{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCommissions(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCommissions(cli)
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
