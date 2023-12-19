package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	couponmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var (
	coupon = couponmwpb.Coupon{
		EntID:                         uuid.NewString(),
		AppID:                         uuid.NewString(),
		Name:                          uuid.NewString(),
		Message:                       uuid.NewString(),
		CouponType:                    types.CouponType_FixAmount,
		CouponTypeStr:                 types.CouponType_FixAmount.String(),
		Denomination:                  decimal.RequireFromString("100").String(),
		Circulation:                   decimal.RequireFromString("100").String(),
		DurationDays:                  365,
		IssuedBy:                      uuid.NewString(),
		StartAt:                       uint32(time.Now().Unix()),
		EndAt:                         uint32(time.Now().Add(24 * time.Hour).Unix()),
		CouponConstraint:              types.CouponConstraint_Normal,
		CouponConstraintStr:           types.CouponConstraint_Normal.String(),
		CouponScope:                   types.CouponScope_Whitelist,
		CouponScopeStr:                types.CouponScope_Whitelist.String(),
		Allocated:                     decimal.NewFromInt(0).String(),
		Threshold:                     decimal.NewFromInt(0).String(),
		CashableProbabilityPerMillion: decimal.RequireFromString("0.0001").String(),
	}

	ret = npool.CouponCoin{
		EntID:              uuid.NewString(),
		AppID:              coupon.AppID,
		CouponID:           coupon.EntID,
		CouponName:         coupon.Name,
		CouponDenomination: coupon.Denomination,
		CoinTypeID:         uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	info, err := couponmwcli.CreateCoupon(context.Background(), &couponmwpb.CouponReq{
		EntID:                         &coupon.EntID,
		AppID:                         &coupon.AppID,
		Name:                          &coupon.Name,
		Message:                       &coupon.Message,
		CouponType:                    &coupon.CouponType,
		Denomination:                  &coupon.Denomination,
		Circulation:                   &coupon.Circulation,
		DurationDays:                  &coupon.DurationDays,
		IssuedBy:                      &coupon.IssuedBy,
		StartAt:                       &coupon.StartAt,
		EndAt:                         &coupon.EndAt,
		CouponScope:                   &coupon.CouponScope,
		CashableProbabilityPerMillion: &coupon.CashableProbabilityPerMillion,
	})
	if assert.Nil(t, err) {
		coupon.ID = info.ID
		coupon.CreatedAt = info.CreatedAt
		coupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &coupon, info)
	}

	return func(*testing.T) {
		_, _ = couponmwcli.DeleteCoupon(context.Background(), coupon.ID)
	}
}

func createCouponCoin(t *testing.T) {
	info, err := CreateCouponCoin(context.Background(), &npool.CouponCoinReq{
		EntID:      &ret.EntID,
		AppID:      &ret.AppID,
		CoinTypeID: &ret.CoinTypeID,
		CouponID:   &ret.CouponID,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getCouponCoin(t *testing.T) {
	info, err := GetCouponCoin(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getCouponCoins(t *testing.T) {
	infos, total, err := GetCouponCoins(context.Background(), &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		CouponID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, &ret, infos[0])
	}
}

func existCouponCoinConds(t *testing.T) {
	conds := &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CouponID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}
	exist, err := ExistCouponCoinConds(context.Background(), conds)
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteCouponCoin(t *testing.T) {
	info, err := DeleteCouponCoin(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetCouponCoin(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCouponCoin(t *testing.T) {
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCouponCoin", createCouponCoin)
	t.Run("getCouponCoin", getCouponCoin)
	t.Run("getCouponCoins", getCouponCoins)
	t.Run("existCouponCoinConds", existCouponCoinConds)
	t.Run("deleteCouponCoin", deleteCouponCoin)
}
