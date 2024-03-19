package config

import (
	"context"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
)

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppGoodCommissionConfig.UpdateOneID(*h.ID),
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
