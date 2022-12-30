package commission

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	goodorderpercent "github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"
)

func UpdateCommission(ctx context.Context, in *npool.CommissionReq) (*npool.Commission, error) {
	switch in.GetSettleType() {
	case mgrpb.SettleType_GoodOrderPercent:
		return goodorderpercent.UpdateGoodOrderPercent(ctx, in)
	case mgrpb.SettleType_LimitedOrderPercent:
	case mgrpb.SettleType_AmountThreshold:
	default:
		return nil, fmt.Errorf("unknown settle type")
	}
	return nil, fmt.Errorf("not implemented")
}
