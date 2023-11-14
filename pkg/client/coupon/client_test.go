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

func createDiscount(t *testing.T) {
	info, err := CreateCoupon(context.Background(), &npool.CouponReq{
		EntID:            &ret.EntID,
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
		CouponScope:      &ret.CouponScope,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateDiscount(t *testing.T) {
	ret.Denomination = "10.02"
	ret.Circulation = "200.4"
	ret.CouponScope = types.CouponScope_AllGood
	ret.CouponScopeStr = types.CouponScope_AllGood.String()
	ret.Allocated = "1"

	info, err := UpdateCoupon(context.Background(), &npool.CouponReq{
		ID:           &ret.ID,
		Denomination: &ret.Denomination,
		Circulation:  &ret.Circulation,
		CouponScope:  &ret.CouponScope,
	})
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
