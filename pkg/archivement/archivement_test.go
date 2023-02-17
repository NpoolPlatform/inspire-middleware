//nolint:dupl
package archivement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	detailmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/archivement/detail"
	generalmgrcli "github.com/NpoolPlatform/inspire-manager/pkg/client/archivement/general"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
	generalmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

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

var ret = &detailmgrpb.Detail{
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

var req = &detailmgrpb.DetailReq{
	ID:                     &ret.ID,
	AppID:                  &ret.AppID,
	UserID:                 &ret.UserID,
	DirectContributorID:    &ret.DirectContributorID,
	GoodID:                 &ret.GoodID,
	OrderID:                &ret.OrderID,
	SelfOrder:              &ret.SelfOrder,
	PaymentID:              &ret.PaymentID,
	CoinTypeID:             &ret.CoinTypeID,
	PaymentCoinTypeID:      &ret.PaymentCoinTypeID,
	PaymentCoinUSDCurrency: &ret.PaymentCoinUSDCurrency,
	Units:                  &ret.Units,
	Amount:                 &ret.Amount,
	USDAmount:              &ret.USDAmount,
	Commission:             &ret.Commission,
}

func bookKeeping(t *testing.T) {
	err := BookKeeping(context.Background(), req)
	assert.Nil(t, err)

	info, err := detailmgrcli.GetDetail(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, ret, info)
	}

	g, err := generalmgrcli.GetGeneralOnly(context.Background(), &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.UserID,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.GoodID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.UserID, ret.UserID)
		assert.Equal(t, g.GoodID, ret.GoodID)
		assert.Equal(t, g.CoinTypeID, ret.CoinTypeID)
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.TotalUnits, ret.Units)
		assert.Equal(t, g.SelfUnits, ret.Units)
		assert.Equal(t, g.TotalAmount, ret.USDAmount)
		assert.Equal(t, g.SelfAmount, ret.USDAmount)
		assert.Equal(t, g.TotalCommission, ret.Commission)
		assert.Equal(t, g.SelfCommission, ret.Commission)
	}
}

func bookKeepingV2(t *testing.T) {
	err := BookKeepingV2(context.Background(), []*detailmgrpb.DetailReq{req})
	assert.Nil(t, err)

	info, err := detailmgrcli.GetDetail(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, ret, info)
	}

	g, err := generalmgrcli.GetGeneralOnly(context.Background(), &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.UserID,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.GoodID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.UserID, ret.UserID)
		assert.Equal(t, g.GoodID, ret.GoodID)
		assert.Equal(t, g.CoinTypeID, ret.CoinTypeID)
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.TotalUnits, ret.Units)
		assert.Equal(t, g.SelfUnits, ret.Units)
		assert.Equal(t, g.TotalAmount, ret.USDAmount)
		assert.Equal(t, g.SelfAmount, ret.USDAmount)
		assert.Equal(t, g.TotalCommission, ret.Commission)
		assert.Equal(t, g.SelfCommission, ret.Commission)
	}
}

func expropriate(t *testing.T) {
	err := Expropriate(context.Background(), ret.OrderID)
	assert.Nil(t, err)

	_, err = detailmgrcli.GetDetail(context.Background(), ret.ID)
	assert.NotNil(t, err)

	g, err := generalmgrcli.GetGeneralOnly(context.Background(), &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.AppID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.UserID,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.GoodID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.UserID, ret.UserID)
		assert.Equal(t, g.GoodID, ret.GoodID)
		assert.Equal(t, g.CoinTypeID, ret.CoinTypeID)
		assert.Equal(t, g.AppID, ret.AppID)
		assert.Equal(t, g.TotalUnits, "0")
		assert.Equal(t, g.SelfUnits, "0")
		assert.Equal(t, g.TotalAmount, "0")
		assert.Equal(t, g.SelfAmount, "0")
		assert.Equal(t, g.TotalCommission, "0")
		assert.Equal(t, g.SelfCommission, "0")
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("bookKeeping", bookKeeping)
	t.Run("bookKeepingV2", bookKeepingV2)
	t.Run("expropriate", expropriate)
}
