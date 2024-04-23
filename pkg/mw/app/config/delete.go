package config

import (
	"context"
	"time"

	appconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) DeleteAppConfig(ctx context.Context) error {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := appconfigcrud.UpdateSet(
			tx.AppConfig.UpdateOneID(*h.ID),
			&appconfigcrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
