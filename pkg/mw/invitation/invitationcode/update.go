package invitationcode

import (
	"context"
	"fmt"

	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

func (h *Handler) UpdateInvitationCode(ctx context.Context, in *npool.InvitationCodeReq) (*npool.InvitationCode, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.UpdateSet(
			cli.InvitationCode.UpdateOneID(*h.ID),
			&invitationcodecrud.Req{
				Disabled: h.Disabled,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetInvitationCode(ctx)
}
