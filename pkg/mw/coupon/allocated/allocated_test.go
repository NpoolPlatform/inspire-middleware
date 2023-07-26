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
		ID:               uuid.NewString(),
		CouponType:       types.CouponType_FixAmount,
		AppID:            uuid.NewString(),
		UserID:           uuid.NewString(),
		Denomination:     decimal.RequireFromString("12.25").String(),
		Circulation:      decimal.RequireFromString("12.25").String(),
		Allocated:        decimal.RequireFromString("12.25").String(),
		StartAt:          uint32(time.Now().Unix()),
		DurationDays:     27,
		CouponID:         uuid.NewString(),
		CouponName:       uuid.NewString(),
		Message:          uuid.NewString(),
		CouponConstraint: types.CouponConstraint_Normal,
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.CouponTypeStr = ret.CouponType.String()
	ret.CouponConstraintStr = ret.CouponConstraint.String()

	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithID(&ret.CouponID),
		coupon1.WithCouponType(&ret.CouponType),
		coupon1.WithAppID(&ret.AppID),
		coupon1.WithDenomination(&ret.Denomination),
		coupon1.WithCirculation(&ret.Circulation),
		coupon1.WithIssuedBy(&ret.UserID),
		coupon1.WithStartAt(&ret.StartAt),
		coupon1.WithDurationDays(&ret.DurationDays),
		coupon1.WithMessage(&ret.Message),
		coupon1.WithName(&ret.CouponName),
	)
	assert.Nil(t, err)

	_, _ = h1.CreateCoupon(context.Background())

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func creatCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithCouponID(&ret.CouponID),
		WithUserID(&ret.UserID),
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
	orderID := uuid.NewString()
	ret.UsedByOrderID = &orderID
	ret.Used = true

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithUsed(&ret.Used),
		WithUsedByOrderID(ret.UsedByOrderID),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.UsedAt = info.UsedAt
		assert.Equal(t, info, &ret)
	}
}

func getCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
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
		WithID(&ret.ID),
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

	t.Run("creatCoupon", creatCoupon)
	t.Run("updateCoupon", updateCoupon)
	t.Run("getCoupon", getCoupon)
	t.Run("getCoupons", getCoupons)
	t.Run("deleteCoupon", deleteCoupon)
}
