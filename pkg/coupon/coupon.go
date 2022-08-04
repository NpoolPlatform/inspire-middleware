//nolint:dupl
package coupon

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"

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
		stm := cli.
			CouponAllocated.
			Query().
			Where(
				couponallocated.ID(uuid.MustParse(id)),
			)

		switch couponType {
		case npool.CouponType_FixAmount:
			fallthrough //nolint
		case npool.CouponType_Discount:
			fallthrough //nolint
		case npool.CouponType_SpecialOffer:
			return join(stm, couponType).
				Scan(ctx, &infos)
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

	return post(infos[0]), nil
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
		stm := cli.
			CouponAllocated.
			Query().
			Where(
				couponallocated.IDIn(uids...),
			)

		switch couponType {
		case npool.CouponType_FixAmount:
			fallthrough //nolint
		case npool.CouponType_Discount:
			fallthrough //nolint
		case npool.CouponType_SpecialOffer:
			return join(stm, couponType).
				Scan(ctx, &infos)
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
		infos[i] = post(info)
	}

	return infos, nil
}

func join(stm *ent.CouponAllocatedQuery, couponType npool.CouponType) *ent.CouponAllocatedSelect {
	stm1 := stm.
		Select(
			couponallocated.FieldID,
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
						t1.C(couponpool.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(couponpool.FieldName), "name"),
						sql.As(t1.C(couponpool.FieldMessage), "message"),
						sql.As(t1.C(couponpool.FieldStart), "start"),
						sql.As(t1.C(couponpool.FieldDurationDays), "duration_days"),
						sql.As(t1.C(couponpool.FieldDenomination), "value"),
					)
			})
	case npool.CouponType_SpecialOffer:
		return stm1.
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(userspecialreduction.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(couponallocated.FieldCouponID),
						t1.C(couponpool.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(couponpool.FieldMessage), "message"),
						sql.As(t1.C(couponpool.FieldStart), "start"),
						sql.As(t1.C(couponpool.FieldDurationDays), "duration_days"),
						sql.As(t1.C(couponpool.FieldDenomination), "value"),
					)
			})
	}
	return nil
}

func post(info *npool.Coupon) *npool.Coupon {
	info.End = info.Start + info.DurationDays*secondsPerDay
	now := uint32(time.Now().Unix())

	if info.Start <= now && now <= info.End {
		info.Valid = true
	}
	if info.End < now {
		info.Expired = true
	}

	return info
}
