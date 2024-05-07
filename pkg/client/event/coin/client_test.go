package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event/coin"

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

var ret = &npool.EventCoin{
	EntID:        uuid.NewString(),
	AppID:        uuid.NewString(),
	EventID:      uuid.NewString(),
	CoinConfigID: uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createEventCoin(t *testing.T) {
	err := CreateEventCoin(context.Background(), &npool.EventCoinReq{
		EntID:        &ret.EntID,
		AppID:        &ret.AppID,
		EventID:      &ret.EventID,
		CoinConfigID: &ret.CoinConfigID,
	})
	if assert.Nil(t, err) {
		info, err := GetEventCoin(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func getEventCoin(t *testing.T) {
	info, err := GetEventCoin(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getEventCoins(t *testing.T) {
	infos, total, err := GetEventCoins(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinConfigID},
		EntIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getEventCoinOnly(t *testing.T) {
	info, err := GetEventCoinOnly(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EventID},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinConfigID},
		EntIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteEventCoin(t *testing.T) {
	err := DeleteEventCoin(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetEventCoin(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestEventCoin(t *testing.T) {
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

	t.Run("createEventCoin", createEventCoin)
	t.Run("getEventCoin", getEventCoin)
	t.Run("getEventCoins", getEventCoins)
	t.Run("getEventCoinOnly", getEventCoinOnly)
	t.Run("deleteEventCoin", deleteEventCoin)
}
