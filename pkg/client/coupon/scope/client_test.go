package scope

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	couponmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon"
	couponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/scope"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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
		CouponType:         coupon.CouponType,
		CouponTypeStr:      coupon.CouponTypeStr,
		CouponScope:        coupon.CouponScope,
		CouponScopeStr:     coupon.CouponScopeStr,
		CouponName:         coupon.Name,
		CouponDenomination: coupon.Denomination,
	}
)

func setup(t *testing.T) func(*testing.T) {
	info, err := couponmwcli.CreateCoupon(context.Background(), &couponmwpb.CouponReq{
		ID:           &coupon.ID,
		AppID:        &coupon.AppID,
		Name:         &coupon.Name,
		Message:      &coupon.Message,
		CouponType:   &coupon.CouponType,
		Denomination: &coupon.Denomination,
		Circulation:  &coupon.Circulation,
		DurationDays: &coupon.DurationDays,
		IssuedBy:     &coupon.IssuedBy,
		StartAt:      &coupon.StartAt,
		CouponScope:  &coupon.CouponScope,
	})
	if assert.Nil(t, err) {
		coupon.CreatedAt = info.CreatedAt
		coupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &coupon, info)
	}

	return func(*testing.T) {}
}

func createScope(t *testing.T) {
	info, err := CreateScope(context.Background(), &npool.ScopeReq{
		ID:          &ret.ID,
		AppID:       &ret.AppID,
		AppGoodID:   &ret.AppGoodID,
		CouponID:    &ret.CouponID,
		CouponScope: &ret.CouponScope,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getScope(t *testing.T) {
	info, err := GetScope(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getScopes(t *testing.T) {
	infos, total, err := GetScopes(context.Background(), &npool.Conds{
		ID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
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

func existScopeConds(t *testing.T) {
	conds := &npool.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}

	exist, err := ExistScopeConds(context.Background(), conds)
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteScope(t *testing.T) {
	info, err := DeleteScope(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetScope(context.Background(), ret.ID)
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

	t.Run("createScope", createScope)
	t.Run("getScope", getScope)
	t.Run("getScopes", getScopes)
	t.Run("existScopeConds", existScopeConds)
	t.Run("deleteScope", deleteScope)
}
