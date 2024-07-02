package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	invitationcodemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

func (h *Handler) validateInvitationCode(ctx context.Context) error {
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
		return wlog.WrapError(err)
	}
	exist, err := h1.ExistInvitationCodeConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if !exist {
		return wlog.Errorf("invatationcode not exist")
	}

	return nil
}
