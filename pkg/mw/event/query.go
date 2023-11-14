package event

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"
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
	stmCount  *ent.EventSelect
	stmSelect *ent.EventSelect
	infos     []*npool.Event
	total     uint32
}

func (h *queryHandler) selectEvent(stm *ent.EventQuery) *ent.EventSelect {
	return stm.Select(
		entevent.FieldID,
	)
}

func (h *queryHandler) queryEvent(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}

	stm := cli.Event.Query().Where(entevent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entevent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entevent.EntID(*h.EntID))
	}
	h.selectEvent(stm)
	return nil
}

func (h *queryHandler) queryEvents(cli *ent.Client) (*ent.EventSelect, error) {
	stm, err := eventcrud.SetQueryConds(cli.Event.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectEvent(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entevent.Table)
	s.LeftJoin(t).
		On(
			s.C(entevent.FieldID),
			t.C(entevent.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entevent.FieldEntID), "ent_id"),
			sql.As(t.C(entevent.FieldAppID), "app_id"),
			sql.As(t.C(entevent.FieldEventType), "event_type"),
			sql.As(t.C(entevent.FieldCouponIds), "coupon_ids"),
			sql.As(t.C(entevent.FieldCredits), "credits"),
			sql.As(t.C(entevent.FieldCreditsPerUsd), "credits_per_usd"),
			sql.As(t.C(entevent.FieldMaxConsecutive), "max_consecutive"),
			sql.As(t.C(entevent.FieldGoodID), "good_id"),
			sql.As(t.C(entevent.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entevent.FieldInviterLayers), "inviter_layers"),
			sql.As(t.C(entevent.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entevent.FieldCreatedAt), "created_at"),
			sql.As(t.C(entevent.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return err
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
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Event{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEvent(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
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

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryEvents(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryEvents(cli)
		if err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)
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
