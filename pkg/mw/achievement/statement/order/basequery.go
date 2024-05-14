package orderstatement

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderStatementSelect
}

func (h *baseQueryHandler) selectOrderStatement(stm *ent.OrderStatementQuery) *ent.OrderStatementSelect {
	return stm.Select(entorderstatement.FieldID)
}

func (h *baseQueryHandler) queryOrderStatement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OrderStatement.Query().Where(entorderstatement.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderstatement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderstatement.EntID(*h.EntID))
	}
	h.stmSelect = h.selectOrderStatement(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderStatements(cli *ent.Client) (*ent.OrderStatementSelect, error) {
	stm, err := orderstatementcrud.SetQueryConds(cli.OrderStatement.Query(), h.OrderStatementConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderStatement(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderstatement.Table)
	s.AppendSelect(
		t.C(entorderstatement.FieldID),
		t.C(entorderstatement.FieldEntID),
		t.C(entorderstatement.FieldAppID),
		t.C(entorderstatement.FieldUserID),
		t.C(entorderstatement.FieldGoodID),
		t.C(entorderstatement.FieldAppGoodID),
		t.C(entorderstatement.FieldOrderID),
		t.C(entorderstatement.FieldOrderUserID),
		t.C(entorderstatement.FieldGoodCoinTypeID),
		t.C(entorderstatement.FieldUnits),
		t.C(entorderstatement.FieldGoodValueUsd),
		t.C(entorderstatement.FieldPaymentAmountUsd),
		t.C(entorderstatement.FieldCommissionAmountUsd),
		t.C(entorderstatement.FieldAppConfigID),
		t.C(entorderstatement.FieldCommissionConfigID),
		t.C(entorderstatement.FieldCommissionConfigType),
		t.C(entorderstatement.FieldCreatedAt),
		t.C(entorderstatement.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
