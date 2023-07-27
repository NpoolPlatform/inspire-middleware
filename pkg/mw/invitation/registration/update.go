package registration

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"
)

func (h *Handler) UpdateRegistration(ctx context.Context) (*npool.Registration, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := registrationcrud.UpdateSet(
			cli.Registration.UpdateOneID(*h.ID),
			&registrationcrud.Req{
				InviterID: h.InviterID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRegistration(ctx)
}