package user

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievementuser "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievementuser"
)

func (h *Handler) DeleteAchievementUser(ctx context.Context) error {
	var err error
	info, err := h.GetAchievementUser(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
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
}
