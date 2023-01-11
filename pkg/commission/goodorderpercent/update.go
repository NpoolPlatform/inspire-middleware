package goodorderpercent

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	gopmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/commission/goodorderpercent"
	gopmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"
)

func UpdateGoodOrderPercent(ctx context.Context, in *npool.CommissionReq) (*npool.Commission, error) {
	_, err := gopmgrcli.UpdateOrderPercent(ctx, &gopmgrpb.OrderPercentReq{
		ID:      in.ID,
		AppID:   in.AppID,
		UserID:  in.UserID,
		GoodID:  in.GoodID,
		Percent: in.Percent,
		StartAt: in.StartAt,
	})
	if err != nil {
		return nil, err
	}

	return GetGoodOrderPercent(ctx, in.GetID())
}
