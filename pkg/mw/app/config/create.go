package config

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	appconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entappconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appconfig"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	"github.com/google/uuid"
)

func (h *Handler) CreateAppConfig(ctx context.Context) (*npool.AppConfig, error) {
	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateAppConfig, *h.AppID)
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

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := tx.
			AppConfig.
			Update().
			Where(
				entappconfig.AppID(*h.AppID),
				entappconfig.EndAt(0),
				entappconfig.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}

		if _, err := appconfigcrud.CreateSet(
			tx.AppConfig.Create(),
			&appconfigcrud.Req{
				EntID:            h.EntID,
				AppID:            h.AppID,
				SettleMode:       h.SettleMode,
				SettleAmountType: h.SettleAmountType,
				SettleInterval:   h.SettleInterval,
				CommissionType:   h.CommissionType,
				SettleBenefit:    h.SettleBenefit,
				StartAt:          h.StartAt,
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
