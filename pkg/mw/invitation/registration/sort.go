package registration

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

func (h *Handler) GetSortedInviters(ctx context.Context) ([]*npool.Registration, []string, error) {
	if h.AppID == nil {
		return nil, nil, fmt.Errorf("invalid appid")
	}
	if h.InviteeID == nil {
		return nil, nil, fmt.Errorf("invalid inviteeid")
	}

	h.Limit = constant.DefaultRowLimit
	h.Offset = 0
	h.Conds = &registrationcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		InviteeIDs: &cruder.Cond{Op: cruder.IN, Val: []uuid.UUID{*h.InviteeID}},
	}

	inviters := []*npool.Registration{}
	for {
		_inviters, _, err := h.GetSuperiores(ctx)
		if err != nil {
			return nil, nil, err
		}
		if len(_inviters) == 0 {
			break
		}
		inviters = append(inviters, _inviters...)
		h.Offset += h.Limit
	}

	inviteeMap := map[string]struct{}{}
	for _, inviter := range inviters {
		inviteeMap[inviter.InviteeID] = struct{}{}
	}

	inviterCount := len(inviters)
	_inviters := []*npool.Registration{}

	for i, inviter := range inviters {
		_, ok := inviteeMap[inviter.InviterID]
		if !ok {
			_inviters = append(_inviters, inviter)
			inviters = append(inviters[0:i], inviters[i+1:]...)
			break
		}
	}

	if inviterCount == 0 {
		_inviters = append(_inviters, &npool.Registration{
			AppID:     h.AppID.String(),
			InviterID: uuid.Nil.String(),
			InviteeID: h.InviteeID.String(),
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

	inviterIDs := []string{h.InviteeID.String()}
	if inviterCount > 0 {
		inviterIDs = []string{_inviters[0].InviterID}
		for _, inviter := range _inviters {
			inviterIDs = append(inviterIDs, inviter.InviteeID)
		}
	}

	return _inviters, inviterIDs, nil
}
