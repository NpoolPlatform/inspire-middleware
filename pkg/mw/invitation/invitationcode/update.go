package invitationcode

import (
	"context"

	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

func (h *Handler) UpdateInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
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
