package accounting

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

//nolint
func (h *Handler) Accounting(ctx context.Context) ([]*npool.Commission, error) {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithAppID(&h.AppID),
		registration1.WithInviteeID(&h.UserID),
	)
	if err != nil {
		return nil, err
	}

	inviters, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, err
	}

	h1, err := commission1.NewHandler(
		ctx,
		commission1.WithConds(&commmwpb.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: inviterIDs},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
			SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
			EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
			StartAt:    &basetypes.Uint32Val{Op: cruder.LT, Value: h.OrderCreatedAt},
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

	_comms := []*npool.Commission{}
	if hasCommission {
		_comms, err = commission1.Accounting(
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

	}

	commMap := map[string]*npool.Commission{}
	for _, comm := range _comms {
		commMap[comm.UserID] = comm
	}

	if hasCommission {
		_details, _, err := archivement1.GetDetails(ctx, &detailmgrpb.Conds{
			AppID:   &commonpb.StringVal{Op: cruder.EQ, Value: appID},
			UserIDs: &commonpb.StringSliceVal{Op: cruder.IN, Value: inviterIDs},
			GoodID:  &commonpb.StringVal{Op: cruder.EQ, Value: goodID},
			OrderID: &commonpb.StringVal{Op: cruder.EQ, Value: orderID},
		}, 0, int32(len(inviterIDs)))
		if err != nil {
			return nil, err
		}

		detailMap := map[string]*detailmgrpb.Detail{}
		for _, detail := range _details {
			detailMap[detail.UserID] = detail
		}

		for _, inviter := range inviterIDs {
			detail, ok := detailMap[inviter]
			if !ok {
				continue
			}

			commission := decimal.RequireFromString(detail.Commission)

			if commission.Cmp(decimal.NewFromInt(0)) == 0 {
				continue
			}

			comm, ok := commMap[inviter]
			if !ok {
				continue
			}

			commission1 := decimal.RequireFromString(comm.Amount)

			if commission1.Cmp(commission) == 0 {
				continue
			}

			return nil, fmt.Errorf("order %v of user %v's commission exist", orderID, inviter)
		}
	}

	currency := paymentCoinUSDCurrency.String()
	amount := goodValue.Div(paymentCoinUSDCurrency).String()
	usdAmount := goodValue.String()

	details := []*detailmgrpb.DetailReq{}
	for _, inviter := range inviters {
		commission := decimal.NewFromInt(0).String()
		comm, ok := commMap[inviter.InviterID]
		if ok && hasCommission {
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
