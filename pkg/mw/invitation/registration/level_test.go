//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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

var reg1 = npool.Registration{
	ID:        uuid.NewString(),
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _reg1 = npool.RegistrationReq{
	ID:        &reg1.ID,
	AppID:     &reg1.AppID,
	InviterID: &reg1.InviterID,
	InviteeID: &reg1.InviteeID,
}

var reg2 = npool.Registration{
	ID:        uuid.NewString(),
	AppID:     reg1.AppID,
	InviterID: reg1.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg2 = npool.RegistrationReq{
	ID:        &reg2.ID,
	AppID:     &reg2.AppID,
	InviterID: &reg2.InviterID,
	InviteeID: &reg2.InviteeID,
}

var reg3 = npool.Registration{
	ID:        uuid.NewString(),
	AppID:     reg1.AppID,
	InviterID: reg2.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg3 = npool.RegistrationReq{
	ID:        &reg3.ID,
	AppID:     &reg3.AppID,
	InviterID: &reg3.InviterID,
	InviteeID: &reg3.InviteeID,
}

var reg4 = npool.Registration{
	ID:        uuid.NewString(),
	AppID:     reg1.AppID,
	InviterID: reg3.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg4 = npool.RegistrationReq{
	ID:        &reg4.ID,
	AppID:     &reg4.AppID,
	InviterID: &reg4.InviterID,
	InviteeID: &reg4.InviteeID,
}

var reg5 = npool.Registration{
	ID:        uuid.NewString(),
	AppID:     reg1.AppID,
	InviterID: reg4.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg5 = npool.RegistrationReq{
	ID:        &reg5.ID,
	AppID:     &reg5.AppID,
	InviterID: &reg5.InviterID,
	InviteeID: &reg5.InviteeID,
}

func setupSuperior(t *testing.T) func(*testing.T) { //nolint
	_h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg1.AppID, true),
		invitationcode1.WithUserID(_reg1.InviterID, true),
	)
	assert.Nil(t, err)

	_info1, err := _h1.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info1)

	h1, err := NewHandler(
		context.Background(),
		WithID(_reg1.ID, true),
		WithAppID(_reg1.AppID, true),
		WithInviterID(_reg1.InviterID, true),
		WithInviteeID(_reg1.InviteeID, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateRegistration(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info1)
	reg1.CreatedAt = info1.CreatedAt
	reg1.UpdatedAt = info1.UpdatedAt

	_h2, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg2.AppID, true),
		invitationcode1.WithUserID(_reg2.InviterID, true),
	)
	assert.Nil(t, err)

	_info2, err := _h2.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info2)

	h2, err := NewHandler(
		context.Background(),
		WithID(_reg2.ID, true),
		WithAppID(_reg2.AppID, true),
		WithInviterID(_reg2.InviterID, true),
		WithInviteeID(_reg2.InviteeID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateRegistration(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info2)
	reg2.CreatedAt = info2.CreatedAt
	reg2.UpdatedAt = info2.UpdatedAt

	_h3, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg3.AppID, true),
		invitationcode1.WithUserID(_reg3.InviterID, true),
	)
	assert.Nil(t, err)

	_info3, err := _h3.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info3)

	h3, err := NewHandler(
		context.Background(),
		WithID(_reg3.ID, true),
		WithAppID(_reg3.AppID, true),
		WithInviterID(_reg3.InviterID, true),
		WithInviteeID(_reg3.InviteeID, true),
	)
	assert.Nil(t, err)

	info3, err := h3.CreateRegistration(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info3)
	reg3.CreatedAt = info3.CreatedAt
	reg3.UpdatedAt = info3.UpdatedAt

	_h4, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg4.AppID, true),
		invitationcode1.WithUserID(_reg4.InviterID, true),
	)
	assert.Nil(t, err)

	_info4, err := _h4.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info4)

	h4, err := NewHandler(
		context.Background(),
		WithID(_reg4.ID, true),
		WithAppID(_reg4.AppID, true),
		WithInviterID(_reg4.InviterID, true),
		WithInviteeID(_reg4.InviteeID, true),
	)
	assert.Nil(t, err)

	info4, err := h4.CreateRegistration(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info4)
	reg4.CreatedAt = info4.CreatedAt
	reg4.UpdatedAt = info4.UpdatedAt

	_h5, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg5.AppID, true),
		invitationcode1.WithUserID(_reg5.InviterID, true),
	)
	assert.Nil(t, err)

	_info5, err := _h5.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _info5)

	h5, err := NewHandler(
		context.Background(),
		WithID(_reg5.ID, true),
		WithAppID(_reg5.AppID, true),
		WithInviterID(_reg5.InviterID, true),
		WithInviteeID(_reg5.InviteeID, true),
	)
	assert.Nil(t, err)

	info5, err := h5.CreateRegistration(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info5)
	reg5.CreatedAt = info5.CreatedAt
	reg5.UpdatedAt = info5.UpdatedAt

	return func(*testing.T) {
		_, _ = _h1.DeleteInvitationCode(context.Background())
		_, _ = _h2.DeleteInvitationCode(context.Background())
		_, _ = _h3.DeleteInvitationCode(context.Background())
		_, _ = _h4.DeleteInvitationCode(context.Background())
		_, _ = _h5.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteRegistration(context.Background())
		_, _ = h2.DeleteRegistration(context.Background())
		_, _ = h3.DeleteRegistration(context.Background())
		_, _ = h4.DeleteRegistration(context.Background())
		_, _ = h5.DeleteRegistration(context.Background())
	}
}

//nolint
func getSuperiores(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			InviteeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{reg5.InviteeID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetSuperiores(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 5, len(infos))
		assert.Equal(t, uint32(5), total)

		found := false
		for _, info := range infos {
			if info.ID == reg1.ID {
				assert.Equal(t, &reg1, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg2.ID {
				assert.Equal(t, &reg2, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg3.ID {
				assert.Equal(t, &reg3, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg4.ID {
				assert.Equal(t, &reg4, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg5.ID {
				assert.Equal(t, &reg5, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func getSubordinates(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			InviterIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{reg1.InviterID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetSubordinates(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 5, len(infos))
		assert.Equal(t, uint32(5), total)

		found := false
		for _, info := range infos {
			if info.ID == reg1.ID {
				assert.Equal(t, &reg1, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg2.ID {
				assert.Equal(t, &reg2, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg3.ID {
				assert.Equal(t, &reg3, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg4.ID {
				assert.Equal(t, &reg4, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, info := range infos {
			if info.ID == reg5.ID {
				assert.Equal(t, &reg5, info)
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func TestSuperior(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupSuperior(t)
	defer teardown(t)

	t.Run("getSuperiores", getSuperiores)
	t.Run("getSubordinates", getSubordinates)
}
