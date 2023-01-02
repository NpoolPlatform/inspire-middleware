package goodorderpercent

import (
	"context"
	"fmt"

	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	accmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

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

	upperPercent := decimal.NewFromInt(90) //nolint
	for _, inviter := range inviters {
		comm, ok := commMap[inviter.InviteeID]
		if !ok {
			upperPercent = decimal.NewFromInt(0)
			continue
		}

		percent, err := decimal.NewFromString(comm.GetPercent())
		if err != nil {
			return nil, err
		}

		if upperPercent.Cmp(percent) < 0 {
			return nil, fmt.Errorf("invalid commission")
		}
	}

	_comms := []*accmwpb.Commission{}
	lowerPercent := decimal.NewFromInt(0)

	for i := len(inviters) - 1; i >= 0; i-- {
		comm, ok := commMap[inviters[i].InviteeID]
		if !ok {
			break
		}

		percent, err := decimal.NewFromString(comm.GetPercent())
		if err != nil {
			return nil, err
		}

		_comms = append(_comms, &accmwpb.Commission{
			AppID:  inviters[i].AppID,
			UserID: inviters[i].InviteeID,
			Amount: amount.Mul(percent.Sub(lowerPercent)).String(),
		})
	}

	return _comms, nil
}
