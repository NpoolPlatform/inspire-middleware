package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"

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

var ret = &npool.AppGoodCommissionConfig{
	EntID:           uuid.NewString(),
	AppID:           uuid.NewString(),
	GoodID:          uuid.NewString(),
	AppGoodID:       uuid.NewString(),
	SettleType:      types.SettleType_GoodOrderPayment,
	SettleTypeStr:   types.SettleType_GoodOrderPayment.String(),
	Invites:         uint32(0),
	AmountOrPercent: percent,
	ThresholdAmount: decimal.NewFromInt(0).String(),
	StartAt:         uint32(time.Now().Unix()) + 10000,
}

func create(t *testing.T) {
	info, err := CreateCommissionConfig(context.Background(), &npool.AppGoodCommissionConfigReq{
		EntID:           &ret.EntID,
		AppID:           &ret.AppID,
		GoodID:          &ret.GoodID,
		AppGoodID:       &ret.AppGoodID,
		SettleType:      &ret.SettleType,
		ThresholdAmount: &ret.ThresholdAmount,
		Invites:         &ret.Invites,
		AmountOrPercent: &ret.AmountOrPercent,
		StartAt:         &ret.StartAt,
	})
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	ret.AmountOrPercent = "13"
	ret.StartAt += 10000
	ret.ThresholdAmount = decimal.NewFromInt(10).String()

	info, err := UpdateCommissionConfig(context.Background(), &npool.AppGoodCommissionConfigReq{
		ID:              &ret.ID,
		StartAt:         &ret.StartAt,
		AmountOrPercent: &ret.AmountOrPercent,
		ThresholdAmount: &ret.ThresholdAmount,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getCommissions(t *testing.T) {
	infos, total, err := GetCommissionConfigs(context.Background(), &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
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
	info, err := GetCommissionConfigOnly(context.Background(), &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
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
