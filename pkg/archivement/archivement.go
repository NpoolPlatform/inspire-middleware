package archivement

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/gw/v1/archivement"

	"github.com/shopspring/decimal"
)

func GetCoinArchivements(
	ctx context.Context,
	appID, userID string,
	withSub bool,
	start, end uint32,
	offset, limit int32,
) (
	totalAmount, selfAmount decimal.Decimal,
	archivements []*npool.CoinArchivement,
	total uint32,
	erro error,
) {
	return decimal.NewFromInt(1), decimal.NewFromInt(1), nil, 0, nil
}
