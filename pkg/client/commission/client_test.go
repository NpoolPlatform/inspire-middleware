package commission

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var percent = "10"

var ret = &npool.Commission{
	ID:                  uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	AppGoodID:           uuid.NewString(),
	SettleType:          types.SettleType_GoodOrderPayment,
	SettleTypeStr:       types.SettleType_GoodOrderPayment.String(),
	SettleMode:          types.SettleMode_SettleWithPaymentAmount,
	SettleModeStr:       types.SettleMode_SettleWithPaymentAmount.String(),
	SettleAmountType:    types.SettleAmountType_SettleByPercent,
	SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
	SettleInterval:      types.SettleInterval_SettleEveryOrder,
	SettleIntervalStr:   types.SettleInterval_SettleEveryOrder.String(),
	AmountOrPercent:     percent,
	Threshold:           decimal.NewFromInt(0).String(),
	StartAt:             uint32(time.Now().Unix()) + 10000,
}

var req = &npool.CommissionReq{
	ID:               &ret.ID,
	AppID:            &ret.AppID,
	UserID:           &ret.UserID,
	GoodID:           &ret.GoodID,
	AppGoodID:        &ret.AppGoodID,
	SettleType:       &ret.SettleType,
	SettleMode:       &ret.SettleMode,
	SettleAmountType: &ret.SettleAmountType,
	SettleInterval:   &ret.SettleInterval,
	AmountOrPercent:  &ret.AmountOrPercent,
	StartAt:          &ret.StartAt,
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
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GetGoodID()},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.GetAppGoodID()},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func getCommissionOnly(t *testing.T) {
	info, err := GetCommissionOnly(context.Background(), &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GetGoodID()},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.GetAppGoodID()},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func TestCommission(t *testing.T) {
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

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getCommissionOnly", getCommissionOnly)
	t.Run("getCommissions", getCommissions)
}
