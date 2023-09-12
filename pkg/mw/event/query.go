package event

import (
	"context"
	"encoding/json"
	"fmt"

	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entevent "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/event"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.EventSelect
	infos     []*npool.Event
	total     uint32
}

func (h *queryHandler) queryEvent(cli *ent.Client) {
	h.stmSelect = cli.
		Event.
		Query().
		Where(
			entevent.ID(*h.ID),
			entevent.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) queryEvents(ctx context.Context, cli *ent.Client) error {
	stm, err := eventcrud.SetQueryConds(cli.Event.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.stmSelect = stm.Select()
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.EventType = basetypes.UsedFor(basetypes.UsedFor_value[info.EventTypeStr])
		_ = json.Unmarshal([]byte(info.CouponIDsStr), &info.CouponIDs)
		if info.GoodID != nil && *info.GoodID == uuid.Nil.String() {
			info.GoodID = nil
		}
		if info.AppGoodID != nil && *info.AppGoodID == uuid.Nil.String() {
			info.AppGoodID = nil
		}
		amount, err := decimal.NewFromString(info.Credits)
		if err != nil {
			info.Credits = decimal.NewFromInt(0).String()
		} else {
			info.Credits = amount.String()
		}
		amount, err = decimal.NewFromString(info.CreditsPerUSD)
		if err != nil {
			info.CreditsPerUSD = decimal.NewFromInt(0).String()
		} else {
			info.CreditsPerUSD = amount.String()
		}
	}
}

func (h *Handler) GetEvent(ctx context.Context) (*npool.Event, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Event{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryEvent(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetEvents(ctx context.Context) ([]*npool.Event, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Event{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEvents(ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetEventOnly(ctx context.Context) (*npool.Event, error) {
	const rowLimit = 2
	h.Limit = rowLimit
	infos, _, err := h.GetEvents(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos[0], nil
}
