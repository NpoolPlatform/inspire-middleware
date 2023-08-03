package commission

import (
	"context"
	"fmt"
	"time"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	commissioncrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/commission"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/google/uuid"
)

func (h *Handler) CreateCommission(ctx context.Context) (*npool.Commission, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateCommission, *h.AppID, *h.UserID)
	if h.GoodID != nil {
		key = fmt.Sprintf("%v:%v", key, *h.GoodID)
	}
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stm := tx.
			Commission.
			Update().
			Where(
				entcommission.AppID(*h.AppID),
				entcommission.UserID(*h.UserID),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			)
		if h.GoodID != nil {
			stm.Where(
				entcommission.GoodID(*h.GoodID),
			)
		}

		if _, err := stm.
			SetEndAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}

		if _, err := commissioncrud.CreateSet(
			tx.Commission.Create(),
			&commissioncrud.Req{
				ID:              h.ID,
				AppID:           h.AppID,
				UserID:          h.UserID,
				GoodID:          h.GoodID,
				SettleType:      h.SettleType,
				SettleMode:      h.SettleMode,
				SettleInterval:  h.SettleInterval,
				AmountOrPercent: h.AmountOrPercent,
				StartAt:         h.StartAt,
				Threshold:       h.Threshold,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCommission(ctx)
}
