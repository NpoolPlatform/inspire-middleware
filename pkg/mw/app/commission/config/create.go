package config

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommissionconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appcommissionconfig"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"

	"github.com/google/uuid"
)

func (h *Handler) CreateCommissionConfig(ctx context.Context) (*npool.AppCommissionConfig, error) {
	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateAppCommissionConfig, *h.AppID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}
	if h.StartAt == nil {
		startAt := uint32(time.Now().Unix())
		h.StartAt = &startAt
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			AppCommissionConfig.
			Update().
			Where(
				entcommissionconfig.AppID(*h.AppID),
				entcommissionconfig.ThresholdAmount(*h.ThresholdAmount),
				entcommissionconfig.Invites(*h.Invites),
				entcommissionconfig.Level(*h.Level),
				entcommissionconfig.SettleType(h.SettleType.String()),
				entcommissionconfig.EndAt(0),
				entcommissionconfig.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}

		if _, err := commissionconfigcrud.CreateSet(
			tx.AppCommissionConfig.Create(),
			&commissionconfigcrud.Req{
				EntID:           h.EntID,
				AppID:           h.AppID,
				ThresholdAmount: h.ThresholdAmount,
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				Invites:         h.Invites,
				SettleType:      h.SettleType,
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

	return nil, nil
}
