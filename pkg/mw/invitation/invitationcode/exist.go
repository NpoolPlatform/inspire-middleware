package invitationcode

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entinvitationcode "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/invitationcode"
)

func (h *Handler) ExistInvitationCode(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_exist, err := cli.
			InvitationCode.
			Query().
			Where(
				entinvitationcode.ID(*h.ID),
				entinvitationcode.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}

func (h *Handler) ExistInvitationCodeConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := invitationcodecrud.SetQueryConds(cli.InvitationCode.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}
