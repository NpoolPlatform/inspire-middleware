package calculate

import (
	"context"

	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/calculate/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

//nolint
func (h *Handler) Calculate(ctx context.Context) ([]*statementmwpb.Statement, error) {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return nil, err
	}

	handler.AppID = &h.AppID
	handler.InviteeID = &h.UserID

	inviters, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, err
	}

	h1, err := commission1.NewHandler(
		ctx,
		commission1.WithConds(&commmwpb.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: inviterIDs},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
			SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
			EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
			StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
		}),
		commission1.WithOffset(0),
		commission1.WithLimit(int32(len(inviterIDs))),
	)
	if err != nil {
		return nil, err
	}

	comms, _, err := h1.GetCommissions(ctx)
	if err != nil {
		return nil, err
	}

	_comms := []*commission2.Commission{}
	if h.HasCommission {
		handler, err := commission2.NewHandler(
			ctx,
			commission2.WithSettleType(h.SettleType),
			commission2.WithSettleMode(h.SettleMode),
			commission2.WithInviters(inviters),
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
	}

	commMap := map[string]*commission2.Commission{}
	for _, comm := range _comms {
		commMap[comm.UserID] = comm
	}

	statements := []*statementmwpb.Statement{}
	for _, inviter := range inviters {
		commission := decimal.NewFromInt(0).String()
		comm, ok := commMap[inviter.InviterID]
		if ok && h.HasCommission {
			commission = comm.Amount
		}

		statements = append(statements, &statementmwpb.Statement{
			AppID:                  h.AppID.String(),
			UserID:                 inviter.InviterID,
			DirectContributorID:    inviter.InviteeID,
			GoodID:                 h.GoodID.String(),
			OrderID:                h.OrderID.String(),
			SelfOrder:              false,
			PaymentID:              h.PaymentID.String(),
			CoinTypeID:             h.CoinTypeID.String(),
			PaymentCoinTypeID:      h.PaymentCoinTypeID.String(),
			PaymentCoinUSDCurrency: h.PaymentCoinUSDCurrency.String(),
			Units:                  h.Units.String(),
			Amount:                 h.GoodValue.Div(h.PaymentCoinUSDCurrency).String(),
			USDAmount:              h.GoodValue.String(),
			Commission:             commission,
		})
	}

	commission := decimal.NewFromInt(0).String()
	comm, ok := commMap[h.UserID.String()]
	if ok && h.HasCommission {
		commission = comm.Amount
	}

	statements = append(statements, &statementmwpb.Statement{
		AppID:                  h.AppID.String(),
		UserID:                 h.UserID.String(),
		GoodID:                 h.GoodID.String(),
		OrderID:                h.OrderID.String(),
		SelfOrder:              true,
		PaymentID:              h.PaymentID.String(),
		CoinTypeID:             h.CoinTypeID.String(),
		PaymentCoinTypeID:      h.PaymentCoinTypeID.String(),
		PaymentCoinUSDCurrency: h.PaymentCoinUSDCurrency.String(),
		Units:                  h.Units.String(),
		Amount:                 h.GoodValue.Div(h.PaymentCoinUSDCurrency).String(),
		USDAmount:              h.GoodValue.String(),
		Commission:             commission,
	})

	return statements, nil
}
