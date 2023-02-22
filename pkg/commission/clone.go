package commission

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodordervaluepercent"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

	"github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"
)

func CloneCommissions(ctx context.Context, appID, fromGoodID, toGoodID string, settleType mgrpb.SettleType) error {
	switch settleType {
	case mgrpb.SettleType_GoodOrderPercent:
		return goodorderpercent.CloneGoodOrderPercents(ctx, appID, fromGoodID, toGoodID)
	case mgrpb.SettleType_GoodOrderValuePercent:
		return goodordervaluepercent.CloneGoodOrderValuePercents(ctx, appID, fromGoodID, toGoodID)
	case mgrpb.SettleType_LimitedOrderPercent:
	case mgrpb.SettleType_AmountThreshold:
	default:
		return fmt.Errorf("unknown settle type")
	}
	return fmt.Errorf("not implemented")
}
