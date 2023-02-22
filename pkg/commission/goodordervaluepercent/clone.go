package goodordervaluepercent

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entgop "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/goodordervaluepercent"

	"github.com/google/uuid"
)

func CloneGoodOrderValuePercents(ctx context.Context, appID, fromGoodID, toGoodID string) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		gops, err := cli.
			GoodOrderValuePercent.
			Query().
			Where(
				entgop.AppID(uuid.MustParse(appID)),
				entgop.GoodID(uuid.MustParse(fromGoodID)),
				entgop.DeletedAt(0),
				entgop.EndAt(0),
			).
			All(_ctx)
		if err != nil {
			return err
		}

		for _, gop := range gops {
			_gop, err := cli.
				GoodOrderValuePercent.
				Query().
				Where(
					entgop.AppID(gop.AppID),
					entgop.GoodID(uuid.MustParse(toGoodID)),
					entgop.UserID(gop.UserID),
					entgop.DeletedAt(0),
					entgop.EndAt(0),
				).
				Only(_ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return err
				}
			}
			if _gop != nil {
				continue
			}

			_, err = cli.
				GoodOrderValuePercent.
				Create().
				SetAppID(gop.AppID).
				SetUserID(gop.UserID).
				SetGoodID(uuid.MustParse(toGoodID)).
				SetPercent(gop.Percent).
				SetStartAt(gop.StartAt).
				SetEndAt(0).
				Save(_ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
