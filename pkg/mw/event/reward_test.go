package event

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	coinconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	taskconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/task/config"
	coinconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
	taskconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	coinConfig = coinconfigmwpb.CoinConfig{
		EntID:      uuid.NewString(),
		AppID:      uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		MaxValue:   decimal.RequireFromString("20").String(),
		Allocated:  decimal.RequireFromString("0").String(),
	}
	taskConfig = taskconfigmwpb.TaskConfig{
		EntID:            uuid.NewString(),
		AppID:            uuid.NewString(),
		EventID:          uuid.NewString(),
		Name:             uuid.NewString(),
		TaskDesc:         uuid.NewString(),
		StepGuide:        uuid.NewString(),
		RecommendMessage: uuid.NewString(),
		Index:            uint32(1),
		LastTaskID:       uuid.NewString(),
		MaxRewardCount:   uint32(10),
		CooldownSecord:   uint32(10),
		TaskType:         types.TaskType_BaseTask,
		TaskTypeStr:      types.TaskType_BaseTask.String(),
	}
	eventcoupon = couponmwpb.Coupon{
		EntID:               uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		AppID:               uuid.NewString(),
		Denomination:        decimal.RequireFromString("12.25").String(),
		Circulation:         decimal.RequireFromString("12.25").String(),
		IssuedBy:            uuid.NewString(),
		StartAt:             uint32(time.Now().Unix()),
		EndAt:               uint32(time.Now().Add(24 * time.Hour).Unix()),
		DurationDays:        234,
		Message:             uuid.NewString(),
		Name:                uuid.NewString(),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Allocated:           decimal.NewFromInt(0).String(),
		Threshold:           decimal.NewFromInt(0).String(),
		CashableProbability: decimal.RequireFromString("0.0001").String(),
	}
	eventRet = npool.Event{
		EntID:          uuid.NewString(),
		AppID:          uuid.NewString(),
		EventType:      basetypes.UsedFor_Signup,
		EventTypeStr:   basetypes.UsedFor_Signup.String(),
		CouponIDs:      []string{eventcoupon.EntID},
		Credits:        decimal.RequireFromString("12.25").String(),
		CreditsPerUSD:  decimal.RequireFromString("12.25").String(),
		MaxConsecutive: 1,
		InviterLayers:  2,
	}
)

//nolint:funlen,dupl
func resetup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&eventcoupon.EntID, true),
		coupon1.WithAppID(&eventcoupon.AppID, true),
		coupon1.WithName(&eventcoupon.Name, true),
		coupon1.WithMessage(&eventcoupon.Message, true),
		coupon1.WithCouponType(&eventcoupon.CouponType, true),
		coupon1.WithDenomination(&eventcoupon.Denomination, true),
		coupon1.WithCouponScope(&eventcoupon.CouponScope, true),
		coupon1.WithCirculation(&eventcoupon.Circulation, true),
		coupon1.WithDurationDays(&eventcoupon.DurationDays, true),
		coupon1.WithIssuedBy(&eventcoupon.IssuedBy, true),
		coupon1.WithStartAt(&eventcoupon.StartAt, true),
		coupon1.WithEndAt(&eventcoupon.EndAt, true),
		coupon1.WithCashableProbability(&eventcoupon.CashableProbability, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		eventcoupon.ID = info.ID
		eventcoupon.CreatedAt = info.CreatedAt
		eventcoupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &coupon, info)
		h1.ID = &info.ID
	}

	h2, err := taskconfig1.NewHandler(
		context.Background(),
		taskconfig1.WithEntID(&taskConfig.EntID, true),
		taskconfig1.WithAppID(&taskConfig.AppID, true),
		taskconfig1.WithEventID(&taskConfig.EventID, true),
		taskconfig1.WithName(&taskConfig.Name, true),
		taskconfig1.WithTaskDesc(&taskConfig.TaskDesc, true),
		taskconfig1.WithStepGuide(&taskConfig.StepGuide, true),
		taskconfig1.WithRecommendMessage(&taskConfig.RecommendMessage, true),
		taskconfig1.WithIndex(&taskConfig.Index, true),
		taskconfig1.WithLastTaskID(&taskConfig.LastTaskID, true),
		taskconfig1.WithMaxRewardCount(&taskConfig.MaxRewardCount, true),
		taskconfig1.WithCooldownSecord(&taskConfig.CooldownSecord, true),
		taskconfig1.WithTaskType(&taskConfig.TaskType, true),
	)
	assert.Nil(t, err)

	err = h2.CreateTaskConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h2.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			taskConfig.ID = info.ID
			taskConfig.CreatedAt = info.CreatedAt
			taskConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &taskConfig, info)
			h2.ID = &info.ID
		}
	}

	h3, err := coinconfig1.NewHandler(
		context.Background(),
		coinconfig1.WithEntID(&coinConfig.EntID, true),
		coinconfig1.WithAppID(&coinConfig.AppID, true),
		coinconfig1.WithCoinTypeID(&coinConfig.CoinTypeID, true),
		coinconfig1.WithMaxValue(&coinConfig.MaxValue, true),
		coinconfig1.WithAllocated(&coinConfig.Allocated, true),
	)
	assert.Nil(t, err)

	err = h3.CreateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h2.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			coinConfig.ID = info.ID
			coinConfig.CreatedAt = info.CreatedAt
			coinConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &coinConfig, info)
			h3.ID = &info.ID
		}
	}

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&eventRet.EntID, true),
		WithAppID(&eventRet.AppID, true),
		WithEventType(&eventRet.EventType, true),
		WithCouponIDs(eventRet.CouponIDs, true),
		WithCredits(&eventRet.Credits, true),
		WithCreditsPerUSD(&eventRet.CreditsPerUSD, true),
		WithMaxConsecutive(&eventRet.MaxConsecutive, true),
		WithInviterLayers(&eventRet.InviterLayers, true),
	)
	assert.Nil(t, err)

	info2, err := handler.CreateEvent(context.Background())
	if assert.Nil(t, err) {
		eventRet.ID = info2.ID
		eventRet.CreatedAt = info2.CreatedAt
		eventRet.UpdatedAt = info2.UpdatedAt
		eventRet.CouponIDsStr = info2.CouponIDsStr
		assert.Equal(t, info2, &ret)
		handler.ID = &info2.ID
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
		_ = h2.DeleteTaskConfig(context.Background())
		_ = h3.DeleteCoinConfig(context.Background())
		_, _ = handler.DeleteEvent(context.Background())
	}
}

func rewardEvent(t *testing.T) {
	userID := uuid.NewString()
	eventType := basetypes.UsedFor_KYCApproved
	consecutive := uint32(1)
	amount := decimal.NewFromInt(10).String()
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&eventRet.AppID, true),
		WithUserID(&userID, true),
		WithEventType(&eventType, true),
		WithConsecutive(&consecutive, true),
		WithAmount(&amount, false),
	)
	assert.Nil(t, err)

	credit, err := handler.RewardEvent(context.Background())
	if assert.Nil(t, err) {
		fmt.Println("credit: ", credit)
	}
}

func TestReward(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := resetup(t)
	defer teardown(t)

	t.Run("rewardEvent", rewardEvent)
}
