package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

var ret = &npool.Coupon{
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

var req = &npool.CouponReq{
	ID:               &ret.ID,
	CouponType:       &ret.CouponType,
	AppID:            &ret.AppID,
	Denomination:     &ret.Denomination,
	Circulation:      &ret.Circulation,
	IssuedBy:         &ret.IssuedBy,
	StartAt:          &ret.StartAt,
	DurationDays:     &ret.DurationDays,
	Message:          &ret.Message,
	Name:             &ret.Name,
	CouponConstraint: &ret.CouponConstraint,
}

func createDiscount(t *testing.T) {
	info, err := CreateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateDiscount(t *testing.T) {
	denomination := "10.02"
	circulation := "200.4"

	req.Denomination = &denomination
	req.Circulation = &circulation

	ret.Denomination = denomination
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
	req.Denomination = nil
	req.Circulation = nil

	info, err = UpdateCoupon(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getDiscount(t *testing.T) {
	infos, total, err := GetCoupons(context.Background(), &npool.Conds{
		CouponType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponType)},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createDiscount", createDiscount)
	t.Run("updateDiscount", updateDiscount)
	t.Run("getDiscount", getDiscount)
}
