package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	coinconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	coinconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
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
	ret = npool.CoinAllocated{
		EntID:        uuid.NewString(),
		AppID:        uuid.NewString(),
		CoinConfigID: uuid.NewString(),
		CoinTypeID:   uuid.NewString(),
		UserID:       uuid.NewString(),
		Value:        decimal.RequireFromString("2.25").String(),
	}
	coinConfig = coinconfigmwpb.CoinConfig{
		EntID:      ret.CoinConfigID,
		AppID:      ret.AppID,
		CoinTypeID: ret.CoinTypeID,
		MaxValue:   decimal.RequireFromString("20").String(),
		Allocated:  decimal.RequireFromString("0").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := coinconfig1.NewHandler(
		context.Background(),
		coinconfig1.WithEntID(&coinConfig.EntID, true),
		coinconfig1.WithAppID(&coinConfig.AppID, true),
		coinconfig1.WithCoinTypeID(&coinConfig.CoinTypeID, true),
		coinconfig1.WithMaxValue(&coinConfig.MaxValue, true),
		coinconfig1.WithAllocated(&coinConfig.Allocated, true),
	)
	assert.Nil(t, err)

	err = h1.CreateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetCoinConfig(context.Background())
		if assert.Nil(t, err) {
			coinConfig.ID = info.ID
			coinConfig.CreatedAt = info.CreatedAt
			coinConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &coinConfig, info)
			h1.ID = &info.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteCoinConfig(context.Background())
	}
}

func createCoinAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCoinConfigID(&ret.CoinConfigID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithUserID(&ret.UserID, true),
		WithValue(&ret.Value, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoinAllocated(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoinAllocated(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func getCoinAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCoinAllocated(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCoinAllocateds(t *testing.T) {
	conds := &npool.Conds{
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinConfigID},
		CoinTypeID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		EntIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCoinAllocateds(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCoinAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCoinAllocated(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCoinAllocated(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoinAllocated(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCoinAllocated", createCoinAllocated)
	t.Run("getCoinAllocated", getCoinAllocated)
	t.Run("getCoinAllocateds", getCoinAllocateds)
	t.Run("deleteCoinAllocated", deleteCoinAllocated)
}
