package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"

	"github.com/google/uuid"
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

var ret2 = &npool.AchievementUser{
	AppID:                ret.AppID,
	UserID:               ret.UserID,
	TotalCommission:      "3030.3",
	SelfCommission:       "3030.3",
	DirectConsumeAmount:  "20000",
	DirectInvites:        "0",
	IndirectInvites:      "0",
	InviteeConsumeAmount: "0",
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

func getAchievementUsers(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetAchievementUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			ret2.ID = infos[0].ID
			ret2.EntID = infos[0].EntID
			ret2.CreatedAt = infos[0].CreatedAt
			ret2.UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, infos[0], ret2)
		}
	}
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getAchievementUsers", getAchievementUsers)
}
