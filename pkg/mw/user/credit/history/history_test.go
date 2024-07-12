package history

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
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
	ret = npool.UserCreditHistory{
		EntID:   uuid.NewString(),
		AppID:   uuid.NewString(),
		UserID:  uuid.NewString(),
		TaskID:  uuid.NewString(),
		EventID: uuid.NewString(),
		Credits: decimal.RequireFromString("11.25").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserCreditHistory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithTaskID(&ret.TaskID, true),
		WithEventID(&ret.EventID, true),
		WithCredits(&ret.Credits, true),
	)
	assert.Nil(t, err)

	err = handler.CreateUserCreditHistory(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserCreditHistory(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func getUserCreditHistory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetUserCreditHistory(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUserCreditHistories(t *testing.T) {
	conds := &npool.Conds{
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		TaskID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TaskID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		EntIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetUserCreditHistories(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteUserCreditHistory(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteUserCreditHistory(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetUserCreditHistory(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserCreditHistory(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createUserCreditHistory", createUserCreditHistory)
	t.Run("getUserCreditHistory", getUserCreditHistory)
	t.Run("getUserCreditHistories", getUserCreditHistories)
	t.Run("deleteUserCreditHistory", deleteUserCreditHistory)
}
