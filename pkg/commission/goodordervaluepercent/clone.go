package goodordervaluepercent

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entgop "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/goodordervaluepercent"

	"github.com/google/uuid"
)

func CloneGoodOrderValuePercents(ctx context.Context, appID, fromGoodID, toGoodID, value string) error {
	val, err := decimal.NewFromString(value)
	if err != nil {
		return err
	}
	if val.Cmp(decimal.NewFromInt(0)) == 0 {
		return nil
	}
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

			if gop.Percent.Cmp(decimal.NewFromInt(0)) == 0 {
				continue
			}

			_, err = cli.
				GoodOrderValuePercent.
				Create().
				SetAppID(gop.AppID).
				SetUserID(gop.UserID).
				SetGoodID(uuid.MustParse(toGoodID)).
				SetPercent(
					gop.
						Percent.
						Mul(val).
						Div(decimal.NewFromInt(100)). //nolint
						RoundUp(0),
				).
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
