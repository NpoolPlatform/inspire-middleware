package config

import (
	"context"
	"time"

	appconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
)

func (h *Handler) DeleteAppConfig(ctx context.Context) (*npool.AppConfig, error) {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
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
	if err != nil {
		return nil, err
	}

	return info, nil
}
