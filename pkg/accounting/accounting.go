package accounting

import (
	"context"

	"github.com/shopspring/decimal"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

//nolint
func Accounting(
	ctx context.Context,
	appID, userID, goodID, orderID string,
	paymentID, coinTypeID, paymentCoinTypeID string,
	paymentCoinUSDCurrency decimal.Decimal,
	units string,
	settleType commmgrpb.SettleType,
	paymentAmount decimal.Decimal,
	goodValue decimal.Decimal,
	hasCommission bool,
) (
	[]*npool.Commission,
	error,
) {
	inviters, inviterIDs, err := registration1.GetInviters(ctx, appID, userID)
	if err != nil {
		return nil, err
	}

	comms, _, err := commission1.GetCommissions(ctx, &commmwpb.Conds{
		AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
		UserIDs:    &commonpb.StringSliceVal{Op: cruder.IN, Value: inviterIDs},
		GoodID:     &commonpb.StringVal{Op: cruder.EQ, Value: goodID},
		SettleType: &commonpb.Int32Val{Op: cruder.EQ, Value: int32(settleType)},
		EndAt:      &commonpb.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
	}, int32(0), int32(len(inviterIDs)))
	if err != nil {
		return nil, err
	}

	_comms, err := commission1.Accounting(
		ctx,
		settleType,
		inviters,
		comms,
		paymentAmount,
		goodValue,
	)
	if err != nil {
		return nil, err
	}

	commMap := map[string]*npool.Commission{}
	for _, comm := range _comms {
		commMap[comm.UserID] = comm
	}

	currency := paymentCoinUSDCurrency.String()
	amount := paymentAmount.String()
	usdAmount := paymentAmount.Mul(paymentCoinUSDCurrency).String()

	details := []*detailmgrpb.DetailReq{}
	for _, inviter := range inviters {
		commission := decimal.NewFromInt(0).String()
		comm, ok := commMap[inviter.InviterID]
		if ok {
			commission = comm.Amount
		}

		selfOrder := false

		details = append(details, &detailmgrpb.DetailReq{
			AppID:                  &appID,
			UserID:                 &inviter.InviterID,
			DirectContributorID:    &inviter.InviteeID,
			GoodID:                 &goodID,
			OrderID:                &orderID,
			SelfOrder:              &selfOrder,
			PaymentID:              &paymentID,
			CoinTypeID:             &coinTypeID,
			PaymentCoinTypeID:      &paymentCoinTypeID,
			PaymentCoinUSDCurrency: &currency,
			Units:                  &units,
			Amount:                 &amount,
			USDAmount:              &usdAmount,
			Commission:             &commission,
		})
	}

	commission := decimal.NewFromInt(0).String()
	selfOrder := true
	comm, ok := commMap[userID]
	if ok && hasCommission {
		commission = comm.Amount
	}

	details = append(details, &detailmgrpb.DetailReq{
		AppID:                  &appID,
		UserID:                 &userID,
		GoodID:                 &goodID,
		OrderID:                &orderID,
		SelfOrder:              &selfOrder,
		PaymentID:              &paymentID,
		CoinTypeID:             &coinTypeID,
		PaymentCoinTypeID:      &paymentCoinTypeID,
		PaymentCoinUSDCurrency: &currency,
		Units:                  &units,
		Amount:                 &amount,
		USDAmount:              &usdAmount,
		Commission:             &commission,
	})

	if len(details) > 0 {
		err = archivement1.BookKeepingV2(ctx, details)
		if err != nil {
			return nil, err
		}
	}

	return _comms, nil
}
