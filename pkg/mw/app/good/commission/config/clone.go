package config

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommissionconfig "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/appgoodcommissionconfig"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

func (h *Handler) CloneCommissionConfigs(ctx context.Context) error {
	if *h.FromAppGoodID == *h.ToAppGoodID {
		return nil
	}

	key1 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCloneCommission, *h.AppID, *h.FromAppGoodID)
	if err := redis2.TryLock(key1, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key1)
	}()
	key2 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCloneCommission, *h.AppID, *h.ToAppGoodID)
	if err := redis2.TryLock(key2, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key2)
	}()

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			AppGoodCommissionConfig.
			Query().
			Where(
				entcommissionconfig.AppID(*h.AppID),
				entcommissionconfig.GoodID(*h.FromGoodID),
				entcommissionconfig.AppGoodID(*h.FromAppGoodID),
				entcommissionconfig.EndAt(0),
				entcommissionconfig.DeletedAt(0),
			).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(infos) == 0 {
			return fmt.Errorf("commission not found")
		}

		percent := decimal.NewFromInt(1)
		if h.ScalePercent != nil {
			percent = h.ScalePercent.Div(decimal.NewFromInt(100)) //nolint
		}

		cs := []*ent.AppGoodCommissionConfigCreate{}
		for _, info := range infos {
			info1, err := cli.
				AppGoodCommissionConfig.
				Query().
				Where(
					entcommissionconfig.AppID(*h.AppID),
					entcommissionconfig.GoodID(*h.ToGoodID),
					entcommissionconfig.AppGoodID(*h.ToAppGoodID),
					entcommissionconfig.SettleType(info.SettleType),
					entcommissionconfig.EndAt(0),
					entcommissionconfig.DeletedAt(0),
				).
				Only(_ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return err
				}
			}
			if info1 != nil {
				if _, err := cli.
					AppGoodCommissionConfig.
					UpdateOneID(info1.ID).
					SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
					SetSettleType(info.SettleType).
					SetThresholdAmount(info.ThresholdAmount).
					Save(_ctx); err != nil {
					return err
				}
				continue
			}

			c := cli.
				AppGoodCommissionConfig.
				Create().
				SetAppID(info.AppID).
				SetGoodID(*h.ToGoodID).
				SetAppGoodID(*h.ToAppGoodID).
				SetSettleType(info.SettleType).
				SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
				SetStartAt(now).
				SetThresholdAmount(info.ThresholdAmount)
			cs = append(cs, c)
		}
		if _, err := cli.
			AppGoodCommissionConfig.
			CreateBulk(cs...).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
