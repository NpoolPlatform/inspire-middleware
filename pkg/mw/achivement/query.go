package achivement

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachivement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achivement"

	achivementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achivement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AchivementSelect
	infos     []*npool.Achivement
	total     uint32
}

func (h *queryHandler) selectAchivement(stm *ent.AchivementQuery) {
	h.stmSelect = stm.Select(
		entachivement.FieldID,
		entachivement.FieldAppID,
		entachivement.FieldUserID,
		entachivement.FieldGoodID,
		entachivement.FieldCoinTypeID,
		entachivement.FieldTotalAmount,
		entachivement.FieldSelfAmount,
		entachivement.FieldTotalUnitsV1,
		entachivement.FieldSelfUnitsV1,
		entachivement.FieldTotalCommission,
		entachivement.FieldSelfCommission,
		entachivement.FieldCreatedAt,
		entachivement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAchivements(ctx context.Context, cli *ent.Client) error {
	stm, err := achivementcrud.SetQueryConds(cli.Achivement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAchivement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.TotalAmount)
		if err != nil {
			info.TotalAmount = decimal.NewFromInt(0).String()
		} else {
			info.TotalAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfAmount)
		if err != nil {
			info.SelfAmount = decimal.NewFromInt(0).String()
		} else {
			info.SelfAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.TotalUnits)
		if err != nil {
			info.TotalUnits = decimal.NewFromInt(0).String()
		} else {
			info.TotalUnits = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfUnits)
		if err != nil {
			info.SelfUnits = decimal.NewFromInt(0).String()
		} else {
			info.SelfUnits = amount.String()
		}
		amount, err = decimal.NewFromString(info.TotalCommission)
		if err != nil {
			info.TotalCommission = decimal.NewFromInt(0).String()
		} else {
			info.TotalCommission = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfCommission)
		if err != nil {
			info.SelfCommission = decimal.NewFromInt(0).String()
		} else {
			info.SelfCommission = amount.String()
		}
	}
}

func (h *Handler) GetAchivements(ctx context.Context) ([]*npool.Achivement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Achivement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAchivements(_ctx, cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
