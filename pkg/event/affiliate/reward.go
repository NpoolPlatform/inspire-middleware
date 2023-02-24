package affiliate

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/shopspring/decimal"
)

func RewardEvent(
	ctx context.Context,
	appID, userID string,
	eventType basetypes.UsedFor,
	goodID *string,
	consecutive uint32,
	amount decimal.Decimal,
) (decimal.Decimal, error) {
	return decimal.Decimal{}, fmt.Errorf("NOT IMPLEMENTED")
}
