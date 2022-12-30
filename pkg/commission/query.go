package commission

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	goodorderpercent "github.com/NpoolPlatform/inspire-middleware/pkg/commission/goodorderpercent"
)

func GetCommissions(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Commission, uint32, error) {
	switch mgrpb.SettleType(conds.GetSettleType().GetValue()) {
	case mgrpb.SettleType_GoodOrderPercent:
		return goodorderpercent.GetGoodOrderPercents(ctx, conds, offset, limit)
	case mgrpb.SettleType_LimitedOrderPercent:
		fallthrough //nolint
	case mgrpb.SettleType_AmountThreshold:
	default:
		return nil, 0, fmt.Errorf("invalid settle type")
	}
	return nil, 0, fmt.Errorf("not implemented")
}

func GetCommissionOnly(ctx context.Context, conds *npool.Conds) (*npool.Commission, error) {
	switch mgrpb.SettleType(conds.GetSettleType().GetValue()) {
	case mgrpb.SettleType_GoodOrderPercent:
		return goodorderpercent.GetGoodOrderPercentOnly(ctx, conds)
	case mgrpb.SettleType_LimitedOrderPercent:
		fallthrough //nolint
	case mgrpb.SettleType_AmountThreshold:
	default:
		return nil, fmt.Errorf("invalid settle type")
	}
	return nil, fmt.Errorf("not implemented")
}
