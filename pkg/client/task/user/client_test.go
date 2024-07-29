package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"

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

var ret = &npool.TaskUser{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	TaskID:         uuid.NewString(),
	EventID:        uuid.NewString(),
	TaskState:      types.TaskState_Done,
	TaskStateStr:   types.TaskState_Done.String(),
	RewardState:    types.RewardState_Issued,
	RewardStateStr: types.RewardState_Issued.String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createTaskUser(t *testing.T) {
	err := CreateTaskUser(context.Background(), &npool.TaskUserReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		UserID:      &ret.UserID,
		TaskID:      &ret.TaskID,
		EventID:     &ret.EventID,
		TaskState:   &ret.TaskState,
		RewardState: &ret.RewardState,
	})
	if assert.Nil(t, err) {
		info, err := GetTaskUser(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateTaskUser(t *testing.T) {
	err := UpdateTaskUser(context.Background(), &npool.TaskUserReq{
		ID:          &ret.ID,
		EntID:       &ret.EntID,
		TaskState:   &ret.TaskState,
		RewardState: &ret.RewardState,
	})
	if assert.Nil(t, err) {
		info, err := GetTaskUser(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getTaskUser(t *testing.T) {
	info, err := GetTaskUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getTaskUsers(t *testing.T) {
	infos, total, err := GetTaskUsers(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		TaskID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:      &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		TaskState:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskState)},
		RewardState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.RewardState)},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getTaskUserOnly(t *testing.T) {
	info, err := GetTaskUserOnly(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		TaskID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs:      &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		TaskState:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskState)},
		RewardState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.RewardState)},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteTaskUser(t *testing.T) {
	err := DeleteTaskUser(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetTaskUser(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTaskUser(t *testing.T) {
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

	t.Run("createTaskUser", createTaskUser)
	t.Run("updateTaskUser", updateTaskUser)
	t.Run("getTaskUser", getTaskUser)
	t.Run("getTaskUsers", getTaskUsers)
	t.Run("getTaskUserOnly", getTaskUserOnly)
	t.Run("deleteTaskUser", deleteTaskUser)
}
