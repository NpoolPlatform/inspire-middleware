package accounting

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"
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
) (
	[]*npool.Commission,
	error,
) {
	offset := int32(0)
	limit := int32(100)

	inviters := []*regmgrpb.Registration{}
	for {
		_inviters, _, err := registration1.GetSuperiores(ctx, &regmgrpb.Conds{
			AppID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			InviteeIDs: &commonpb.StringSliceVal{
				Op:    cruder.IN,
				Value: []string{userID},
			},
		}, offset, limit)
		if err != nil {
			return nil, err
		}
		if len(_inviters) == 0 {
			break
		}

		inviters = append(inviters, _inviters...)

		offset += limit
	}

	inviteeMap := map[string]struct{}{}
	for _, inviter := range inviters {
		inviteeMap[inviter.InviteeID] = struct{}{}
	}

	inviterCount := len(inviters)
	_inviters := []*regmgrpb.Registration{}

	for i, inviter := range inviters {
		_, ok := inviteeMap[inviter.InviterID]
		if !ok {
			_inviters = append(_inviters, inviter)
			inviters = append(inviters[0:i], inviters[i+1:]...)
			break
		}
	}

	if inviterCount == 0 {
		_inviters = append(_inviters, &regmgrpb.Registration{
			AppID:     appID,
			InviterID: uuid1.InvalidUUIDStr,
			InviteeID: userID,
		})
	}

	if len(_inviters) == 0 {
		return nil, fmt.Errorf("invalid top inviter")
	}

	for {
		if inviterCount == 0 || len(inviters) == 0 {
			break
		}

		if len(inviters) == 1 {
			if _inviters[len(_inviters)-1].InviteeID != inviters[0].InviterID {
				return nil, fmt.Errorf("mismatch registration")
			}
			_inviters = append(_inviters, inviters[0])
			break
		}

		for i, inviter := range inviters {
			if _inviters[len(_inviters)-1].InviteeID == inviter.InviterID {
				_inviters = append(_inviters, inviter)
				inviters = append(inviters[0:i], inviters[i+1:]...)
				break
			}
		}
	}

	inviterIDs := []string{userID}
	if inviterCount > 0 {
		inviterIDs = []string{_inviters[0].InviterID}
		for _, inviter := range _inviters {
			inviterIDs = append(inviterIDs, inviter.InviteeID)
		}
	}

	comms, _, err := commission1.GetCommissions(ctx, &commmwpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: inviterIDs,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: goodID,
		},
		SettleType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(settleType),
		},
		EndAt: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(0),
		},
	}, int32(0), int32(len(inviterIDs)))
	if err != nil {
		return nil, err
	}

	_comms, err := commission1.Accounting(
		ctx,
		settleType,
		_inviters,
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
	for _, inviter := range _inviters {
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
	if ok {
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
