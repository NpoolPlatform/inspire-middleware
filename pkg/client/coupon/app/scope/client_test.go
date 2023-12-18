package scope

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	couponmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon"
	scopemwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon/scope"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"
	scopemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

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
		Allocated:                     "0",
		CouponScope:                   types.CouponScope_Whitelist,
		CouponScopeStr:                types.CouponScope_Whitelist.String(),
		Threshold:                     decimal.NewFromInt(0).String(),
		CashableProbabilityPerMillion: decimal.RequireFromString("0.0001").String(),
	}

	scope = scopemwpb.Scope{
		EntID:              uuid.NewString(),
		GoodID:             uuid.NewString(),
		CouponID:           coupon.EntID,
		CouponType:         coupon.CouponType,
		CouponTypeStr:      coupon.CouponTypeStr,
		CouponScope:        coupon.CouponScope,
		CouponScopeStr:     coupon.CouponScopeStr,
		CouponName:         coupon.Name,
		CouponDenomination: coupon.Denomination,
		CouponCirculation:  coupon.Circulation,
	}

	ret = npool.Scope{
		EntID:              uuid.NewString(),
		AppID:              coupon.AppID,
		AppGoodID:          uuid.NewString(),
		CouponID:           scope.CouponID,
		CouponName:         coupon.Name,
		CouponType:         scope.CouponType,
		CouponTypeStr:      scope.CouponType.String(),
		CouponScope:        scope.CouponScope,
		CouponScopeStr:     scope.CouponScope.String(),
		CouponDenomination: scope.CouponDenomination,
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

	_scope, err := scopemwcli.CreateScope(context.Background(), &scopemwpb.ScopeReq{
		EntID:       &scope.EntID,
		GoodID:      &scope.GoodID,
		CouponID:    &coupon.EntID,
		CouponScope: &coupon.CouponScope,
	})
	if assert.Nil(t, err) {
		scope.ID = _scope.ID
		scope.CreatedAt = _scope.CreatedAt
		scope.UpdatedAt = _scope.UpdatedAt
		assert.Equal(t, &scope, _scope)
	}

	return func(*testing.T) {
		_, _ = couponmwcli.DeleteCoupon(context.Background(), coupon.ID)
		_, _ = scopemwcli.DeleteScope(context.Background(), scope.ID)
	}
}

func createAppGoodScope(t *testing.T) {
	info, err := CreateAppGoodScope(context.Background(), &npool.ScopeReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		AppGoodID:   &ret.AppGoodID,
		CouponID:    &ret.CouponID,
		CouponScope: &ret.CouponScope,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getAppGoodScope(t *testing.T) {
	info, err := GetAppGoodScope(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getAppGoodScopes(t *testing.T) {
	infos, total, err := GetAppGoodScopes(context.Background(), &npool.Conds{
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, &ret, infos[0])
	}
}

func existAppGoodScopeConds(t *testing.T) {
	conds := &npool.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}
	exist, err := ExistAppGoodScopeConds(context.Background(), conds)
	assert.Nil(t, err)
	assert.True(t, exist)
}

func verifyCouponScope(t *testing.T) {
	err := VerifyCouponScopes(context.Background(), []*npool.ScopeReq{{
		AppID:       &ret.AppID,
		GoodID:      &scope.GoodID,
		AppGoodID:   &ret.AppGoodID,
		CouponID:    &ret.CouponID,
		CouponScope: &ret.CouponScope,
	}})
	assert.Nil(t, err)
}

func deleteAppGoodScope(t *testing.T) {
	info, err := DeleteAppGoodScope(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetAppGoodScope(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
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

	t.Run("createAppGoodScope", createAppGoodScope)
	t.Run("getAppGoodScope", getAppGoodScope)
	t.Run("getAppGoodScopes", getAppGoodScopes)
	t.Run("existAppGoodScopeConds", existAppGoodScopeConds)
	t.Run("verifyCouponScope", verifyCouponScope)
	t.Run("deleteAppGoodScope", deleteAppGoodScope)
}
