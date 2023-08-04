package commission

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

//nolint:funlen
func (h *Handler) CloneCommissions(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if h.FromGoodID == nil {
		return fmt.Errorf("invalid fromgoodid")
	}
	if h.ToGoodID == nil {
		return fmt.Errorf("invalid togoodid")
	}
	if h.ScalePercent != nil && h.ScalePercent.Cmp(decimal.NewFromInt(0)) <= 0 {
		return nil
	}
	if *h.FromGoodID == *h.ToGoodID {
		return nil
	}

	key1 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCloneCommission, *h.AppID, *h.FromGoodID)
	if err := redis2.TryLock(key1, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key1)
	}()
	key2 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCloneCommission, *h.AppID, *h.ToGoodID)
	if err := redis2.TryLock(key2, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key2)
	}()

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		infos, err := cli.
			Commission.
			Query().
			Where(
				entcommission.AppID(*h.AppID),
				entcommission.GoodID(*h.FromGoodID),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			).
			All(_ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Infos: %v, AppID: %v, GoodID: %v\n", infos, *h.AppID, *h.FromGoodID)

		percent := decimal.NewFromInt(1)
		if h.ScalePercent != nil {
			percent = h.ScalePercent.Div(decimal.NewFromInt(100)) //nolint
		}

		cs := []*ent.CommissionCreate{}
		for _, info := range infos {
			info1, err := cli.
				Commission.
				Query().
				Where(
					entcommission.AppID(*h.AppID),
					entcommission.UserID(info.UserID),
					entcommission.GoodID(*h.ToGoodID),
					entcommission.EndAt(0),
					entcommission.DeletedAt(0),
				).
				Only(_ctx)
			if err != nil {
				return err
			}
			if info1 != nil {
				if _, err := cli.
					Commission.
					UpdateOneID(info1.ID).
					SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
					SetSettleType(info.SettleType).
					SetSettleMode(info.SettleMode).
					SetSettleInterval(info.SettleInterval).
					SetThreshold(info.Threshold).
					Save(_ctx); err != nil {
					return err
				}
				continue
			}

			c := cli.
				Commission.
				Create().
				SetAppID(info.AppID).
				SetUserID(info.UserID).
				SetGoodID(*h.ToGoodID).
				SetSettleType(info.SettleType).
				SetSettleMode(info.SettleMode).
				SetSettleInterval(info.SettleInterval).
				SetAmountOrPercent(info.AmountOrPercent.Mul(percent)).
				SetStartAt(now).
				SetThreshold(info.Threshold)
			cs = append(cs, c)
		}
		if _, err := cli.
			Commission.
			CreateBulk(cs...).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
