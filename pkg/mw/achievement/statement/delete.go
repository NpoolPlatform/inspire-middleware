package statement

import (
	"context"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entstatement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
)

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			Statement.
			Update().
			Where(
				entstatement.ID(*h.ID),
				entstatement.DeletedAt(0),
			).
			SetDeletedAt(uint32(time.Now().Unix())).
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
