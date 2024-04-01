package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var ret = &npool.AppConfig{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	CommissionType:      types.CommissionType_LayeredCommission,
	CommissionTypeStr:   types.CommissionType_LayeredCommission.String(),
	SettleMode:          types.SettleMode_SettleWithPaymentAmount,
	SettleModeStr:       types.SettleMode_SettleWithPaymentAmount.String(),
	SettleAmountType:    types.SettleAmountType_SettleByPercent,
	SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
	SettleInterval:      types.SettleInterval_SettleEveryOrder,
	SettleIntervalStr:   types.SettleInterval_SettleEveryOrder.String(),
	SettleBenefit:       false,
	StartAt:             uint32(time.Now().Unix()) + 10000,
}

func create(t *testing.T) {
	_, err := CreateAppConfig(context.Background(), &npool.AppConfigReq{
		EntID:            &ret.EntID,
		AppID:            &ret.AppID,
		CommissionType:   &ret.CommissionType,
		SettleMode:       &ret.SettleMode,
		SettleAmountType: &ret.SettleAmountType,
		SettleInterval:   &ret.SettleInterval,
		SettleBenefit:    &ret.SettleBenefit,
		StartAt:          &ret.StartAt,
	})
	if assert.Nil(t, err) {
		info, err := GetAppConfigOnly(context.Background(), &npool.Conds{
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
	ret.StartAt += 10000

	info, err := UpdateAppConfig(context.Background(), &npool.AppConfigReq{
		ID:      &ret.ID,
		StartAt: &ret.StartAt,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getAppConfigs(t *testing.T) {
	infos, total, err := GetAppConfigs(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, int32(0), int32(100))
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, ret, infos[0])
	}
}

func getAppConfigOnly(t *testing.T) {
	info, err := GetAppConfigOnly(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteAppConfig(t *testing.T) {
	info, err := DeleteAppConfig(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetAppConfig(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAppConfig(t *testing.T) {
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
	t.Run("getAppConfigOnly", getAppConfigOnly)
	t.Run("getAppConfigs", getAppConfigs)
	t.Run("deleteAppConfig", deleteAppConfig)
}
