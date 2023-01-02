package accounting

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	comm1 "github.com/NpoolPlatform/inspire-middleware/pkg/commission"
	reg1 "github.com/NpoolPlatform/inspire-middleware/pkg/invitation/registration"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

//nolint
func Accounting(
	ctx context.Context,
	appID, userID, goodID, orderID string,
	settleType commmgrpb.SettleType,
	paymentAmount decimal.Decimal,
	goodValue decimal.Decimal,
) (
	[]*npool.Commission,
	error,
) {
	offset := int32(0)
	limit := int32(100)

	inviters := []*regmgrpb.Registration{}
	for {
		_inviters, _, err := reg1.GetSuperiores(ctx, &regmgrpb.Conds{
			AppID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: appID,
			},
			InviterIDs: &commonpb.StringSliceVal{
				Op:    cruder.IN,
				Value: []string{userID},
			},
		}, offset, limit)
		if err != nil {
			return nil, err
		}
		if len(_inviters) == 0 {
			break
		}

		inviters = append(inviters, _inviters...)

		offset += limit
	}

	_inviters := []*regmgrpb.Registration{}
	for i, inviter := range inviters {
		if inviter.InviterID == uuid1.InvalidUUIDStr || inviter.InviterID == "" {
			_inviters = append(_inviters, inviter)
			inviters = append(inviters[0:i], inviters[i+1:]...)
			break
		}
	}

	if len(_inviters) == 0 {
		return nil, fmt.Errorf("invalid top inviter")
	}

	for {
		if len(inviters) == 1 {
			if _inviters[len(_inviters)-1].InviteeID == inviters[0].InviterID {
				return nil, fmt.Errorf("mismatch registration")
			}
			_inviters = append(_inviters, inviters[0])
			break
		}

		for i, inviter := range inviters {
			if _inviters[len(_inviters)-1].InviteeID == inviter.InviterID {
				_inviters = append(_inviters, inviter)
				inviters = append(inviters[0:i], inviters[i+1:]...)
				break
			}
		}
	}

	inviterIDs := []string{}
	for _, inviter := range _inviters {
		inviterIDs = append(inviterIDs, inviter.InviteeID)
	}

	comms, _, err := comm1.GetCommissions(ctx, &commmwpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: inviterIDs,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: goodID,
		},
		SettleType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(settleType),
		},
		EndAt: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(0),
		},
	}, int32(0), int32(len(inviterIDs)))
	if err != nil {
		return nil, err
	}

	_comms, err := comm1.Accounting(
		ctx,
		settleType,
		_inviters,
		comms,
		paymentAmount,
		goodValue,
	)
	if err != nil {
		return nil, err
	}

	// TODO: record user archivement

	return _comms, nil
}
