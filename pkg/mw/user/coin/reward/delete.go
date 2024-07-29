package reward

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	rewardcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/user/coin/reward"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteUserCoinReward(ctx context.Context, cli *ent.Client) error {
	if _, err := rewardcrud.UpdateSet(
		cli.UserCoinReward.UpdateOneID(*h.ID),
		&rewardcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteUserCoinReward(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetUserCoinReward(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteUserCoinReward(_ctx, cli)
	})
}
