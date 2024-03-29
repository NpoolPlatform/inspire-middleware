package invitationcode

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

var ret = &npool.InvitationCode{
	EntID:    uuid.NewString(),
	AppID:    uuid.NewString(),
	UserID:   uuid.NewString(),
	Disabled: false,
}

func create(t *testing.T) {
	info, err := CreateInvitationCode(context.Background(), &npool.InvitationCodeReq{
		EntID:  &ret.EntID,
		AppID:  &ret.AppID,
		UserID: &ret.UserID,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.InvitationCode = info.InvitationCode
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	ret.Disabled = true

	info, err := UpdateInvitationCode(context.Background(), &npool.InvitationCodeReq{
		ID:       &ret.ID,
		Disabled: &ret.Disabled,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getOnly(t *testing.T) {
	info, err := GetInvitationCodeOnly(context.Background(), &npool.Conds{
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		InvitationCode: &basetypes.StringVal{Op: cruder.EQ, Value: ret.InvitationCode},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func TestInvitationCode(t *testing.T) {
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

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getOnly", getOnly)
}
