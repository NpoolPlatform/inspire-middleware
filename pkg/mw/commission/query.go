package commission

import (
	"context"
	"fmt"

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
	stmSelect *ent.CommissionSelect
	infos     []*npool.Commission
	total     uint32
}

func (h *queryHandler) queryCommission(cli *ent.Client) {
	h.stmSelect = cli.
		Commission.
		Query().
		Where(
			entcommission.ID(*h.ID),
			entcommission.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) queryCommissions(ctx context.Context, cli *ent.Client) error {
	stm, err := commissioncrud.SetQueryConds(cli.Commission.Query(), h.Conds)
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
		info.SettleType = types.SettleType(types.SettleType_value[info.SettleTypeStr])
		info.SettleMode = types.SettleMode(types.SettleMode_value[info.SettleModeStr])
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
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Commission{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCommission(cli)
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

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCommissions(ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
