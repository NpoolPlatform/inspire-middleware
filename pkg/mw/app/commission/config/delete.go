package config

import (
	"context"
	"time"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
)

func (h *Handler) DeleteCommissionConfig(ctx context.Context) (*npool.AppCommissionConfig, error) {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	h.ID = &info.ID

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
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
	if err != nil {
		return nil, err
	}

	return info, nil
}
