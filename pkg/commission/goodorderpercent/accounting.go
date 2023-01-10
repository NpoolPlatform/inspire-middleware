package goodorderpercent

import (
	"context"
	"fmt"

	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	accmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"

	"github.com/shopspring/decimal"
)

func Accounting(
	ctx context.Context,
	inviters []*regmgrpb.Registration,
	comms []*npool.Commission,
	amount decimal.Decimal,
) (
	[]*accmwpb.Commission,
	error,
) {
	commMap := map[string]*npool.Commission{}
	for _, comm := range comms {
		commMap[comm.UserID] = comm
	}

	_comms := []*accmwpb.Commission{}

	for _, inviter := range inviters {
		if inviter.InviterID == uuid1.InvalidUUIDStr {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, ok := commMap[inviter.InviteeID]
		if ok {
			percent1, err = decimal.NewFromString(comm1.GetPercent())
			if err != nil {
				return nil, err
			}
		}

		comm2, ok := commMap[inviter.InviterID]
		if ok {
			percent2, err = decimal.NewFromString(comm2.GetPercent())
			if err != nil {
				return nil, err
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("invalid commission")
		}

		_comms = append(_comms, &accmwpb.Commission{
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, ok := commMap[inviters[len(inviters)-1].InviteeID]
	if !ok {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetPercent())
	if err != nil {
		return nil, err
	}

	_comms = append(_comms, &accmwpb.Commission{
		AppID:  inviters[len(inviters)-1].AppID,
		UserID: inviters[len(inviters)-1].InviteeID,
		Amount: amount.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
	})

	return _comms, nil
}
