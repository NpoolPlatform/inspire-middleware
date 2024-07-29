package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/task/user"
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
	ret = npool.TaskUser{
		EntID:          uuid.NewString(),
		AppID:          uuid.NewString(),
		UserID:         uuid.NewString(),
		TaskID:         uuid.NewString(),
		EventID:        uuid.NewString(),
		TaskState:      types.TaskState_InProgress,
		TaskStateStr:   types.TaskState_InProgress.String(),
		RewardState:    types.RewardState_UnIssued,
		RewardStateStr: types.RewardState_UnIssued.String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createTaskUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithTaskID(&ret.TaskID, true),
		WithEventID(&ret.EventID, true),
		WithTaskState(&ret.TaskState, true),
		WithRewardState(&ret.RewardState, true),
	)
	assert.Nil(t, err)

	err = handler.CreateTaskUser(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetTaskUser(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateTaskUser(t *testing.T) {
	ret.TaskState = types.TaskState_InProgress
	ret.TaskStateStr = types.TaskState_InProgress.String()
	ret.RewardState = types.RewardState_UnIssued
	ret.RewardStateStr = types.RewardState_UnIssued.String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithTaskState(&ret.TaskState, true),
		WithRewardState(&ret.RewardState, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateTaskUser(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetTaskUser(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getTaskUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetTaskUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTaskUsers(t *testing.T) {
	conds := &npool.Conds{
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		TaskID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		EventID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		EntIDs:      &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		TaskState:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TaskState)},
		RewardState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.RewardState)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetTaskUsers(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteTaskUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteTaskUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetTaskUser(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTaskUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createTaskUser", createTaskUser)
	t.Run("updateTaskUser", updateTaskUser)
	t.Run("getTaskUser", getTaskUser)
	t.Run("getTaskUsers", getTaskUsers)
	t.Run("deleteTaskUser", deleteTaskUser)
}
