package invitationcode

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	codegenerator "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode/generator"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
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
	ret = npool.InvitationCode{
		ID:     uuid.NewString(),
		AppID:  uuid.NewString(),
		UserID: uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createInvitationCode(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, len(info.InvitationCode), codegenerator.InvitationCodeLen)
		ret.InvitationCode = info.InvitationCode
		assert.Equal(t, info, &ret)
	}
}

func updateInvitationCode(t *testing.T) {
	ret.Disabled = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithDisabled(&ret.Disabled, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getInvitationCode(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetInvitationCode(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getInvitationCodes(t *testing.T) {
	conds := &npool.Conds{
		ID:             &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		InvitationCode: &basetypes.StringVal{Op: cruder.EQ, Value: ret.InvitationCode},
		Disabled:       &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Disabled},
		UserIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetInvitationCodes(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteInvitationCode(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteInvitationCode(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestInvitationCode(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createInvitationCode", createInvitationCode)
	t.Run("updateInvitationCode", updateInvitationCode)
	t.Run("getInvitationCode", getInvitationCode)
	t.Run("getInvitationCodes", getInvitationCodes)
	t.Run("deleteInvitationCode", deleteInvitationCode)
}
