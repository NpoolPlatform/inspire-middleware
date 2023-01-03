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

func create(t *testing.T) {
	info, err := CreateRegistration(context.Background(), req)
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
