package registration

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entregistration "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/registration"

	// registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RegistrationSelect
	infos     []*npool.Registration
	total     uint32
}

func (h *queryHandler) queryRegistration(cli *ent.Client) {
	h.stmSelect = cli.
		Registration.
		Query().
		Where(
			entregistration.ID(*h.ID),
			entregistration.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetRegistration(ctx context.Context) (*npool.Registration, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Registration{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryRegistration(cli)
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

func (h *Handler) GetRegistrations(ctx context.Context) ([]*npool.Registration, uint32, error) {
	return nil, 0, nil
}
