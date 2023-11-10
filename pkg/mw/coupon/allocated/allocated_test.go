package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
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
		UserID:              uuid.NewString(),
		Denomination:        decimal.RequireFromString("10").String(),
		Circulation:         decimal.RequireFromString("100").String(),
		Allocated:           decimal.RequireFromString("10").String(),
		StartAt:             uint32(time.Now().Unix()),
		DurationDays:        27,
		CouponID:            uuid.NewString(),
		CouponName:          uuid.NewString(),
		Message:             uuid.NewString(),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Valid:               true,
	}

	ret1 = npool.Coupon{
		ID:                  uuid.NewString(),
		AppID:               ret.AppID,
		CouponID:            ret.CouponID,
		CouponName:          ret.CouponName,
		Message:             ret.Message,
		CouponType:          ret.CouponType,
		CouponTypeStr:       ret.CouponType.String(),
		UserID:              uuid.NewString(),
		Denomination:        decimal.RequireFromString("10").String(),
		Circulation:         decimal.RequireFromString("100").String(),
		Allocated:           decimal.RequireFromString("20").String(),
		StartAt:             uint32(time.Now().Unix()),
		DurationDays:        27,
		CouponConstraint:    ret.CouponConstraint,
		CouponConstraintStr: ret.CouponConstraint.String(),
		CouponScope:         ret.CouponScope,
		CouponScopeStr:      ret.CouponScope.String(),
		Valid:               true,
		Used:                false,
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.CouponTypeStr = ret.CouponType.String()
	ret.CouponConstraintStr = ret.CouponConstraint.String()

	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithID(&ret.CouponID, true),
		coupon1.WithCouponType(&ret.CouponType, true),
		coupon1.WithAppID(&ret.AppID, true),
		coupon1.WithDenomination(&ret.Denomination, true),
		coupon1.WithCirculation(&ret.Circulation, true),
		coupon1.WithIssuedBy(&ret.UserID, true),
		coupon1.WithStartAt(&ret.StartAt, true),
		coupon1.WithDurationDays(&ret.DurationDays, true),
		coupon1.WithMessage(&ret.Message, true),
		coupon1.WithName(&ret.CouponName, true),
	)
	assert.Nil(t, err)

	coup, err := h1.CreateCoupon(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, coup)

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithCouponID(&ret.CouponID, true),
		WithUserID(&ret.UserID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.StartAt = info.StartAt
		ret.EndAt = info.EndAt
		assert.Equal(t, info, &ret)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret1.ID, true),
		WithAppID(&ret1.AppID, true),
		WithCouponID(&ret1.CouponID, true),
		WithUserID(&ret1.UserID, true),
	)
	assert.Nil(t, err)

	info, err = handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret1.CreatedAt = info.CreatedAt
		ret1.UpdatedAt = info.UpdatedAt
		ret1.StartAt = info.StartAt
		ret1.EndAt = info.EndAt
		assert.Equal(t, info, &ret1)
		ret.Allocated = info.Allocated
	}
}

func updateCoupon(t *testing.T) {
	orderID := uuid.NewString()
	ret.UsedByOrderID = &orderID
	ret.Used = true

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithUsed(&ret.Used, true),
		WithUsedByOrderID(ret.UsedByOrderID, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.UsedAt = info.UsedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCoupons(t *testing.T) {
	orderID := uuid.NewString()
	ret1.UsedByOrderID = &orderID
	ret1.Used = true

	reqs := []*npool.CouponReq{{
		ID:            &ret1.ID,
		Used:          &ret1.Used,
		UsedByOrderID: ret1.UsedByOrderID,
	}}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(reqs, true),
	)
	assert.Nil(t, err)

	infos, err := handler.UpdateCoupons(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, int(1), len(infos))
		ret1.UsedAt = infos[0].UsedAt
		assert.Equal(t, &ret1, infos[0])
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
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		IDs:           &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
		CouponType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponType)},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CouponID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		UsedByOrderID: &basetypes.StringVal{Op: cruder.EQ, Value: *ret.UsedByOrderID},
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
	t.Run("updateCoupons", updateCoupons)
	t.Run("getCoupon", getCoupon)
	t.Run("getCoupons", getCoupons)
	t.Run("deleteCoupon", deleteCoupon)
}
