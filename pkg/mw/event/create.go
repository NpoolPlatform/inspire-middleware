package event

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/google/uuid"
)

func (h *Handler) CreateEvent(ctx context.Context) (*npool.Event, error) {
	if err := h.validateCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppEvent, *h.AppID, *h.EventType)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
	}
	exist, err := h.ExistEventConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := eventcrud.CreateSet(
			cli.Event.Create(),
			&eventcrud.Req{
				EntID:          h.EntID,
				AppID:          h.AppID,
				EventType:      h.EventType,
				CouponIDs:      h.CouponIDs,
				Credits:        h.Credits,
				CreditsPerUSD:  h.CreditsPerUSD,
				MaxConsecutive: h.MaxConsecutive,
				GoodID:         h.GoodID,
				AppGoodID:      h.AppGoodID,
				InviterLayers:  h.InviterLayers,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetEvent(ctx)
}
