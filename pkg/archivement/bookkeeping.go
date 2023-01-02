package archivement

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entarchivementdetail "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementdetail"
	entarchivementgeneral "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementgeneral"

	detailcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/archivement/detail"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	generalcrud "github.com/NpoolPlatform/inspire-manager/pkg/crud/archivement/general"
	generalmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	commonpb "github.com/NpoolPlatform/message/npool"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	errno "github.com/NpoolPlatform/go-service-framework/pkg/errno"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func TryCreateGeneral(ctx context.Context, appID, userID, goodID, coinTypeID string) (string, error) {
	key := fmt.Sprintf("archivement-general:%v:%v:%v:%v", appID, userID, goodID, coinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	general1, err := generalcrud.RowOnly(ctx, &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: goodID,
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: coinTypeID,
		},
	})
	if err != nil {
		return "", err
	}
	if general1 != nil {
		return general1.ID.String(), nil
	}

	general1, err = generalcrud.Create(ctx, &generalmgrpb.GeneralReq{
		AppID:      &appID,
		UserID:     &userID,
		GoodID:     &goodID,
		CoinTypeID: &coinTypeID,
	})
	if err != nil {
		return "", err
	}

	return general1.ID.String(), nil
}

func detailKey(in *detailmgrpb.DetailReq) string {
	return fmt.Sprintf("archivement-detail:%v:%v:%v:%v",
		in.GetAppID(),
		in.GetUserID(),
		in.GetGoodID(),
		in.GetOrderID(),
	)
}

func BookKeeping(ctx context.Context, in *detailmgrpb.DetailReq) error { //nolint
	val, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		return err
	}
	if val.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("invalid amount")
	}

	var generalID string
	if generalID, err = TryCreateGeneral(
		ctx,
		in.GetAppID(), in.GetUserID(), in.GetGoodID(), in.GetCoinTypeID(),
	); err != nil {
		return err
	}

	key := detailKey(in)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := detailcrud.ExistConds(ctx, &detailmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetUserID(),
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
		OrderID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetOrderID(),
		},
	})
	if err != nil {
		return err
	}
	if exist {
		return errno.ErrAlreadyExists
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c1, err := detailcrud.CreateSet(tx.ArchivementDetail.Create(), in)
		if err != nil {
			return err
		}

		_, err = c1.Save(ctx)
		if err != nil {
			return err
		}

		info, err := tx.
			ArchivementGeneral.
			Query().
			Where(
				entarchivementgeneral.ID(uuid.MustParse(generalID)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		selfUnits := uint32(0)
		selfAmount := decimal.NewFromInt(0).String()
		selfCommission := decimal.NewFromInt(0).String()

		if in.GetSelfOrder() {
			selfUnits += *in.Units
			selfAmount = in.GetUSDAmount()
			selfCommission = in.GetCommission()
		}

		c2, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
			AppID:           in.AppID,
			UserID:          in.UserID,
			GoodID:          in.GoodID,
			CoinTypeID:      in.CoinTypeID,
			TotalUnits:      in.Units,
			SelfUnits:       &selfUnits,
			TotalAmount:     in.USDAmount,
			SelfAmount:      &selfAmount,
			TotalCommission: in.Commission,
			SelfCommission:  &selfCommission,
		})
		if err != nil {
			return err
		}

		_, err = c2.Save(ctx)
		return err
	})
}

func BookKeepingV2(ctx context.Context, in []*detailmgrpb.DetailReq) error { //nolint
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, info := range in {
			err := func(info *detailmgrpb.DetailReq) error {
				key := detailKey(in[0])
				if err := redis2.TryLock(key, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key)
				}()

				key1 := fmt.Sprintf("archivement-general:%v:%v:%v:%v",
					info.GetAppID(),
					info.GetUserID(),
					info.GetGoodID(),
					info.GetCoinTypeID())
				if err := redis2.TryLock(key1, 0); err != nil {
					return err
				}
				defer func() {
					_ = redis2.Unlock(key1)
				}()

				val, err := decimal.NewFromString(info.GetAmount())
				if err != nil {
					return err
				}
				if val.Cmp(decimal.NewFromInt(0)) <= 0 {
					return fmt.Errorf("invalid amount")
				}

				d, err := tx.
					ArchivementDetail.
					Query().
					Where(
						entarchivementdetail.AppID(uuid.MustParse(info.GetAppID())),
						entarchivementdetail.UserID(uuid.MustParse(info.GetUserID())),
						entarchivementdetail.GoodID(uuid.MustParse(info.GetGoodID())),
						entarchivementdetail.OrderID(uuid.MustParse(info.GetOrderID())),
					).
					Only(_ctx)
				if err != nil {
					if !ent.IsNotFound(err) {
						return err
					}
				}
				if d != nil {
					return nil
				}

				c1, err := detailcrud.CreateSet(tx.ArchivementDetail.Create(), info)
				if err != nil {
					return err
				}

				_, err = c1.Save(ctx)
				if err != nil {
					return err
				}

				selfUnits := uint32(0)
				selfAmount := decimal.NewFromInt(0).String()
				selfCommission := decimal.NewFromInt(0).String()

				if info.GetSelfOrder() {
					selfUnits += info.GetUnits()
					selfAmount = info.GetUSDAmount()
					selfCommission = info.GetCommission()
				}

				g, err := tx.
					ArchivementGeneral.
					Query().
					Where(
						entarchivementgeneral.AppID(uuid.MustParse(info.GetAppID())),
						entarchivementgeneral.UserID(uuid.MustParse(info.GetUserID())),
						entarchivementgeneral.GoodID(uuid.MustParse(info.GetGoodID())),
						entarchivementgeneral.CoinTypeID(uuid.MustParse(info.GetCoinTypeID())),
					).
					ForUpdate().
					Only(ctx)
				if err != nil {
					if !ent.IsNotFound(err) {
						return err
					}
				}
				if g == nil {
					c2 := generalcrud.CreateSet(
						tx.ArchivementGeneral.Create(),
						&generalmgrpb.GeneralReq{
							AppID:      info.AppID,
							UserID:     info.UserID,
							GoodID:     info.GoodID,
							CoinTypeID: info.CoinTypeID,
						})

					g, err = c2.Save(ctx)
					if err != nil {
						return err
					}
				}

				c2, err := generalcrud.UpdateSet(g, &generalmgrpb.GeneralReq{
					AppID:           info.AppID,
					UserID:          info.UserID,
					GoodID:          info.GoodID,
					CoinTypeID:      info.CoinTypeID,
					TotalUnits:      info.Units,
					SelfUnits:       &selfUnits,
					TotalAmount:     info.USDAmount,
					SelfAmount:      &selfAmount,
					TotalCommission: info.Commission,
					SelfCommission:  &selfCommission,
				})
				if err != nil {
					return err
				}

				_, err = c2.Save(ctx)
				if err != nil {
					return err
				}

				return nil
			}(info)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
