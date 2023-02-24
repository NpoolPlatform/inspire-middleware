package affiliate

import (
	"context"
	"fmt"

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

	info, err := mgrcli.GetEventOnly(ctx, conds)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("event is invalid")
	}

	if info.InviterLayers == 0 {
		return nil, err
	}

	credits := []*npool.Credit{}

	for i := uint32(0); i < info.InviterLayers; i++ {
		inviterID := inviterIDs[len(inviterIDs)-1-int(i)]
		credit, err := self.RewardEvent(ctx, appID, inviterID, eventType, nil, 0, amount)
		if err != nil {
			return nil, err
		}

		credits = append(credits, credit...)
	}

	return credits, fmt.Errorf("NOT IMPLEMENTED")
}
