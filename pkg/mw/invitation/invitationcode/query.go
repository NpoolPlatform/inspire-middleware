package invitationcode

import (
	"context"
	"fmt"

	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entinvitationcode "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.InvitationCodeSelect
	infos     []*npool.InvitationCode
	total     uint32
}

func (h *queryHandler) queryInvitationCode(cli *ent.Client) {
	h.stmSelect = cli.
		InvitationCode.
		Query().
		Where(
			entinvitationcode.ID(*h.ID),
			entinvitationcode.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) queryInvitationCodes(ctx context.Context, cli *ent.Client) error {
	stm, err := invitationcodecrud.SetQueryConds(cli.InvitationCode.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.stmSelect = stm.Select()
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.InvitationCode{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryInvitationCode(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetInvitationCodes(ctx context.Context) ([]*npool.InvitationCode, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.InvitationCode{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryInvitationCodes(_ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
