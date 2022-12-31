package accounting

import (
	"context"

	"github.com/shopspring/decimal"

	commonmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
)

func Accounting(
	ctx context.Context,
	appID, userID, goodID, orderID string,
	settleType commissionmgrpb.SettleType,
	paymentAmount decimal.Decimal,
	goodValue decimal.Decimal,
) (
	[]*npool.Commission,
	error,
) {
	// TODO: query all inviters
	// TODO: query user settle type commission
	// TODO: calculate commission
	return nil, nil
}
