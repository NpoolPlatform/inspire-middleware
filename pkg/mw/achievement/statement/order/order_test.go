package orderstatement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	// cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	orderstatementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"

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

var ret = &orderstatementmwpb.Statement{
	EntID:                uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	GoodID:               uuid.NewString(),
	AppGoodID:            uuid.NewString(),
	OrderID:              uuid.NewString(),
	OrderUserID:          uuid.NewString(),
	GoodCoinTypeID:       uuid.NewString(),
	Units:                decimal.NewFromInt(10).String(),
	GoodValueUSD:         decimal.NewFromInt(120).String(),
	PaymentAmountUSD:     decimal.NewFromInt(120).String(),
	CommissionAmountUSD:  decimal.NewFromInt(30).String(),
	AppConfigID:          uuid.NewString(),
	CommissionConfigID:   uuid.NewString(),
	CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
}

func setup(t *testing.T) func(*testing.T) { //nolint
	ret.CommissionConfigTypeStr = ret.CommissionConfigType.String()
	return func(*testing.T) {}
}

func createStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithOrderUserID(&ret.OrderUserID, true),
		WithGoodCoinTypeID(&ret.GoodCoinTypeID, true),
		WithUnits(&ret.Units, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithCommissionAmountUSD(&ret.CommissionAmountUSD, true),
		WithAppConfigID(&ret.AppConfigID, true),
		WithCommissionConfigID(&ret.CommissionConfigID, true),
		WithCommissionConfigType(&ret.CommissionConfigType, true),
		// WithPaymentStatements(&ret.PaymentStatements, true),
	)
	assert.Nil(t, err)

	err = handler.CreateStatement(context.Background())
	assert.Nil(t, err)
}

func deleteStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteStatement(context.Background())
	assert.Nil(t, err)
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createStatement", createStatement)
	t.Run("deleteStatement", deleteStatement)
}
