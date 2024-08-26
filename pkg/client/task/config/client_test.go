package config

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"

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

var ret = &npool.TaskConfig{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	EventID:          uuid.NewString(),
	TaskType:         types.TaskType_BaseTask,
	TaskTypeStr:      types.TaskType_BaseTask.String(),
	Name:             uuid.NewString(),
	TaskDesc:         uuid.NewString(),
	StepGuide:        uuid.NewString(),
	RecommendMessage: uuid.NewString(),
	Index:            uint32(1),
	LastTaskID:       uuid.NewString(),
	MaxRewardCount:   uint32(1),
	CooldownSecond:   uint32(1),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createTaskConfig(t *testing.T) {
	err := CreateTaskConfig(context.Background(), &npool.TaskConfigReq{
		EntID:            &ret.EntID,
		AppID:            &ret.AppID,
		EventID:          &ret.EventID,
		TaskType:         &ret.TaskType,
		Name:             &ret.Name,
		TaskDesc:         &ret.TaskDesc,
		StepGuide:        &ret.StepGuide,
		RecommendMessage: &ret.RecommendMessage,
		Index:            &ret.Index,
		LastTaskID:       &ret.LastTaskID,
		MaxRewardCount:   &ret.MaxRewardCount,
		CooldownSecond:   &ret.CooldownSecond,
	})
	if assert.Nil(t, err) {
		info, err := GetTaskConfig(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateTaskConfig(t *testing.T) {
	err := UpdateTaskConfig(context.Background(), &npool.TaskConfigReq{
		ID:               &ret.ID,
		EntID:            &ret.EntID,
		EventID:          &ret.EventID,
		TaskType:         &ret.TaskType,
		Name:             &ret.Name,
		TaskDesc:         &ret.TaskDesc,
		StepGuide:        &ret.StepGuide,
		RecommendMessage: &ret.RecommendMessage,
		Index:            &ret.Index,
		LastTaskID:       &ret.LastTaskID,
		MaxRewardCount:   &ret.MaxRewardCount,
		CooldownSecond:   &ret.CooldownSecond,
	})
	if assert.Nil(t, err) {
		info, err := GetTaskConfig(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getTaskConfig(t *testing.T) {
	info, err := GetTaskConfig(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getTaskConfigs(t *testing.T) {
	infos, total, err := GetTaskConfigs(context.Background(), &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		LastTaskID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.LastTaskID},
		EventID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		Name:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		Index:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Index},
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		TaskType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskType)},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getTaskConfigOnly(t *testing.T) {
	info, err := GetTaskConfigOnly(context.Background(), &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		LastTaskID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.LastTaskID},
		EventID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		Name:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		Index:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Index},
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		TaskType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskType)},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteTaskConfig(t *testing.T) {
	err := DeleteTaskConfig(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetTaskConfig(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTaskConfig(t *testing.T) {
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

	t.Run("createTaskConfig", createTaskConfig)
	t.Run("updateTaskConfig", updateTaskConfig)
	t.Run("getTaskConfig", getTaskConfig)
	t.Run("getTaskConfigs", getTaskConfigs)
	t.Run("getTaskConfigOnly", getTaskConfigOnly)
	t.Run("deleteTaskConfig", deleteTaskConfig)
}
