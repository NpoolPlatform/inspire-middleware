package goodorderpercent

import (
	"context"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	gopmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/commission/goodorderpercent"
	gopmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodorderpercent"
)

func GetGoodOrderPercent(ctx context.Context, id string) (*npool.Commission, error) {
	info, err := gopmgrcli.GetOrderPercent(ctx, id)
	if err != nil {
		return nil, err
	}

	return gop2Comm(info), nil
}

func gop2Comm(info *gopmgrpb.OrderPercent) *npool.Commission {
	return &npool.Commission{
		ID:             info.ID,
		AppID:          info.AppID,
		UserID:         info.UserID,
		GoodID:         &info.GoodID,
		SettleType:     commmgrpb.SettleType_GoodOrderPercent,
		SettleMode:     commmgrpb.SettleMode_SettleWithPaymentAmount,
		SettleInterval: commmgrpb.SettleInterval_SettleEveryOrder,
		Percent:        &info.Percent,
		StartAt:        info.StartAt,
		EndAt:          info.EndAt,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
}
