package config

import (
	"context"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
)

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppCommissionConfig, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppCommissionConfig.UpdateOneID(*h.ID),
			&commissionconfigcrud.Req{
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				ThresholdAmount: h.ThresholdAmount,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCommissionConfig(ctx)
}
