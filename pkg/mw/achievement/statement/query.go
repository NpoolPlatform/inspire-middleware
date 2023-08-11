package statement

import (
	"context"
	"fmt"

	statementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.StatementSelect
	infos     []*npool.Statement
	total     uint32
}

func (h *queryHandler) selectStatement(stm *ent.StatementQuery) {
	h.stmSelect = stm.Select(
		entstatement.FieldID,
		entstatement.FieldAppID,
		entstatement.FieldUserID,
		entstatement.FieldDirectContributorID,
		entstatement.FieldGoodID,
		entstatement.FieldOrderID,
		entstatement.FieldSelfOrder,
		entstatement.FieldPaymentID,
		entstatement.FieldCoinTypeID,
		entstatement.FieldPaymentCoinTypeID,
		entstatement.FieldPaymentCoinUsdCurrency,
		entstatement.FieldUnitsV1,
		entstatement.FieldAmount,
		entstatement.FieldUsdAmount,
		entstatement.FieldCommission,
		entstatement.FieldCreatedAt,
		entstatement.FieldUpdatedAt,
		entstatement.FieldDeletedAt,
	)
}

func (h *queryHandler) queryStatement(cli *ent.Client) {
	h.selectStatement(
		cli.Statement.
			Query().
			Where(
				entstatement.ID(*h.ID),
				entstatement.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := statementcrud.SetQueryConds(cli.Statement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectStatement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.PaymentCoinUSDCurrency)
		if err != nil {
			info.PaymentCoinUSDCurrency = decimal.NewFromInt(0).String()
		} else {
			info.PaymentCoinUSDCurrency = amount.String()
		}
		amount, err = decimal.NewFromString(info.Units)
		if err != nil {
			info.Units = decimal.NewFromInt(0).String()
		} else {
			info.Units = amount.String()
		}
		amount, err = decimal.NewFromString(info.Amount)
		if err != nil {
			info.Amount = decimal.NewFromInt(0).String()
		} else {
			info.Amount = amount.String()
		}
		amount, err = decimal.NewFromString(info.USDAmount)
		if err != nil {
			info.USDAmount = decimal.NewFromInt(0).String()
		} else {
			info.USDAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.Commission)
		if err != nil {
			info.Commission = decimal.NewFromInt(0).String()
		} else {
			info.Commission = amount.String()
		}
	}
}

func (h *Handler) GetStatement(ctx context.Context) (*npool.Statement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryStatement(cli)
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

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(ctx, cli); err != nil {
			return err
		}
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

func (h *Handler) GetStatementOnly(ctx context.Context) (*npool.Statement, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(ctx, cli); err != nil {
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
