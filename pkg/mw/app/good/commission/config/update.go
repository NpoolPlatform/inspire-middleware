package config

import (
	"context"
	"fmt"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
)

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	h.ID = &info.ID
	if h.ThresholdAmount != nil && h.Invites != nil {
		endAt := uint32(0)
		h.Conds = &commissionconfigcrud.Conds{
			EntID:           &cruder.Cond{Op: cruder.NEQ, Val: info.EntID},
			AppID:           &cruder.Cond{Op: cruder.EQ, Val: info.AppID},
			ThresholdAmount: &cruder.Cond{Op: cruder.EQ, Val: *h.ThresholdAmount},
			Invites:         &cruder.Cond{Op: cruder.EQ, Val: *h.Invites},
			EndAt:           &cruder.Cond{Op: cruder.EQ, Val: endAt},
		}

		exist, err := h.ExistCommissionConfigs(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("exist same config")
		}
	}

	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppGoodCommissionConfig.UpdateOneID(*h.ID),
			&commissionconfigcrud.Req{
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				ThresholdAmount: h.ThresholdAmount,
				Invites:         h.Invites,
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
