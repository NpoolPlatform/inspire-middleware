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
			Select(
				couponallocated.FieldID,
				couponallocated.FieldCreateAt,
			).
			Where(
				couponallocated.ID(uuid.MustParse(id)),
			)

		switch couponType {
		case npool.CouponType_FixAmount:
			err = stm.
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
				}).
				Scan(ctx, &infos)
		case npool.CouponType_Discount:
			err = stm.
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
				}).
				Scan(ctx, &infos)
		case npool.CouponType_SpecialOffer:
			err = stm.
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
				}).
				Scan(ctx, &infos)
		case npool.CouponType_ThresholdReduction:
			return fmt.Errorf("NOT IMPLEMENTED")
		default:
			return fmt.Errorf("UNKNOWN coupon")
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	coupon := infos[0]
	coupon.End = coupon.Start + coupon.DurationDays*secondsPerDay
	now := uint32(time.Now().Unix())

	if coupon.Start <= now && now <= coupon.End {
		coupon.Valid = true
	}
	if coupon.End < now {
		coupon.Expired = true
	}

	return infos[0], nil
}
