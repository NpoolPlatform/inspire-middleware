package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
}

func (h *rewardHandler) condGood() {
	switch *h.EventType {
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		h.Conds.GoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID}
	}
}

func (h *rewardHandler) getEvent(ctx context.Context) (*npool.Event, error) {
	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: uint32(*h.EventType)},
	}
	if h.GoodID != nil {
		h.condGood()
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
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithID(&id),
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

		coups = append(coups, _coupon)
	}

	for _, coup := range coups {
		userID := h.UserID.String()

		handler, err := allocated1.NewHandler(
			ctx,
			allocated1.WithAppID(&coup.AppID),
			allocated1.WithUserID(&userID),
			allocated1.WithCouponID(&coup.ID),
			allocated1.WithCouponType(&coup.CouponType),
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

	if err := h.allocateCoupons(ctx, ev); err != nil {
		return nil, err
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
	amount := h.Amount.String()

	for ; i < ev.InviterLayers && j >= 0; i++ {
		handler, err := NewHandler(
			ctx,
			WithAppID(&appID),
			WithUserID(&inviterIDs[j]),
			WithEventType(h.EventType),
			WithGoodID(&goodID),
			WithConsecutive(h.Consecutive),
			WithAmount(&amount),
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
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}
	if h.EventType == nil {
		return nil, fmt.Errorf("invalid eventtype")
	}
	if h.Consecutive == nil {
		return nil, fmt.Errorf("invalid consecutive")
	}
	if h.Amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	handler := &rewardHandler{
		Handler: h,
	}

	switch *h.EventType {
	case basetypes.UsedFor_Signup:
		fallthrough //nolint
	case basetypes.UsedFor_Purchase:
		return handler.rewardSelf(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.rewardAffiliate(ctx)
	default:
		return nil, fmt.Errorf("not implemented")
	}
}
