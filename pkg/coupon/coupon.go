//nolint:dupl
package coupon

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/discountpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userspecialreduction"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/coupon"

	"github.com/google/uuid"
)

const secondsPerDay = 24 * 60 * 60

func GetCoupon(ctx context.Context, id string, couponType npool.CouponType) (info *npool.Coupon, err error) {
	infos := []*npool.Coupon{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		switch couponType {
		case npool.CouponType_FixAmount:
			fallthrough //nolint
		case npool.CouponType_Discount:
			stm := cli.
				CouponAllocated.
				Query().
				Where(
					couponallocated.ID(uuid.MustParse(id)),
				)
			return join(stm, couponType).
				Scan(ctx, &infos)
		case npool.CouponType_SpecialOffer:
			infos, err = special(ctx, cli, []uuid.UUID{uuid.MustParse(id)})
			return err
		case npool.CouponType_ThresholdReduction:
			return fmt.Errorf("NOT IMPLEMENTED")
		default:
			return fmt.Errorf("UNKNOWN coupon")
		}
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	return post(infos[0], couponType), nil
}

func GetManyCoupons(ctx context.Context, ids []string, couponType npool.CouponType) (infos []*npool.Coupon, err error) {
	uids := []uuid.UUID{}
	for _, id := range ids {
		if _, err := uuid.Parse(id); err != nil {
			return nil, err
		}
		uids = append(uids, uuid.MustParse(id))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		switch couponType {
		case npool.CouponType_FixAmount:
			fallthrough //nolint
		case npool.CouponType_Discount:
			stm := cli.
				CouponAllocated.
				Query().
				Where(
					couponallocated.IDIn(uids...),
				)
			return join(stm, couponType).
				Scan(ctx, &infos)
		case npool.CouponType_SpecialOffer:
			infos, err = special(ctx, cli, uids)
			return err
		case npool.CouponType_ThresholdReduction:
			return fmt.Errorf("NOT IMPLEMENTED")
		default:
			return fmt.Errorf("UNKNOWN coupon")
		}
	})
	if err != nil {
		return nil, err
	}

	for i, info := range infos {
		infos[i] = post(info, couponType)
	}

	return infos, nil
}

func join(stm *ent.CouponAllocatedQuery, couponType npool.CouponType) *ent.CouponAllocatedSelect {
	stm1 := stm.
		Select(
			couponallocated.FieldID,
			couponallocated.FieldAppID,
			couponallocated.FieldUserID,
			couponallocated.FieldCreateAt,
		)

	switch couponType {
	case npool.CouponType_FixAmount:
		return stm1.
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(couponpool.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(couponallocated.FieldCouponID),
						t1.C(couponpool.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(couponpool.FieldID), "coupon_id"),
						sql.As(t1.C(couponpool.FieldName), "name"),
						sql.As(t1.C(couponpool.FieldMessage), "message"),
						sql.As(t1.C(couponpool.FieldStart), "start"),
						sql.As(t1.C(couponpool.FieldDurationDays), "duration_days"),
						sql.As(t1.C(couponpool.FieldDenomination), "value"),
					)
			})
	case npool.CouponType_Discount:
		return stm1.
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(discountpool.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(couponallocated.FieldCouponID),
						t1.C(discountpool.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(couponpool.FieldID), "coupon_id"),
						sql.As(t1.C(discountpool.FieldName), "name"),
						sql.As(t1.C(discountpool.FieldMessage), "message"),
						sql.As(t1.C(discountpool.FieldStart), "start"),
						sql.As(t1.C(discountpool.FieldDurationDays), "duration_days"),
						sql.As(t1.C(discountpool.FieldDiscount), "value"),
					)
			})
	}
	return nil
}

func post(info *npool.Coupon, couponType npool.CouponType) *npool.Coupon {
	info.End = info.Start + info.DurationDays*secondsPerDay
	now := uint32(time.Now().Unix())

	if info.Start <= now && now <= info.End {
		if info.CreatedAt >= info.Start && info.CreatedAt <= info.End {
			info.Valid = true
		}
	}
	if info.CreatedAt+info.DurationDays*secondsPerDay < now {
		info.Expired = true
	}

	const accuracy = 1000000000000
	amount := func(samount string) string {
		damount, err := decimal.NewFromString(samount)
		if err != nil {
			logger.Sugar().Errorw(
				"post",
				"ID", info.ID,
				"CouponID", info.CouponID,
				"CouponType", couponType,
				"Value", samount,
			)
			return decimal.NewFromInt(0).String()
		}
		return damount.Div(decimal.NewFromInt(accuracy)).String()
	}

	switch couponType {
	case npool.CouponType_FixAmount:
		fallthrough //nolint
	case npool.CouponType_SpecialOffer:
		info.Value = amount(info.Value)
	}

	return info
}

func special(ctx context.Context, cli *ent.Client, uids []uuid.UUID) (coupons []*npool.Coupon, err error) {
	err = cli.
		UserSpecialReduction.
		Query().
		Where(
			userspecialreduction.IDIn(uids...),
		).
		Select(
			couponallocated.FieldID,
			couponallocated.FieldAppID,
			couponallocated.FieldUserID,
			couponallocated.FieldCreateAt,
		).
		Modify(func(s *sql.Selector) {
			s.AppendSelect(
				sql.As(s.C(couponpool.FieldID), "coupon_id"),
				sql.As(s.C(userspecialreduction.FieldMessage), "message"),
				sql.As(s.C(userspecialreduction.FieldStart), "start"),
				sql.As(s.C(userspecialreduction.FieldDurationDays), "duration_days"),
				sql.As(s.C(userspecialreduction.FieldAmount), "value"),
			)
		}).
		Scan(ctx, &coupons)
	return coupons, err
}
