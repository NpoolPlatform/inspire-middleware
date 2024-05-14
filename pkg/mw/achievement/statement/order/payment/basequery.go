package orderpaymentstatement

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderpaymentstatementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/statement/order/payment"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entorderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderPaymentStatementSelect
}

func (h *baseQueryHandler) selectOrderPaymentStatement(stm *ent.OrderPaymentStatementQuery) *ent.OrderPaymentStatementSelect {
	return stm.Select(entorderstatement.FieldID)
}

func (h *baseQueryHandler) queryOrderPaymentStatements(cli *ent.Client) (*ent.OrderPaymentStatementSelect, error) {
	stm, err := orderpaymentstatementcrud.SetQueryConds(cli.OrderPaymentStatement.Query(), h.PaymentStatementConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderPaymentStatement(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderpaymentstatement.Table)
	s.Join(t).
		On(
			s.C(entorderpaymentstatement.FieldID),
			t.C(entorderpaymentstatement.FieldID),
		).
		AppendSelect(
			t.C(entorderpaymentstatement.FieldEntID),
			t.C(entorderpaymentstatement.FieldStatementID),
			t.C(entorderpaymentstatement.FieldPaymentCoinTypeID),
			t.C(entorderpaymentstatement.FieldAmount),
			t.C(entorderpaymentstatement.FieldCommissionAmount),
			t.C(entorderpaymentstatement.FieldCreatedAt),
			t.C(entorderpaymentstatement.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinOrderStatement(s *sql.Selector) error {
	t := sql.Table(entorderstatement.Table)
	s.Join(t).
		On(
			s.C(entorderpaymentstatement.FieldStatementID),
			t.C(entorderstatement.FieldEntID),
		)
	if h.OrderStatementConds.AppID != nil {
		id, ok := h.OrderStatementConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatement.FieldAppID), id),
		)
	}
	if h.OrderStatementConds.UserID != nil {
		id, ok := h.OrderStatementConds.UserID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid userid")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatement.FieldUserID), id),
		)
	}
	s.AppendSelect(
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
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOrderStatement(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStatement", "Error", err)
		}
	})
}
