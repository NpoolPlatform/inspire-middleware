package allocated

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcouponallocated "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponallocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	// cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CouponAllocatedSelect
	infos     []*npool.Coupon
	total     uint32
}

func (h *queryHandler) selectCoupon(stm *ent.CouponAllocatedQuery) *ent.CouponAllocatedSelect {
	return stm.Select(
		entcouponallocated.FieldID,
	)
}

func (h *queryHandler) queryCoupon(cli *ent.Client) {
	stm := cli.
		CouponAllocated.
		Query().
		Where(
			entcouponallocated.ID(*h.ID),
			entcouponallocated.DeletedAt(0),
		)
	h.stmSelect = h.selectCoupon(stm)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetCoupon(ctx context.Context) (*npool.Coupon, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Coupon{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCoupon(cli)
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

	return handler.infos[0], nil
}
