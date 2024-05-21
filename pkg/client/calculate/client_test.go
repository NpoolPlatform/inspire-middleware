package calculate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	appconfigmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/config"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	ivcodemwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/invitationcode"
	regmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	appconfig1 "github.com/NpoolPlatform/inspire-middleware/pkg/client/app/config"
	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/client/commission"
	ivcodemwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/invitation/invitationcode"
	registra1 "github.com/NpoolPlatform/inspire-middleware/pkg/client/invitation/registration"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	"github.com/shopspring/decimal"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/calculate"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var reg1 = regmwpb.Registration{
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _reg1 = regmwpb.RegistrationReq{
	AppID:     &reg1.AppID,
	InviterID: &reg1.InviterID,
	InviteeID: &reg1.InviteeID,
}

var reg2 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg1.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg2 = regmwpb.RegistrationReq{
	AppID:     &reg2.AppID,
	InviterID: &reg2.InviterID,
	InviteeID: &reg2.InviteeID,
}

var reg3 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg2.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg3 = regmwpb.RegistrationReq{
	AppID:     &reg3.AppID,
	InviterID: &reg3.InviterID,
	InviteeID: &reg3.InviteeID,
}

var reg4 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg3.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg4 = regmwpb.RegistrationReq{
	AppID:     &reg4.AppID,
	InviterID: &reg4.InviterID,
	InviteeID: &reg4.InviteeID,
}

var reg5 = regmwpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg4.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg5 = regmwpb.RegistrationReq{
	AppID:     &reg5.AppID,
	InviterID: &reg5.InviterID,
	InviteeID: &reg5.InviteeID,
}

var config1 = appconfigmwpb.AppConfig{
	EntID:            uuid.NewString(),
	AppID:            reg1.AppID,
	CommissionType:   types.CommissionType_LegacyCommission,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	SettleBenefit:    false,
	StartAt:          uint32(time.Now().Unix()),
	MaxLevel:         uint32(5),
}

var _config1 = appconfigmwpb.AppConfigReq{
	EntID:            &config1.EntID,
	AppID:            &config1.AppID,
	CommissionType:   &config1.CommissionType,
	SettleAmountType: &config1.SettleAmountType,
	SettleMode:       &config1.SettleMode,
	SettleInterval:   &config1.SettleInterval,
	SettleBenefit:    &config1.SettleBenefit,
	StartAt:          &config1.StartAt,
	MaxLevel:         &config1.MaxLevel,
}

var percent1 = "30"
var comm1 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg1.InviterID,
	GoodID:           uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent1,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm1 = commmwpb.CommissionReq{
	AppID:            &comm1.AppID,
	UserID:           &comm1.UserID,
	GoodID:           &comm1.GoodID,
	AppGoodID:        &comm1.AppGoodID,
	SettleType:       &comm1.SettleType,
	AmountOrPercent:  &comm1.AmountOrPercent,
	SettleAmountType: &comm1.SettleAmountType,
	SettleMode:       &comm1.SettleMode,
	SettleInterval:   &comm1.SettleInterval,
	StartAt:          &comm1.StartAt,
}

var percent2 = "25"
var comm2 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg2.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent2,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm2 = commmwpb.CommissionReq{
	AppID:            &comm2.AppID,
	UserID:           &comm2.UserID,
	GoodID:           &comm2.GoodID,
	AppGoodID:        &comm2.AppGoodID,
	SettleType:       &comm2.SettleType,
	AmountOrPercent:  &comm2.AmountOrPercent,
	SettleAmountType: &comm2.SettleAmountType,
	SettleMode:       &comm2.SettleMode,
	SettleInterval:   &comm2.SettleInterval,
	StartAt:          &comm2.StartAt,
}

var percent3 = "20"
var comm3 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg3.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent3,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm3 = commmwpb.CommissionReq{
	AppID:            &comm3.AppID,
	UserID:           &comm3.UserID,
	GoodID:           &comm3.GoodID,
	AppGoodID:        &comm3.AppGoodID,
	SettleType:       &comm3.SettleType,
	AmountOrPercent:  &comm3.AmountOrPercent,
	SettleAmountType: &comm3.SettleAmountType,
	SettleMode:       &comm3.SettleMode,
	SettleInterval:   &comm3.SettleInterval,
	StartAt:          &comm3.StartAt,
}

var percent4 = "15"
var comm4 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg4.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent4,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm4 = commmwpb.CommissionReq{
	AppID:            &comm4.AppID,
	UserID:           &comm4.UserID,
	GoodID:           &comm4.GoodID,
	AppGoodID:        &comm4.AppGoodID,
	SettleType:       &comm4.SettleType,
	AmountOrPercent:  &comm4.AmountOrPercent,
	SettleAmountType: &comm4.SettleAmountType,
	SettleMode:       &comm4.SettleMode,
	SettleInterval:   &comm4.SettleInterval,
	StartAt:          &comm4.StartAt,
}

var percent5 = "12.4"
var comm5 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent5,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm5 = commmwpb.CommissionReq{
	AppID:            &comm5.AppID,
	UserID:           &comm5.UserID,
	GoodID:           &comm5.GoodID,
	AppGoodID:        &comm5.AppGoodID,
	SettleType:       &comm5.SettleType,
	AmountOrPercent:  &comm5.AmountOrPercent,
	SettleAmountType: &comm5.SettleAmountType,
	SettleMode:       &comm5.SettleMode,
	SettleInterval:   &comm5.SettleInterval,
	StartAt:          &comm5.StartAt,
}

var percent6 = "7"
var comm6 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviteeID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	AmountOrPercent:  percent6,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleInterval:   types.SettleInterval_SettleEveryOrder,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm6 = commmwpb.CommissionReq{
	AppID:            &comm6.AppID,
	UserID:           &comm6.UserID,
	GoodID:           &comm6.GoodID,
	AppGoodID:        &comm6.AppGoodID,
	SettleType:       &comm6.SettleType,
	AmountOrPercent:  &comm6.AmountOrPercent,
	SettleAmountType: &comm6.SettleAmountType,
	SettleMode:       &comm6.SettleMode,
	SettleInterval:   &comm6.SettleInterval,
	StartAt:          &comm6.StartAt,
}

//nolint
func calculate(t *testing.T) {
	_, err := ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemwpb.InvitationCodeReq{
		AppID:  &reg1.AppID,
		UserID: &reg1.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemwpb.InvitationCodeReq{
		AppID:  &reg2.AppID,
		UserID: &reg2.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemwpb.InvitationCodeReq{
		AppID:  &reg3.AppID,
		UserID: &reg3.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemwpb.InvitationCodeReq{
		AppID:  &reg4.AppID,
		UserID: &reg4.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemwpb.InvitationCodeReq{
		AppID:  &reg5.AppID,
		UserID: &reg5.InviterID,
	})
	assert.Nil(t, err)

	_, err = registra1.CreateRegistration(context.Background(), &_reg1)
	assert.Nil(t, err)
	_, err = registra1.CreateRegistration(context.Background(), &_reg2)
	assert.Nil(t, err)
	_, err = registra1.CreateRegistration(context.Background(), &_reg3)
	assert.Nil(t, err)
	_, err = registra1.CreateRegistration(context.Background(), &_reg4)
	assert.Nil(t, err)
	_, err = registra1.CreateRegistration(context.Background(), &_reg5)
	assert.Nil(t, err)

	_, err = commission1.CreateCommission(context.Background(), &_comm1)
	assert.Nil(t, err)
	_, err = commission1.CreateCommission(context.Background(), &_comm2)
	assert.Nil(t, err)
	_, err = commission1.CreateCommission(context.Background(), &_comm3)
	assert.Nil(t, err)
	_, err = commission1.CreateCommission(context.Background(), &_comm4)
	assert.Nil(t, err)
	_, err = commission1.CreateCommission(context.Background(), &_comm5)
	assert.Nil(t, err)
	_, err = commission1.CreateCommission(context.Background(), &_comm6)
	assert.Nil(t, err)

	_, err = appconfig1.CreateAppConfig(context.Background(), &_config1)
	assert.Nil(t, err)

	orderID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	paymentCoinUSDCurrency := decimal.RequireFromString("12.345")
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000)
	paymentAmountUSD := decimal.NewFromInt(2000)
	goodValueUSD := decimal.NewFromInt(30000)
	settleType := types.SettleType_GoodOrderPayment

	comms, err := Calculate(
		context.Background(),
		&npool.CalculateRequest{
			AppID:                  comm6.AppID,
			UserID:                 comm6.UserID,
			GoodID:                 comm6.GetGoodID(),
			AppGoodID:              comm6.GetAppGoodID(),
			OrderID:                orderID,
			GoodCoinTypeID:         coinTypeID,
			PaymentCoinTypeID:      paymentCoinTypeID,
			PaymentCoinUSDCurrency: paymentCoinUSDCurrency.String(),
			Units:                  units,
			SettleType:             settleType,
			PaymentAmount:          paymentAmount.String(),
			PaymentAmountUSD:       paymentAmountUSD.String(),
			GoodValueUSD:           goodValueUSD.String(),
			HasCommission:          true,
			OrderCreatedAt:         uint32(time.Now().Unix()),
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, len(comms), 6)

		found := false
		for _, comm := range comms {
			if *comm.UserID == comm1.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm2.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm3.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm4.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.RequireFromString("2.6").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm5.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.RequireFromString("5.4").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if *comm.UserID == comm6.UserID {
				assert.Equal(t, *comm.PaymentStatements[0].CommissionAmount, paymentAmount.Mul(decimal.NewFromInt(7).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func TestCalculate(t *testing.T) {
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

	t.Run("calculate", calculate)
}
