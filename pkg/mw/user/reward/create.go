package reward

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) createUserReward(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create userreward: %v", err)
	}
	return nil
}

func (h *Handler) CreateUserReward(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	sql := handler.constructCreateSQL()
	handler.sql = sql
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createUserReward(_ctx, tx)
	})
}
