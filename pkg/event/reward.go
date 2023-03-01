package event

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	affiliate "github.com/NpoolPlatform/inspire-middleware/pkg/event/affiliate"
	self "github.com/NpoolPlatform/inspire-middleware/pkg/event/self"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/shopspring/decimal"
)

func RewardEvent(
	ctx context.Context,
	appID, userID string,
	eventType basetypes.UsedFor,
	goodID *string,
	consecutive uint32,
	amount decimal.Decimal,
) ([]*npool.Credit, error) {
	switch eventType {
	case basetypes.UsedFor_Purchase:
		return self.RewardEvent(ctx, appID, userID, eventType, goodID, consecutive, amount)
	case basetypes.UsedFor_AffiliatePurchase:
		return affiliate.RewardEvent(ctx, appID, userID, eventType, amount)
	}
	return []*npool.Credit{}, fmt.Errorf("NOT IMPLEMENTED")
}
