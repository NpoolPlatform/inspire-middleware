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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

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

var retDiscount = &npool.Coupon{
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
	Allocated:        "0",
}

var reqDiscount = &npool.CouponReq{
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

func createDiscount(t *testing.T) {
	info, err := CreateCoupon(context.Background(), reqDiscount)
	if assert.Nil(t, err) {
		retDiscount.CreatedAt = info.CreatedAt
		retDiscount.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retDiscount, info)
	}
}

func updateDiscount(t *testing.T) {
	value := "10.02" //nolint
	circulation := "200.4"

	reqDiscount.Value = &value
	reqDiscount.Circulation = &circulation

	retDiscount.Value = value
	retDiscount.Circulation = circulation
	info, err := UpdateCoupon(context.Background(), reqDiscount)
	if assert.Nil(t, err) {
		retDiscount.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retDiscount, info)
	}

	allocated := "2"

	reqDiscount.Allocated = &allocated
	retDiscount.Allocated = allocated

	info, err = UpdateCoupon(context.Background(), reqDiscount)
	if assert.Nil(t, err) {
		retDiscount.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retDiscount, info)
	}

	allocated = "4"

	reqDiscount.Allocated = &allocated
	retDiscount.Allocated = "6"
	reqDiscount.Value = nil
	reqDiscount.Circulation = nil

	info, err = UpdateCoupon(context.Background(), reqDiscount)
	if assert.Nil(t, err) {
		retDiscount.UpdatedAt = info.UpdatedAt
		assert.Equal(t, retDiscount, info)
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
			Value: retDiscount.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, retDiscount, infos[0])
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
	t.Run("getDiscount", getDiscount)
}
