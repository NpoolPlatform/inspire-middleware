package config

import (
	"context"
	"time"

	configcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coin/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCoinConfig(ctx context.Context, cli *ent.Client) error {
	if _, err := configcrud.UpdateSet(
		cli.CoinConfig.UpdateOneID(*h.ID),
		&configcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCoinConfig(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCoinConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteCoinConfig(_ctx, cli)
	})
}
