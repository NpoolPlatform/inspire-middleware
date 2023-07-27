package commission

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entcommission "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CommissionSelect
	infos     []*npool.Commission
	total     uint32
}

func (h *queryHandler) queryCommission(cli *ent.Client) {
	h.stmSelect = cli.
		Commission.
		Query().
		Where(
			entcommission.ID(*h.ID),
			entcommission.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetCommission(ctx context.Context) (*npool.Commission, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Commission{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCommission(cli)
		return handler.scan(ctx)
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
