package affiliate

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	mgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/event"
	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	self "github.com/NpoolPlatform/inspire-middleware/pkg/event/self"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/shopspring/decimal"
)

func RewardEvent(
	ctx context.Context,
	appID, userID string,
	eventType basetypes.UsedFor,
	goodID *string,
	amount decimal.Decimal,
) ([]*npool.Credit, error) {
	_, inviterIDs, err := registration1.GetInviters(ctx, appID, userID)
	if err != nil {
		return nil, err
	}
	if len(inviterIDs) == 0 {
		return nil, err
	}

	conds := &mgrpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: appID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(eventType)},
	}
	if goodID != nil && eventType == basetypes.UsedFor_AffiliatePurchase {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *goodID}
	}

	info, err := mgrcli.GetEventOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if info.InviterLayers == 0 {
		return nil, err
	}

	credits := []*npool.Credit{}

	i := uint32(0)
	const layerIgnore = 2
	j := len(inviterIDs) - layerIgnore

	for ; i < info.InviterLayers && j >= 0; i++ {
		credit, err := self.RewardEvent(ctx, appID, inviterIDs[j], eventType, goodID, 0, amount)
		if err != nil {
			return nil, err
		}

		j--
		if len(credit) == 0 {
			continue
		}

		credits = append(credits, credit...)
	}

	return credits, nil
}
