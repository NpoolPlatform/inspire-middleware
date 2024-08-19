package event

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	coinconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	taskconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"
	taskuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinallocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	couponallocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	taskconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
	taskusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
	"github.com/google/uuid"

	"github.com/shopspring/decimal"
)

type calculateHandler struct {
	*Handler
	taskConfig           *taskconfigmwpb.TaskConfig
	addCredits           decimal.Decimal
	coinPreUSDAmount     decimal.Decimal
	couponAmount         decimal.Decimal
	couponCashableAmount decimal.Decimal
}

//nolint:dupl
func (h *calculateHandler) condGood() error {
	switch *h.EventType {
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		if h.GoodID == nil {
			return wlog.Errorf("need goodid")
		}
		if h.AppGoodID == nil {
			return wlog.Errorf("need appgoodid")
		}
		h.Conds.GoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID}
		h.Conds.AppGoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID}
	}
	return nil
}

func (h *calculateHandler) getEvent(ctx context.Context) (*npool.Event, error) {
	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
	}
	if err := h.condGood(); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetEventOnly(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}

func (h *calculateHandler) calculateCredits(ev *npool.Event) (decimal.Decimal, error) {
	credits, err := decimal.NewFromString(ev.Credits)
	if err != nil {
		return decimal.NewFromInt(0), wlog.WrapError(err)
	}

	_credits, err := decimal.NewFromString(ev.CreditsPerUSD)
	if err != nil {
		return decimal.NewFromInt(0), wlog.WrapError(err)
	}

	credits = credits.Add(_credits.Mul(*h.Amount))

	return credits, nil
}

func (h *calculateHandler) calculateCoinRewards(ctx context.Context, ev *npool.Event) ([]*coinallocatedmwpb.CoinAllocated, error) {
	coinRewards := []*coinallocatedmwpb.CoinAllocated{}
	for _, eventCoin := range ev.Coins {
		_id := eventCoin.CoinConfigID
		handler, err := coinconfig1.NewHandler(
			ctx,
			coinconfig1.WithEntID(&_id, true),
		)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		_coinConfig, err := handler.GetCoinConfig(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if _coinConfig == nil {
			return nil, wlog.Errorf("invalid coinconfig")
		}
		if _coinConfig.MaxValue == _coinConfig.Allocated {
			continue
		}

		userID := h.UserID.String()

		coinValue, err := decimal.NewFromString(eventCoin.CoinValue)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		coinPreUSD, err := decimal.NewFromString(eventCoin.CoinPreUSD)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		amount := decimal.NewFromInt(0)
		if h.Amount != nil {
			amount = *h.Amount
		}
		coinPreUSDAmount := coinPreUSD.Mul(amount)
		h.coinPreUSDAmount = coinPreUSDAmount

		coins := coinValue.Add(coinPreUSDAmount)
		if coins.Cmp(decimal.NewFromInt(0)) == 0 {
			continue
		}
		allocated, err := decimal.NewFromString(_coinConfig.Allocated)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		maxValue, err := decimal.NewFromString(_coinConfig.MaxValue)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		if coins.Add(allocated).Cmp(maxValue) >= 0 {
			continue
		}
		coinsStr := coins.String()

		id := uuid.NewString()
		coinRewards = append(coinRewards, &coinallocatedmwpb.CoinAllocated{
			EntID:        id,
			AppID:        _coinConfig.AppID,
			UserID:       userID,
			CoinConfigID: _coinConfig.EntID,
			CoinTypeID:   _coinConfig.CoinTypeID,
			Value:        coinsStr,
		})
	}

	return coinRewards, nil
}

func (h *calculateHandler) calculateCouponRewards(ctx context.Context, ev *npool.Event) ([]*couponallocatedmwpb.Coupon, error) {
	couponRewards := []*couponallocatedmwpb.Coupon{}
	coups := []*couponmwpb.Coupon{}
	for _, id := range ev.CouponIDs {
		_id := id
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithEntID(&_id, true),
		)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		_coupon, err := handler.GetCoupon(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if _coupon == nil {
			return nil, wlog.Errorf("invalid coupon")
		}

		now := time.Now().Unix()
		if now < int64(_coupon.StartAt) || now > int64(_coupon.EndAt) {
			logger.Sugar().Errorw("coupon can not be issued in current time")
			continue
		}
		coups = append(coups, _coupon)
	}
	for _, coup := range coups {
		// calculate coupon
		userID := h.UserID.String()

		handler, err := allocated1.NewHandler(
			ctx,
			allocated1.WithAppID(&coup.AppID, true),
			allocated1.WithUserID(&userID, true),
			allocated1.WithCouponID(&coup.EntID, true),
		)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		info, err := handler.CalcluateAllocatedCoupon(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		couponAmount, err := decimal.NewFromString(info.Denomination)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		couponCashableAmount := decimal.NewFromInt(0)
		if info.Cashable {
			couponCashableAmount = couponAmount
		}
		h.couponAmount = couponAmount
		h.couponCashableAmount = couponCashableAmount

		if err != nil {
			return nil, wlog.WrapError(err)
		}

		couponRewards = append(couponRewards, info)
	}

	return couponRewards, nil
}

func (h *calculateHandler) calcluateAffiliate(ctx context.Context) ([]*npool.Reward, error) {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.AppID = h.AppID
	handler.InviteeID = h.UserID

	_, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	fmt.Println("inviterIDs: ", inviterIDs)
	if len(inviterIDs) == 0 {
		return nil, nil
	}

	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if ev == nil {
		return nil, nil
	}

	if ev.InviterLayers == 0 {
		return nil, nil
	}

	rewards := []*npool.Reward{}

	i := uint32(0)
	const inviterIgnore = 2
	j := len(inviterIDs) - inviterIgnore

	appID := h.AppID.String()
	goodID := uuid.Nil.String()
	if h.GoodID != nil {
		goodID = h.GoodID.String()
	}
	appGoodID := uuid.Nil.String()
	if h.AppGoodID != nil {
		appGoodID = h.AppGoodID.String()
	}
	amount := "0"
	if h.Amount != nil {
		amount = h.Amount.String()
	}

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
			return nil, wlog.WrapError(err)
		}

		_handler := &calculateHandler{
			Handler: handler,
		}

		reward, err := _handler.calcluateEventRewards(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		rewards = append(rewards, reward...)

		j--
	}

	if h.taskConfig == nil {
		return nil, nil
	}

	return rewards, nil
}

func (h *calculateHandler) validateTask(ctx context.Context, ev *npool.Event) error {
	handler, err := taskconfig1.NewHandler(
		ctx,
		taskconfig1.WithConds(&taskconfigmwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ev.AppID},
			EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ev.EntID},
		}),
		taskconfig1.WithOffset(0),
		taskconfig1.WithLimit(constant.DefaultRowLimit),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	configs, _, err := handler.GetTaskConfigs(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(configs) == 0 {
		return wlog.Errorf("invalid taskconfig")
	}
	h.taskConfig = configs[0]
	// check user has finished this task
	userID := h.UserID.String()
	handler2, err := taskuser1.NewHandler(
		ctx,
		taskuser1.WithConds(&taskusermwpb.Conds{
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ev.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: userID},
			TaskID: &basetypes.StringVal{Op: cruder.EQ, Value: configs[0].EntID},
		}),
		taskuser1.WithOffset(0),
		taskuser1.WithLimit(int32(configs[0].MaxRewardCount+1)),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	taskUsers, _, err := handler2.GetTaskUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(taskUsers) == 0 {
		return nil
	}
	// check user has over the max finish times
	if len(taskUsers) >= int(configs[0].MaxRewardCount) {
		return wlog.Errorf("invalid maxrewardcount")
	}

	// check daily task
	nowTime := time.Now()
	midnightTime := uint32(time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location()).Unix())
	if h.taskConfig.TaskType == types.TaskType_DailyTask && taskUsers[len(taskUsers)-1].CreatedAt > midnightTime {
		return wlog.Errorf("dailytask have been completed")
	}

	// check user next task startat
	now := uint32(time.Now().Unix())
	if taskUsers[len(taskUsers)-1].UpdatedAt+configs[0].CooldownSecord > now {
		return wlog.Errorf("not the right time")
	}
	// check last task exist and finish status
	if configs[0].LastTaskID != uuid.Nil.String() {
		done := types.TaskState_Done
		handler3, err := taskuser1.NewHandler(
			ctx,
			taskuser1.WithConds(&taskusermwpb.Conds{
				AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ev.AppID},
				UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: userID},
				TaskID:    &basetypes.StringVal{Op: cruder.EQ, Value: configs[0].LastTaskID},
				TaskState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(done)},
			}),
			taskuser1.WithOffset(0),
			taskuser1.WithLimit(constant.DefaultRowLimit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		lastTaskUsers, _, err := handler3.GetTaskUsers(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(lastTaskUsers) == 0 {
			return wlog.Errorf("invalid last task")
		}
	}

	return nil
}

func (h *calculateHandler) calcluateEventRewards(ctx context.Context) ([]*npool.Reward, error) {
	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if ev == nil {
		return nil, nil
	}

	if *h.Consecutive > ev.MaxConsecutive {
		return nil, nil
	}

	if err := h.validateTask(ctx, ev); err != nil {
		return nil, wlog.WrapError(err)
	}

	h.addCredits = decimal.NewFromInt(0)
	h.coinPreUSDAmount = decimal.NewFromInt(0)
	h.couponAmount = decimal.NewFromInt(0)
	h.couponCashableAmount = decimal.NewFromInt(0)
	userID := h.UserID.String()
	credits, err := h.calculateCredits(ev)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	h.addCredits = credits

	allocateCoinRewards, err := h.calculateCoinRewards(ctx, ev)
	if err != nil {
		logger.Sugar().Warnw(
			"rewardTask calculateCoinRewards",
			"Event", ev,
			"Error", err,
		)
	}
	coinRewards := []*npool.CoinReward{}
	for _, coin := range allocateCoinRewards {
		coinReward := npool.CoinReward{
			AppID:        coin.AppID,
			UserID:       coin.UserID,
			CoinTypeID:   coin.CoinTypeID,
			CoinConfigID: coin.CoinConfigID,
			CoinRewards:  coin.Value,
		}
		coinRewards = append(coinRewards, &coinReward)
	}
	allocateCouponRewards, err := h.calculateCouponRewards(ctx, ev)
	if err != nil {
		logger.Sugar().Warnw(
			"rewardTask allocateCouponRewards",
			"Event", ev,
			"Error", err,
		)
	}
	couponRewards := []*npool.CouponReward{}
	for _, coupon := range allocateCouponRewards {
		couponReward := npool.CouponReward{
			AppID:        coupon.AppID,
			UserID:       coupon.UserID,
			CouponID:     coupon.CouponID,
			Cashable:     coupon.Cashable,
			Denomination: coupon.Denomination,
		}
		couponRewards = append(couponRewards, &couponReward)
	}

	_rewards := &npool.Reward{
		TaskID:        h.taskConfig.EntID,
		UserID:        userID,
		Credits:       h.addCredits.String(),
		CoinRewards:   coinRewards,
		CouponRewards: couponRewards,
	}

	return []*npool.Reward{_rewards}, nil
}

//nolint:funlen
func (h *Handler) CalcluateEventRewards(ctx context.Context) ([]*npool.Reward, error) {
	handler := &calculateHandler{
		Handler: h,
	}

	switch *h.EventType {
	case basetypes.UsedFor_Signup:
		fallthrough //nolint
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_SimulateOrderProfit:
		fallthrough //nolint
	case basetypes.UsedFor_SetWithdrawAddress:
		fallthrough //nolint
	case basetypes.UsedFor_ConsecutiveLogin:
		fallthrough //nolint
	case basetypes.UsedFor_GoodSocialSharing:
		fallthrough //nolint
	case basetypes.UsedFor_FirstOrderCompleted:
		fallthrough //nolint
	case basetypes.UsedFor_SetAddress:
		fallthrough //nolint
	case basetypes.UsedFor_Set2FA:
		fallthrough //nolint
	case basetypes.UsedFor_FirstBenefit:
		fallthrough //nolint
	case basetypes.UsedFor_WriteComment:
		fallthrough //nolint
	case basetypes.UsedFor_WriteRecommend:
		fallthrough //nolint
	case basetypes.UsedFor_GoodScoring:
		fallthrough //nolint
	case basetypes.UsedFor_SubmitTicket:
		fallthrough //nolint
	case basetypes.UsedFor_IntallApp:
		fallthrough //nolint
	case basetypes.UsedFor_SetNFTAvatar:
		fallthrough //nolint
	case basetypes.UsedFor_SetPersonalImage:
		fallthrough //nolint
	case basetypes.UsedFor_Signin:
		fallthrough //nolint
	case basetypes.UsedFor_KYCApproved:
		fallthrough //nolint
	case basetypes.UsedFor_OrderCompleted:
		fallthrough //nolint
	case basetypes.UsedFor_WithdrawalCompleted:
		fallthrough //nolint
	case basetypes.UsedFor_DepositReceived:
		fallthrough //nolint
	case basetypes.UsedFor_UpdatePassword:
		fallthrough //nolint
	case basetypes.UsedFor_ResetPassword:
		fallthrough //nolint
	case basetypes.UsedFor_InternalTransfer:
		return handler.calcluateEventRewards(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.calcluateAffiliate(ctx)
	default:
		return nil, wlog.Errorf("not implemented")
	}
}