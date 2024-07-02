package invitationcode

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

func (h *Handler) DeleteInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	info, err := h.GetInvitationCode(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, nil
	}
	if h.ID == nil {
		h.ID = &info.ID
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.UpdateSet(
			cli.InvitationCode.UpdateOneID(*h.ID),
			&invitationcodecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
