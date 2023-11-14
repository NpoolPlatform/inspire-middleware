package registration

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"
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

var ret = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}
var updateInviterID = uuid.NewString()

func setup(t *testing.T) func(*testing.T) {
	h, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(&ret.AppID, true),
		invitationcode1.WithUserID(&ret.InviterID, true),
	)
	assert.Nil(t, err)

	info, err := h.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
	h.ID = &info.ID

	h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(&ret.AppID, true),
		invitationcode1.WithUserID(&updateInviterID, true),
	)
	assert.Nil(t, err)

	info, err = h1.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
	h1.ID = &info.ID

	return func(*testing.T) {
		_, _ = h.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteInvitationCode(context.Background())
	}
}

func createRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithInviterID(&ret.InviterID, true),
		WithInviteeID(&ret.InviteeID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateRegistration(t *testing.T) {
	ret.InviterID = updateInviterID
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithInviterID(&ret.InviterID, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateRegistration(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetRegistration(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRegistrations(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		InviterID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.InviterID},
		InviteeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.InviteeID},
		InviterIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.InviterID}},
		InviteeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.InviteeID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetRegistrations(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteRegistration(context.Background())
	if assert.Nil(t, err) {
		ret.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetRegistration(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRegistration(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRegistration", createRegistration)
	t.Run("updateRegistration", updateRegistration)
	t.Run("getRegistration", getRegistration)
	t.Run("getRegistrations", getRegistrations)
	t.Run("deleteRegistration", deleteRegistration)
}
