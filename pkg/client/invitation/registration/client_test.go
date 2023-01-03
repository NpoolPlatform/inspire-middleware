package registration

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	ivcodemwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/invitation/invitationcode"
	ivcodemgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

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
	_, err := ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &ret.AppID,
		UserID: &ret.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &ret1.AppID,
		UserID: &ret1.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &ret2.AppID,
		UserID: &ret2.InviterID,
	})
	assert.Nil(t, err)

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

	_, err := ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &ret.AppID,
		UserID: &inviterID,
	})
	assert.Nil(t, err)

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
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(2))
		assert.Equal(t, len(infos), 2)
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
	infos, total, err := GetSubordinates(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		InviterIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: []string{ret.InviterID},
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(3))

		found := false
		for _, info := range infos {
			if info.ID == ret.ID {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == ret1.ID {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == ret2.ID {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func getSuperiores(t *testing.T) {
	infos, total, err := GetSuperiores(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		InviteeIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: []string{ret2.InviteeID},
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(2))

		found := false
		for _, info := range infos {
			if info.ID == ret.ID {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == ret2.ID {
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func TestRegistration(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("create", create)
	t.Run("GetSubordinates", getSubordinates)
	t.Run("GetSuperiores", getSuperiores)
	t.Run("getRegistrations", getRegistrations)
	t.Run("getRegistrationOnly", getRegistrationOnly)
	t.Run("update", update)
}
