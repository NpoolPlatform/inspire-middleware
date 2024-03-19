package user

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievementuser "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievementuser"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
)

func (h *Handler) DeleteAchievementUser(ctx context.Context) (*npool.AchievementUser, error) {
	var err error
	info, err := h.GetAchievementUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		now := uint32(time.Now().Unix())
		if _, err := tx.
			AchievementUser.
			Update().
			Where(
				entachievementuser.ID(info.ID),
				entachievementuser.DeletedAt(0),
			).
			SetDeletedAt(now).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
