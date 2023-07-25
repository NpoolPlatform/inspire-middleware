package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

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

var discountRet = &npool.Coupon{
	ID:               uuid.NewString(),
	CouponType:       allocatedmgrpb.CouponType_Discount,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	Circulation:      "100.1",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	Name:             "Test coupon name",
}

var discountRetReq = &npool.CouponReq{
	ID:               &discountRet.ID,
	CouponType:       &discountRet.CouponType,
	AppID:            &discountRet.AppID,
	Value:            &discountRet.Value,
	Circulation:      &discountRet.Circulation,
	ReleasedByUserID: &discountRet.ReleasedByUserID,
	StartAt:          &discountRet.StartAt,
	DurationDays:     &discountRet.DurationDays,
	Message:          &discountRet.Message,
	Name:             &discountRet.Name,
}

func createDiscount(t *testing.T) {
	info, err := CreateCoupon(context.Background(), discountRetReq)
	if assert.Nil(t, err) {
		discountRet.CreatedAt = info.CreatedAt
		discountRet.UpdatedAt = info.UpdatedAt
		discountRet.Allocated = info.Allocated
		assert.Equal(t, discountRet, info)
	}
}

func updateDiscount(t *testing.T) {
	value := "10.02" //nolint
	circulation := "200.4"

	discountRetReq.Value = &value
	discountRetReq.Circulation = &circulation

	discountRet.Value = value
	discountRet.Circulation = circulation

	info, err := UpdateCoupon(context.Background(), discountRetReq)
	if assert.Nil(t, err) {
		discountRet.UpdatedAt = info.UpdatedAt
		discountRet.Allocated = info.Allocated
		assert.Equal(t, discountRet, info)
	}

	allocated := "2"

	discountRetReq.Allocated = &allocated
	discountRet.Allocated = allocated

	info, err = UpdateCoupon(context.Background(), discountRetReq)
	if assert.Nil(t, err) {
		discountRet.Allocated = info.Allocated
		discountRet.UpdatedAt = info.UpdatedAt
		assert.Equal(t, discountRet, info)
	}

	allocated = "4"

	discountRetReq.Allocated = &allocated
	discountRet.Allocated = "6"

	info, err = UpdateCoupon(context.Background(), discountRetReq)
	if assert.Nil(t, err) {
		discountRet.UpdatedAt = info.UpdatedAt
		assert.Equal(t, discountRet, info)
	}
}

func getDiscount(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		CouponType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(allocatedmgrpb.CouponType_Discount),
		},
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: discountRet.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, discountRet, infos[0])
	}
}

func TestCouponDiscount(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createDiscount", createDiscount)
	t.Run("updateDiscount", updateDiscount)
	t.Run("getDiscount", getDiscount)
}
