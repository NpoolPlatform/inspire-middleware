package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	ret = npool.Coupon{
		ID:                  uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		AppID:               uuid.NewString(),
		Denomination:        decimal.RequireFromString("12.25").String(),
		Circulation:         decimal.RequireFromString("12.25").String(),
		IssuedBy:            uuid.NewString(),
		StartAt:             uint32(time.Now().Unix()),
		DurationDays:        234,
		Message:             uuid.NewString(),
		Name:                uuid.NewString(),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Threshold:           nil,
		UserID:              nil,
		Allocated:           decimal.NewFromInt(0).String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCouponType(&ret.CouponType, true),
		WithAppID(&ret.AppID, true),
		WithDenomination(&ret.Denomination, true),
		WithCirculation(&ret.Circulation, true),
		WithIssuedBy(&ret.IssuedBy, true),
		WithStartAt(&ret.StartAt, true),
		WithDurationDays(&ret.DurationDays, true),
		WithMessage(&ret.Message, true),
		WithName(&ret.Name, true),
		WithCouponScope(&ret.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCoupon(t *testing.T) {
	ret.CouponScope = types.CouponScope_AllGood
	ret.CouponScopeStr = types.CouponScope_AllGood.String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithDenomination(&ret.Denomination, true),
		WithCirculation(&ret.Circulation, true),
		WithStartAt(&ret.StartAt, true),
		WithDurationDays(&ret.DurationDays, true),
		WithMessage(&ret.Message, true),
		WithName(&ret.Name, true),
		WithCouponScope(&ret.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCoupon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCoupons(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		IDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
		CouponType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponType)},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCoupons(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCoupon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCoupon(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCoupon", createCoupon)
	t.Run("updateCoupon", updateCoupon)
	t.Run("getCoupon", getCoupon)
	t.Run("getCoupons", getCoupons)
	t.Run("deleteCoupon", deleteCoupon)
}
