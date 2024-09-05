package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/shopspring/decimal"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.CoinConfig{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	MaxValue:   decimal.RequireFromString("202.25").String(),
	Allocated:  decimal.RequireFromString("0").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCoinConfig(t *testing.T) {
	err := CreateCoinConfig(context.Background(), &npool.CoinConfigReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		CoinTypeID: &ret.CoinTypeID,
		MaxValue:   &ret.MaxValue,
	})
	if assert.Nil(t, err) {
		info, err := GetCoinConfig(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateCoinConfig(t *testing.T) {
	err := UpdateCoinConfig(context.Background(), &npool.CoinConfigReq{
		ID:        &ret.ID,
		EntID:     &ret.EntID,
		MaxValue:  &ret.MaxValue,
		Allocated: &ret.Allocated,
	})
	if assert.Nil(t, err) {
		info, err := GetCoinConfig(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getCoinConfig(t *testing.T) {
	info, err := GetCoinConfig(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getCoinConfigs(t *testing.T) {
	infos, total, err := GetCoinConfigs(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getCoinConfigOnly(t *testing.T) {
	info, err := GetCoinConfigOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteCoinConfig(t *testing.T) {
	err := DeleteCoinConfig(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetCoinConfig(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoinConfig(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCoinConfig", createCoinConfig)
	t.Run("updateCoinConfig", updateCoinConfig)
	t.Run("getCoinConfig", getCoinConfig)
	t.Run("getCoinConfigs", getCoinConfigs)
	t.Run("getCoinConfigOnly", getCoinConfigOnly)
	t.Run("deleteCoinConfig", deleteCoinConfig)
}
