package commission

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

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

var percent = "10"

var goodID = uuid.NewString()

var ret = &npool.Commission{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	GoodID:         &goodID,
	SettleType:     mgrpb.SettleType_GoodOrderPercent,
	SettleMode:     mgrpb.SettleMode_SettleWithPaymentAmount,
	SettleInterval: mgrpb.SettleInterval_SettleEveryOrder,
	Percent:        &percent,
	StartAt:        uint32(time.Now().Unix()) + 10000,
}

var req = &npool.CommissionReq{
	ID:         &ret.ID,
	AppID:      &ret.AppID,
	UserID:     &ret.UserID,
	GoodID:     ret.GoodID,
	SettleType: &ret.SettleType,
	Percent:    ret.Percent,
	StartAt:    &ret.StartAt,
}

func create(t *testing.T) {
	info, err := CreateCommission(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	info, err := UpdateCommission(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getCommissions(t *testing.T) {
	infos, total, err := GetCommissions(context.Background(), &npool.Conds{
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
			Value: ret.GetGoodID(),
		},
		SettleType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(ret.SettleType),
		},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func getCommissionOnly(t *testing.T) {
	info, err := GetCommissionOnly(context.Background(), &npool.Conds{
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
			Value: ret.GetGoodID(),
		},
		SettleType: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(ret.SettleType),
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func TestCommission(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getCommissionOnly", getCommissionOnly)
	t.Run("getCommissions", getCommissions)
}
