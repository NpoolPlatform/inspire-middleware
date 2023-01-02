package commission

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	accmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	goodorderpercent "github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"

	"github.com/shopspring/decimal"
)

func Accounting(
	ctx context.Context,
	settleType mgrpb.SettleType,
	inviters []*regmgrpb.Registration,
	comms []*npool.Commission,
	paymentAmount decimal.Decimal,
	goodValue decimal.Decimal,
) (
	[]*accmwpb.Commission,
	error,
) {
	switch settleType {
	case mgrpb.SettleType_GoodOrderPercent:
		return goodorderpercent.Accounting(ctx, inviters, comms, paymentAmount)
	case mgrpb.SettleType_LimitedOrderPercent:
		fallthrough //nolint
	case mgrpb.SettleType_AmountThreshold:
	default:
		return nil, fmt.Errorf("invalid settle type")
	}
	return nil, fmt.Errorf("not implemented")
}
