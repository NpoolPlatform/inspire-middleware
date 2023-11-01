package scope

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/scope"
	scopemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
		ID:                  uuid.NewString(),
		AppID:               uuid.NewString(),
		Name:                uuid.NewString(),
		Message:             uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		Denomination:        decimal.RequireFromString("100").String(),
		Circulation:         decimal.RequireFromString("100").String(),
		DurationDays:        365,
		IssuedBy:            uuid.NewString(),
		StartAt:             uint32(time.Now().Unix()),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		Allocated:           "0",
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
	}

	scope = scopemwpb.Scope{
		ID:                 uuid.NewString(),
		GoodID:             uuid.NewString(),
		CouponID:           coupon.ID,
		CouponType:         coupon.CouponType,
		CouponTypeStr:      coupon.CouponTypeStr,
		CouponScope:        coupon.CouponScope,
		CouponScopeStr:     coupon.CouponScopeStr,
		CouponName:         coupon.Name,
		CouponDenomination: coupon.Denomination,
	}

	ret = npool.Scope{
		ID:                 uuid.NewString(),
		AppID:              uuid.NewString(),
		ScopeID:            scope.ID,
		GoodID:             scope.GoodID,
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
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithID(&coupon.ID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&scope.CouponName, true),
		coupon1.WithMessage(&coupon.Message, true),
		coupon1.WithCouponType(&coupon.CouponType, true),
		coupon1.WithDenomination(&coupon.Denomination, true),
		coupon1.WithCouponScope(&coupon.CouponScope, true),
		coupon1.WithCirculation(&coupon.Circulation, true),
		coupon1.WithDurationDays(&coupon.DurationDays, true),
		coupon1.WithIssuedBy(&coupon.IssuedBy, true),
		coupon1.WithStartAt(&coupon.StartAt, true),
	)
	assert.Nil(t, err)

	coup, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		coupon.CreatedAt = coup.CreatedAt
		coupon.UpdatedAt = coup.UpdatedAt
		assert.Equal(t, &coupon, coup)
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithScopeID(&ret.ScopeID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppGoodScopes(t *testing.T) {
	conds := &npool.Conds{
		ID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAppGoodScopes(context.Background())
	if assert.Nil(t, err) {
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
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistAppGoodScopeConds(context.Background())
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAppGoodScope(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createAppGoodScope", createAppGoodScope)
	t.Run("getAppGoodScope", getAppGoodScope)
	t.Run("getAppGoodScopes", getAppGoodScopes)
	t.Run("existAppGoodScope", existAppGoodScopeConds)
	t.Run("deleteAppGoodScope", deleteAppGoodScope)
}
