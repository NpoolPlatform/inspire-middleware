//nolint:dupl
package calculate

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entorderpaymentstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderstatement"
	achievementuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	common1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user/common"
	appConfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"
	achievementusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
	appconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type reconcileCalculateHandler struct {
	*Handler
	inviters           []*registrationmwpb.Registration
	inviterIDs         []string
	appConfig          *appconfigmwpb.AppConfig
	statements         []*ent.OrderStatement
	payments           map[uuid.UUID][]*ent.OrderPaymentStatement
	orderUserStatement *ent.OrderStatement
	infos              []*statementmwpb.StatementReq
}

func (h *reconcileCalculateHandler) getLayeredInviters(ctx context.Context) error {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	handler.AppID = &h.AppID
	handler.InviteeID = &h.UserID

	inviters, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.inviters = inviters
	h.inviterIDs = inviterIDs
	return nil
}

func (h *reconcileCalculateHandler) getAchievementUsers(ctx context.Context) (map[string]*achievementusermwpb.AchievementUser, error) {
	achievementUserMap := map[string]*achievementusermwpb.AchievementUser{}
	handler, err := achievementuser1.NewHandler(
		ctx,
		common1.WithConds(&achievementusermwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.inviterIDs},
		}),
		common1.WithOffset(0),
		common1.WithLimit(int32(len(h.inviterIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	achievementUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(achievementUsers) == 0 {
		return achievementUserMap, nil
	}
	for _, id := range h.inviterIDs {
		for _, achievementUser := range achievementUsers {
			if achievementUser.UserID == id {
				achievementUserMap[id] = achievementUser
			}
		}
	}
	return achievementUserMap, nil
}

func (h *reconcileCalculateHandler) requireOrderStatement(ctx context.Context, cli *ent.Client) error {
	statement, err := cli.
		OrderStatement.
		Query().
		Where(
			entorderstatement.AppID(h.AppID),
			entorderstatement.OrderID(h.OrderID),
			entorderstatement.UserID(h.UserID),
			entorderstatement.OrderUserID(h.UserID),
			entorderstatement.CommissionConfigType(types.CommissionConfigType_LegacyCommissionConfig.String()),
		).
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.GoodID = statement.GoodID
	h.AppGoodID = statement.AppGoodID
	h.GoodCoinTypeID = statement.GoodCoinTypeID
	h.Units = statement.Units
	h.GoodValueUSD = statement.GoodValueUsd
	h.PaymentAmountUSD = statement.PaymentAmountUsd
	h.orderUserStatement = statement
	return nil
}

func (h *reconcileCalculateHandler) getOrderStatements(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := h.requireOrderStatement(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}

		statements, err := cli.
			OrderStatement.
			Query().
			Where(
				entorderstatement.AppID(h.AppID),
				entorderstatement.OrderID(h.OrderID),
			).
			All(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(statements) == 0 {
			return wlog.Errorf("invalid statement")
		}
		h.statements = statements

		if err := h.getOrderPaymentStatements(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *reconcileCalculateHandler) getOrderPaymentStatements(ctx context.Context, cli *ent.Client) error {
	ids := []uuid.UUID{}
	for _, statement := range h.statements {
		ids = append(ids, statement.EntID)
	}
	payments, err := cli.
		OrderPaymentStatement.
		Query().
		Where(
			entorderpaymentstatement.StatementIDIn(ids...),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, payment := range payments {
		h.payments[payment.StatementID] = append(h.payments[payment.StatementID], payment)
	}
	return nil
}

func (h *reconcileCalculateHandler) getAppConfig(ctx context.Context) error {
	h1, err := appConfig1.NewHandler(
		ctx,
		appConfig1.WithConds(&appconfigmwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
			EndAt:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
		}),
		appConfig1.WithOffset(0),
		appConfig1.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	appConfigs, _, err := h1.GetAppConfigs(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(appConfigs) == 0 {
		return wlog.Errorf("invalid app config")
	}

	h.appConfig = appConfigs[0]
	return nil
}

func (h *Handler) ReconcileCalculate(ctx context.Context) ([]*statementmwpb.StatementReq, error) {
	handler := &reconcileCalculateHandler{
		Handler:    h,
		inviters:   []*registrationmwpb.Registration{},
		inviterIDs: []string{},
		statements: []*ent.OrderStatement{},
		payments:   map[uuid.UUID][]*ent.OrderPaymentStatement{},
		infos:      []*statementmwpb.StatementReq{},
	}
	if err := handler.getOrderStatements(ctx); err != nil {
		return nil, err
	}
	if err := handler.getAppConfig(ctx); err != nil {
		return nil, err
	}

	switch handler.appConfig.CommissionType {
	case types.CommissionType_LegacyCommission:
		err := handler.getLayeredInviters(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
	default:
		return nil, wlog.Errorf("invalid commission type")
	}

	achievementUsers, err := handler.getAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	payments, ok := handler.payments[handler.orderUserStatement.EntID]
	if !ok {
		return nil, wlog.Errorf("invalid payment")
	}
	commissions := map[uuid.UUID][]*commission2.Commission{} // cointypeid->[]Commission
	for _, payment := range payments {
		switch handler.appConfig.CommissionType {
		case types.CommissionType_LegacyCommission:
			h2, err := commission1.NewHandler(
				ctx,
				commission1.WithConds(&commmwpb.Conds{
					AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
					UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: handler.inviterIDs},
					GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
					AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppGoodID.String()},
					SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
					StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
					EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
				}),
				commission1.WithOffset(0),
				commission1.WithLimit(int32(len(handler.inviterIDs))),
			)
			if err != nil {
				return nil, wlog.WrapError(err)
			}

			comms, _, err := h2.GetCommissions(ctx)
			if err != nil {
				return nil, wlog.WrapError(err)
			}
			handler, err := commission2.NewHandler(
				ctx,
				commission2.WithSettleType(types.SettleType_GoodOrderPayment),
				commission2.WithSettleAmountType(h.SettleAmountType),
				commission2.WithInviters(handler.inviters),
				commission2.WithAppConfig(handler.appConfig),
				commission2.WithCommissions(comms),
				commission2.WithPaymentAmount(payment.Amount.String()),
				commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
				commission2.WithAchievementUsers(achievementUsers),
				commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
			)
			if err != nil {
				return nil, wlog.WrapError(err)
			}
			_comms, err := handler.Calculate(ctx)
			if err != nil {
				return nil, wlog.WrapError(err)
			}
			_commissions, ok := commissions[payment.PaymentCoinTypeID]
			if !ok {
				_commissions = []*commission2.Commission{}
			}
			_commissions = append(_commissions, _comms...)
			commissions[payment.PaymentCoinTypeID] = _commissions
		default:
			return nil, wlog.Errorf("invalid commission type")
		}
	}

	handler.formalizeStatements(commissions)
	return handler.infos, nil
}

func (h *reconcileCalculateHandler) formalizeStatements(commissions map[uuid.UUID][]*commission2.Commission) {
	_commissions := map[string]*commission2.Commission{}
	for key, comms := range commissions {
		for _, comm := range comms {
			key := fmt.Sprintf("%v-%v-%v", key, comm.UserID, comm.PaymentAmount) // cointypeid-userid-paymentamount
			_commissions[key] = comm
		}
	}

	for _, statement := range h.statements {
		dbPayments, ok := h.payments[statement.EntID]
		if !ok {
			continue
		}
		req := statementmwpb.StatementReq{
			ID:               &statement.ID,
			EntID:            func() *string { s := statement.EntID.String(); return &s }(),
			AppID:            func() *string { s := statement.AppID.String(); return &s }(),
			UserID:           func() *string { s := statement.UserID.String(); return &s }(),
			GoodID:           func() *string { s := statement.GoodID.String(); return &s }(),
			AppGoodID:        func() *string { s := statement.AppGoodID.String(); return &s }(),
			OrderID:          func() *string { s := statement.OrderID.String(); return &s }(),
			OrderUserID:      func() *string { s := statement.OrderUserID.String(); return &s }(),
			GoodCoinTypeID:   func() *string { s := statement.GoodCoinTypeID.String(); return &s }(),
			Units:            func() *string { s := statement.Units.String(); return &s }(),
			GoodValueUSD:     func() *string { s := statement.GoodValueUsd.String(); return &s }(),
			PaymentAmountUSD: func() *string { s := h.PaymentAmountUSD.String(); return &s }(),
			AppConfigID:      &h.appConfig.EntID,
			CommissionConfigType: func() *types.CommissionConfigType {
				s := types.CommissionConfigType(types.CommissionConfigType_value[statement.CommissionConfigType])
				return &s
			}(),
		}
		for _, dbPayment := range dbPayments {
			key := fmt.Sprintf("%v-%v-%v", dbPayment.PaymentCoinTypeID, statement.UserID, dbPayment.Amount)
			comm, ok := _commissions[key]
			if !ok {
				if statement.UserID == h.UserID {
					req.CommissionAmountUSD = func() *string { amount := "0"; return &amount }()
					req.PaymentStatements = append(req.PaymentStatements, &payment.StatementReq{
						Amount:            func() *string { amount := dbPayment.Amount.String(); return &amount }(),
						CommissionAmount:  func() *string { amount := "0"; return &amount }(),
						PaymentCoinTypeID: func() *string { id := dbPayment.PaymentCoinTypeID.String(); return &id }(),
					})
				}
				continue
			}
			req.CommissionAmountUSD = &comm.CommissionAmountUSD
			req.PaymentStatements = append(req.PaymentStatements, &payment.StatementReq{
				Amount:            func() *string { amount := dbPayment.Amount.String(); return &amount }(),
				CommissionAmount:  &comm.Amount,
				PaymentCoinTypeID: func() *string { id := dbPayment.PaymentCoinTypeID.String(); return &id }(),
			})
		}
		h.infos = append(h.infos, &req)
	}
}
