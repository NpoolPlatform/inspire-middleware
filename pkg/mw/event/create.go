//nolint:dupl
package event

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	eventcoincrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event/coin"
	eventcouponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createEventCoupons(ctx context.Context, tx *ent.Tx) error {
	for _, val := range h.CouponIDs {
		id := uuid.New()
		couponID := val
		if _, err := eventcouponcrud.CreateSet(
			tx.EventCoupon.Create(),
			&eventcouponcrud.Req{
				EntID:    &id,
				AppID:    h.AppID,
				EventID:  h.EntID,
				CouponID: &couponID,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createEventCoins(ctx context.Context, tx *ent.Tx) error {
	for _, coin := range h.Coins {
		id := uuid.New()
		coinID := uuid.MustParse(*coin.CoinConfigID)
		coinValue, err := decimal.NewFromString(*coin.CoinValue)
		if err != nil {
			return wlog.WrapError(err)
		}
		coinPreUSD, err := decimal.NewFromString(*coin.CoinPreUSD)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _, err := eventcoincrud.CreateSet(
			tx.EventCoin.Create(),
			&eventcoincrud.Req{
				EntID:        &id,
				AppID:        h.AppID,
				EventID:      h.EntID,
				CoinConfigID: &coinID,
				CoinValue:    &coinValue,
				CoinPreUSD:   &coinPreUSD,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *Handler) CreateEvent(ctx context.Context) (*npool.Event, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := h.validateCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppEvent, *h.AppID, *h.EventType)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
	}
	exist, err := h.ExistEventConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createEventCoupons(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createEventCoins(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if _, err := eventcrud.CreateSet(
			tx.Event.Create(),
			&eventcrud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				EventType:      h.EventType,
				Credits:        h.Credits,
				CreditsPerUSD:  h.CreditsPerUSD,
				MaxConsecutive: h.MaxConsecutive,
				GoodID:         h.GoodID,
				AppGoodID:      h.AppGoodID,
				InviterLayers:  h.InviterLayers,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetEvent(ctx)
}
