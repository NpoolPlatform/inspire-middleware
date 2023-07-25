//nolint:dupl
package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	coupon "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon/coupon"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var retDiscount = &couponmwpb.Coupon{
	ID:               uuid.NewString(),
	CouponType:       mgrpb.CouponType_Discount,
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

var reqDiscount = &couponmwpb.CouponReq{
	ID:               &retDiscount.ID,
	CouponType:       &retDiscount.CouponType,
	AppID:            &retDiscount.AppID,
	Value:            &retDiscount.Value,
	Circulation:      &retDiscount.Circulation,
	ReleasedByUserID: &retDiscount.ReleasedByUserID,
	StartAt:          &retDiscount.StartAt,
	DurationDays:     &retDiscount.DurationDays,
	Message:          &retDiscount.Message,
	Name:             &retDiscount.Name,
}

var retDiscount1 = &npool.Coupon{
	ID:           uuid.NewString(),
	CouponType:   mgrpb.CouponType_Discount,
	AppID:        retDiscount.AppID,
	CouponID:     retDiscount.ID,
	UserID:       uuid.NewString(),
	Value:        retDiscount.Value,
	Circulation:  retDiscount.Circulation,
	DurationDays: retDiscount.DurationDays,
	CouponName:   retDiscount.Name,
	Message:      retDiscount.Message,
}

var reqDiscount1 = &npool.CouponReq{
	ID:         &retDiscount1.ID,
	CouponType: &retDiscount.CouponType,
	AppID:      &retDiscount1.AppID,
	CouponID:   &retDiscount1.CouponID,
	UserID:     &retDiscount1.UserID,
}

func createDiscount(t *testing.T) {
	_, err := coupon.CreateCoupon(context.Background(), reqDiscount)
	assert.Nil(t, err)

	info, err := CreateCoupon(context.Background(), reqDiscount1)
	if assert.Nil(t, err) {
		retDiscount1.CreatedAt = info.CreatedAt
		retDiscount1.UpdatedAt = info.UpdatedAt
		retDiscount1.StartAt = info.StartAt
		retDiscount1.EndAt = info.EndAt
		retDiscount1.Valid = info.Valid
		assert.Equal(t, retDiscount1, info)
	}
}

func updateDiscount(t *testing.T) {
	used := true
	orderID := uuid.NewString()

	retDiscount1.Used = used
	retDiscount1.UsedByOrderID = &orderID

	reqDiscount1.Used = &used
	reqDiscount1.UsedByOrderID = &orderID

	info, err := UpdateCoupon(context.Background(), reqDiscount1)
	if assert.Nil(t, err) {
		retDiscount1.UpdatedAt = info.UpdatedAt
		retDiscount1.UsedAt = info.UsedAt
		assert.Equal(t, retDiscount1, info)
	}
}

func getDiscountCoupon(t *testing.T) {
	info, err := GetCoupon(context.Background(), retDiscount1.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, retDiscount1, info)
	}
}

func getManyDiscountCoupons(t *testing.T) {
	infos, err := GetManyCoupons(context.Background(), []string{retDiscount1.ID})
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, retDiscount1, infos[0])
	}
}

func getDiscountCoupons(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.UserID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, retDiscount1, infos[0])
	}
}

func getDiscountCouponOnly(t *testing.T) {
	info, err := GetCouponOnly(context.Background(), &mgrpb.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.ID,
		},
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.UserID,
		},
		UsedByOrderID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retDiscount1.GetUsedByOrderID(),
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, retDiscount1, info)
	}
}

func TestCouponDiscount(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createDiscount", createDiscount)
	t.Run("updateDiscount", updateDiscount)
	t.Run("getDiscountCoupon", getDiscountCoupon)
	t.Run("getManyDiscountCoupons", getManyDiscountCoupons)
	t.Run("getDiscountCoupons", getDiscountCoupons)
	t.Run("getDiscountCouponOnly", getDiscountCouponOnly)
}
