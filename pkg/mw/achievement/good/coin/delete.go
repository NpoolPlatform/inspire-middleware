package goodcoinachievement

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodcoinachievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/good/coin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodCoinAchievement(ctx context.Context, cli *ent.Client) error {
	_, err := goodcoinachievementcrud.UpdateSet(
		cli.GoodCoinAchievement.UpdateOneID(*h.ID),
		&goodcoinachievementcrud.Req{
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
		return handler.deleteGoodCoinAchievement(_ctx, cli)
	})
}
