package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	usercrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/task/user"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
)

func (h *Handler) ExistTaskUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(
			cli.TaskUser.Query(),
			h.Conds,
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}
