package commission

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Commission struct {
	AppID                   string
	UserID                  string
	DirectContributorUserID *string
	Amount                  string
}

func (h *Handler) Calculate(ctx context.Context) ([]*Commission, error) {
	commMap := map[string]*npool.Commission{}
	for _, comm := range h.Commissions {
		commMap[comm.UserID] = comm
	}

	_comms := []*Commission{}

	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, ok := commMap[inviter.InviteeID]
		if ok {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		comm2, ok := commMap[inviter.InviterID]
		if ok {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v (%v)", inviter.InviterID, percent2, inviter.InviteeID, percent1, comm1.GetGoodID())
		}

		if percent2.Cmp(percent1) == 0 {
			continue
		}

		amount := h.PaymentAmount
		if comm2.SettleMode == types.SettleMode_SettleWithGoodValue {
			amount = h.GoodValue
		}

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return _comms, nil
		}

		_comms = append(_comms, &Commission{
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, ok := commMap[h.Inviters[len(h.Inviters)-1].InviteeID]
	if !ok {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, err
	}

	amount := h.PaymentAmount
	if commLast.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValue
	}

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return _comms, nil
	}

	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		_comms = append(_comms, &Commission{
			AppID:  h.Inviters[len(h.Inviters)-1].AppID,
			UserID: h.Inviters[len(h.Inviters)-1].InviteeID,
			Amount: amount.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	return _comms, nil
}
