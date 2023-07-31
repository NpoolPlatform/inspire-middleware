package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	coupon "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	ID:                  uuid.NewString(),
	CouponType:          types.CouponType_Discount,
	CouponTypeStr:       types.CouponType_Discount.String(),
	AppID:               uuid.NewString(),
	Denomination:        "10.01",
	Circulation:         "100.1",
	IssuedBy:            uuid.NewString(),
	StartAt:             uint32(time.Now().Unix()),
	DurationDays:        30,
	Message:             "Test coupon message",
	Name:                "Test coupon name",
	Allocated:           "0",
	CouponConstraint:    types.CouponConstraint_Normal,
	CouponConstraintStr: types.CouponConstraint_Normal.String(),
}

var req = &couponmwpb.CouponReq{
	ID:           &ret.ID,
	CouponType:   &ret.CouponType,
	AppID:        &ret.AppID,
	Denomination: &ret.Denomination,
	Circulation:  &ret.Circulation,
	IssuedBy:     &ret.IssuedBy,
	StartAt:      &ret.StartAt,
	DurationDays: &ret.DurationDays,
	Message:      &ret.Message,
	Name:         &ret.Name,
}

var ret1 = &npool.Coupon{
	ID:                  uuid.NewString(),
	CouponType:          types.CouponType_Discount,
	CouponTypeStr:       types.CouponType_Discount.String(),
	AppID:               ret.AppID,
	CouponID:            ret.ID,
	UserID:              uuid.NewString(),
	Denomination:        ret.Denomination,
	Circulation:         ret.Circulation,
	DurationDays:        ret.DurationDays,
	CouponName:          ret.Name,
	Message:             ret.Message,
	CouponConstraint:    types.CouponConstraint_Normal,
	CouponConstraintStr: types.CouponConstraint_Normal.String(),
}

var req1 = &npool.CouponReq{
	ID:       &ret1.ID,
	AppID:    &ret1.AppID,
	CouponID: &ret1.CouponID,
	UserID:   &ret1.UserID,
}

func createCoupon(t *testing.T) {
	_, err := coupon.CreateCoupon(context.Background(), req)
	assert.Nil(t, err)

	info, err := CreateCoupon(context.Background(), req1)
	if assert.Nil(t, err) {
		ret1.CreatedAt = info.CreatedAt
		ret1.UpdatedAt = info.UpdatedAt
		ret1.StartAt = info.StartAt
		ret1.EndAt = info.EndAt
		ret1.Valid = info.Valid
		ret.Allocated = decimal.NewFromInt(1).String()
		ret1.Allocated = decimal.NewFromInt(1).String()
		assert.Equal(t, ret1, info)
	}
}

func updateCoupon(t *testing.T) {
	used := true
	orderID := uuid.NewString()

	ret1.Used = used
	ret1.UsedByOrderID = &orderID

	req1.Used = &used
	req1.UsedByOrderID = &orderID

	info, err := UpdateCoupon(context.Background(), req1)
	if assert.Nil(t, err) {
		ret1.UpdatedAt = info.UpdatedAt
		ret1.UsedAt = info.UsedAt
		assert.Equal(t, ret1, info)
	}
}

func getCouponCoupon(t *testing.T) {
	info, err := GetCoupon(context.Background(), ret1.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret1, info)
	}
}

func getCouponCoupons(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret1.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret1.UserID},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, ret1, infos[0])
	}
}

func getCouponCouponOnly(t *testing.T) {
	info, err := GetCouponOnly(context.Background(), &npool.Conds{
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret1.ID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret1.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret1.UserID},
		UsedByOrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret1.GetUsedByOrderID()},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret1, info)
	}
}

func TestCouponCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCoupon", createCoupon)
	t.Run("updateCoupon", updateCoupon)
	t.Run("getCouponCoupon", getCouponCoupon)
	t.Run("getCouponCoupons", getCouponCoupons)
	t.Run("getCouponCouponOnly", getCouponCouponOnly)
}
