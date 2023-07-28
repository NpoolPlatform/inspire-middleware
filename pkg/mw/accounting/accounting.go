package accounting

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	archivement1 "github.com/NpoolPlatform/inspire-middleware/pkg/archivement"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	commission2 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/accounting/commission"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

//nolint
func (h *Handler) Accounting(ctx context.Context) ([]*npool.Commission, error) {
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
		_comms, err = handler.Accounting(ctx)
		if err != nil {
			return nil, err
		}
	}

	commMap := map[string]*npool.Commission{}
	for _, comm := range _comms {
		commMap[comm.UserID] = comm
	}

	if h.HasCommission {
		_details, _, err := archivement1.GetDetails(ctx, &detailmgrpb.Conds{
			AppID:   &commonpb.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserIDs: &commonpb.StringSliceVal{Op: cruder.IN, Value: inviterIDs},
			GoodID:  &commonpb.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
			OrderID: &commonpb.StringVal{Op: cruder.EQ, Value: h.OrderID.String()},
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

			return nil, fmt.Errorf("order %v of user %v's commission exist", h.OrderID, inviter)
		}
	}

	currency := h.PaymentCoinUSDCurrency.String()
	amount := h.GoodValue.Div(h.PaymentCoinUSDCurrency).String()
	usdAmount := h.GoodValue.String()

	appID := h.AppID.String()
	goodID := h.GoodID.String()
	orderID := h.OrderID.String()
	paymentID := h.PaymentID.String()
	coinTypeID := h.CoinTypeID.String()
	paymentCoinTypeID := h.PaymentCoinTypeID.String()
	units := h.Units.String()
	userID := h.UserID.String()

	details := []*detailmgrpb.DetailReq{}
	for _, inviter := range inviters {
		commission := decimal.NewFromInt(0).String()
		comm, ok := commMap[inviter.InviterID]
		if ok && h.HasCommission {
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
	comm, ok := commMap[h.UserID.String()]
	if ok && h.HasCommission {
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
