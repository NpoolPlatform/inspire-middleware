package event

import (
	"context"
	"fmt"
	"time"

	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func (h *Handler) DeleteEvent(ctx context.Context) (*npool.Event, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetEvent(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := eventcrud.UpdateSet(
			cli.Event.UpdateOneID(*h.ID),
			&eventcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
