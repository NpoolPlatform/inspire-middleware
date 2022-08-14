package invitation

import (
	"context"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/apppurchaseamountsetting"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

	"github.com/google/uuid"
)

func GetPercents(
	ctx context.Context, appID string, userIDs []string, activeOnly bool, offset, limit int32,
) (
	infos []*npool.Percent, total uint32, err error,
) {
	users := []uuid.UUID{}
	for _, user := range userIDs {
		users = append(users, uuid.MustParse(user))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppPurchaseAmountSetting.
			Query()

		if activeOnly {
			stm = stm.
				Where(
					apppurchaseamountsetting.AppID(uuid.MustParse(appID)),
					apppurchaseamountsetting.UserIDIn(users...),
					apppurchaseamountsetting.End(0),
				)
		} else {
			stm = stm.
				Where(
					apppurchaseamountsetting.AppID(uuid.MustParse(appID)),
					apppurchaseamountsetting.UserIDIn(users...),
				)
		}

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return stm.
			Select(
				apppurchaseamountsetting.FieldUserID,
				apppurchaseamountsetting.FieldCoinTypeID,
				apppurchaseamountsetting.FieldGoodID,
				apppurchaseamountsetting.FieldPercent,
				apppurchaseamountsetting.FieldStart,
			).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify().
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
