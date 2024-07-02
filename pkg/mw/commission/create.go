package commission

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/google/uuid"
)

func (h *Handler) CreateCommission(ctx context.Context) (*npool.Commission, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateCommission, *h.AppID, *h.UserID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, wlog.WrapError(err)
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
			Commission.
			Update().
			Where(
				entcommission.AppID(*h.AppID),
				entcommission.UserID(*h.UserID),
				entcommission.AppGoodID(*h.AppGoodID),
				entcommission.SettleType(h.SettleType.String()),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			).
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}

		if _, err := commissioncrud.CreateSet(
			tx.Commission.Create(),
			&commissioncrud.Req{
				EntID:            h.EntID,
				AppID:            h.AppID,
				UserID:           h.UserID,
				GoodID:           h.GoodID,
				AppGoodID:        h.AppGoodID,
				SettleType:       h.SettleType,
				SettleMode:       h.SettleMode,
				SettleAmountType: h.SettleAmountType,
				SettleInterval:   h.SettleInterval,
				AmountOrPercent:  h.AmountOrPercent,
				StartAt:          h.StartAt,
				Threshold:        h.Threshold,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetCommission(ctx)
}
