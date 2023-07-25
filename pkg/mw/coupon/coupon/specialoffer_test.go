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

var userID = uuid.NewString()
var sret = &npool.Coupon{
	ID:               uuid.NewString(),
	CouponType:       allocatedmgrpb.CouponType_SpecialOffer,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	Circulation:      "10.01",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	UserID:           &userID,
	Allocated:        "10.01",
}

var sreq = &npool.CouponReq{
	ID:               &sret.ID,
	CouponType:       &sret.CouponType,
	AppID:            &sret.AppID,
	Value:            &sret.Value,
	ReleasedByUserID: &sret.ReleasedByUserID,
	StartAt:          &sret.StartAt,
	DurationDays:     &sret.DurationDays,
	Message:          &sret.Message,
	UserID:           sret.UserID,
}

func createSpecialOffer(t *testing.T) {
	info, err := CreateCoupon(context.Background(), sreq)
	if assert.Nil(t, err) {
		sret.CreatedAt = info.CreatedAt
		sret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, sret, info)
	}
}

func updateSpecialOffer(t *testing.T) {
	value := "10.02"

	sreq.Value = &value
	sret.Value = value
	sret.Circulation = value
	sret.Allocated = value

	info, err := UpdateCoupon(context.Background(), sreq)
	if assert.Nil(t, err) {
		sret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, sret, info)
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
			Value: sret.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, sret, infos[0])
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
