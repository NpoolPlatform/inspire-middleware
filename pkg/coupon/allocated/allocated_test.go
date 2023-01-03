package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

	coupon "github.com/NpoolPlatform/inspire-middleware/pkg/coupon/coupon"

	// cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	// commonpb "github.com/NpoolPlatform/message/npool"

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

var ret = &couponmwpb.Coupon{
	ID:               uuid.NewString(),
	CouponType:       allocatedmgrpb.CouponType_FixAmount,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	Circulation:      "100.1",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	Name:             "Test coupon name",
	Allocated:        "0",
}

var req = &couponmwpb.CouponReq{
	ID:               &ret.ID,
	CouponType:       &ret.CouponType,
	AppID:            &ret.AppID,
	Value:            &ret.Value,
	Circulation:      &ret.Circulation,
	ReleasedByUserID: &ret.ReleasedByUserID,
	StartAt:          &ret.StartAt,
	DurationDays:     &ret.DurationDays,
	Message:          &ret.Message,
	Name:             &ret.Name,
}

var ret1 = &npool.Coupon{
	ID:           uuid.NewString(),
	CouponType:   allocatedmgrpb.CouponType_FixAmount,
	AppID:        uuid.NewString(),
	CouponID:     ret.ID,
	UserID:       uuid.NewString(),
	Value:        ret.Value,
	Circulation:  ret.Circulation,
	DurationDays: ret.DurationDays,
	CouponName:   ret.Name,
	Message:      ret.Message,
}

var req1 = &npool.CouponReq{
	ID:         &ret1.ID,
	CouponType: &ret.CouponType,
	AppID:      &ret1.AppID,
	CouponID:   &ret1.CouponID,
	UserID:     &ret1.UserID,
}

func create(t *testing.T) {
	_, err := coupon.CreateCoupon(context.Background(), req)
	assert.Nil(t, err)

	info, err := CreateCoupon(context.Background(), req1)
	if assert.Nil(t, err) {
		ret1.CreatedAt = info.CreatedAt
		ret1.UpdatedAt = info.UpdatedAt
		ret1.StartAt = info.StartAt
		ret1.EndAt = info.EndAt
		ret1.Valid = info.Valid
		assert.Equal(t, ret1, info)
	}
}

func getCoupon(t *testing.T) {
}

func TestCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("getCoupon", getCoupon)
}
