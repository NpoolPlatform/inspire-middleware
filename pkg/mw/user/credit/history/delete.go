package history

import (
	"context"
	"time"

	historycrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/user/credit/history"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteUserCreditHistory(ctx context.Context, cli *ent.Client) error {
	if _, err := historycrud.UpdateSet(
		cli.UserCreditHistory.UpdateOneID(*h.ID),
		&historycrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteUserCreditHistory(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetUserCreditHistory(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteUserCreditHistory(_ctx, cli)
	})
}
