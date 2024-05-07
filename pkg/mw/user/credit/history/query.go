package history

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	devicecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/user/credit/history"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entusercredithistory "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/usercredithistory"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/credit/history"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UserCreditHistorySelect
	stmCount  *ent.UserCreditHistorySelect
	infos     []*npool.UserCreditHistory
	total     uint32
}

func (h *queryHandler) selectUserCreditHistory(stm *ent.UserCreditHistoryQuery) {
	h.stmSelect = stm.Select(entusercredithistory.FieldID)
}

func (h *queryHandler) queryUserCreditHistory(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.UserCreditHistory.Query().Where(entusercredithistory.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entusercredithistory.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entusercredithistory.EntID(*h.EntID))
	}
	h.selectUserCreditHistory(stm)
	return nil
}

func (h *queryHandler) queryUserCreditHistories(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.UserCreditHistory.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectUserCreditHistory(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entusercredithistory.Table)
	s.LeftJoin(t1).
		On(
			s.C(entusercredithistory.FieldEntID),
			t1.C(entusercredithistory.FieldEntID),
		).
		AppendSelect(
			t1.C(entusercredithistory.FieldEntID),
			t1.C(entusercredithistory.FieldAppID),
			t1.C(entusercredithistory.FieldUserID),
			t1.C(entusercredithistory.FieldTaskID),
			t1.C(entusercredithistory.FieldEventID),
			t1.C(entusercredithistory.FieldCredits),
			t1.C(entusercredithistory.FieldCreatedAt),
			t1.C(entusercredithistory.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Credits)
		if err != nil {
			info.Credits = decimal.NewFromInt(0).String()
		} else {
			info.Credits = amount.String()
		}
	}
}

func (h *Handler) GetUserCreditHistory(ctx context.Context) (*npool.UserCreditHistory, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserCreditHistory(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
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
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetUserCreditHistories(ctx context.Context) ([]*npool.UserCreditHistory, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUserCreditHistories(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
