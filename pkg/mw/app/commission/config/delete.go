package config

import (
	"context"
	"time"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) DeleteCommissionConfig(ctx context.Context) error {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppCommissionConfig.UpdateOneID(*h.ID),
			&commissionconfigcrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
