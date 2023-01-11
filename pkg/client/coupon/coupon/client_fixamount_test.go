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

var ret = &npool.Coupon{
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

var req = &npool.CouponReq{
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

func create(t *testing.T) {
	info, err := CreateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	value := "10.02"
	circulation := "200.4"

	req.Value = &value
	req.Circulation = &circulation

	ret.Value = value
	ret.Circulation = circulation

	info, err := UpdateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	allocated := "2"

	req.Allocated = &allocated
	ret.Allocated = allocated

	info, err = UpdateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}

	allocated = "4"

	req.Allocated = &allocated
	ret.Allocated = "6"
	req.Value = nil
	req.Circulation = nil

	info, err = UpdateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func get(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		CouponType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(allocatedmgrpb.CouponType_FixAmount),
		},
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func TestCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("create", create)
	t.Run("update", update)
	t.Run("get", get)
}
