package invitationcode

import (
	"context"
	"fmt"

	invitationcodecrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	codegenerator "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode/generator"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateInvitationCode, *h.AppID, *h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &invitationcodecrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	exist, err := h.ExistInvitationCodeConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	var code string
	for {
		code, err = codegenerator.Generate()
		if err != nil {
			return nil, err
		}
		h.Conds = &invitationcodecrud.Conds{
			AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			InvitationCode: &cruder.Cond{Op: cruder.EQ, Val: code},
		}
		key1 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateInvitationCode, *h.AppID, code)
		if err := redis2.TryLock(key1, 0); err != nil {
			return nil, err
		}
		exist, err := h.ExistInvitationCodeConds(ctx)
		if err != nil {
			_ = redis2.Unlock(key1)
			return nil, err
		}
		_ = redis2.Unlock(key1)
		if exist {
			continue
		}
		break
	}

	key1 := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateInvitationCode, *h.AppID, code)
	if err := redis2.TryLock(key1, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key1)
	}()

	exist, err = h.ExistInvitationCodeConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := invitationcodecrud.CreateSet(
			cli.InvitationCode.Create(),
			&invitationcodecrud.Req{
				ID:             h.ID,
				AppID:          h.AppID,
				UserID:         h.UserID,
				InvitationCode: &code,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetInvitationCode(ctx)
}
