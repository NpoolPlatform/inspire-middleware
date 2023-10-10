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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"
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

	ret = npool.Scope{
		ID:                 uuid.NewString(),
		AppID:              coupon.AppID,
		AppGoodID:          uuid.NewString(),
		CouponID:           coupon.ID,
		CouponScope:        types.CouponScope_Whitelist,
		CouponScopeStr:     types.CouponScope_Whitelist.String(),
		CouponName:         coupon.Name,
		CouponType:         coupon.CouponType,
		CouponTypeStr:      coupon.CouponTypeStr,
		CouponDenomination: coupon.Denomination,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithID(&coupon.ID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&ret.CouponName, true),
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

func createScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithCouponID(&ret.CouponID, true),
		WithCouponScope(&ret.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateScope(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateScope(t *testing.T) {
}

func getScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getScopes(t *testing.T) {
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

	infos, _, err := handler.GetScopes(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, &ret, infos[0])
	}
}

func existScopeConds(t *testing.T) {
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

	exist, err := handler.ExistScopeConds(context.Background())
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetScope(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createScope", createScope)
	t.Run("updateScope", updateScope)
	t.Run("getScope", getScope)
	t.Run("getScopes", getScopes)
	t.Run("existScope", existScopeConds)
	t.Run("deleteScope", deleteScope)
}
