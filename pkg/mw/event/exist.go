package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entevent "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/event"
)

func (h *Handler) ExistEvent(ctx context.Context) (bool, error) {
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
			return wlog.WrapError(err)
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}

func (h *Handler) ExistEventConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := eventcrud.SetQueryConds(cli.Event.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}
