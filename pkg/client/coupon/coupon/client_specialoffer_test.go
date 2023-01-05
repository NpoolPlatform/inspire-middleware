//nolint:dupl
package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	allocatedmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/coupon"

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

var retSpecialOffer = &npool.Coupon{
	ID:               uuid.NewString(),
	CouponType:       allocatedmgrpb.CouponType_SpecialOffer,
	AppID:            uuid.NewString(),
	Value:            "10.01",
	ReleasedByUserID: uuid.NewString(),
	StartAt:          uint32(time.Now().Unix()),
	DurationDays:     30,
	Message:          "Test coupon message",
	Name:             "Test coupon name",
	UserID:           &userID,
}

var reqSpecialOffer = &npool.CouponReq{
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

func createSpecialOffer(t *testing.T) {
	info, err := CreateCoupon(context.Background(), reqSpecialOffer)
	if assert.Nil(t, err) {
		retSpecialOffer.CreatedAt = info.CreatedAt
		retSpecialOffer.UpdatedAt = info.UpdatedAt
		retSpecialOffer.Name = info.Name
		assert.Equal(t, retSpecialOffer, info)
	}
}

func updateSpecialOffer(t *testing.T) {
	value := "10.02"

	reqSpecialOffer.Value = &value

	retSpecialOffer.Value = value

	info, err := UpdateCoupon(context.Background(), reqSpecialOffer)
	fmt.Println("*****************err")
	fmt.Println(err)
	if assert.Nil(t, err) {
		retSpecialOffer.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retSpecialOffer, info)
	}

	info, err = UpdateCoupon(context.Background(), reqSpecialOffer)
	if assert.Nil(t, err) {
		retSpecialOffer.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retSpecialOffer, info)
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
			Value: retSpecialOffer.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, retSpecialOffer, infos[0])
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
	t.Run("getSpecialOffer", getSpecialOffer)
}
