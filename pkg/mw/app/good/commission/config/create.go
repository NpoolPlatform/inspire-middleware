package config

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	commissionconfigcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/app/good/commission/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommissionconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodcommissionconfig"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

	"github.com/google/uuid"
)

func (h *Handler) CreateCommissionConfig(ctx context.Context) (*npool.AppGoodCommissionConfig, error) {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppGoodCommissionConfig, *h.AppID, *h.AppGoodID)
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
			AppGoodCommissionConfig.
			Update().
			Where(
				entcommissionconfig.AppID(*h.AppID),
				entcommissionconfig.AppGoodID(*h.AppGoodID),
				entcommissionconfig.ThresholdAmount(*h.ThresholdAmount),
				entcommissionconfig.Invites(*h.Invites),
				entcommissionconfig.SettleType(h.SettleType.String()),
				entcommissionconfig.EndAt(0),
				entcommissionconfig.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}

		if _, err := commissionconfigcrud.CreateSet(
			tx.AppGoodCommissionConfig.Create(),
			&commissionconfigcrud.Req{
				EntID:           h.EntID,
				AppID:           h.AppID,
				GoodID:          h.GoodID,
				AppGoodID:       h.AppGoodID,
				ThresholdAmount: h.ThresholdAmount,
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				Invites:         h.Invites,
				SettleType:      h.SettleType,
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
