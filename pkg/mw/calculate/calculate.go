package calculate

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	achievementuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	appcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	appconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	appgoodcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
	paymentmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"
	achievementusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
	appcommissionconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
	appconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
	appgoodcommissionconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type calculateHandler struct {
	*Handler
	inviters   []*registrationmwpb.Registration
	inviterIDs []string
}

func (h *calculateHandler) getLayeredInviters(ctx context.Context) error {
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

func (h *calculateHandler) getDirectInviters(ctx context.Context) error {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithConds(&registrationmwpb.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			InviteeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
		}),
		registration1.WithOffset(0),
		registration1.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	inviters, _, err := handler.GetRegistrations(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(inviters) == 0 {
		return nil
	}
	h.inviters = inviters
	h.inviterIDs = append(h.inviterIDs, inviters[0].InviterID)
	return nil
}

func (h *calculateHandler) getAchievementUsers(ctx context.Context) (map[string]*achievementusermwpb.AchievementUser, error) {
	achievementUserMap := map[string]*achievementusermwpb.AchievementUser{}
	handler, err := achievementuser1.NewHandler(
		ctx,
		achievementuser1.WithConds(&achievementusermwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.inviterIDs},
		}),
		achievementuser1.WithOffset(0),
		achievementuser1.WithLimit(int32(len(h.inviterIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	achievmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(achievmentUsers) == 0 {
		return achievementUserMap, nil
	}
	for _, id := range h.inviterIDs {
		for _, achievementUser := range achievmentUsers {
			if achievementUser.UserID == id {
				achievementUserMap[id] = achievementUser
			}
		}
	}
	return achievementUserMap, nil
}

//nolint
func (h *Handler) Calculate(ctx context.Context) ([]*statementmwpb.StatementReq, error) {
	h1, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithConds(&appconfigmwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
			EndAt:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
		}),
		appconfig1.WithOffset(0),
		appconfig1.WithLimit(1),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appconfigs, _, err := h1.GetAppConfigs(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &calculateHandler{
		Handler:    h,
		inviters:   []*registrationmwpb.Registration{},
		inviterIDs: []string{},
	}

	commissionConfigType := types.CommissionConfigType_WithoutCommissionConfig

	if len(appconfigs) == 0 {
		return handler.generateStatements(map[string]map[string]*commission2.Commission{}, uuid.Nil.String(), commissionConfigType)
	}
	appconfig := appconfigs[0]

	switch appconfig.CommissionType {
	case types.CommissionType_LegacyCommission:
		commissionConfigType = types.CommissionConfigType_LegacyCommissionConfig
		fallthrough //nolint
	case types.CommissionType_LayeredCommission:
		err := handler.getLayeredInviters(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
	case types.CommissionType_DirectCommission:
		err := handler.getDirectInviters(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if len(handler.inviters) == 0 {
			return handler.generateStatements(map[string]map[string]*commission2.Commission{}, appconfig.EntID, commissionConfigType)
		}
	case types.CommissionType_WithoutCommission:
	default:
		return nil, wlog.Errorf("invalid commissiontype")
	}

	achievementUsers, err := handler.getAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	commMap := map[string]map[string]*commission2.Commission{} // userid->cointypeid->commission
	if h.HasCommission {
		for _, payment := range h.Payments {
			_comms := []*commission2.Commission{}

			switch appconfig.CommissionType {
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
					commission2.WithSettleType(h.SettleType),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appconfig),
					commission2.WithCommissions(comms),
					commission2.WithPaymentAmount(payment.Amount),
					commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
					commission2.WithAchievementUsers(achievementUsers),
					commission2.WithPaymentCoinUSDCurrency(payment.CoinUSDCurrency),
					commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				_comms, err = handler.Calculate(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
			case types.CommissionType_LayeredCommission:
				fallthrough //nolint
			case types.CommissionType_DirectCommission:
				h2, err := appgoodcommissionconfig1.NewHandler(
					ctx,
					appgoodcommissionconfig1.WithConds(&appgoodcommissionconfigmwpb.Conds{
						AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
						GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
						AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppGoodID.String()},
						SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
						EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
						StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
						Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
					}),
					appgoodcommissionconfig1.WithOffset(0),
					appgoodcommissionconfig1.WithLimit(0),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}

				goodcomms, _, err := h2.GetCommissionConfigs(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				if len(goodcomms) > 0 {
					handler, err := commission2.NewHandler(
						ctx,
						commission2.WithSettleType(h.SettleType),
						commission2.WithSettleAmountType(h.SettleAmountType),
						commission2.WithInviters(handler.inviters),
						commission2.WithAppConfig(appconfig),
						commission2.WithAppGoodCommissionConfigs(goodcomms),
						commission2.WithPaymentAmount(payment.Amount),
						commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
						commission2.WithAchievementUsers(achievementUsers),
						commission2.WithPaymentCoinUSDCurrency(payment.CoinUSDCurrency),
						commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
					)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					_comms, err = handler.CalculateByAppGoodCommConfig(ctx)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					break
				}

				h3, err := appcommissionconfig1.NewHandler(
					ctx,
					appcommissionconfig1.WithConds(&appcommissionconfigmwpb.Conds{
						AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
						SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
						EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
						StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
						Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
					}),
					appcommissionconfig1.WithOffset(0),
					appcommissionconfig1.WithLimit(0),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}

				appcomms, _, err := h3.GetCommissionConfigs(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				if len(appcomms) > 0 {
					handler, err := commission2.NewHandler(
						ctx,
						commission2.WithSettleType(h.SettleType),
						commission2.WithSettleAmountType(h.SettleAmountType),
						commission2.WithInviters(handler.inviters),
						commission2.WithAppConfig(appconfig),
						commission2.WithAppCommissionConfigs(appcomms),
						commission2.WithPaymentAmount(payment.Amount),
						commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
						commission2.WithAchievementUsers(achievementUsers),
						commission2.WithPaymentCoinUSDCurrency(payment.CoinUSDCurrency),
						commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
					)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					_comms, err = handler.CalculateByAppCommConfig(ctx)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
				}
			case types.CommissionType_WithoutCommission:
			default:
				return nil, wlog.Errorf("invalid commissiontype")
			}

			for _, _com := range _comms {
				com, ok := commMap[_com.UserID][payment.CoinTypeID]
				if !ok {
					commMap[com.UserID][payment.CoinTypeID] = com
				} else {
					oldAmount, err := decimal.NewFromString(com.Amount)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					amount, err := decimal.NewFromString(_com.Amount)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					com.Amount = oldAmount.Add(amount).String()

					oldCommAmountUSD, err := decimal.NewFromString(com.CommissionAmountUSD)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					commAmountUSD, err := decimal.NewFromString(_com.CommissionAmountUSD)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					com.CommissionAmountUSD = oldCommAmountUSD.Add(commAmountUSD).String()

					commMap[_com.UserID][payment.CoinTypeID] = com
				}
			}

		}
	}

	return handler.generateStatements(commMap, appconfigs[0].EntID, commissionConfigType)
}

//nolint
func (h *calculateHandler) generateStatements(
	userCoinCommMap map[string]map[string]*commission2.Commission,
	appConfigID string,
	commissionConfigType types.CommissionConfigType,
) ([]*statementmwpb.StatementReq, error) {
	statements := []*statementmwpb.StatementReq{}
	for _, inviter := range h.inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}

		commissionAmountUSD := decimal.NewFromInt(0).String()
		commissionConfigID := uuid.Nil.String()

		payments := []*paymentmwpb.StatementReq{}

		userCommMap, ok := userCoinCommMap[inviter.InviterID]
		if ok {
			for key, value := range userCommMap {
				payments = append(payments, &paymentmwpb.StatementReq{
					PaymentCoinTypeID: &key,
					Amount:            &value.PaymentAmount,
					CommissionAmount:  &value.Amount,
				})
				commissionConfigID = value.CommissionConfigID
				commissionConfigType = value.CommissionConfigType
				if h.HasCommission {
					commissionAmountUSD = value.CommissionAmountUSD
				}
			}
		}

		statements = append(statements, &statementmwpb.StatementReq{
			AppID: func() *string {
				id := h.AppID.String()
				return &id
			}(),
			UserID:      &inviter.InviterID,
			OrderUserID: &inviter.InviteeID,
			GoodID: func() *string {
				id := h.GoodID.String()
				return &id
			}(),
			AppGoodID: func() *string {
				id := h.AppGoodID.String()
				return &id
			}(),
			OrderID: func() *string {
				id := h.OrderID.String()
				return &id
			}(),

			GoodCoinTypeID: func() *string {
				id := h.GoodCoinTypeID.String()
				return &id
			}(),
			Units: func() *string {
				units := h.Units.String()
				return &units
			}(),
			GoodValueUSD: func() *string {
				goodValueUSD := h.GoodValueUSD.String()
				return &goodValueUSD
			}(),
			PaymentAmountUSD: func() *string {
				paymentAmountUSD := h.PaymentAmountUSD.String()
				return &paymentAmountUSD
			}(),
			CommissionAmountUSD:  &commissionAmountUSD,
			AppConfigID:          &appConfigID,
			CommissionConfigID:   &commissionConfigID,
			CommissionConfigType: &commissionConfigType,
			PaymentStatements:    payments,
		})
	}

	commissionAmountUSD := decimal.NewFromInt(0).String()
	commissionConfigID := uuid.Nil.String()
	payments := []*paymentmwpb.StatementReq{}

	userCommMap, ok := userCoinCommMap[h.UserID.String()]
	if ok {
		for key, value := range userCommMap {
			payments = append(payments, &paymentmwpb.StatementReq{
				PaymentCoinTypeID: &key,
				Amount:            &value.PaymentAmount,
				CommissionAmount:  &value.Amount,
			})
			commissionConfigID = value.CommissionConfigID
			commissionConfigType = value.CommissionConfigType
			if h.HasCommission {
				commissionAmountUSD = value.CommissionAmountUSD
			}
		}
	}

	statements = append(statements, &statementmwpb.StatementReq{
		AppID: func() *string {
			id := h.AppID.String()
			return &id
		}(),
		UserID: func() *string {
			id := h.UserID.String()
			return &id
		}(),
		OrderUserID: func() *string {
			id := h.UserID.String()
			return &id
		}(),
		GoodID: func() *string {
			id := h.GoodID.String()
			return &id
		}(),
		AppGoodID: func() *string {
			id := h.AppGoodID.String()
			return &id
		}(),
		OrderID: func() *string {
			id := h.OrderID.String()
			return &id
		}(),

		Units: func() *string {
			units := h.Units.String()
			return &units
		}(),
		GoodValueUSD: func() *string {
			goodValueUSD := h.GoodValueUSD.String()
			return &goodValueUSD
		}(),
		PaymentAmountUSD: func() *string {
			paymentAmountUSD := h.PaymentAmountUSD.String()
			return &paymentAmountUSD
		}(),
		CommissionAmountUSD:  &commissionAmountUSD,
		AppConfigID:          &appConfigID,
		CommissionConfigID:   &commissionConfigID,
		CommissionConfigType: &commissionConfigType,
		PaymentStatements:    payments,
	})

	return statements, nil
}
