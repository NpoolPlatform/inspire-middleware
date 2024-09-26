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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"

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

var ret = &npool.UserReward{
	EntID:                uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	ActionCredits:        decimal.RequireFromString("11.25").String(),
	CouponAmount:         decimal.RequireFromString("11.25").String(),
	CouponCashableAmount: decimal.RequireFromString("11.25").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserReward(t *testing.T) {
	err := AddUserReward(context.Background(), &npool.UserRewardReq{
		EntID:                &ret.EntID,
		AppID:                &ret.AppID,
		UserID:               &ret.UserID,
		ActionCredits:        &ret.ActionCredits,
		CouponAmount:         &ret.CouponAmount,
		CouponCashableAmount: &ret.CouponCashableAmount,
	})
	if assert.Nil(t, err) {
		info, err := GetUserReward(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

//nolint:dupl
func addUserReward(t *testing.T) {
	addActionCredits := decimal.RequireFromString("22.25").String()
	addCouponAmount := decimal.RequireFromString("22.25").String()
	addCouponCashableAmount := decimal.RequireFromString("22.25").String()

	ret.ActionCredits = decimal.RequireFromString("33.5").String()
	ret.CouponAmount = decimal.RequireFromString("33.5").String()
	ret.CouponCashableAmount = decimal.RequireFromString("33.5").String()
	err := AddUserReward(context.Background(), &npool.UserRewardReq{
		ID:                   &ret.ID,
		EntID:                &ret.EntID,
		AppID:                &ret.AppID,
		UserID:               &ret.UserID,
		ActionCredits:        &addActionCredits,
		CouponAmount:         &addCouponAmount,
		CouponCashableAmount: &addCouponCashableAmount,
	})
	if assert.Nil(t, err) {
		info, err := GetUserReward(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

//nolint:dupl
func subUserReward(t *testing.T) {
	subActionCredits := decimal.RequireFromString("22.25").String()
	subCouponAmount := decimal.RequireFromString("22.25").String()
	subCouponCashableAmount := decimal.RequireFromString("22.25").String()

	ret.ActionCredits = decimal.RequireFromString("11.25").String()
	ret.CouponAmount = decimal.RequireFromString("11.25").String()
	ret.CouponCashableAmount = decimal.RequireFromString("11.25").String()
	err := SubUserReward(context.Background(), &npool.UserRewardReq{
		ID:                   &ret.ID,
		EntID:                &ret.EntID,
		AppID:                &ret.AppID,
		UserID:               &ret.UserID,
		ActionCredits:        &subActionCredits,
		CouponAmount:         &subCouponAmount,
		CouponCashableAmount: &subCouponCashableAmount,
	})
	if assert.Nil(t, err) {
		info, err := GetUserReward(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getUserReward(t *testing.T) {
	info, err := GetUserReward(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getUserRewards(t *testing.T) {
	infos, total, err := GetUserRewards(context.Background(), &npool.Conds{
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

func getUserRewardOnly(t *testing.T) {
	info, err := GetUserRewardOnly(context.Background(), &npool.Conds{
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

func deleteUserReward(t *testing.T) {
	err := DeleteUserReward(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetUserReward(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserReward(t *testing.T) {
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

	t.Run("createUserReward", createUserReward)
	t.Run("addUserReward", addUserReward)
	t.Run("subUserReward", subUserReward)
	t.Run("getUserReward", getUserReward)
	t.Run("getUserRewards", getUserRewards)
	t.Run("getUserRewardOnly", getUserRewardOnly)
	t.Run("deleteUserReward", deleteUserReward)
}
