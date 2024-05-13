package event

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	coinallocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"
	coinconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	allocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon/allocated"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	taskconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"
	taskuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/user"
	usercredithistory1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/credit/history"
	userreward1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/user/reward"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinallocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	taskconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
	taskusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
	userrewardmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
	"github.com/google/uuid"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
	taskConfig           *taskconfigmwpb.TaskConfig
	addCredits           decimal.Decimal
	coinPreUSDAmount     decimal.Decimal
	couponAmount         decimal.Decimal
	couponCashableAmount decimal.Decimal
}

func (h *rewardHandler) condGood() error {
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

func (h *rewardHandler) allocatedCredits(ctx context.Context, ev *npool.Event) (decimal.Decimal, error) {
	credits, err := decimal.NewFromString(ev.Credits)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	_credits, err := decimal.NewFromString(ev.CreditsPerUSD)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	credits = credits.Add(_credits.Mul(*h.Amount))
	creditsStr := credits.String()

	userID := h.UserID.String()
	handler, err := usercredithistory1.NewHandler(
		ctx,
		usercredithistory1.WithAppID(&ev.AppID, true),
		usercredithistory1.WithUserID(&userID, true),
		usercredithistory1.WithTaskID(&h.taskConfig.EntID, true),
		usercredithistory1.WithEventID(&ev.EntID, true),
		usercredithistory1.WithCredits(&creditsStr, true),
	)
	if err != nil {
		return decimal.NewFromInt(0), err
	}
	if err := handler.CreateUserCreditHistory(ctx); err != nil {
		return decimal.NewFromInt(0), err
	}
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
			return wlog.Errorf("invalid coupon")
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

		info, err := handler.CreateCoupon(ctx)
		if err != nil {
			return err
		}
		couponAmount, err := decimal.NewFromString(info.Allocated)
		if err != nil {
			return err
		}
		couponCashableAmount := decimal.NewFromInt(0)
		if info.Cashable {
			couponCashableAmount = couponAmount
		}
		h.couponAmount = couponAmount
		h.couponCashableAmount = couponCashableAmount

		if err != nil {
			return err
		}
	}

	return nil
}

//nolint:funlen
func (h *rewardHandler) allocateCoins(ctx context.Context, ev *npool.Event) ([]*coinallocatedmwpb.CoinAllocated, error) {
	coinRewardIDs := []string{}
	for _, eventCoin := range ev.Coins {
		_id := eventCoin.CoinConfigID
		handler, err := coinconfig1.NewHandler(
			ctx,
			coinconfig1.WithEntID(&_id, true),
		)
		if err != nil {
			return nil, err
		}
		_coinConfig, err := handler.GetCoinConfig(ctx)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		coinPreUSD, err := decimal.NewFromString(eventCoin.CoinPreUSD)
		if err != nil {
			return nil, err
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
			return nil, err
		}

		maxValue, err := decimal.NewFromString(_coinConfig.MaxValue)
		if err != nil {
			return nil, err
		}

		if coins.Add(allocated).Cmp(maxValue) >= 0 {
			continue
		}
		coinsStr := coins.String()

		id := uuid.NewString()
		handler2, err := coinallocated1.NewHandler(
			ctx,
			coinallocated1.WithEntID(&id, true),
			coinallocated1.WithAppID(&_coinConfig.AppID, true),
			coinallocated1.WithUserID(&userID, true),
			coinallocated1.WithCoinConfigID(&_coinConfig.EntID, true),
			coinallocated1.WithCoinTypeID(&_coinConfig.CoinTypeID, true),
			coinallocated1.WithValue(&coinsStr, true),
		)
		if err != nil {
			return nil, wlog.Errorf("handler coinallocated1: %v", err)
		}
		coinRewardIDs = append(coinRewardIDs, id)

		if err := handler2.CreateCoinAllocated(ctx); err != nil {
			return nil, err
		}
	}

	handler3, err := coinallocated1.NewHandler(
		ctx,
		coinallocated1.WithConds(&coinallocatedmwpb.Conds{
			EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinRewardIDs},
		},
		),
		coinallocated1.WithOffset(0),
		coinallocated1.WithLimit(int32(len(coinRewardIDs))),
	)
	if err != nil {
		return nil, err
	}
	coinRewards, _, err := handler3.GetCoinAllocateds(ctx)
	if err != nil {
		return nil, err
	}

	return coinRewards, nil
}

func (h *rewardHandler) calculateCoinRewards(ctx context.Context, ev *npool.Event) ([]*coinallocatedmwpb.CoinAllocated, error) {
	coinRewards := []*coinallocatedmwpb.CoinAllocated{}
	for _, eventCoin := range ev.Coins {
		_id := eventCoin.CoinConfigID
		handler, err := coinconfig1.NewHandler(
			ctx,
			coinconfig1.WithEntID(&_id, true),
		)
		if err != nil {
			return nil, err
		}
		_coinConfig, err := handler.GetCoinConfig(ctx)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		coinPreUSD, err := decimal.NewFromString(eventCoin.CoinPreUSD)
		if err != nil {
			return nil, err
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
			return nil, err
		}

		maxValue, err := decimal.NewFromString(_coinConfig.MaxValue)
		if err != nil {
			return nil, err
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

func (h *rewardHandler) rewardSelf(ctx context.Context) (*npool.Reward, error) {
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

	credits, err := h.allocatedCredits(ctx, ev)
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

	_rewards := &npool.Reward{
		Credits:     _credits,
		CoinRewards: []*npool.CoinReward{},
	}

	return _rewards, nil
}

func (h *rewardHandler) rewardAffiliate(ctx context.Context) (*npool.Reward, error) {
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

		reward, err := _handler.rewardSelf(ctx)
		if err != nil {
			return nil, err
		}

		j--
		if len(reward.Credits) == 0 {
			continue
		}

		credits = append(credits, reward.Credits...)
	}

	_rewards := &npool.Reward{
		Credits:     credits,
		CoinRewards: []*npool.CoinReward{},
	}

	return _rewards, nil
}

func (h *rewardHandler) validateTask(ctx context.Context, ev *npool.Event) error {
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
		return err
	}
	configs, _, err := handler.GetTaskConfigs(ctx)
	if err != nil {
		return err
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
		return err
	}
	taskUsers, _, err := handler2.GetTaskUsers(ctx)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if len(taskUsers) == 0 {
		return nil
	}
	// check user has over the max finish times
	if len(taskUsers) >= int(configs[0].MaxRewardCount) {
		return wlog.Errorf("invalid maxrewardcount")
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
			return err
		}
		lastTaskUsers, _, err := handler3.GetTaskUsers(ctx)
		if err != nil {
			return err
		}
		if len(lastTaskUsers) == 0 {
			return wlog.Errorf("invalid last task")
		}
	}

	return nil
}

func (h *rewardHandler) rewardTask(ctx context.Context) (*npool.Reward, error) {
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

	if err := h.validateTask(ctx, ev); err != nil {
		return nil, err
	}

	h.addCredits = decimal.NewFromInt(0)
	h.coinPreUSDAmount = decimal.NewFromInt(0)
	h.couponAmount = decimal.NewFromInt(0)
	h.couponCashableAmount = decimal.NewFromInt(0)
	credits, err := h.allocatedCredits(ctx, ev)
	if err != nil {
		return nil, err
	}

	// We don't care about result of allocate coupon
	if err := h.allocateCoupons(ctx, ev); err != nil {
		logger.Sugar().Warnw(
			"rewardTask allocateCoupons",
			"Event", ev,
			"Error", err,
		)
	}
	allocateCoinRewards, err := h.allocateCoins(ctx, ev)
	if err != nil {
		logger.Sugar().Warnw(
			"rewardTask allocateCoins",
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
	h.addCredits = credits

	if err := h.createTaskUser(ctx, ev); err != nil {
		return nil, err
	}
	if err := h.createOrUpdateUserReward(ctx, ev); err != nil {
		return nil, err
	}

	coinRewards := []*npool.CoinReward{}
	for _, coin := range allocateCoinRewards {
		coinReward := npool.CoinReward{
			AppID:       coin.AppID,
			UserID:      coin.UserID,
			CoinTypeID:  coin.CoinTypeID,
			CoinRewards: coin.Value,
		}
		coinRewards = append(coinRewards, &coinReward)
	}

	_rewards := &npool.Reward{
		Credits:     _credits,
		CoinRewards: coinRewards,
	}

	return _rewards, nil
}

func (h *rewardHandler) createTaskUser(ctx context.Context, ev *npool.Event) error {
	userID := h.UserID.String()
	taskState := types.TaskState_Done
	rewardState := types.RewardState_Issued
	handler, err := taskuser1.NewHandler(
		ctx,
		taskuser1.WithAppID(&ev.AppID, true),
		taskuser1.WithUserID(&userID, true),
		taskuser1.WithTaskID(&h.taskConfig.EntID, true),
		taskuser1.WithEventID(&ev.EntID, true),
		taskuser1.WithTaskState(&taskState, true),
		taskuser1.WithRewardState(&rewardState, true),
	)
	if err != nil {
		return err
	}
	if err := handler.CreateTaskUser(ctx); err != nil {
		return err
	}
	return nil
}

func (h *rewardHandler) createOrUpdateUserReward(ctx context.Context, ev *npool.Event) error {
	userID := h.UserID.String()
	handler, err := userreward1.NewHandler(
		ctx,
		userreward1.WithConds(&userrewardmwpb.Conds{
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ev.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: userID},
		}),
		userreward1.WithOffset(0),
		userreward1.WithLimit(constant.DefaultRowLimit),
	)
	if err != nil {
		return err
	}
	userrewards, _, err := handler.GetUserRewards(ctx)
	if err != nil {
		return err
	}
	if len(userrewards) == 0 {
		credits := h.addCredits.String()
		couponAmount := h.couponAmount.String()
		couponCashableAmount := h.couponCashableAmount.String()
		handler2, err := userreward1.NewHandler(
			ctx,
			userreward1.WithAppID(&ev.AppID, true),
			userreward1.WithUserID(&userID, true),
			userreward1.WithActionCredits(&credits, true),
			userreward1.WithCouponAmount(&couponAmount, true),
			userreward1.WithCouponCashableAmount(&couponCashableAmount, true),
		)
		if err != nil {
			return err
		}
		if err := handler2.CreateUserReward(ctx); err != nil {
			return err
		}
		return nil
	}

	oldActionCredits, err := decimal.NewFromString(userrewards[0].ActionCredits)
	if err != nil {
		return err
	}
	actionCredits := oldActionCredits.Add(h.addCredits).String()

	oldCouponAmount, err := decimal.NewFromString(userrewards[0].CouponAmount)
	if err != nil {
		return err
	}
	couponAmount := oldCouponAmount.Add(h.couponAmount).String()
	oldCouponCashableAmount, err := decimal.NewFromString(userrewards[0].CouponCashableAmount)
	if err != nil {
		return err
	}
	couponCashableAmount := oldCouponCashableAmount.Add(h.couponAmount).String()

	handler2, err := userreward1.NewHandler(
		ctx,
		userreward1.WithID(&userrewards[0].ID, true),
		userreward1.WithActionCredits(&actionCredits, true),
		userreward1.WithCouponAmount(&couponAmount, true),
		userreward1.WithCouponCashableAmount(&couponCashableAmount, true),
	)
	if err != nil {
		return err
	}
	if err := handler2.UpdateUserReward(ctx); err != nil {
		return err
	}
	return nil
}

func (h *rewardHandler) calcluateEventRewards(ctx context.Context) (*npool.Reward, error) {
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

	if err := h.validateTask(ctx, ev); err != nil {
		return nil, err
	}

	h.addCredits = decimal.NewFromInt(0)
	h.coinPreUSDAmount = decimal.NewFromInt(0)
	h.couponAmount = decimal.NewFromInt(0)
	h.couponCashableAmount = decimal.NewFromInt(0)
	credits, err := h.calculateCredits(ev)
	if err != nil {
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
			AppID:       coin.AppID,
			UserID:      coin.UserID,
			CoinTypeID:  coin.CoinTypeID,
			CoinRewards: coin.Value,
		}
		coinRewards = append(coinRewards, &coinReward)
	}

	_rewards := &npool.Reward{
		TaskID:      h.taskConfig.EntID,
		Credits:     _credits,
		CoinRewards: coinRewards,
	}

	return _rewards, nil
}

func (h *Handler) CalcluateEventRewards(ctx context.Context) (*npool.Reward, error) {
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
	case basetypes.UsedFor_NewLogin:
		fallthrough //nolint
	case basetypes.UsedFor_SetWithdrawAddress:
		return handler.calcluateEventRewards(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.rewardAffiliate(ctx)
	default:
		return nil, wlog.Errorf("not implemented")
	}
}

func (h *Handler) RewardEvent(ctx context.Context) (*npool.Reward, error) {
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
	case basetypes.UsedFor_NewLogin:
		fallthrough //nolint
	case basetypes.UsedFor_SetWithdrawAddress:
		return handler.rewardTask(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.rewardAffiliate(ctx)
	default:
		return nil, wlog.Errorf("not implemented")
	}
}
