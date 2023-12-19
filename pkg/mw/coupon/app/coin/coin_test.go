package coin

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
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/app/coin"
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
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&coupon.EntID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&coupon.Name, true),
		coupon1.WithMessage(&coupon.Message, true),
		coupon1.WithCouponType(&coupon.CouponType, true),
		coupon1.WithDenomination(&coupon.Denomination, true),
		coupon1.WithCouponScope(&coupon.CouponScope, true),
		coupon1.WithCirculation(&coupon.Circulation, true),
		coupon1.WithDurationDays(&coupon.DurationDays, true),
		coupon1.WithIssuedBy(&coupon.IssuedBy, true),
		coupon1.WithStartAt(&coupon.StartAt, true),
		coupon1.WithEndAt(&coupon.EndAt, true),
		coupon1.WithCashableProbabilityPerMillion(&coupon.CashableProbabilityPerMillion, true),
	)
	assert.Nil(t, err)

	coup, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		coupon.ID = coup.ID
		coupon.CreatedAt = coup.CreatedAt
		coupon.UpdatedAt = coup.UpdatedAt
		assert.Equal(t, &coupon, coup)
		h1.ID = &coup.ID
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createCouponCoin(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithCouponID(&ret.CouponID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCouponCoin(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getCouponCoin(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCouponCoin(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCouponCoins(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		CouponID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCouponCoins(context.Background())
	if assert.Nil(t, err) {
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
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistCouponCoinConds(context.Background())
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteCouponCoin(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCouponCoin(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCouponCoin(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCouponCoin", createCouponCoin)
	t.Run("getCouponCoin", getCouponCoin)
	t.Run("getCouponCoins", getCouponCoins)
	t.Run("existCouponCoin", existCouponCoinConds)
	t.Run("deleteCouponCoin", deleteCouponCoin)
}
