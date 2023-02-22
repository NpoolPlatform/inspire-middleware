package goodordervaluepercent

import (
	"context"

	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	accmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	percent1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission/percent"
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
	return percent1.Accounting(ctx, inviters, comms, amount)
}
