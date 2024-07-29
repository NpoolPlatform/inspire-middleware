package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/shopspring/decimal"

	coinconfigmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coin/config"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	coinconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/config"
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

var (
	ret = &npool.CoinAllocated{
		EntID:        uuid.NewString(),
		AppID:        uuid.NewString(),
		CoinConfigID: uuid.NewString(),
		CoinTypeID:   uuid.NewString(),
		UserID:       uuid.NewString(),
		Value:        decimal.RequireFromString("12.25").String(),
	}
	coinConfig = coinconfigmwpb.CoinConfig{
		EntID:      ret.CoinConfigID,
		AppID:      ret.AppID,
		CoinTypeID: ret.CoinTypeID,
		MaxValue:   decimal.RequireFromString("50").String(),
		Allocated:  decimal.RequireFromString("0").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	err := coinconfigmwcli.CreateCoinConfig(context.Background(), &coinconfigmwpb.CoinConfigReq{
		EntID:      &coinConfig.EntID,
		AppID:      &coinConfig.AppID,
		CoinTypeID: &coinConfig.CoinTypeID,
		MaxValue:   &coinConfig.MaxValue,
		Allocated:  &coinConfig.Allocated,
	})
	if assert.Nil(t, err) {
		info, err := coinconfigmwcli.GetCoinConfigOnly(context.Background(), &coinconfigmwpb.Conds{
			EntID: &basetypes.StringVal{Op: cruder.EQ, Value: coinConfig.EntID},
		})
		if assert.Nil(t, err) {
			coinConfig.ID = info.ID
			coinConfig.CreatedAt = info.CreatedAt
			coinConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &coinConfig, info)
		}
	}
	return func(*testing.T) {
		_ = coinconfigmwcli.DeleteCoinConfig(context.Background(), &coinConfig.ID, &coinConfig.EntID)
	}
}

func createCoinAllocated(t *testing.T) {
	err := CreateCoinAllocated(context.Background(), &npool.CoinAllocatedReq{
		EntID:        &ret.EntID,
		AppID:        &ret.AppID,
		CoinConfigID: &ret.CoinConfigID,
		CoinTypeID:   &ret.CoinTypeID,
		UserID:       &ret.UserID,
		Value:        &ret.Value,
	})
	if assert.Nil(t, err) {
		info, err := GetCoinAllocated(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func getCoinAllocated(t *testing.T) {
	info, err := GetCoinAllocated(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getCoinAllocateds(t *testing.T) {
	infos, total, err := GetCoinAllocateds(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinConfigID},
		CoinTypeID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getCoinAllocatedOnly(t *testing.T) {
	info, err := GetCoinAllocatedOnly(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinConfigID},
		CoinTypeID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteCoinAllocated(t *testing.T) {
	err := DeleteCoinAllocated(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetCoinAllocated(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoinAllocated(t *testing.T) {
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

	t.Run("createCoinAllocated", createCoinAllocated)
	t.Run("getCoinAllocated", getCoinAllocated)
	t.Run("getCoinAllocateds", getCoinAllocateds)
	t.Run("getCoinAllocatedOnly", getCoinAllocatedOnly)
	t.Run("deleteCoinAllocated", deleteCoinAllocated)
}
