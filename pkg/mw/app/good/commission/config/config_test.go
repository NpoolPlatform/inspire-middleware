package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
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
	ret = npool.AppGoodCommissionConfig{
		EntID:           uuid.NewString(),
		AppID:           uuid.NewString(),
		GoodID:          uuid.NewString(),
		AppGoodID:       uuid.NewString(),
		SettleType:      types.SettleType_GoodOrderPayment,
		SettleTypeStr:   types.SettleType_GoodOrderPayment.String(),
		AmountOrPercent: decimal.RequireFromString("12.25").String(),
		ThresholdAmount: decimal.RequireFromString("12.26").String(),
		Invites:         uint32(1),
		StartAt:         uint32(time.Now().Unix()),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithSettleType(&ret.SettleType, true),
		WithThresholdAmount(&ret.ThresholdAmount, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithInvites(&ret.Invites, true),
		WithStartAt(&ret.StartAt, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCommissionConfig(t *testing.T) {
	ret.AmountOrPercent = "13"
	ret.StartAt += 10000
	ret.ThresholdAmount = decimal.NewFromInt(10).String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithStartAt(&ret.StartAt, true),
		WithThresholdAmount(&ret.ThresholdAmount, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCommissionConfigs(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCommissionConfigs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCommissionConfig(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCommission(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCommissionConfig", createCommissionConfig)
	t.Run("updateCommissionConfig", updateCommissionConfig)
	t.Run("getCommissionConfig", getCommissionConfig)
	t.Run("getCommissionConfigs", getCommissionConfigs)
	t.Run("deleteCommissionConfig", deleteCommissionConfig)
}
