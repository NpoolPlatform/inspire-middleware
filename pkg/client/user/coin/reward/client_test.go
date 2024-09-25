package reward

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/coin/reward"

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

var ret = &npool.UserCoinReward{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	UserID:      uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
	CoinRewards: decimal.RequireFromString("12.25").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserCoinReward(t *testing.T) {
	err := CreateUserCoinReward(context.Background(), &npool.UserCoinRewardReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		UserID:      &ret.UserID,
		CoinTypeID:  &ret.CoinTypeID,
		CoinRewards: &ret.CoinRewards,
	})
	if assert.Nil(t, err) {
		info, err := GetUserCoinReward(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateUserCoinReward(t *testing.T) {
	err := UpdateUserCoinReward(context.Background(), &npool.UserCoinRewardReq{
		ID:          &ret.ID,
		EntID:       &ret.EntID,
		CoinRewards: &ret.CoinRewards,
	})
	if assert.Nil(t, err) {
		info, err := GetUserCoinReward(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getUserCoinReward(t *testing.T) {
	info, err := GetUserCoinReward(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getUserCoinRewards(t *testing.T) {
	infos, total, err := GetUserCoinRewards(context.Background(), &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getUserCoinRewardOnly(t *testing.T) {
	info, err := GetUserCoinRewardOnly(context.Background(), &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteUserCoinReward(t *testing.T) {
	err := DeleteUserCoinReward(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetUserCoinReward(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserCoinReward(t *testing.T) {
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

	t.Run("createUserCoinReward", createUserCoinReward)
	t.Run("updateUserCoinReward", updateUserCoinReward)
	t.Run("getUserCoinReward", getUserCoinReward)
	t.Run("getUserCoinRewards", getUserCoinRewards)
	t.Run("getUserCoinRewardOnly", getUserCoinRewardOnly)
	t.Run("deleteUserCoinReward", deleteUserCoinReward)
}
