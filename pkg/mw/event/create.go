package event

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	eventcoin1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coin"
	eventcoupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/event/coupon"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createEventCoupons(ctx context.Context) error {
	appID := h.AppID.String()
	eventID := h.EntID.String()
	for _, couponID := range h.CouponIDs {
		id := uuid.NewString()
		couponIDStr := couponID.String()
		handler, err := eventcoupon1.NewHandler(
			ctx,
			eventcoupon1.WithEntID(&id, true),
			eventcoupon1.WithAppID(&appID, true),
			eventcoupon1.WithEventID(&eventID, true),
			eventcoupon1.WithCouponID(&couponIDStr, true),
		)
		if err != nil {
			return err
		}
		if err := handler.CreateEventCoupon(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) createEventCoins(ctx context.Context) error {
	appID := h.AppID.String()
	eventID := h.EntID.String()
	for _, coin := range h.Coins {
		id := uuid.NewString()
		handler, err := eventcoin1.NewHandler(
			ctx,
			eventcoin1.WithEntID(&id, true),
			eventcoin1.WithAppID(&appID, true),
			eventcoin1.WithEventID(&eventID, true),
			eventcoin1.WithCoinConfigID(coin.CoinConfigID, true),
			eventcoin1.WithCoinValue(coin.CoinValue, true),
			eventcoin1.WithCoinPreUSD(coin.CoinPreUSD, true),
		)
		if err != nil {
			return err
		}
		if err := handler.CreateEventCoin(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) CreateEvent(ctx context.Context) (*npool.Event, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := h.validateCoupons(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppEvent, *h.AppID, *h.EventType)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
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
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createEventCoupons(ctx); err != nil {
			return err
		}
		if err := handler.createEventCoins(ctx); err != nil {
			return err
		}
		if _, err := eventcrud.CreateSet(
			cli.Event.Create(),
			&eventcrud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				EventType:      h.EventType,
				CouponIDs:      h.CouponIDs,
				Credits:        h.Credits,
				CreditsPerUSD:  h.CreditsPerUSD,
				MaxConsecutive: h.MaxConsecutive,
				GoodID:         h.GoodID,
				AppGoodID:      h.AppGoodID,
				InviterLayers:  h.InviterLayers,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetEvent(ctx)
}
