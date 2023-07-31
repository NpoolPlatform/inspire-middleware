//nolint:dupl
package achivement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achivement/statement"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement/statement"

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
	ID:                     uuid.NewString(),
	AppID:                  uuid.NewString(),
	UserID:                 uuid.NewString(),
	DirectContributorID:    uuid.NewString(),
	GoodID:                 uuid.NewString(),
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

var ret1 = &npool.Achivement{
	AppID:           ret.AppID,
	UserID:          ret.UserID,
	GoodID:          ret.GoodID,
	CoinTypeID:      ret.CoinTypeID,
	TotalAmount:     ret.Amount,
	SelfAmount:      ret.Amount,
	TotalUnits:      ret.Units,
	SelfUnits:       ret.Units,
	TotalCommission: ret.Commission,
	SelfCommission:  ret.Commission,
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithID(&ret.ID),
		statement1.WithAppID(&ret.AppID),
		statement1.WithUserID(&ret.UserID),
		statement1.WithDirectContributorID(&ret.DirectContributorID),
		statement1.WithGoodID(&ret.GoodID),
		statement1.WithOrderID(&ret.OrderID),
		statement1.WithSelfOrder(&ret.SelfOrder),
		statement1.WithPaymentID(&ret.PaymentID),
		statement1.WithCoinTypeID(&ret.CoinTypeID),
		statement1.WithPaymentCoinTypeID(&ret.PaymentCoinTypeID),
		statement1.WithPaymentCoinUSDCurrency(&ret.PaymentCoinUSDCurrency),
		statement1.WithUnits(&ret.Units),
		statement1.WithAmount(&ret.Amount),
		statement1.WithUSDAmount(&ret.USDAmount),
		statement1.WithCommission(&ret.Commission),
	)
	assert.Nil(t, err)

	info, err := h1.CreateStatement(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	return func(*testing.T) {
		_, _ = h1.DeleteStatement(context.Background())
	}
}

func getAchivements(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		WithOffset(0),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetAchivements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			ret1.ID = infos[0].ID
			ret1.CreatedAt = infos[0].CreatedAt
			ret1.UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, ret1, infos[0])
		}
	}
}

func expropriateAchivement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID),
	)
	assert.Nil(t, err)

	err = handler.ExpropriateAchivement(context.Background())
	assert.Nil(t, err)

	h1, err := statement1.NewHandler(
		context.Background(),
		statement1.WithID(&ret.ID),
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
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		WithOffset(0),
	)
	assert.Nil(t, err)

	infos, total, err := h2.GetAchivements(context.Background())
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
			assert.Equal(t, ret1, infos[0])
		}
	}
}

func TestAchivement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getAchivements", getAchivements)
	t.Run("expropriateAchivement", expropriateAchivement)
}
