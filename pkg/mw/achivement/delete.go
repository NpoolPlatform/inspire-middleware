package achivement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachivement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achivement"
	entstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/statement"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement/statement"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	info         *npool.Achivement
	statementIDs []uuid.UUID
}

func (h *deleteHandler) getStatementIDs(ctx context.Context) error {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithConds(&statementmwpb.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.info.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.info.UserID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.info.GoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.info.CoinTypeID},
		}),
		statement1.WithLimit(constant.DefaultRowLimit),
	)
	if err != nil {
		return err
	}

	for {
		statements, _, err := handler.GetStatements(ctx)
		if err != nil {
			return err
		}
		if len(statements) == 0 {
			break
		}
		for _, statement := range statements {
			h.statementIDs = append(h.statementIDs, uuid.MustParse(statement.ID))
		}
		handler.Offset += handler.Limit
	}

	return nil
}

func (h *Handler) DeleteAchivement(ctx context.Context) (*npool.Achivement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &deleteHandler{
		Handler:      h,
		statementIDs: []uuid.UUID{},
	}

	var err error
	handler.info, err = h.GetAchivement(ctx)
	if err != nil {
		return nil, err
	}
	if handler.info == nil {
		return nil, nil
	}
	if err := handler.getStatementIDs(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			Statement.
			Update().
			Where(
				entstatement.IDIn(handler.statementIDs...),
				entstatement.DeletedAt(0),
			).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}
		if _, err := tx.
			Achivement.
			Update().
			Where(
				entachivement.ID(uuid.MustParse(handler.info.ID)),
				entachivement.DeletedAt(0),
			).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return handler.info, nil
}
