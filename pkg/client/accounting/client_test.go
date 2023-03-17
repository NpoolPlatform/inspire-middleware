package accounting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
	ivcodemgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/invitationcode"
	regmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/client/commission"
	ivcodemwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/invitation/invitationcode"
	registra1 "github.com/NpoolPlatform/inspire-middleware/pkg/client/invitation/registration"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/accounting"

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

var reg1 = regmgrpb.Registration{
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _reg1 = regmgrpb.RegistrationReq{
	AppID:     &reg1.AppID,
	InviterID: &reg1.InviterID,
	InviteeID: &reg1.InviteeID,
}

var reg2 = regmgrpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg1.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg2 = regmgrpb.RegistrationReq{
	AppID:     &reg2.AppID,
	InviterID: &reg2.InviterID,
	InviteeID: &reg2.InviteeID,
}

var reg3 = regmgrpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg2.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg3 = regmgrpb.RegistrationReq{
	AppID:     &reg3.AppID,
	InviterID: &reg3.InviterID,
	InviteeID: &reg3.InviteeID,
}

var reg4 = regmgrpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg3.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg4 = regmgrpb.RegistrationReq{
	AppID:     &reg4.AppID,
	InviterID: &reg4.InviterID,
	InviteeID: &reg4.InviteeID,
}

var reg5 = regmgrpb.Registration{
	AppID:     reg1.AppID,
	InviterID: reg4.InviteeID,
	InviteeID: uuid.NewString(),
}

var _reg5 = regmgrpb.RegistrationReq{
	AppID:     &reg5.AppID,
	InviterID: &reg5.InviterID,
	InviteeID: &reg5.InviteeID,
}

var percent1 = "30"
var goodID = uuid.NewString()

var comm1 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg1.InviterID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent1,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm1 = commmwpb.CommissionReq{
	AppID:      &comm1.AppID,
	UserID:     &comm1.UserID,
	GoodID:     comm1.GoodID,
	SettleType: &comm1.SettleType,
	Percent:    comm1.Percent,
	StartAt:    &comm1.StartAt,
}

var percent2 = "25"
var comm2 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg2.InviterID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent2,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm2 = commmwpb.CommissionReq{
	AppID:      &comm2.AppID,
	UserID:     &comm2.UserID,
	GoodID:     comm2.GoodID,
	SettleType: &comm2.SettleType,
	Percent:    comm2.Percent,
	StartAt:    &comm2.StartAt,
}

var percent3 = "20"
var comm3 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg3.InviterID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent3,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm3 = commmwpb.CommissionReq{
	AppID:      &comm3.AppID,
	UserID:     &comm3.UserID,
	GoodID:     comm3.GoodID,
	SettleType: &comm3.SettleType,
	Percent:    comm3.Percent,
	StartAt:    &comm3.StartAt,
}

var percent4 = "15"
var comm4 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg4.InviterID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent4,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm4 = commmwpb.CommissionReq{
	AppID:      &comm4.AppID,
	UserID:     &comm4.UserID,
	GoodID:     comm4.GoodID,
	SettleType: &comm4.SettleType,
	Percent:    comm4.Percent,
	StartAt:    &comm4.StartAt,
}

var percent5 = "12.4"
var comm5 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg5.InviterID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent5,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm5 = commmwpb.CommissionReq{
	AppID:      &comm5.AppID,
	UserID:     &comm5.UserID,
	GoodID:     comm5.GoodID,
	SettleType: &comm5.SettleType,
	Percent:    comm5.Percent,
	StartAt:    &comm5.StartAt,
}

var percent6 = "7"
var comm6 = commmwpb.Commission{
	AppID:      reg1.AppID,
	UserID:     reg5.InviteeID,
	GoodID:     &goodID,
	SettleType: commmgrpb.SettleType_GoodOrderPercent,
	Percent:    &percent6,
	StartAt:    uint32(time.Now().Unix()),
}

var _comm6 = commmwpb.CommissionReq{
	AppID:      &comm6.AppID,
	UserID:     &comm6.UserID,
	GoodID:     comm6.GoodID,
	SettleType: &comm6.SettleType,
	Percent:    comm6.Percent,
	StartAt:    &comm6.StartAt,
}

//nolint
func accounting(t *testing.T) {
	_, err := ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &reg1.AppID,
		UserID: &reg1.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &reg2.AppID,
		UserID: &reg2.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &reg3.AppID,
		UserID: &reg3.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
		AppID:  &reg4.AppID,
		UserID: &reg4.InviterID,
	})
	assert.Nil(t, err)

	_, err = ivcodemwcli.CreateInvitationCode(context.Background(), &ivcodemgrpb.InvitationCodeReq{
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

	orderID := uuid.NewString()
	paymentID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	paymentCoinUSDCurrency := decimal.RequireFromString("12.345")
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000)
	goodValue := decimal.NewFromInt(3000)
	settleType := commmgrpb.SettleType_GoodOrderPercent

	comms, err := Accounting(
		context.Background(),
		&npool.AccountingRequest{
			AppID:                  comm6.AppID,
			UserID:                 comm6.UserID,
			GoodID:                 comm6.GetGoodID(),
			OrderID:                orderID,
			PaymentID:              paymentID,
			CoinTypeID:             coinTypeID,
			PaymentCoinTypeID:      paymentCoinTypeID,
			PaymentCoinUSDCurrency: paymentCoinUSDCurrency.String(),
			Units:                  units,
			SettleType:             settleType,
			PaymentAmount:          paymentAmount.String(),
			GoodValue:              goodValue.String(),
			HasCommission:          true,
			OrderCreatedAt:         uint32(time.Now().Unix()),
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, len(comms), 6)

		found := false
		for _, comm := range comms {
			if comm.UserID == comm1.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm2.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm3.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm4.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.RequireFromString("2.6").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm5.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.RequireFromString("5.4").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm6.UserID {
				assert.Equal(t, comm.Amount, paymentAmount.Mul(decimal.NewFromInt(7).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)
	}
}

func TestAccounting(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("accounting", accounting)
}
