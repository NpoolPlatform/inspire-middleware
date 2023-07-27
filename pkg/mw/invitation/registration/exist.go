package registration

import (
	"context"
	"fmt"

	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entregistration "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/registration"
)

func (h *Handler) ExistRegistration(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	exist := false
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Registration.
			Query().
			Where(
				entregistration.ID(*h.ID),
				entregistration.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return false, nil
}

func (h *Handler) ExistRegistrationConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := registrationcrud.SetQueryConds(cli.Registration.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
