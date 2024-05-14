package goodachievement

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodAchievement(ctx context.Context, cli *ent.Client) error {
	_, err := goodachievementcrud.UpdateSet(
		cli.GoodAchievement.UpdateOneID(*h.ID),
		&goodachievementcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteAchievement(ctx context.Context) error {
	info, err := h.GetAchievement(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteGoodAchievement(_ctx, cli)
	})
}
