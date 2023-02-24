package registration

import (
	"context"
	"fmt"

	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"
)

func GetInviters(ctx context.Context, appID, userID string) ([]*regmgrpb.Registration, []string, error) { //nolint
	offset := int32(0)
	const limit = int32(100)

	inviters := []*regmgrpb.Registration{}
	for {
		_inviters, _, err := GetSuperiores(ctx, &regmgrpb.Conds{
			AppID:      &commonpb.StringVal{Op: cruder.EQ, Value: appID},
			InviteeIDs: &commonpb.StringSliceVal{Op: cruder.IN, Value: []string{userID}},
		}, offset, limit)
		if err != nil {
			return nil, nil, err
		}
		if len(_inviters) == 0 {
			break
		}

		inviters = append(inviters, _inviters...)
		offset += limit
	}

	inviteeMap := map[string]struct{}{}
	for _, inviter := range inviters {
		inviteeMap[inviter.InviteeID] = struct{}{}
	}

	inviterCount := len(inviters)
	_inviters := []*regmgrpb.Registration{}

	for i, inviter := range inviters {
		_, ok := inviteeMap[inviter.InviterID]
		if !ok {
			_inviters = append(_inviters, inviter)
			inviters = append(inviters[0:i], inviters[i+1:]...)
			break
		}
	}

	if inviterCount == 0 {
		_inviters = append(_inviters, &regmgrpb.Registration{
			AppID:     appID,
			InviterID: uuid1.InvalidUUIDStr,
			InviteeID: userID,
		})
	}

	if len(_inviters) == 0 {
		return nil, nil, fmt.Errorf("invalid top inviter")
	}

	for {
		if inviterCount == 0 || len(inviters) == 0 {
			break
		}

		if len(inviters) == 1 {
			if _inviters[len(_inviters)-1].InviteeID != inviters[0].InviterID {
				return nil, nil, fmt.Errorf("mismatch registration")
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

	inviterIDs := []string{userID}
	if inviterCount > 0 {
		inviterIDs = []string{_inviters[0].InviterID}
		for _, inviter := range _inviters {
			inviterIDs = append(inviterIDs, inviter.InviteeID)
		}
	}

	return _inviters, inviterIDs, nil
}
