package event

import (
	"context"
	"fmt"

	eventcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/event"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func (h *Handler) UpdateEvent(ctx context.Context) (*npool.Event, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	if err := h.validateCoupons(ctx); err != nil {
		return nil, err
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
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetEvent(ctx)
}
