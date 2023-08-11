package event

import (
	"context"
	"fmt"

	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entevent "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/event"
)

func (h *Handler) ExistEvent(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_exist, err := cli.
			Event.
			Query().
			Where(
				entevent.ID(*h.ID),
				entevent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistEventConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := eventcrud.SetQueryConds(cli.Event.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
