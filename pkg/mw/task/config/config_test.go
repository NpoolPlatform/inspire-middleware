package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/config"
	"github.com/google/uuid"
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
	ret = npool.TaskConfig{
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
		MaxRewardCount:   uint32(10),
		CooldownSecord:   uint32(10),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createTaskConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEventID(&ret.EventID, true),
		WithTaskType(&ret.TaskType, true),
		WithName(&ret.Name, true),
		WithTaskDesc(&ret.TaskDesc, true),
		WithStepGuide(&ret.StepGuide, true),
		WithRecommendMessage(&ret.RecommendMessage, true),
		WithIndex(&ret.Index, true),
		WithLastTaskID(&ret.LastTaskID, true),
		WithMaxRewardCount(&ret.MaxRewardCount, true),
		WithCooldownSecord(&ret.CooldownSecord, true),
	)
	assert.Nil(t, err)

	err = handler.CreateTaskConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateTaskConfig(t *testing.T) {
	ret.EventID = uuid.NewString()
	ret.TaskType = types.TaskType_DailyTask
	ret.TaskTypeStr = types.TaskType_DailyTask.String()
	ret.Name = uuid.NewString()
	ret.TaskDesc = uuid.NewString()
	ret.StepGuide = uuid.NewString()
	ret.RecommendMessage = uuid.NewString()
	ret.Index = uint32(2)
	ret.LastTaskID = uuid.NewString()
	ret.MaxRewardCount = uint32(5)
	ret.CooldownSecord = uint32(5)

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEventID(&ret.EventID, true),
		WithTaskType(&ret.TaskType, true),
		WithName(&ret.Name, true),
		WithTaskDesc(&ret.TaskDesc, true),
		WithStepGuide(&ret.StepGuide, true),
		WithRecommendMessage(&ret.RecommendMessage, true),
		WithIndex(&ret.Index, true),
		WithLastTaskID(&ret.LastTaskID, true),
		WithMaxRewardCount(&ret.MaxRewardCount, true),
		WithCooldownSecord(&ret.CooldownSecord, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateTaskConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getTaskConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetTaskConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTaskConfigs(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		Name:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
		LastTaskID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.LastTaskID},
		EventID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		TaskType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskType)},
		Index:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Index},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetTaskConfigs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteTaskConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteTaskConfig(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetTaskConfig(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTaskConfig(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createTaskConfig", createTaskConfig)
	t.Run("updateTaskConfig", updateTaskConfig)
	t.Run("getTaskConfig", getTaskConfig)
	t.Run("getTaskConfigs", getTaskConfigs)
	t.Run("deleteTaskConfig", deleteTaskConfig)
}
