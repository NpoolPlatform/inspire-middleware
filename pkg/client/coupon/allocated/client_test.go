package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	couponmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

var coupon = &couponmwpb.Coupon{
	EntID:               uuid.NewString(),
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
	CouponScope:         types.CouponScope_Whitelist,
	CouponScopeStr:      types.CouponScope_Whitelist.String(),
}

var ret = &npool.Coupon{
	EntID:               uuid.NewString(),
	CouponType:          types.CouponType_Discount,
	CouponTypeStr:       types.CouponType_Discount.String(),
	AppID:               coupon.AppID,
	CouponID:            coupon.EntID,
	UserID:              uuid.NewString(),
	Denomination:        coupon.Denomination,
	Circulation:         coupon.Circulation,
	DurationDays:        coupon.DurationDays,
	CouponName:          coupon.Name,
	Message:             coupon.Message,
	CouponConstraint:    types.CouponConstraint_Normal,
	CouponConstraintStr: types.CouponConstraint_Normal.String(),
	CouponScope:         types.CouponScope_Whitelist,
	CouponScopeStr:      types.CouponScope_Whitelist.String(),
	Valid:               true,
	Allocated:           "1",
}

func createCoupon(t *testing.T) {
	info1, err := couponmwcli.CreateCoupon(context.Background(), &couponmwpb.CouponReq{
		EntID:        &coupon.EntID,
		CouponType:   &coupon.CouponType,
		AppID:        &coupon.AppID,
		Denomination: &coupon.Denomination,
		Circulation:  &coupon.Circulation,
		IssuedBy:     &coupon.IssuedBy,
		StartAt:      &coupon.StartAt,
		DurationDays: &coupon.DurationDays,
		Message:      &coupon.Message,
		Name:         &coupon.Name,
	})
	if assert.Nil(t, err) {
		coupon.ID = info1.ID
		coupon.CreatedAt = info1.CreatedAt
		coupon.UpdatedAt = info1.UpdatedAt
		assert.Equal(t, coupon, info1)
	}

	info, err := CreateCoupon(context.Background(), &npool.CouponReq{
		EntID:    &ret.EntID,
		AppID:    &ret.AppID,
		CouponID: &ret.CouponID,
		UserID:   &ret.UserID,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.StartAt = info.StartAt
		ret.EndAt = info.EndAt
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
		coupon.Allocated = info.Allocated
	}
}

func updateCoupon(t *testing.T) {
	used := true
	orderID := uuid.NewString()

	ret.Used = used
	ret.UsedByOrderID = orderID

	info, err := UpdateCoupon(context.Background(), &npool.CouponReq{
		ID:            &ret.ID,
		Used:          &ret.Used,
		UsedByOrderID: &ret.UsedByOrderID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.UsedAt = info.UsedAt
		assert.Equal(t, ret, info)
	}
}

func getCouponCoupon(t *testing.T) {
	info, err := GetCoupon(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getCouponCoupons(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, ret, infos[0])
	}
}

func getCouponCouponOnly(t *testing.T) {
	info, err := GetCouponOnly(context.Background(), &npool.Conds{
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		UsedByOrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GetUsedByOrderID()},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCoupon", createCoupon)
	t.Run("updateCoupon", updateCoupon)
	t.Run("getCouponCoupon", getCouponCoupon)
	t.Run("getCouponCoupons", getCouponCoupons)
	t.Run("getCouponCouponOnly", getCouponCouponOnly)
}
