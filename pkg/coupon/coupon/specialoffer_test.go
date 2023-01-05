package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var userID = uuid.NewString()
var specialOffer = &npool.Coupon{
	ID:               uuid.NewString(),
	CouponType:       allocatedmgrpb.CouponType_SpecialOffer,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	UserID:           &userID,
}

var specialOfferReq = &npool.CouponReq{
	ID:               &specialOffer.ID,
	CouponType:       &specialOffer.CouponType,
	AppID:            &specialOffer.AppID,
	Value:            &specialOffer.Value,
	ReleasedByUserID: &specialOffer.ReleasedByUserID,
	StartAt:          &specialOffer.StartAt,
	DurationDays:     &specialOffer.DurationDays,
	Message:          &specialOffer.Message,
	UserID:           specialOffer.UserID,
}

func createSpecialOffer(t *testing.T) {
	info, err := CreateCoupon(context.Background(), specialOfferReq)
	if assert.Nil(t, err) {
		specialOffer.CreatedAt = info.CreatedAt
		specialOffer.UpdatedAt = info.UpdatedAt
		assert.Equal(t, specialOffer, info)
	}
}

func updateSpecialOffer(t *testing.T) {
	value := "10.02"

	specialOfferReq.Value = &value
	specialOffer.Value = value

	info, err := UpdateCoupon(context.Background(), specialOfferReq)
	if assert.Nil(t, err) {
		specialOffer.UpdatedAt = info.UpdatedAt
		assert.Equal(t, specialOffer, info)
	}
}

func getSpecialOffer(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		CouponType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(allocatedmgrpb.CouponType_SpecialOffer),
		},
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: specialOffer.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, specialOffer, infos[0])
	}
}

func TestCouponSpecialOffer(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createSpecialOffer", createSpecialOffer)
	t.Run("updateSpecialOffer", updateSpecialOffer)
	t.Run("getSpecialOffer", getSpecialOffer)
}
