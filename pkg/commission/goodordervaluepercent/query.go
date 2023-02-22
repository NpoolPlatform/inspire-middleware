package goodordervaluepercent

import (
	"context"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	gopmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/commission/goodordervaluepercent"
	gopmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"
)

func GetGoodOrderValuePercent(ctx context.Context, id string) (*npool.Commission, error) {
	info, err := gopmgrcli.GetOrderValuePercent(ctx, id)
	if err != nil {
		return nil, err
	}

	return gop2Comm(info), nil
}

func GetGoodOrderValuePercents(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Commission, uint32, error) {
	infos, total, err := gopmgrcli.GetOrderValuePercents(ctx, &gopmgrpb.Conds{
		ID:      conds.ID,
		AppID:   conds.AppID,
		UserID:  conds.UserID,
		GoodID:  conds.GoodID,
		EndAt:   conds.EndAt,
		UserIDs: conds.UserIDs,
		GoodIDs: conds.GoodIDs,
	}, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return gops2Comms(infos), total, nil
}

func GetGoodOrderValuePercentOnly(ctx context.Context, conds *npool.Conds) (*npool.Commission, error) {
	info, err := gopmgrcli.GetOrderValuePercentOnly(ctx, &gopmgrpb.Conds{
		ID:     conds.ID,
		AppID:  conds.AppID,
		UserID: conds.UserID,
		GoodID: conds.GoodID,
		EndAt:  conds.EndAt,
	})
	if err != nil {
		return nil, err
	}

	return gop2Comm(info), nil
}

func gop2Comm(info *gopmgrpb.OrderValuePercent) *npool.Commission {
	if info == nil {
		return nil
	}

	return &npool.Commission{
		ID:             info.ID,
		AppID:          info.AppID,
		UserID:         info.UserID,
		GoodID:         &info.GoodID,
		SettleType:     commmgrpb.SettleType_GoodOrderValuePercent,
		SettleMode:     commmgrpb.SettleMode_SettleWithPaymentAmount,
		SettleInterval: commmgrpb.SettleInterval_SettleEveryOrder,
		Percent:        &info.Percent,
		StartAt:        info.StartAt,
		EndAt:          info.EndAt,
		CreatedAt:      info.CreatedAt,
		UpdatedAt:      info.UpdatedAt,
	}
}

func gops2Comms(infos []*gopmgrpb.OrderValuePercent) []*npool.Commission {
	comms := []*npool.Commission{}

	for _, info := range infos {
		comms = append(comms, gop2Comm(info))
	}

	return comms
}
