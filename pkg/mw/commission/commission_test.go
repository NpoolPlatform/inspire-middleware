package commission

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
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
	ret = npool.Commission{
		ID:                  uuid.NewString(),
		AppID:               uuid.NewString(),
		UserID:              uuid.NewString(),
		GoodID:              uuid.NewString(),
		SettleType:          types.SettleType_GoodOrderPayment,
		SettleTypeStr:       types.SettleType_GoodOrderPayment.String(),
		SettleMode:          types.SettleMode_SettleWithGoodValue,
		SettleModeStr:       types.SettleMode_SettleWithGoodValue.String(),
		SettleAmountType:    types.SettleAmountType_SettleByPercent,
		SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
		SettleInterval:      types.SettleInterval_SettleYearly,
		SettleIntervalStr:   types.SettleInterval_SettleYearly.String(),
		AmountOrPercent:     decimal.RequireFromString("12.25").String(),
		StartAt:             uint32(time.Now().Unix()),
		Threshold:           decimal.RequireFromString("12.26").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithUserID(&ret.UserID),
		WithGoodID(&ret.GoodID),
		WithSettleType(&ret.SettleType),
		WithSettleMode(&ret.SettleMode),
		WithSettleAmountType(&ret.SettleAmountType),
		WithSettleInterval(&ret.SettleInterval),
		WithAmountOrPercent(&ret.AmountOrPercent),
		WithStartAt(&ret.StartAt),
		WithThreshold(&ret.Threshold),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAmountOrPercent(&ret.AmountOrPercent),
		WithStartAt(&ret.StartAt),
		WithThreshold(&ret.Threshold),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCommission(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetCommission(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCommissions(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCommissions(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCommission(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCommission(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCommission(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCommission", createCommission)
	t.Run("updateCommission", updateCommission)
	t.Run("getCommission", getCommission)
	t.Run("getCommissions", getCommissions)
	t.Run("deleteCommission", deleteCommission)
}
