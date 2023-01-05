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
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

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

var userID = uuid.NewString()

var retSpecialOffer = &couponmwpb.Coupon{
	ID:               uuid.NewString(),
	CouponType:       mgrpb.CouponType_SpecialOffer,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	Name:             "Test coupon name",
	Allocated:        "0",
	UserID:           &userID,
}

var reqSpecialOffer = &couponmwpb.CouponReq{
	ID:               &retSpecialOffer.ID,
	CouponType:       &retSpecialOffer.CouponType,
	AppID:            &retSpecialOffer.AppID,
	Value:            &retSpecialOffer.Value,
	ReleasedByUserID: &retSpecialOffer.ReleasedByUserID,
	StartAt:          &retSpecialOffer.StartAt,
	DurationDays:     &retSpecialOffer.DurationDays,
	Message:          &retSpecialOffer.Message,
	Name:             &retSpecialOffer.Name,
	UserID:           retSpecialOffer.UserID,
}

var retSpecialOffer1 = &npool.Coupon{
	ID:           uuid.NewString(),
	CouponType:   mgrpb.CouponType_SpecialOffer,
	AppID:        retSpecialOffer.AppID,
	CouponID:     retSpecialOffer.ID,
	UserID:       uuid.NewString(),
	Value:        retSpecialOffer.Value,
	Circulation:  retSpecialOffer.Circulation,
	DurationDays: retSpecialOffer.DurationDays,
	CouponName:   retSpecialOffer.Name,
	Message:      retSpecialOffer.Message,
}

var reqSpecialOffer1 = &npool.CouponReq{
	ID:         &retSpecialOffer1.ID,
	CouponType: &retSpecialOffer.CouponType,
	AppID:      &retSpecialOffer1.AppID,
	CouponID:   &retSpecialOffer1.CouponID,
	UserID:     &retSpecialOffer1.UserID,
}

func createSpecialOffer(t *testing.T) {
	_, err := coupon.CreateCoupon(context.Background(), reqSpecialOffer)
	assert.Nil(t, err)

	info, err := CreateCoupon(context.Background(), reqSpecialOffer1)
	if assert.Nil(t, err) {
		retSpecialOffer1.CreatedAt = info.CreatedAt
		retSpecialOffer1.UpdatedAt = info.UpdatedAt
		retSpecialOffer1.StartAt = info.StartAt
		retSpecialOffer1.EndAt = info.EndAt
		retSpecialOffer1.Valid = info.Valid
		retSpecialOffer1.CouponName = info.CouponName
		assert.Equal(t, retSpecialOffer1, info)
	}
}

func updateSpecialOffer(t *testing.T) {
	used := true
	orderID := uuid.NewString()

	retSpecialOffer1.Used = used
	retSpecialOffer1.UsedByOrderID = &orderID

	reqSpecialOffer1.Used = &used
	reqSpecialOffer1.UsedByOrderID = &orderID

	info, err := UpdateCoupon(context.Background(), reqSpecialOffer1)
	if assert.Nil(t, err) {
		retSpecialOffer1.UpdatedAt = info.UpdatedAt
		retSpecialOffer1.UsedAt = info.UsedAt
		retSpecialOffer1.CouponName = info.CouponName
		assert.Equal(t, retSpecialOffer1, info)
	}
}

func getSpecialOfferCoupon(t *testing.T) {
	info, err := GetCoupon(context.Background(), retSpecialOffer1.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, retSpecialOffer1, info)
	}
}

func getManySpecialOfferCoupons(t *testing.T) {
	infos, err := GetManyCoupons(context.Background(), []string{retSpecialOffer1.ID})
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		retSpecialOffer1.CouponName = infos[0].CouponName
		assert.Equal(t, retSpecialOffer1, infos[0])
	}
}

func getSpecialOfferCoupons(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.UserID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		retSpecialOffer1.CouponName = infos[0].CouponName
		assert.Equal(t, retSpecialOffer1, infos[0])
	}
}

func getSpecialOfferCouponOnly(t *testing.T) {
	info, err := GetCouponOnly(context.Background(), &mgrpb.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.ID,
		},
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.UserID,
		},
		UsedByOrderID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: retSpecialOffer1.GetUsedByOrderID(),
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, retSpecialOffer1, info)
	}
}

func TestCouponSpecialOffer(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createSpecialOffer", createSpecialOffer)
	t.Run("updateSpecialOffer", updateSpecialOffer)
	t.Run("getSpecialOfferCoupon", getSpecialOfferCoupon)
	t.Run("getManySpecialOfferCoupons", getManySpecialOfferCoupons)
	t.Run("getSpecialOfferCoupons", getSpecialOfferCoupons)
	t.Run("getSpecialOfferCouponOnly", getSpecialOfferCouponOnly)
}
