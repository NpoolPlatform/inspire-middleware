package archivement

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entarchivementgeneral "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/archivementgeneral"

	archivementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/archivement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ArchivementGeneralSelect
	infos     []*npool.Archivement
	total     uint32
}

func (h *queryHandler) selectArchivement(stm *ent.ArchivementGeneralQuery) {
	h.stmSelect = stm.Select(
		entarchivementgeneral.FieldID,
		entarchivementgeneral.FieldAppID,
		entarchivementgeneral.FieldUserID,
		entarchivementgeneral.FieldGoodID,
		entarchivementgeneral.FieldCoinTypeID,
		entarchivementgeneral.FieldTotalAmount,
		entarchivementgeneral.FieldSelfAmount,
		entarchivementgeneral.FieldTotalUnitsV1,
		entarchivementgeneral.FieldSelfUnitsV1,
		entarchivementgeneral.FieldTotalCommission,
		entarchivementgeneral.FieldSelfCommission,
		entarchivementgeneral.FieldCreatedAt,
		entarchivementgeneral.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryArchivements(ctx context.Context, cli *ent.Client) error {
	stm, err := archivementcrud.SetQueryConds(cli.ArchivementGeneral.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectArchivement(stm)
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

func (h *Handler) GetArchivements(ctx context.Context) ([]*npool.Archivement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Archivement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryArchivements(_ctx, cli); err != nil {
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
