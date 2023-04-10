package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"
	allocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/event"

	allocated "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated"
	coupon "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"
	registration "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
}

func (h *rewardHandler) condGood(conds *mgrpb.Conds) {
	switch h.EventType {
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
}

func (h *rewardHandler) getEvent(ctx context.Context) (*mgrpb.Event, error) {
	conds := &mgrpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.EventType)},
	}
	if h.GoodID != nil {
		h.condGood(conds)
	}
	info, err := mgrcli.GetEventOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *rewardHandler) calculateCredits(ev *mgrpb.Event) (decimal.Decimal, error) {
	credits, err := decimal.NewFromString(ev.Credits)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	_credits, err := decimal.NewFromString(ev.CreditsPerUSD)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	amount, err := decimal.NewFromString(h.Amount)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	credits = credits.Add(_credits.Mul(amount))
	return credits, nil
}

func (h *rewardHandler) allocateCoupons(ctx context.Context, ev *mgrpb.Event) error {
	coups := []*couponmwpb.Coupon{}
	for _, coup := range ev.Coupons {
		_coupon, err := coupon.GetCoupon(ctx, coup.ID, coup.CouponType)
		if err != nil {
			return err
		}
		if _coupon == nil {
			return fmt.Errorf("invalid coupon")
		}

		coups = append(coups, _coupon)
	}

	for _, coup := range coups {
		if _, err := allocated.CreateCoupon(
			ctx,
			&allocatedmwpb.CouponReq{
				AppID:      &h.AppID,
				UserID:     &h.UserID,
				CouponID:   &coup.ID,
				CouponType: &coup.CouponType,
			},
		); err != nil {
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

	if h.Consecutive > ev.MaxConsecutive {
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
			AppID:   h.AppID,
			UserID:  h.UserID,
			Credits: credits.String(),
		})
	}

	return _credits, nil
}

func (h *rewardHandler) rewardAffiliate(ctx context.Context) ([]*npool.Credit, error) {
	_, inviterIDs, err := registration.GetInviters(ctx, h.AppID, h.UserID)
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

	for ; i < ev.InviterLayers && j >= 0; i++ {
		handler, err := NewHandler(
			ctx,
			WithAppID(h.AppID),
			WithUserID(inviterIDs[j]),
			WithEventType(h.EventType),
			WithGoodID(h.GoodID),
			WithConsecutive(h.Consecutive),
			WithAmount(h.Amount),
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

	switch h.EventType {
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
