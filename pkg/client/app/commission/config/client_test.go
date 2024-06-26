package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	appconfigmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/app/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
	appconfigmw "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

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

var (
	percent = "10"
	ret     = &npool.AppCommissionConfig{
		EntID:           uuid.NewString(),
		AppID:           uuid.NewString(),
		SettleType:      types.SettleType_GoodOrderPayment,
		SettleTypeStr:   types.SettleType_GoodOrderPayment.String(),
		Invites:         uint32(0),
		AmountOrPercent: percent,
		ThresholdAmount: decimal.NewFromInt(0).String(),
		StartAt:         uint32(time.Now().Unix()) + 10000,
		Disabled:        false,
		Level:           uint32(1),
	}

	appConfigRet = appconfigmw.AppConfig{
		EntID:               uuid.NewString(),
		AppID:               ret.AppID,
		SettleMode:          types.SettleMode_SettleWithGoodValue,
		SettleModeStr:       types.SettleMode_SettleWithGoodValue.String(),
		SettleAmountType:    types.SettleAmountType_SettleByPercent,
		SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
		SettleInterval:      types.SettleInterval_SettleYearly,
		SettleIntervalStr:   types.SettleInterval_SettleYearly.String(),
		CommissionType:      types.CommissionType_LayeredCommission,
		CommissionTypeStr:   types.CommissionType_LayeredCommission.String(),
		SettleBenefit:       false,
		StartAt:             uint32(time.Now().Unix()),
		MaxLevel:            uint32(5),
	}
)

func setup(t *testing.T) func(*testing.T) {
	_, err := appconfigmwcli.CreateAppConfig(context.Background(), &appconfigmw.AppConfigReq{
		EntID:            &appConfigRet.EntID,
		AppID:            &appConfigRet.AppID,
		SettleMode:       &appConfigRet.SettleMode,
		SettleAmountType: &appConfigRet.SettleAmountType,
		SettleInterval:   &appConfigRet.SettleInterval,
		CommissionType:   &appConfigRet.CommissionType,
		SettleBenefit:    &appConfigRet.SettleBenefit,
		StartAt:          &appConfigRet.StartAt,
		MaxLevel:         &appConfigRet.MaxLevel,
	})
	if assert.Nil(t, err) {
		info, err := appconfigmwcli.GetAppConfigOnly(context.Background(), &appconfigmw.Conds{
			EntID: &basetypes.StringVal{Op: cruder.EQ, Value: appConfigRet.EntID},
		})
		if assert.Nil(t, err) {
			appConfigRet.ID = info.ID
			appConfigRet.CreatedAt = info.CreatedAt
			appConfigRet.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &appConfigRet, info)
		}
	}
	return func(*testing.T) {
		_, _ = appconfigmwcli.DeleteAppConfig(context.Background(), &appConfigRet.ID, &appConfigRet.EntID)
	}
}

func create(t *testing.T) {
	_, err := CreateCommissionConfig(context.Background(), &npool.AppCommissionConfigReq{
		EntID:           &ret.EntID,
		AppID:           &ret.AppID,
		SettleType:      &ret.SettleType,
		ThresholdAmount: &ret.ThresholdAmount,
		Invites:         &ret.Invites,
		AmountOrPercent: &ret.AmountOrPercent,
		StartAt:         &ret.StartAt,
		Disabled:        &ret.Disabled,
		Level:           &ret.Level,
	})
	if assert.Nil(t, err) {
		info, err := GetCommissionConfigOnly(context.Background(), &npool.Conds{
			EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		})
		if assert.Nil(t, err) {
			ret.ID = info.ID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func update(t *testing.T) {
	ret.AmountOrPercent = "13"
	ret.StartAt += 10000
	ret.ThresholdAmount = decimal.NewFromInt(10).String()

	_, err := UpdateCommissionConfig(context.Background(), &npool.AppCommissionConfigReq{
		ID:              &ret.ID,
		StartAt:         &ret.StartAt,
		AmountOrPercent: &ret.AmountOrPercent,
		ThresholdAmount: &ret.ThresholdAmount,
	})
	if assert.Nil(t, err) {
		info, err := GetCommissionConfig(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getCommissions(t *testing.T) {
	infos, total, err := GetCommissionConfigs(context.Background(), &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		Level:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Level},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func getCommissionOnly(t *testing.T) {
	info, err := GetCommissionConfigOnly(context.Background(), &npool.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		Level:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Level},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteCommission(t *testing.T) {
	err := DeleteCommissionConfig(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetCommissionConfig(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getCommissionOnly", getCommissionOnly)
	t.Run("getCommissions", getCommissions)
	t.Run("deleteCommission", deleteCommission)
}
