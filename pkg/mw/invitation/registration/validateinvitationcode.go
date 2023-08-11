package registration

import (
	"context"
	"fmt"

	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	invitationcodemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

func (h *Handler) validateInvitationCode(ctx context.Context) error {
	if h.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if h.InviterID == nil {
		return fmt.Errorf("invalid inviterid")
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
		return err
	}
	exist, err := h1.ExistInvitationCodeConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid inviterid")
	}

	return nil
}
