package config

import (
	"context"
	"fmt"

	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) UpdateCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid appgoodcommissionconfig")
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
	if h.Level == nil {
		h.Level = &info.Level
	}

	endAt := uint32(0)
	entID := uuid.MustParse(info.EntID)
	appID := uuid.MustParse(info.AppID)
	goodID := uuid.MustParse(info.GoodID)
	appgoodID := uuid.MustParse(info.AppGoodID)
	h.Conds = &commissionconfigcrud.Conds{
		EntID:           &cruder.Cond{Op: cruder.NEQ, Val: entID},
		AppID:           &cruder.Cond{Op: cruder.EQ, Val: appID},
		GoodID:          &cruder.Cond{Op: cruder.EQ, Val: goodID},
		AppGoodID:       &cruder.Cond{Op: cruder.EQ, Val: appgoodID},
		ThresholdAmount: &cruder.Cond{Op: cruder.EQ, Val: *h.ThresholdAmount},
		Invites:         &cruder.Cond{Op: cruder.EQ, Val: *h.Invites},
		Level:           &cruder.Cond{Op: cruder.EQ, Val: *h.Level},
		EndAt:           &cruder.Cond{Op: cruder.EQ, Val: endAt},
	}

	exist, err := h.ExistCommissionConfigs(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("exist same config")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppGoodCommissionConfig.UpdateOneID(*h.ID),
			&commissionconfigcrud.Req{
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				ThresholdAmount: h.ThresholdAmount,
				Invites:         h.Invites,
				Disabled:        h.Disabled,
				Level:           h.Level,
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
