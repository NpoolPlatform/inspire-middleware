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

func GetGoodOrderPercents(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Commission, uint32, error) {
	infos, total, err := gopmgrcli.GetOrderPercents(ctx, &gopmgrpb.Conds{
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

func GetGoodOrderPercentOnly(ctx context.Context, conds *npool.Conds) (*npool.Commission, error) {
	info, err := gopmgrcli.GetOrderPercentOnly(ctx, &gopmgrpb.Conds{
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

func gop2Comm(info *gopmgrpb.OrderPercent) *npool.Commission {
	if info == nil {
		return nil
	}

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

func gops2Comms(infos []*gopmgrpb.OrderPercent) []*npool.Commission {
	comms := []*npool.Commission{}

	for _, info := range infos {
		comms = append(comms, gop2Comm(info))
	}

	return comms
}
