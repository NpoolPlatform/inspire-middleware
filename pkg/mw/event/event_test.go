package event

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	coupon1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coupon"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
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
	ret = npool.Event{
		ID:             uuid.NewString(),
		AppID:          uuid.NewString(),
		EventType:      basetypes.UsedFor_Signup,
		EventTypeStr:   basetypes.UsedFor_Signup.String(),
		CouponIDs:      []string{uuid.NewString()},
		Credits:        decimal.RequireFromString("12.25").String(),
		CreditsPerUSD:  decimal.RequireFromString("12.25").String(),
		MaxConsecutive: 1,
		InviterLayers:  2,
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.EventTypeStr = ret.EventType.String()
	b, _ := json.Marshal(ret.CouponIDs)
	ret.CouponIDsStr = string(b)
	goodID := uuid.NewString()
	ret.GoodID = &goodID

	couponType := types.CouponType_FixAmount
	denomination := decimal.RequireFromString("12.25").String()
	circulation := decimal.RequireFromString("1225").String()
	userID := uuid.NewString()
	name := uuid.NewString()
	message := uuid.NewString()

	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithID(&ret.CouponIDs[0]),
		coupon1.WithCouponType(&couponType),
		coupon1.WithAppID(&ret.AppID),
		coupon1.WithDenomination(&denomination),
		coupon1.WithCirculation(&circulation),
		coupon1.WithIssuedBy(&userID),
		coupon1.WithMessage(&message),
		coupon1.WithName(&name),
	)
	assert.Nil(t, err)

	_, _ = h1.CreateCoupon(context.Background())

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithEventType(&ret.EventType),
		WithCouponIDs(ret.CouponIDs),
		WithCredits(&ret.Credits),
		WithCreditsPerUSD(&ret.CreditsPerUSD),
		WithMaxConsecutive(&ret.MaxConsecutive),
		WithGoodID(ret.GoodID),
		WithInviterLayers(&ret.InviterLayers),
	)
	assert.Nil(t, err)

	info, err := handler.CreateEvent(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCouponIDs(ret.CouponIDs),
		WithCredits(&ret.Credits),
		WithCreditsPerUSD(&ret.CreditsPerUSD),
		WithMaxConsecutive(&ret.MaxConsecutive),
		WithInviterLayers(&ret.InviterLayers),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateEvent(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetEvent(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getEvents(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		IDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.EventType)},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: *ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetEvents(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteEvent(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetEvent(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestEvent(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createEvent", createEvent)
	t.Run("updateEvent", updateEvent)
	t.Run("getEvent", getEvent)
	t.Run("getEvents", getEvents)
	t.Run("deleteEvent", deleteEvent)
}