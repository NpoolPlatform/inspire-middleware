package event

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func (h *Handler) UpdateEvent(ctx context.Context) (*npool.Event, error) {
	if err := h.validateCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := eventcrud.UpdateSet(
			cli.Event.UpdateOneID(*h.ID),
			&eventcrud.Req{
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
