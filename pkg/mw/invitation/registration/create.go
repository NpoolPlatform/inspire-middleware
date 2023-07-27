package registration

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	invitationcodemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

func (h *Handler) CreateRegistration(ctx context.Context, in *npool.RegistrationReq) (*npool.Registration, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.InviterID == nil {
		return nil, fmt.Errorf("invalid inviterid")
	}
	if h.InviteeID == nil {
		return nil, fmt.Errorf("invalid inviteeid")
	}

	h1, err := invitationcode1.NewHandler(
		ctx,
		invitationcode1.WithConds(&invitationcodemwpb.Conds{
			AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.InviterID.String()},
			Disabled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		}),
		invitationcode1.WithLimit(0),
	)
	if err != nil {
		return nil, err
	}
	exist, err := h1.ExistInvitationCodeConds(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid inviterid")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRegistration, *h.AppID, *h.InviteeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &registrationcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		InviteeID: &cruder.Cond{Op: cruder.EQ, Val: *h.InviteeID},
	}
	exist, err = h.ExistRegistrationConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := registrationcrud.CreateSet(
			cli.Registration.Create(),
			&registrationcrud.Req{
				ID:        h.ID,
				AppID:     h.AppID,
				InviterID: h.InviterID,
				InviteeID: h.InviteeID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRegistration(ctx)
}
