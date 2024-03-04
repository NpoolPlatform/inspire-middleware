package event

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
}

func (h *rewardHandler) condGood() error {
	switch *h.EventType {
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		if h.GoodID == nil {
			return fmt.Errorf("need goodid")
		}
		if h.AppGoodID == nil {
			return fmt.Errorf("need appgoodid")
		}
		h.Conds.GoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID}
		h.Conds.AppGoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID}
	}
	return nil
}

func (h *rewardHandler) getEvent(ctx context.Context) (*npool.Event, error) {
	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
	}
	if err := h.condGood(); err != nil {
		return nil, err
	}
	info, err := h.GetEventOnly(ctx)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *rewardHandler) calculateCredits(ev *npool.Event) (decimal.Decimal, error) {
	credits, err := decimal.NewFromString(ev.Credits)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	_credits, err := decimal.NewFromString(ev.CreditsPerUSD)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	credits = credits.Add(_credits.Mul(*h.Amount))
	return credits, nil
}

func (h *rewardHandler) allocateCoupons(ctx context.Context, ev *npool.Event) error {
	coups := []*couponmwpb.Coupon{}
	for _, id := range ev.CouponIDs {
		_id := id
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithEntID(&_id, true),
		)
		if err != nil {
			return err
		}
		_coupon, err := handler.GetCoupon(ctx)
		if err != nil {
			return err
		}
		if _coupon == nil {
			return fmt.Errorf("invalid coupon")
		}

		now := time.Now().Unix()
		if now < int64(_coupon.StartAt) || now > int64(_coupon.EndAt) {
			logger.Sugar().Errorw("coupon can not be issued in current time")
			continue
		}
		coups = append(coups, _coupon)
	}

	for _, coup := range coups {
		userID := h.UserID.String()

		handler, err := allocated1.NewHandler(
			ctx,
			allocated1.WithAppID(&coup.AppID, true),
			allocated1.WithUserID(&userID, true),
			allocated1.WithCouponID(&coup.EntID, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.CreateCoupon(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h *rewardHandler) rewardSelf(ctx context.Context) ([]*npool.Credit, error) {
	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, err
	}
	if ev == nil {
		return nil, nil
	}

	if *h.Consecutive > ev.MaxConsecutive {
		return nil, nil
	}

	credits, err := h.calculateCredits(ev)
	if err != nil {
		return nil, err
	}

	// We don't care about result of allocate coupon
	if err := h.allocateCoupons(ctx, ev); err != nil {
		logger.Sugar().Warnw(
			"rewardSelf",
			"Event", ev,
			"Error", err,
		)
	}

	_credits := []*npool.Credit{}
	if credits.Cmp(decimal.NewFromInt(0)) > 0 {
		_credits = append(_credits, &npool.Credit{
			AppID:   h.AppID.String(),
			UserID:  h.UserID.String(),
			Credits: credits.String(),
		})
	}

	return _credits, nil
}

func (h *rewardHandler) rewardAffiliate(ctx context.Context) ([]*npool.Credit, error) {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return nil, err
	}
	handler.AppID = h.AppID
	handler.InviteeID = h.UserID

	_, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, err
	}
	if len(inviterIDs) == 0 {
		return nil, nil
	}

	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, err
	}
	if ev == nil {
		return nil, nil
	}

	if ev.InviterLayers == 0 {
		return nil, nil
	}

	credits := []*npool.Credit{}
	i := uint32(0)
	const inviterIgnore = 2
	j := len(inviterIDs) - inviterIgnore

	appID := h.AppID.String()
	goodID := h.GoodID.String()
	appGoodID := h.AppGoodID.String()
	amount := h.Amount.String()

	for ; i < ev.InviterLayers && j >= 0; i++ {
		handler, err := NewHandler(
			ctx,
			WithAppID(&appID, true),
			WithUserID(&inviterIDs[j], true),
			WithEventType(h.EventType, true),
			WithGoodID(&goodID, true),
			WithAppGoodID(&appGoodID, true),
			WithConsecutive(h.Consecutive, true),
			WithAmount(&amount, true),
		)
		if err != nil {
			return nil, err
		}

		_handler := &rewardHandler{
			Handler: handler,
		}

		credit, err := _handler.rewardSelf(ctx)
		if err != nil {
			return nil, err
		}

		j--
		if len(credit) == 0 {
			continue
		}

		credits = append(credits, credit...)
	}

	return credits, nil
}

func (h *Handler) RewardEvent(ctx context.Context) ([]*npool.Credit, error) {
	handler := &rewardHandler{
		Handler: h,
	}

	switch *h.EventType {
	case basetypes.UsedFor_Signup:
		fallthrough //nolint
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_SimulateOrderProfit:
		return handler.rewardSelf(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.rewardAffiliate(ctx)
	default:
		return nil, fmt.Errorf("not implemented")
	}
}
