package allocated

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	converter "github.com/NpoolPlatform/inspire-manager/pkg/converter/coupon/allocated"
	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/coupon/allocated"

	fixamount "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/allocated/fixamount"
)

func CreateCoupon(ctx context.Context, in *npool.CouponReq) (*npool.Coupon, error) {
	var id string

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stm, err := crud.CreateSet(tx.CouponAllocated.Create(), &allocatedmgrpb.AllocatedReq{
			ID:         in.ID,
			CouponType: in.CouponType,
			AppID:      in.AppID,
			UserID:     in.UserID,
			CouponID:   in.CouponID,
		})
		if err != nil {
			return err
		}

		_info, err := stm.Save(_ctx)
		if err != nil {
			return err
		}

		id = _info.ID.String()
		info := converter.Ent2Grpc(_info)

		switch in.GetCouponType() {
		case allocatedmgrpb.CouponType_FixAmount:
			_, err = fixamount.CreateFixAmount(
				ctx,
				in.GetID(),
				tx,
				func(_ctx context.Context) (*allocatedmgrpb.Allocated, error) {
					return info, nil
				})
		case allocatedmgrpb.CouponType_Discount:
		case allocatedmgrpb.CouponType_SpecialOffer:
		case allocatedmgrpb.CouponType_ThresholdFixAmount:
		case allocatedmgrpb.CouponType_ThresholdDiscount:
		case allocatedmgrpb.CouponType_GoodFixAmount:
		case allocatedmgrpb.CouponType_GoodDiscount:
		case allocatedmgrpb.CouponType_GoodThresholdFixAmount:
		case allocatedmgrpb.CouponType_GoodThresholdDiscount:
		default:
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return GetCoupon(ctx, id, in.GetCouponType())
}
