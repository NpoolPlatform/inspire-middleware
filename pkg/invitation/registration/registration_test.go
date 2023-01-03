package registration

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

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

var ret = &mgrpb.Registration{
	ID:        uuid.NewString(),
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var req = &mgrpb.RegistrationReq{
	ID:        &ret.ID,
	AppID:     &ret.AppID,
	InviterID: &ret.InviterID,
	InviteeID: &ret.InviteeID,
}

var ret1 = &mgrpb.Registration{
	ID:        uuid.NewString(),
	AppID:     ret.AppID,
	InviterID: ret.InviterID,
	InviteeID: uuid.NewString(),
}

var req1 = &mgrpb.RegistrationReq{
	ID:        &ret1.ID,
	AppID:     &ret.AppID,
	InviterID: &ret.InviterID,
	InviteeID: &ret1.InviteeID,
}

var ret2 = &mgrpb.Registration{
	ID:        uuid.NewString(),
	AppID:     ret.AppID,
	InviterID: ret.InviteeID,
	InviteeID: uuid.NewString(),
}

var req2 = &mgrpb.RegistrationReq{
	ID:        &ret2.ID,
	AppID:     &ret.AppID,
	InviterID: &ret.InviteeID,
	InviteeID: &ret2.InviteeID,
}

func create(t *testing.T) {
	info, err := CreateRegistration(context.Background(), req1)
	if assert.Nil(t, err) {
		ret1.CreatedAt = info.CreatedAt
		ret1.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret1, info)
	}

	info, err = CreateRegistration(context.Background(), req2)
	if assert.Nil(t, err) {
		ret2.CreatedAt = info.CreatedAt
		ret2.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret2, info)
	}

	info, err = CreateRegistration(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	inviterID := uuid.NewString()

	req.InviterID = &inviterID
	ret.InviterID = inviterID

	info, err := UpdateRegistration(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getRegistrations(t *testing.T) {
	infos, total, err := GetRegistrations(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		InviterID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.InviterID,
		},
	}, int32(0), int32(1))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func getRegistrationOnly(t *testing.T) {
	info, err := GetRegistrationOnly(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		InviteeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.InviteeID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getSubordinates(t *testing.T) {

}

func getSuperiores(t *testing.T) {

}

func TestRegistration(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction { //nolint:staticcheck
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getRegistrations", getRegistrations)
	t.Run("getRegistrationOnly", getRegistrationOnly)
	t.Run("GetSubordinates", getSubordinates)
	t.Run("GetSuperiores", getSuperiores)
}