package achievement

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"
)

type deleteHandler struct {
	*Handler
	info         *npool.Achievement
	statementIDs []uint32
}

func (h *Handler) DeleteAchievement(ctx context.Context) (*npool.Achievement, error) {
	handler := &deleteHandler{
		Handler:      h,
		statementIDs: []uint32{},
	}

	var err error
	handler.info, err = h.GetAchievement(ctx)
	if err != nil {
		return nil, err
	}
	if handler.info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		now := uint32(time.Now().Unix())
		if _, err := tx.
			Achievement.
			Update().
			Where(
				entachievement.ID(handler.info.ID),
				entachievement.DeletedAt(0),
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

	return handler.info, nil
}
