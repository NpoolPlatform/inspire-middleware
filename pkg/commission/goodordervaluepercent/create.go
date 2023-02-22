package goodordervaluepercent

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entgop "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/goodordervaluepercent"

	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/commission/goodordervaluepercent"
	gopmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission/goodordervaluepercent"

	"github.com/google/uuid"
)

func CreateGoodOrderValuePercent(ctx context.Context, in *npool.CommissionReq) (*npool.Commission, error) {
	var info *ent.GoodOrderValuePercent
	var id string
	var err error

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.
			GoodOrderValuePercent.
			Query().
			ForUpdate().
			Where(
				entgop.AppID(uuid.MustParse(in.GetAppID())),
				entgop.UserID(uuid.MustParse(in.GetUserID())),
				entgop.GoodID(uuid.MustParse(in.GetGoodID())),
				entgop.EndAt(0),
			).
			Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}

		if info != nil {
			_, err := info.
				Update().
				SetEndAt(uint32(time.Now().Unix())).
				Save(_ctx)
			if err != nil {
				return err
			}
		}

		startAt := in.StartAt
		now := uint32(time.Now().Unix())

		if startAt == nil {
			startAt = &now
		}

		c, err := crud.CreateSet(tx.GoodOrderValuePercent.Create(), &gopmgrpb.OrderValuePercentReq{
			ID:      in.ID,
			AppID:   in.AppID,
			UserID:  in.UserID,
			GoodID:  in.GoodID,
			Percent: in.Percent,
			StartAt: startAt,
		})
		if err != nil {
			return err
		}

		info, err = c.Save(_ctx)
		if err != nil {
			return err
		}

		id = info.ID.String()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetGoodOrderValuePercent(ctx, id)
}
