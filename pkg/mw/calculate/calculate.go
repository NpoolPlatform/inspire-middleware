package calculate

import (
	"context"
	"fmt"
	"sort"

	achievementuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	appcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/commission/config"
	appconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/config"
	appgoodcommissionconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/app/good/commission/config"
	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
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

	if len(inviters) == 0 {
		return nil
	}
	h.inviters = inviters
	h.inviterIDs = append(h.inviterIDs, inviters[0].InviterID)
	return nil
}

func sortAppGoodCommissionConfig(byValue []*appgoodcommissionconfigmwpb.AppGoodCommissionConfig) {
	sort.Slice(byValue, func(i, j int) bool {
		decI, errI := decimal.NewFromString(byValue[i].ThresholdAmount)
		if errI != nil {
			return false
		}
		decJ, errJ := decimal.NewFromString(byValue[j].ThresholdAmount)
		if errJ != nil {
			return false
		}

		if byValue[i].Invites != byValue[j].Invites {
			return byValue[i].Invites < byValue[j].Invites
		}
		return decI.LessThan(decJ)
	})
}

func sortAppCommissionConfig(byValue []*appcommissionconfigmwpb.AppCommissionConfig) {
	sort.Slice(byValue, func(i, j int) bool {
		decI, errI := decimal.NewFromString(byValue[i].ThresholdAmount)
		if errI != nil {
			return false
		}
		decJ, errJ := decimal.NewFromString(byValue[j].ThresholdAmount)
		if errJ != nil {
			return false
		}

		if byValue[i].Invites != byValue[j].Invites {
			return byValue[i].Invites < byValue[j].Invites
		}
		return decI.LessThan(decJ)
	})
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
		return nil, err
	}
	achievmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return nil, err
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

	handler := &calculateHandler{
		Handler:    h,
		inviters:   []*registrationmwpb.Registration{},
		inviterIDs: []string{},
	}

	commissionConfigType := types.CommissionConfigType_WithoutCommissionConfig

	if len(appconfigs) == 0 {
		return handler.generateStatements(map[string]*commission2.Commission{}, uuid.Nil.String(), commissionConfigType)
	}
	appconfig := appconfigs[0]

	switch appconfig.CommissionType {
	case types.CommissionType_LegacyCommission:
		commissionConfigType = types.CommissionConfigType_LegacyCommissionConfig
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
		if len(handler.inviters) == 0 {
			return handler.generateStatements(map[string]*commission2.Commission{}, appconfig.EntID, commissionConfigType)
		}
	case types.CommissionType_WithoutCommission:
	default:
		return nil, fmt.Errorf("invalid commissiontype")
	}

	achievementUsers, err := handler.getAchievementUsers(ctx)
	if err != nil {
		return nil, err
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
				commission2.WithAchievementUsers(achievementUsers),
				commission2.WithPaymentCoinUSDCurrency(h.PaymentCoinUSDCurrency.String()),
				commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
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
					Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
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
			if len(goodcomms) > 0 {
				sortAppGoodCommissionConfig(goodcomms)
				handler, err := commission2.NewHandler(
					ctx,
					commission2.WithSettleType(h.SettleType),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appconfig),
					commission2.WithAppGoodCommissionConfigs(goodcomms),
					commission2.WithPaymentAmount(h.PaymentAmount.String()),
					commission2.WithGoodValue(h.GoodValue.String()),
					commission2.WithAchievementUsers(achievementUsers),
					commission2.WithPaymentCoinUSDCurrency(h.PaymentCoinUSDCurrency.String()),
					commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
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
					Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
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
			if len(appcomms) > 0 {
				sortAppCommissionConfig(appcomms)
				handler, err := commission2.NewHandler(
					ctx,
					commission2.WithSettleType(h.SettleType),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appconfig),
					commission2.WithAppCommissionConfigs(appcomms),
					commission2.WithPaymentAmount(h.PaymentAmount.String()),
					commission2.WithGoodValue(h.GoodValue.String()),
					commission2.WithAchievementUsers(achievementUsers),
					commission2.WithPaymentCoinUSDCurrency(h.PaymentCoinUSDCurrency.String()),
					commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
				)
				if err != nil {
					return nil, err
				}
				_comms, err = handler.CalculateByAppCommConfig(ctx)
				if err != nil {
					return nil, err
				}
			}
		case types.CommissionType_WithoutCommission:
		default:
			return nil, fmt.Errorf("invalid commissiontype")
		}
	}

	commMap := map[string]*commission2.Commission{}
	fmt.Println("============ final commission result ============")
	for _, comm := range _comms {
		fmt.Println("====== commission result: ", comm)
		commMap[comm.UserID] = comm
	}

	return handler.generateStatements(commMap, appconfigs[0].EntID, commissionConfigType)
}

func (h *calculateHandler) generateStatements(
	commMap map[string]*commission2.Commission,
	appConfigID string,
	commissionConfigType types.CommissionConfigType,
) ([]*statementmwpb.Statement, error) {
	statements := []*statementmwpb.Statement{}
	for _, inviter := range h.inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}

		commission := decimal.NewFromInt(0).String()
		commissionConfigID := uuid.Nil.String()
		comm, ok := commMap[inviter.InviterID]
		if ok {
			commissionConfigID = comm.CommissionConfigID
			commissionConfigType = comm.CommissionConfigType
		}
		if ok && h.HasCommission {
			commission = comm.Amount
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
			AppConfigID:            appConfigID,
			CommissionConfigID:     commissionConfigID,
			CommissionConfigType:   commissionConfigType,
		})
	}

	commission := decimal.NewFromInt(0).String()
	commissionConfigID := uuid.Nil.String()
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
		AppConfigID:            appConfigID,
		CommissionConfigID:     commissionConfigID,
		CommissionConfigType:   commissionConfigType,
	})

	fmt.Println("=================== final statements ===================")
	for _, item := range statements {
		fmt.Println("statement: ", item)
	}

	return statements, nil
}
