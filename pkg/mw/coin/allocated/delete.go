package allocated

import (
	"context"
	"time"

	allocatedcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coin/allocated"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCoinAllocated(ctx context.Context, cli *ent.Client) error {
	if _, err := allocatedcrud.UpdateSet(
		cli.CoinAllocated.UpdateOneID(*h.ID),
		&allocatedcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCoinAllocated(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCoinAllocated(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteCoinAllocated(_ctx, cli)
	})
}
