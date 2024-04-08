package config

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommissionconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appcommissionconfig"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.AppCommissionConfigSelect
	stmSelect *ent.AppCommissionConfigSelect
	infos     []*npool.AppCommissionConfig
	total     uint32
}

func (h *queryHandler) selectCommissionConfig(stm *ent.AppCommissionConfigQuery) *ent.AppCommissionConfigSelect {
	return stm.Select(
		entcommissionconfig.FieldID,
	)
}

func (h *queryHandler) queryCommissionConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}

	stm := cli.AppCommissionConfig.Query().Where(entcommissionconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcommissionconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcommissionconfig.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCommissionConfig(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcommissionconfig.Table)
	s.LeftJoin(t).
		On(
			s.C(entcommissionconfig.FieldID),
			t.C(entcommissionconfig.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcommissionconfig.FieldEntID), "ent_id"),
			sql.As(t.C(entcommissionconfig.FieldAppID), "app_id"),
			sql.As(t.C(entcommissionconfig.FieldThresholdAmount), "threshold_amount"),
			sql.As(t.C(entcommissionconfig.FieldAmountOrPercent), "amount_or_percent"),
			sql.As(t.C(entcommissionconfig.FieldStartAt), "start_at"),
			sql.As(t.C(entcommissionconfig.FieldEndAt), "end_at"),
			sql.As(t.C(entcommissionconfig.FieldInvites), "invites"),
			sql.As(t.C(entcommissionconfig.FieldSettleType), "settle_type"),
			sql.As(t.C(entcommissionconfig.FieldDisabled), "disabled"),
			sql.As(t.C(entcommissionconfig.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcommissionconfig.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryCommissionConfigs(cli *ent.Client) (*ent.AppCommissionConfigSelect, error) {
	stm, err := commissionconfigcrud.SetQueryConds(cli.AppCommissionConfig.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCommissionConfig(stm), nil
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
		amount, err := decimal.NewFromString(info.AmountOrPercent)
		if err != nil {
			info.AmountOrPercent = decimal.NewFromInt(0).String()
		} else {
			info.AmountOrPercent = amount.String()
		}
		amount, err = decimal.NewFromString(info.ThresholdAmount)
		if err != nil {
			info.ThresholdAmount = decimal.NewFromInt(0).String()
		} else {
			info.ThresholdAmount = amount.String()
		}
	}
}

func (h *Handler) GetCommissionConfig(ctx context.Context) (*npool.AppCommissionConfig, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppCommissionConfig{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCommissionConfig(cli); err != nil {
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

func (h *Handler) GetCommissionConfigs(ctx context.Context) ([]*npool.AppCommissionConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.AppCommissionConfig{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCommissionConfigs(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCommissionConfigs(cli)
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
