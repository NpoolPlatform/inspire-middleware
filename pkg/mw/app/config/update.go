package config

import (
	"context"
	"fmt"

	appconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
)

func (h *Handler) UpdateAppConfig(ctx context.Context) (*npool.AppConfig, error) {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid appconfig")
	}
	h.ID = &info.ID

	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := appconfigcrud.UpdateSet(
			tx.AppConfig.UpdateOneID(*h.ID),
			&appconfigcrud.Req{
				StartAt: h.StartAt,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAppConfig(ctx)
}