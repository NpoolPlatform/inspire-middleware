//nolint:dupl
package achievement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

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

var ret = &statementmwpb.Statement{
	EntID:                  uuid.NewString(),
	AppID:                  uuid.NewString(),
	UserID:                 uuid.NewString(),
	DirectContributorID:    uuid.NewString(),
	GoodID:                 uuid.NewString(),
	AppGoodID:              uuid.NewString(),
	OrderID:                uuid.NewString(),
	SelfOrder:              true,
	PaymentID:              uuid.NewString(),
	CoinTypeID:             uuid.NewString(),
	PaymentCoinTypeID:      uuid.NewString(),
	PaymentCoinUSDCurrency: "10.101",
	Units:                  "10",
	Amount:                 "10000",
	USDAmount:              "20000",
	Commission:             "300",
}

var ret1 = &npool.Achievement{
	AppID:           ret.AppID,
	UserID:          ret.UserID,
	GoodID:          ret.GoodID,
	AppGoodID:       ret.AppGoodID,
	CoinTypeID:      ret.CoinTypeID,
	TotalAmount:     ret.USDAmount,
	SelfAmount:      ret.USDAmount,
	TotalUnits:      ret.Units,
	SelfUnits:       ret.Units,
	TotalCommission: "3030.3",
	SelfCommission:  "3030.3",
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&ret.EntID, true),
		statement1.WithAppID(&ret.AppID, true),
		statement1.WithUserID(&ret.UserID, true),
		statement1.WithDirectContributorID(&ret.DirectContributorID, true),
		statement1.WithGoodID(&ret.GoodID, true),
		statement1.WithAppGoodID(&ret.AppGoodID, true),
		statement1.WithOrderID(&ret.OrderID, true),
		statement1.WithSelfOrder(&ret.SelfOrder, true),
		statement1.WithPaymentID(&ret.PaymentID, true),
		statement1.WithCoinTypeID(&ret.CoinTypeID, true),
		statement1.WithPaymentCoinTypeID(&ret.PaymentCoinTypeID, true),
		statement1.WithPaymentCoinUSDCurrency(&ret.PaymentCoinUSDCurrency, true),
		statement1.WithUnits(&ret.Units, true),
		statement1.WithAmount(&ret.Amount, true),
		statement1.WithUSDAmount(&ret.USDAmount, true),
		statement1.WithCommission(&ret.Commission, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateStatement(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		h1.ID = &info.ID
	}

	return func(*testing.T) {
		_, _ = h1.DeleteStatement(context.Background())
	}
}

func getAchievements(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetAchievements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			ret1.ID = infos[0].ID
			ret1.EntID = infos[0].EntID
			ret1.CreatedAt = infos[0].CreatedAt
			ret1.UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, infos[0], ret1)
		}
	}
}

func expropriateAchievement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID, true),
	)
	assert.Nil(t, err)

	err = handler.ExpropriateAchievement(context.Background())
	assert.Nil(t, err)

	h1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := h1.GetStatement(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)

	h2, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, total, err := h2.GetAchievements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			ret1.TotalAmount = decimal.NewFromInt(0).String()
			ret1.SelfAmount = decimal.NewFromInt(0).String()
			ret1.TotalUnits = decimal.NewFromInt(0).String()
			ret1.SelfUnits = decimal.NewFromInt(0).String()
			ret1.TotalCommission = decimal.NewFromInt(0).String()
			ret1.SelfCommission = decimal.NewFromInt(0).String()
			ret1.CreatedAt = infos[0].CreatedAt
			ret1.UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, infos[0], ret1)
		}
	}
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getAchievements", getAchievements)
	t.Run("expropriateAchievement", expropriateAchievement)
}
