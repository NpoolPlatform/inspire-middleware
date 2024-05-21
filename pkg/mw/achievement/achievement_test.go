//nolint:dupl
package achievement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	orderstatementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
	paymentmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order/payment"

	achievement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/good"
	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement/order"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/good"

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
	EntID:                   uuid.NewString(),
	AppID:                   uuid.NewString(),
	UserID:                  uuid.NewString(),
	GoodID:                  uuid.NewString(),
	AppGoodID:               uuid.NewString(),
	OrderID:                 uuid.NewString(),
	OrderUserID:             uuid.NewString(),
	GoodCoinTypeID:          uuid.NewString(),
	Units:                   decimal.NewFromInt(10).String(),
	GoodValueUSD:            decimal.NewFromInt(120).String(),
	PaymentAmountUSD:        decimal.NewFromInt(120).String(),
	CommissionAmountUSD:     decimal.NewFromInt(30).String(),
	AppConfigID:             uuid.NewString(),
	CommissionConfigID:      uuid.NewString(),
	CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
	CommissionConfigTypeStr: types.CommissionConfigType_LegacyCommissionConfig.String(),
}

var ret1 = &npool.Achievement{
	AppID:              ret.AppID,
	UserID:             ret.UserID,
	GoodID:             ret.GoodID,
	AppGoodID:          ret.AppGoodID,
	TotalAmountUSD:     ret.PaymentAmountUSD,
	SelfAmountUSD:      "0",
	TotalUnits:         ret.Units,
	SelfUnits:          "0",
	TotalCommissionUSD: ret.CommissionAmountUSD,
	SelfCommissionUSD:  "0",
}

func createStatement(t *testing.T) {
	paymentCoinTypeID := uuid.NewString()
	coinUSDCurrency := "1.23"
	amount := "1000"
	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&ret.EntID, true),
		statement1.WithAppID(&ret.AppID, true),
		statement1.WithUserID(&ret.UserID, true),
		statement1.WithGoodID(&ret.GoodID, true),
		statement1.WithAppGoodID(&ret.AppGoodID, true),
		statement1.WithOrderID(&ret.OrderID, true),
		statement1.WithOrderUserID(&ret.OrderUserID, true),
		statement1.WithGoodCoinTypeID(&ret.GoodCoinTypeID, true),
		statement1.WithUnits(&ret.Units, true),
		statement1.WithGoodValueUSD(&ret.GoodValueUSD, true),
		statement1.WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		statement1.WithCommissionAmountUSD(&ret.CommissionAmountUSD, true),
		statement1.WithAppConfigID(&ret.AppConfigID, true),
		statement1.WithCommissionConfigID(&ret.CommissionConfigID, true),
		statement1.WithCommissionConfigType(&ret.CommissionConfigType, true),
		statement1.WithPaymentStatements([]*paymentmwpb.StatementReq{{
			PaymentCoinTypeID: &paymentCoinTypeID,
			CoinUSDCurrency:   &coinUSDCurrency,
			Amount:            &amount,
		},
		}, true),
	)
	assert.Nil(t, err)

	err = handler.CreateStatement(context.Background())
	assert.Nil(t, err)
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

	h2, err := achievement1.NewHandler(
		context.Background(),
		achievement1.WithConds(&npool.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			UserIDs:   &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
	)
	assert.Nil(t, err)

	infos, total, err := h2.GetAchievements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret1.TotalAmountUSD = decimal.NewFromInt(0).String()
		ret1.SelfAmountUSD = decimal.NewFromInt(0).String()
		ret1.TotalUnits = decimal.NewFromInt(0).String()
		ret1.SelfUnits = decimal.NewFromInt(0).String()
		ret1.TotalCommissionUSD = decimal.NewFromInt(0).String()
		ret1.SelfCommissionUSD = decimal.NewFromInt(0).String()
		ret1.CreatedAt = infos[0].CreatedAt
		ret1.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], ret1)
	}
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createStatement", createStatement)
	t.Run("expropriateAchievement", expropriateAchievement)
}
