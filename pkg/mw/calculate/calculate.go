package calculate

import (
	"context"
	"fmt"

	appcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	appconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	appgoodcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
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
		return err
	}

	handler.AppID = &h.AppID
	handler.InviteeID = &h.UserID

	inviters, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return err
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
		return err
	}

	inviters, _, err := handler.GetRegistrations(ctx)
	if err != nil {
		return err
	}
	h.inviters = inviters
	h.inviterIDs = []string{h.UserID.String()}
	return nil
}

//nolint
func (h *Handler) Calculate(ctx context.Context) ([]*statementmwpb.Statement, error) {
	h1, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithConds(&appconfigmwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			EndAt:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
			StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
		}),
		appconfig1.WithOffset(0),
		appconfig1.WithLimit(1),
	)
	if err != nil {
		return nil, err
	}

	appconfigs, _, err := h1.GetAppConfigs(ctx)
	if err != nil {
		return nil, err
	}
	if appconfigs == nil {
		return nil, fmt.Errorf("invalid appconfig")
	}
	appconfig := appconfigs[0]

	handler := &calculateHandler{
		Handler:    h,
		inviters:   []*registrationmwpb.Registration{},
		inviterIDs: []string{},
	}

	switch appconfig.CommissionType {
	case types.CommissionType_LegacyCommission:
		fallthrough //nolint
	case types.CommissionType_LayeredCommission:
		err := handler.getLayeredInviters(ctx)
		if err != nil {
			return nil, err
		}
	case types.CommissionType_DirectCommission:
		err := handler.getDirectInviters(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid commissiontype")
	}

	_comms := []*commission2.Commission{}
	if h.HasCommission {
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
					EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
					StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
				}),
				commission1.WithOffset(0),
				commission1.WithLimit(int32(len(handler.inviterIDs))),
			)
			if err != nil {
				return nil, err
			}

			comms, _, err := h2.GetCommissions(ctx)
			if err != nil {
				return nil, err
			}
			handler, err := commission2.NewHandler(
				ctx,
				commission2.WithSettleType(h.SettleType),
				commission2.WithSettleAmountType(h.SettleAmountType),
				commission2.WithInviters(handler.inviters),
				commission2.WithAppConfig(appconfig),
				commission2.WithCommissions(comms),
				commission2.WithPaymentAmount(h.PaymentAmount.String()),
				commission2.WithGoodValue(h.GoodValue.String()),
			)
			if err != nil {
				return nil, err
			}
			_comms, err = handler.Calculate(ctx)
			if err != nil {
				return nil, err
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
				}),
				appgoodcommissionconfig1.WithOffset(0),
				appgoodcommissionconfig1.WithLimit(0),
			)
			if err != nil {
				return nil, err
			}

			goodcomms, _, err := h2.GetCommissionConfigs(ctx)
			if err != nil {
				return nil, err
			}
			if goodcomms != nil {
				handler, err := commission2.NewHandler(
					ctx,
					commission2.WithSettleType(h.SettleType),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appconfig),
					commission2.WithAppGoodCommissionConfigs(goodcomms),
					commission2.WithPaymentAmount(h.PaymentAmount.String()),
					commission2.WithGoodValue(h.GoodValue.String()),
				)
				if err != nil {
					return nil, err
				}
				_comms, err = handler.CalculateByAppGoodCommConfig(ctx)
				if err != nil {
					return nil, err
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
				}),
				appcommissionconfig1.WithOffset(0),
				appcommissionconfig1.WithLimit(0),
			)
			if err != nil {
				return nil, err
			}

			appcomms, _, err := h3.GetCommissionConfigs(ctx)
			if err != nil {
				return nil, err
			}

			if appcomms != nil {
				handler, err := commission2.NewHandler(
					ctx,
					commission2.WithSettleType(h.SettleType),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appconfig),
					commission2.WithAppCommissionConfigs(appcomms),
					commission2.WithPaymentAmount(h.PaymentAmount.String()),
					commission2.WithGoodValue(h.GoodValue.String()),
				)
				if err != nil {
					return nil, err
				}
				_comms, err = handler.CalculateByAppCommConfig(ctx)
				if err != nil {
					return nil, err
				}
			}
		default:
			return nil, fmt.Errorf("invalid commissiontype")
		}
	}

	commMap := map[string]*commission2.Commission{}
	for _, comm := range _comms {
		commMap[comm.UserID] = comm
	}

	statements := []*statementmwpb.Statement{}
	for _, inviter := range handler.inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}

		commission := decimal.NewFromInt(0).String()
		commissionConfigID := uuid.Nil.String()
		commissionConfigType := types.CommissionConfigType_DefaultCommissionConfigType
		comm, ok := commMap[inviter.InviterID]
		if ok && h.HasCommission {
			commission = comm.Amount
			commissionConfigID = comm.CommissionConfigID
			commissionConfigType = comm.CommissionConfigType
		}

		statements = append(statements, &statementmwpb.Statement{
			AppID:                  h.AppID.String(),
			UserID:                 inviter.InviterID,
			DirectContributorID:    inviter.InviteeID,
			GoodID:                 h.GoodID.String(),
			AppGoodID:              h.AppGoodID.String(),
			OrderID:                h.OrderID.String(),
			SelfOrder:              false,
			PaymentID:              h.PaymentID.String(),
			CoinTypeID:             h.CoinTypeID.String(),
			PaymentCoinTypeID:      h.PaymentCoinTypeID.String(),
			PaymentCoinUSDCurrency: h.PaymentCoinUSDCurrency.String(),
			Units:                  h.Units.String(),
			Amount:                 h.GoodValue.String(),
			USDAmount:              h.GoodValueUSD.String(),
			Commission:             commission,
			AppConfigID:            appconfig.EntID,
			CommissionConfigID:     commissionConfigID,
			CommissionConfigType:   commissionConfigType,
		})
	}

	commission := decimal.NewFromInt(0).String()
	commissionConfigID := uuid.Nil.String()
	commissionConfigType := types.CommissionConfigType_DefaultCommissionConfigType
	comm, ok := commMap[h.UserID.String()]
	if ok && h.HasCommission {
		commission = comm.Amount
		commissionConfigID = comm.CommissionConfigID
		commissionConfigType = comm.CommissionConfigType
	}

	statements = append(statements, &statementmwpb.Statement{
		AppID:                  h.AppID.String(),
		UserID:                 h.UserID.String(),
		GoodID:                 h.GoodID.String(),
		AppGoodID:              h.AppGoodID.String(),
		OrderID:                h.OrderID.String(),
		SelfOrder:              true,
		PaymentID:              h.PaymentID.String(),
		CoinTypeID:             h.CoinTypeID.String(),
		PaymentCoinTypeID:      h.PaymentCoinTypeID.String(),
		PaymentCoinUSDCurrency: h.PaymentCoinUSDCurrency.String(),
		Units:                  h.Units.String(),
		Amount:                 h.GoodValue.String(),
		USDAmount:              h.GoodValueUSD.String(),
		Commission:             commission,
		AppConfigID:            appconfig.EntID,
		CommissionConfigID:     commissionConfigID,
		CommissionConfigType:   commissionConfigType,
	})

	return statements, nil
}
