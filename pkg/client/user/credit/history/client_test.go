package history

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"

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

var ret = &npool.UserCreditHistory{
	EntID:   uuid.NewString(),
	AppID:   uuid.NewString(),
	UserID:  uuid.NewString(),
	TaskID:  uuid.NewString(),
	EventID: uuid.NewString(),
	Credits: decimal.RequireFromString("12.25").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserCreditHistory(t *testing.T) {
	err := CreateUserCreditHistory(context.Background(), &npool.UserCreditHistoryReq{
		EntID:   &ret.EntID,
		AppID:   &ret.AppID,
		UserID:  &ret.UserID,
		TaskID:  &ret.TaskID,
		EventID: &ret.EventID,
		Credits: &ret.Credits,
	})
	if assert.Nil(t, err) {
		info, err := GetUserCreditHistory(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func getUserCreditHistory(t *testing.T) {
	info, err := GetUserCreditHistory(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getUserCreditHistories(t *testing.T) {
	infos, total, err := GetUserCreditHistories(context.Background(), &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TaskID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getUserCreditHistoryOnly(t *testing.T) {
	info, err := GetUserCreditHistoryOnly(context.Background(), &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TaskID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteUserCreditHistory(t *testing.T) {
	err := DeleteUserCreditHistory(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetUserCreditHistory(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserCreditHistory(t *testing.T) {
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

	t.Run("createUserCreditHistory", createUserCreditHistory)
	t.Run("getUserCreditHistory", getUserCreditHistory)
	t.Run("getUserCreditHistories", getUserCreditHistories)
	t.Run("getUserCreditHistoryOnly", getUserCreditHistoryOnly)
	t.Run("deleteUserCreditHistory", deleteUserCreditHistory)
}
