package statement

import (
	"context"
	"fmt"
	"time"

	statementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
)

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			Statement.
			Update().
			Where(
				entstatement.ID(*h.ID),
				entstatement.DeletedAt(0),
			).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *Handler) DeleteStatements(ctx context.Context) ([]*npool.Statement, error) {
	ids := []uint32{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return nil, fmt.Errorf("invalid statement id")
		}
		ids = append(ids, *req.ID)
	}

	h.Conds = &statementcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Limit = int32(len(ids))
	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) != len(h.Reqs) {
		return nil, fmt.Errorf("statement not found")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			Statement.
			Update().
			Where(
				entstatement.IDIn(ids...),
				entstatement.DeletedAt(0),
			).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return infos, nil
}
