package statement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entarchivementdetail "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/archivementdetail"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/statement"
)

func (h *Handler) DeleteStatement(ctx context.Context) (*npool.Statement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetStatement(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			ArchivementDetail.
			Update().
			Where(
				entarchivementdetail.ID(*h.ID),
				entarchivementdetail.DeletedAt(0),
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
