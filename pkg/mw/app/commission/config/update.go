package config

import (
	"context"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
	"github.com/shopspring/decimal"
)

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppCommissionConfig, error) {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	h.ID = &info.ID
	if h.ThresholdAmount == nil {
		_amount, err := decimal.NewFromString(info.ThresholdAmount)
		if err != nil {
			return nil, err
		}
		h.ThresholdAmount = &_amount
	}
	if h.Invites == nil {
		h.Invites = &info.Invites
	}

	endAt := uint32(0)
	h.Conds = &commissionconfigcrud.Conds{
		EntID:           &cruder.Cond{Op: cruder.NEQ, Val: info.EntID},
		AppID:           &cruder.Cond{Op: cruder.EQ, Val: info.AppID},
		ThresholdAmount: &cruder.Cond{Op: cruder.EQ, Val: *h.ThresholdAmount},
		Invites:         &cruder.Cond{Op: cruder.EQ, Val: *h.Invites},
		EndAt:           &cruder.Cond{Op: cruder.EQ, Val: endAt},
	}

	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppCommissionConfig.UpdateOneID(*h.ID),
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
