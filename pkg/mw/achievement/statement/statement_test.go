package statement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
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
	ret = npool.Statement{
		ID:                     uuid.NewString(),
		AppID:                  uuid.NewString(),
		UserID:                 uuid.NewString(),
		GoodID:                 uuid.NewString(),
		DirectContributorID:    uuid.NewString(),
		OrderID:                uuid.NewString(),
		SelfOrder:              true,
		PaymentID:              uuid.NewString(),
		CoinTypeID:             uuid.NewString(),
		PaymentCoinTypeID:      uuid.NewString(),
		PaymentCoinUSDCurrency: decimal.RequireFromString("12.25").String(),
		Units:                  decimal.RequireFromString("12.25").String(),
		Amount:                 decimal.RequireFromString("12.25").String(),
		USDAmount:              decimal.RequireFromString("12.25").String(),
		Commission:             decimal.RequireFromString("12.25").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithUserID(&ret.UserID),
		WithDirectContributorID(&ret.DirectContributorID),
		WithGoodID(&ret.GoodID),
		WithOrderID(&ret.OrderID),
		WithSelfOrder(&ret.SelfOrder),
		WithPaymentID(&ret.PaymentID),
		WithCoinTypeID(&ret.CoinTypeID),
		WithPaymentCoinTypeID(&ret.PaymentCoinTypeID),
		WithPaymentCoinUSDCurrency(&ret.PaymentCoinUSDCurrency),
		WithUnits(&ret.Units),
		WithAmount(&ret.Amount),
		WithUSDAmount(&ret.USDAmount),
		WithCommission(&ret.Commission),
	)
	assert.Nil(t, err)

	info, err := handler.CreateStatement(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetStatement(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getStatements(t *testing.T) {
	conds := &npool.Conds{
		ID:                  &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:               &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:              &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		DirectContributorID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.DirectContributorID},
		GoodID:              &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		OrderID:             &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		SelfOrder:           &basetypes.BoolVal{Op: cruder.EQ, Value: ret.SelfOrder},
		PaymentID:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.PaymentID},
		CoinTypeID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		PaymentCoinTypeID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.PaymentCoinTypeID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetStatements(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteStatement(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetStatement(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestStatement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createStatement", createStatement)
	t.Run("getStatement", getStatement)
	t.Run("getStatements", getStatements)
	t.Run("deleteStatement", deleteStatement)
}
