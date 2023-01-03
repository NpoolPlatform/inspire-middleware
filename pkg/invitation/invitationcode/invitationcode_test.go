package invitationcode

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

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

var ret = &mgrpb.InvitationCode{
	ID:     uuid.NewString(),
	AppID:  uuid.NewString(),
	UserID: uuid.NewString(),
}

var req = &mgrpb.InvitationCodeReq{
	ID:     &ret.ID,
	AppID:  &ret.AppID,
	UserID: &ret.UserID,
}

func create(t *testing.T) {
	info, err := CreateInvitationCode(context.Background(), req)
	if assert.Nil(t, err) {
		ret.InvitationCode = info.InvitationCode
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {

}

func getOnly(t *testing.T) {

}

func TestInvitationCode(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction { //nolint:staticcheck
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getOnly", getOnly)
}
