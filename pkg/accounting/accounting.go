package accounting

import (
	"context"

	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
)

func Accounting(
	ctx context.Context,
	appID, userID, goodID, orderID string,
	paymentAmount decimal.Decimal,
	goodValue decimal.Decimal,
) (
	[]*npool.Commission,
	error,
) {
	return nil, nil
}
