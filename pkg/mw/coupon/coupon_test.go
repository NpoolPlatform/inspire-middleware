package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
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
		ID:           uuid.NewString(),
		CouponType:   types.CouponType_FixAmount,
		AppID:        uuid.NewString(),
		Denomination: decimal.RequireFromString("12.25").String(),
		Circulation:  decimal.RequireFromString("12.25").String(),
		IssuedBy:     uuid.NewString(),
		StartAt:      uint32(time.Now().Unix()),
		DurationDays: 234,
		Message:      uuid.NewString(),
		Name:         uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func creatCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCouponType(&ret.CouponType),
		WithAppID(&ret.AppID),
		WithDenomination(&ret.Denomination),
		WithCirculation(&ret.Circulation),
		WithIssuedBy(&ret.IssuedBy),
		WithStartAt(&ret.StartAt),
		WithDurationDays(&ret.DurationDays),
		WithMessage(&ret.Message),
		WithName(&ret.Name),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

/*
func updateCoupon(t *testing.T) {
	ret.Coupon = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCoupon(&ret.Coupon),
		WithDescription(&ret.Description),
		WithDefault(&ret.Default),
		WithGenesis(&ret.Genesis),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
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
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
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
		assert.NotEqual(t, len(infos), 0)
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
*/

func TestCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("creatCoupon", creatCoupon)
	// t.Run("updateCoupon", updateCoupon)
	// t.Run("getCoupon", getCoupon)
	// t.Run("getCoupons", getCoupons)
	// t.Run("deleteCoupon", deleteCoupon)
}
