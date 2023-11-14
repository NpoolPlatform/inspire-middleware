package invitationcode

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entinvitationcode "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/invitationcode"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.InvitationCodeSelect
	stmSelect *ent.InvitationCodeSelect
	infos     []*npool.InvitationCode
	total     uint32
}

func (h *queryHandler) selectInvitationCode(stm *ent.InvitationCodeQuery) *ent.InvitationCodeSelect {
	return stm.Select(
		entinvitationcode.FieldID,
	)
}

func (h *queryHandler) queryInvitationCode(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.InvitationCode.Query().Where(entinvitationcode.DeletedAt(0))
	if h.ID != nil {
		h.stmSelect.Where(entinvitationcode.ID(*h.ID))
	}
	if h.EntID != nil {
		h.stmSelect.Where(entinvitationcode.EntID(*h.EntID))
	}
	h.selectInvitationCode(stm)
	return nil
}

func (h *queryHandler) queryInvitationCodes(cli *ent.Client) (*ent.InvitationCodeSelect, error) {
	stm, err := invitationcodecrud.SetQueryConds(cli.InvitationCode.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectInvitationCode(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entinvitationcode.Table)
	s.LeftJoin(t).
		On(
			s.C(entinvitationcode.FieldID),
			t.C(entinvitationcode.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entinvitationcode.FieldEntID), "ent_id"),
			sql.As(t.C(entinvitationcode.FieldAppID), "app_id"),
			sql.As(t.C(entinvitationcode.FieldUserID), "user_id"),
			sql.As(t.C(entinvitationcode.FieldInvitationCode), "invitation_code"),
			sql.As(t.C(entinvitationcode.FieldDisabled), "disable"),
			sql.As(t.C(entinvitationcode.FieldCreatedAt), "created_at"),
			sql.As(t.C(entinvitationcode.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.InvitationCode{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryInvitationCode(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
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

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryInvitationCodes(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryInvitationCodes(cli)
		if err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)
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
