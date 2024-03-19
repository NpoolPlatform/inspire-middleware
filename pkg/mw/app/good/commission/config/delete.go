package config

import (
	"context"
	"time"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
)

func (h *Handler) DeleteCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppGoodCommissionConfig.UpdateOneID(*h.ID),
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
