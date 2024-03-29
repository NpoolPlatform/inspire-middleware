package calculate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	commmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"
	regmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	commission1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/commission"
	invitationcode1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/invitationcode"
	registration1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/invitation/registration"

	"github.com/NpoolPlatform/inspire-middleware/pkg/testinit"
	"github.com/shopspring/decimal"

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

var percent1 = "30"

var comm1 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg1.InviterID,
	GoodID:           uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent1,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm1 = commmwpb.CommissionReq{
	AppID:            &comm1.AppID,
	UserID:           &comm1.UserID,
	GoodID:           &comm1.GoodID,
	AppGoodID:        &comm1.AppGoodID,
	SettleType:       &comm1.SettleType,
	SettleMode:       &comm1.SettleMode,
	SettleAmountType: &comm1.SettleAmountType,
	AmountOrPercent:  &comm1.AmountOrPercent,
	StartAt:          &comm1.StartAt,
}

var percent2 = "25"
var comm2 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg2.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent2,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm2 = commmwpb.CommissionReq{
	AppID:            &comm2.AppID,
	UserID:           &comm2.UserID,
	GoodID:           &comm2.GoodID,
	AppGoodID:        &comm2.AppGoodID,
	SettleType:       &comm2.SettleType,
	SettleMode:       &comm2.SettleMode,
	SettleAmountType: &comm2.SettleAmountType,
	AmountOrPercent:  &comm2.AmountOrPercent,
	StartAt:          &comm2.StartAt,
}

var percent3 = "20"
var comm3 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg3.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent3,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm3 = commmwpb.CommissionReq{
	AppID:            &comm3.AppID,
	UserID:           &comm3.UserID,
	GoodID:           &comm3.GoodID,
	AppGoodID:        &comm3.AppGoodID,
	SettleType:       &comm3.SettleType,
	SettleMode:       &comm2.SettleMode,
	SettleAmountType: &comm3.SettleAmountType,
	AmountOrPercent:  &comm3.AmountOrPercent,
	StartAt:          &comm3.StartAt,
}

var percent4 = "15"
var comm4 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg4.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent4,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm4 = commmwpb.CommissionReq{
	AppID:            &comm4.AppID,
	UserID:           &comm4.UserID,
	GoodID:           &comm4.GoodID,
	AppGoodID:        &comm4.AppGoodID,
	SettleType:       &comm4.SettleType,
	SettleMode:       &comm1.SettleMode,
	SettleAmountType: &comm4.SettleAmountType,
	AmountOrPercent:  &comm4.AmountOrPercent,
	StartAt:          &comm4.StartAt,
}

var percent5 = "12.4"
var comm5 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviterID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent5,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm5 = commmwpb.CommissionReq{
	AppID:            &comm5.AppID,
	UserID:           &comm5.UserID,
	GoodID:           &comm5.GoodID,
	AppGoodID:        &comm5.AppGoodID,
	SettleType:       &comm5.SettleType,
	SettleMode:       &comm5.SettleMode,
	SettleAmountType: &comm5.SettleAmountType,
	AmountOrPercent:  &comm5.AmountOrPercent,
	StartAt:          &comm5.StartAt,
}

var percent6 = "7"
var comm6 = commmwpb.Commission{
	AppID:            reg1.AppID,
	UserID:           reg5.InviteeID,
	GoodID:           comm1.GoodID,
	AppGoodID:        comm1.AppGoodID,
	SettleType:       types.SettleType_GoodOrderPayment,
	SettleMode:       types.SettleMode_SettleWithPaymentAmount,
	SettleAmountType: types.SettleAmountType_SettleByPercent,
	AmountOrPercent:  percent6,
	StartAt:          uint32(time.Now().Unix()),
}

var _comm6 = commmwpb.CommissionReq{
	AppID:            &comm6.AppID,
	UserID:           &comm6.UserID,
	GoodID:           &comm6.GoodID,
	AppGoodID:        &comm6.AppGoodID,
	SettleType:       &comm6.SettleType,
	SettleMode:       &comm6.SettleMode,
	SettleAmountType: &comm6.SettleAmountType,
	AmountOrPercent:  &comm6.AmountOrPercent,
	StartAt:          &comm6.StartAt,
}

func setup(t *testing.T) func(*testing.T) { //nolint
	_h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg1.AppID, true),
		invitationcode1.WithUserID(_reg1.InviterID, true),
	)
	assert.Nil(t, err)

	_info1, err := _h1.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h1.ID = &_info1.ID
	}

	h1, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg1.AppID, true),
		registration1.WithInviterID(_reg1.InviterID, true),
		registration1.WithInviteeID(_reg1.InviteeID, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h1.ID = &info1.ID
	}

	_h2, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg2.AppID, true),
		invitationcode1.WithUserID(_reg2.InviterID, true),
	)
	assert.Nil(t, err)

	_info2, err := _h2.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h2.ID = &_info2.ID
	}

	h2, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg2.AppID, true),
		registration1.WithInviterID(_reg2.InviterID, true),
		registration1.WithInviteeID(_reg2.InviteeID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h2.ID = &info2.ID
	}

	_h3, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg3.AppID, true),
		invitationcode1.WithUserID(_reg3.InviterID, true),
	)
	assert.Nil(t, err)

	_info3, err := _h3.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h3.ID = &_info3.ID
	}

	h3, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg3.AppID, true),
		registration1.WithInviterID(_reg3.InviterID, true),
		registration1.WithInviteeID(_reg3.InviteeID, true),
	)
	assert.Nil(t, err)

	info3, err := h3.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h3.ID = &info3.ID
	}

	_h4, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg4.AppID, true),
		invitationcode1.WithUserID(_reg4.InviterID, true),
	)
	assert.Nil(t, err)

	_info4, err := _h4.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h4.ID = &_info4.ID
	}

	h4, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg4.AppID, true),
		registration1.WithInviterID(_reg4.InviterID, true),
		registration1.WithInviteeID(_reg4.InviteeID, true),
	)
	assert.Nil(t, err)

	info4, err := h4.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h4.ID = &info4.ID
	}

	_h5, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_reg5.AppID, true),
		invitationcode1.WithUserID(_reg5.InviterID, true),
	)
	assert.Nil(t, err)

	_info5, err := _h5.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h5.ID = &_info5.ID
	}

	h5, err := registration1.NewHandler(
		context.Background(),
		registration1.WithAppID(_reg5.AppID, true),
		registration1.WithInviterID(_reg5.InviterID, true),
		registration1.WithInviteeID(_reg5.InviteeID, true),
	)
	assert.Nil(t, err)

	info5, err := h5.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		h5.ID = &info5.ID
	}

	h6, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm1.AppID, true),
		commission1.WithUserID(_comm1.UserID, true),
		commission1.WithGoodID(_comm1.GoodID, true),
		commission1.WithAppGoodID(_comm1.AppGoodID, true),
		commission1.WithSettleType(_comm1.SettleType, true),
		commission1.WithSettleMode(_comm1.SettleMode, true),
		commission1.WithSettleAmountType(_comm1.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm1.AmountOrPercent, true),
		commission1.WithStartAt(_comm1.StartAt, true),
	)
	assert.Nil(t, err)

	info6, err := h6.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h6.ID = &info6.ID
	}

	h7, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm2.AppID, true),
		commission1.WithUserID(_comm2.UserID, true),
		commission1.WithGoodID(_comm2.GoodID, true),
		commission1.WithAppGoodID(_comm2.AppGoodID, true),
		commission1.WithSettleType(_comm2.SettleType, true),
		commission1.WithSettleMode(_comm2.SettleMode, true),
		commission1.WithSettleAmountType(_comm2.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm2.AmountOrPercent, true),
		commission1.WithStartAt(_comm2.StartAt, true),
	)
	assert.Nil(t, err)

	info7, err := h7.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h7.ID = &info7.ID
	}

	h8, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm3.AppID, true),
		commission1.WithUserID(_comm3.UserID, true),
		commission1.WithGoodID(_comm3.GoodID, true),
		commission1.WithAppGoodID(_comm3.AppGoodID, true),
		commission1.WithSettleType(_comm3.SettleType, true),
		commission1.WithSettleMode(_comm3.SettleMode, true),
		commission1.WithSettleAmountType(_comm3.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm3.AmountOrPercent, true),
		commission1.WithStartAt(_comm3.StartAt, true),
	)
	assert.Nil(t, err)

	info8, err := h8.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h8.ID = &info8.ID
	}

	h9, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm4.AppID, true),
		commission1.WithUserID(_comm4.UserID, true),
		commission1.WithGoodID(_comm4.GoodID, true),
		commission1.WithAppGoodID(_comm4.AppGoodID, true),
		commission1.WithSettleType(_comm4.SettleType, true),
		commission1.WithSettleMode(_comm4.SettleMode, true),
		commission1.WithSettleAmountType(_comm4.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm4.AmountOrPercent, true),
		commission1.WithStartAt(_comm4.StartAt, true),
	)
	assert.Nil(t, err)

	info9, err := h9.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h9.ID = &info9.ID
	}

	h10, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm5.AppID, true),
		commission1.WithUserID(_comm5.UserID, true),
		commission1.WithGoodID(_comm5.GoodID, true),
		commission1.WithAppGoodID(_comm5.AppGoodID, true),
		commission1.WithSettleType(_comm5.SettleType, true),
		commission1.WithSettleMode(_comm5.SettleMode, true),
		commission1.WithSettleAmountType(_comm5.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm5.AmountOrPercent, true),
		commission1.WithStartAt(_comm5.StartAt, true),
	)
	assert.Nil(t, err)

	info10, err := h10.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h10.ID = &info10.ID
	}

	h11, err := commission1.NewHandler(
		context.Background(),
		commission1.WithAppID(_comm6.AppID, true),
		commission1.WithUserID(_comm6.UserID, true),
		commission1.WithGoodID(_comm6.GoodID, true),
		commission1.WithAppGoodID(_comm6.AppGoodID, true),
		commission1.WithSettleType(_comm6.SettleType, true),
		commission1.WithSettleMode(_comm6.SettleMode, true),
		commission1.WithSettleAmountType(_comm6.SettleAmountType, true),
		commission1.WithAmountOrPercent(_comm6.AmountOrPercent, true),
		commission1.WithStartAt(_comm6.StartAt, true),
	)
	assert.Nil(t, err)

	info11, err := h11.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		h11.ID = &info11.ID
	}

	return func(*testing.T) {
		_, _ = _h1.DeleteInvitationCode(context.Background())
		_, _ = _h2.DeleteInvitationCode(context.Background())
		_, _ = _h3.DeleteInvitationCode(context.Background())
		_, _ = _h4.DeleteInvitationCode(context.Background())
		_, _ = _h5.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteRegistration(context.Background())
		_, _ = h2.DeleteRegistration(context.Background())
		_, _ = h3.DeleteRegistration(context.Background())
		_, _ = h4.DeleteRegistration(context.Background())
		_, _ = h5.DeleteRegistration(context.Background())
		_, _ = h6.DeleteCommission(context.Background())
		_, _ = h7.DeleteCommission(context.Background())
		_, _ = h8.DeleteCommission(context.Background())
		_, _ = h9.DeleteCommission(context.Background())
		_, _ = h10.DeleteCommission(context.Background())
		_, _ = h11.DeleteCommission(context.Background())
	}
}

//nolint
func calculate(t *testing.T) {
	orderID := uuid.NewString()
	paymentID := uuid.NewString()
	coinTypeID := uuid.NewString()
	paymentCoinTypeID := uuid.NewString()
	paymentCoinUSDCurrency := decimal.RequireFromString("12.345").String()
	units := decimal.NewFromInt(10).String()
	paymentAmount := decimal.NewFromInt(2000).String()
	goodValue := decimal.NewFromInt(3000).String()
	settleType := types.SettleType_GoodOrderPayment
	settleAmount := types.SettleAmountType_SettleByPercent
	hasCommission := true
	orderCreatedAt := uint32(time.Now().Unix())

	handler, err := NewHandler(
		context.Background(),
		WithAppID(comm6.AppID),
		WithUserID(comm6.UserID),
		WithGoodID(comm6.GoodID),
		WithAppGoodID(comm6.AppGoodID),
		WithOrderID(orderID),
		WithPaymentID(paymentID),
		WithCoinTypeID(coinTypeID),
		WithPaymentCoinTypeID(paymentCoinTypeID),
		WithPaymentCoinUSDCurrency(paymentCoinUSDCurrency),
		WithUnits(units),
		WithPaymentAmount(paymentAmount),
		WithGoodValue(goodValue),
		WithHasCommission(hasCommission),
		WithOrderCreatedAt(orderCreatedAt),
		WithSettleType(settleType),
		WithSettleAmountType(settleAmount),
	)
	assert.Nil(t, err)

	comms, err := handler.Calculate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 6, len(comms))

		_paymentAmount := decimal.RequireFromString(paymentAmount)
		found := false
		for _, comm := range comms {
			if comm.UserID == comm1.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm2.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm3.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.NewFromInt(5).Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm4.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.RequireFromString("2.6").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm5.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.RequireFromString("5.4").Div(decimal.NewFromInt(100))).String())
				found = true
				break
			}
		}
		assert.Equal(t, found, true)

		found = false
		for _, comm := range comms {
			if comm.UserID == comm6.UserID {
				assert.Equal(t, comm.Commission, _paymentAmount.Mul(decimal.NewFromInt(7).Div(decimal.NewFromInt(100))).String())
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("calculate", calculate)
}
