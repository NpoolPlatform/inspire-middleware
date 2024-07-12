//nolint:dupl
package event

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	eventcoincrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event/coin"
	eventcouponcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	enteventcoin "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/eventcoin"
	enteventcoupon "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/eventcoupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	eventcoinmw "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) removeEventCoupons(ctx context.Context, tx *ent.Tx) error {
	if h.RemoveCouponIDs == nil || !*h.RemoveCouponIDs {
		return nil
	}
	// get current event coupons
	infos, err := tx.
		EventCoupon.
		Query().
		Where(
			enteventcoupon.EventID(*h.EntID),
			enteventcoupon.AppID(*h.AppID),
			enteventcoupon.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	// delete event coupons
	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if _, err := eventcouponcrud.UpdateSet(
			info.Update(),
			&eventcouponcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) removeEventCoins(ctx context.Context, tx *ent.Tx) error {
	if h.RemoveCoins == nil || !*h.RemoveCoins {
		return nil
	}

	// get current event coins
	infos, err := tx.
		EventCoin.
		Query().
		Where(
			enteventcoin.EventID(*h.EntID),
			enteventcoin.AppID(*h.AppID),
			enteventcoin.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	// delete event coins
	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if _, err := eventcoincrud.UpdateSet(
			info.Update(),
			&eventcoincrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) updateEventCoupons(ctx context.Context, tx *ent.Tx) error {
	if h.RemoveCoins != nil && *h.RemoveCoins {
		return nil
	}
	if len(h.CouponIDs) == 0 {
		return nil
	}
	// get db data
	infos, err := tx.
		EventCoupon.
		Query().
		Where(
			enteventcoupon.EventID(*h.EntID),
			enteventcoupon.AppID(*h.AppID),
			enteventcoupon.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	existCouponIDsMap := map[string]string{}
	inputCouponIDsMap := map[string]string{}
	newCouponsIDs := []string{}
	removeEventCouponsIDs := []uint32{}
	// check input and db data
	for _, info := range infos {
		existCouponIDsMap[info.CouponID.String()] = info.CouponID.String()
	}
	for _, id := range h.CouponIDs {
		inputCouponIDsMap[id.String()] = id.String()
		_, ok := existCouponIDsMap[id.String()]
		if !ok {
			newCouponsIDs = append(newCouponsIDs, id.String())
		}
	}
	for _, info := range infos {
		_, ok := inputCouponIDsMap[info.CouponID.String()]
		if !ok {
			removeEventCouponsIDs = append(removeEventCouponsIDs, info.ID)
		}
	}

	// remove not in input but in db data
	now := uint32(time.Now().Unix())
	for _, id := range removeEventCouponsIDs {
		if _, err := eventcouponcrud.UpdateSet(
			tx.EventCoupon.UpdateOneID(id),
			&eventcouponcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}

	// create in input but not in db data
	for _, id := range newCouponsIDs {
		entID := uuid.New()
		couponID := uuid.MustParse(id)
		if _, err := eventcouponcrud.CreateSet(
			tx.EventCoupon.Create(),
			&eventcouponcrud.Req{
				EntID:    &entID,
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

//nolint:funlen
func (h *updateHandler) updateEventCoins(ctx context.Context, tx *ent.Tx) error {
	if h.RemoveCoins != nil && *h.RemoveCoins {
		return nil
	}
	if len(h.Coins) == 0 {
		return nil
	}
	// get db data
	infos, err := tx.
		EventCoin.
		Query().
		Where(
			enteventcoin.EventID(*h.EntID),
			enteventcoin.AppID(*h.AppID),
			enteventcoin.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	existCoinsMap := map[string]*ent.EventCoin{}
	inputCoinsMap := map[string]*eventcoinmw.EventCoinReq{}
	newCoins := []*eventcoinmw.EventCoinReq{}
	updateCoins := []*eventcoinmw.EventCoinReq{}
	removeEventCoinIDs := []uint32{}
	// check input and db data
	for _, info := range infos {
		existCoinsMap[info.CoinConfigID.String()] = info
	}
	for _, coin := range h.Coins {
		inputCoinsMap[*coin.CoinConfigID] = coin
		existCoin, ok := existCoinsMap[*coin.CoinConfigID]
		if ok {
			coin.ID = &existCoin.ID
			updateCoins = append(updateCoins, coin)
			continue
		}
		newCoins = append(newCoins, coin)
	}
	for _, info := range infos {
		_, ok := inputCoinsMap[info.CoinConfigID.String()]
		if !ok {
			removeEventCoinIDs = append(removeEventCoinIDs, info.ID)
		}
	}

	// remove not in input but in db data
	now := uint32(time.Now().Unix())
	for _, id := range removeEventCoinIDs {
		if _, err := eventcoincrud.UpdateSet(
			tx.EventCoin.UpdateOneID(id),
			&eventcoincrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}

	// create in input but not in db data
	for _, coin := range newCoins {
		entID := uuid.New()
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
				EntID:        &entID,
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

	// update in input and in db data
	for _, coin := range updateCoins {
		coinValue, err := decimal.NewFromString(*coin.CoinValue)
		if err != nil {
			return wlog.WrapError(err)
		}
		coinPreUSD, err := decimal.NewFromString(*coin.CoinPreUSD)
		if err != nil {
			return wlog.WrapError(err)
		}
		id := coin.ID
		if _, err := eventcoincrud.UpdateSet(
			tx.EventCoin.UpdateOneID(*id),
			&eventcoincrud.Req{
				CoinValue:  &coinValue,
				CoinPreUSD: &coinPreUSD,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *Handler) UpdateEvent(ctx context.Context) (*npool.Event, error) {
	info, err := h.GetEvent(ctx)
	if err != nil {
		return nil, nil
	}
	if info == nil {
		return nil, wlog.Errorf("invalid event")
	}
	id := uuid.MustParse(info.EntID)
	h.EntID = &id
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID
	if err := h.validateCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.removeEventCoins(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.removeEventCoupons(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateEventCoins(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateEventCoupons(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if _, err := eventcrud.UpdateSet(
			tx.Event.UpdateOneID(*h.ID),
			&eventcrud.Req{
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
