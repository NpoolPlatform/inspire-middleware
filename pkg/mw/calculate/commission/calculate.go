package commission

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	appcommissionconfig "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/commission/config"
	appgoodcommissionconfig "github.com/NpoolPlatform/message/npool/inspire/mw/v1/app/good/commission/config"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/commission"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Commission struct {
	AppConfigID             string
	CommissionConfigID      string
	CommissionConfigType    types.CommissionConfigType
	AppID                   string
	UserID                  string
	DirectContributorUserID *string
	Amount                  string
}

type calculateHandler struct {
	*Handler
}

//nolint:funlen
func (h *Handler) Calculate(ctx context.Context) ([]*Commission, error) {
	commMap := map[string]*npool.Commission{}
	fmt.Println("=========== calculate commission config ================")
	for _, comm := range h.Commissions {
		commMap[comm.UserID] = comm
		fmt.Println("====== commission config: ", comm)
	}

	_comms := []*Commission{}

	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, ok := commMap[inviter.InviteeID]
		if ok {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm1: ", comm1)

		comm2, ok := commMap[inviter.InviterID]
		if ok {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm2: ", comm2)

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v (%v)", inviter.InviterID, percent2, inviter.InviteeID, percent1, comm1.GetGoodID())
		}

		if percent2.Cmp(percent1) == 0 {
			fmt.Println("----------- percent2 == percent1 ----------")
			commissionConfigID := uuid.Nil.String()
			if comm2 != nil {
				commissionConfigID = comm2.EntID
			}
			fmt.Println("------commissionConfigID: ", commissionConfigID)
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      commissionConfigID,
				CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				Amount:                  "0",
			})
			continue
		}

		amount := h.PaymentAmount
		if comm2.SettleMode == types.SettleMode_SettleWithGoodValue {
			amount = h.GoodValue
		}

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, ok := commMap[h.Inviters[len(h.Inviters)-1].InviteeID]
	if !ok {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, err
	}

	amount := h.PaymentAmount
	if commLast.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValue
	}

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		Amount:               amountLast,
	})

	return _comms, nil
}

func (h *calculateHandler) getInvites(userID string) uint32 {
	achivmentUser, ok := h.AchievementUsers[userID]
	if !ok {
		return uint32(0)
	}

	return achivmentUser.DirectInvites + achivmentUser.IndirectInvites
}

//nolint:dupl
func (h *calculateHandler) getAppGoodCommLevelConf(userID string) (*appgoodcommissionconfig.AppGoodCommissionConfig, bool, error) {
	invites := h.getInvites(userID)
	_comm := &appgoodcommissionconfig.AppGoodCommissionConfig{}
	useful := false
	amount := h.PaymentAmount.Mul(h.PaymentCoinUSDCurrency)
	if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValueUSD
	}
	consumeAmount := h.PaymentAmount.Mul(h.PaymentCoinUSDCurrency)
	achivmentUser, ok := h.AchievementUsers[userID]
	if ok {
		directConsumeAmount, err := decimal.NewFromString(achivmentUser.DirectConsumeAmount)
		if err != nil {
			return nil, false, err
		}
		inviteeConsumeAmount, err := decimal.NewFromString(achivmentUser.InviteeConsumeAmount)
		if err != nil {
			return nil, false, err
		}
		consumeAmount = directConsumeAmount.Add(inviteeConsumeAmount).Add(amount)
	}

	percent := decimal.NewFromInt(0)
	for i, comm := range h.AppGoodCommissionConfigs {
		if i == 0 {
			_comm = comm
		}
		if invites < comm.Invites {
			break
		}
		thresholdAmount, err := decimal.NewFromString(comm.GetThresholdAmount())
		if err != nil {
			return nil, false, err
		}
		if consumeAmount.Cmp(thresholdAmount) < 0 {
			break
		}
		_percent, err := decimal.NewFromString(comm.GetAmountOrPercent())
		if err != nil {
			return nil, false, err
		}
		if _percent.Cmp(percent) > 0 {
			percent = _percent
			_comm = comm
			useful = true
		}
	}
	if percent.Cmp(decimal.NewFromInt(0)) < 0 {
		return _comm, false, nil
	}
	return _comm, useful, nil
}

//nolint:dupl
func (h *calculateHandler) getAppCommLevelConf(userID string) (*appcommissionconfig.AppCommissionConfig, bool, error) {
	invites := h.getInvites(userID)
	_comm := &appcommissionconfig.AppCommissionConfig{}
	useful := false
	amount := h.PaymentAmount.Mul(h.PaymentCoinUSDCurrency)
	if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValueUSD
	}
	consumeAmount := h.PaymentAmount.Mul(h.PaymentCoinUSDCurrency)
	achivmentUser, ok := h.AchievementUsers[userID]
	if ok {
		directConsumeAmount, err := decimal.NewFromString(achivmentUser.DirectConsumeAmount)
		if err != nil {
			return nil, false, err
		}
		inviteeConsumeAmount, err := decimal.NewFromString(achivmentUser.InviteeConsumeAmount)
		if err != nil {
			return nil, false, err
		}
		consumeAmount = directConsumeAmount.Add(inviteeConsumeAmount).Add(amount)
	}

	percent := decimal.NewFromInt(0)
	for i, comm := range h.AppCommissionConfigs {
		if i == 0 {
			_comm = comm
		}
		if invites < comm.Invites {
			break
		}
		thresholdAmount, err := decimal.NewFromString(comm.GetThresholdAmount())
		if err != nil {
			return nil, false, err
		}
		if consumeAmount.Cmp(thresholdAmount) < 0 {
			break
		}
		_percent, err := decimal.NewFromString(comm.GetAmountOrPercent())
		if err != nil {
			return nil, false, err
		}
		if _percent.Cmp(percent) > 0 {
			percent = _percent
			_comm = comm
			useful = true
		}
	}
	if percent.Cmp(decimal.NewFromInt(0)) < 0 {
		return _comm, false, nil
	}
	return _comm, useful, nil
}

//nolint:dupl,funlen
func (h *Handler) CalculateByAppCommConfig(ctx context.Context) ([]*Commission, error) {
	fmt.Println("=========== calculate by app commission config ================")
	_comms := []*Commission{}
	handler := &calculateHandler{
		Handler: h,
	}
	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, comm1Useful, err := handler.getAppCommLevelConf(inviter.InviteeID)
		if err != nil {
			return nil, err
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm1: ", comm1)

		if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
			percent1 = decimal.NewFromInt(0)
		}

		comm2, comm2Useful, err := handler.getAppCommLevelConf(inviter.InviterID)
		if err != nil {
			return nil, err
		}

		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm2: ", comm2)

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			fmt.Println("----------- percent2 == percent1 ----------")
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm2.EntID,
				CommissionConfigType:    types.CommissionConfigType_AppCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				Amount:                  "0",
			})
			continue
		}

		amount := h.PaymentAmount
		if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
			amount = h.GoodValue
		}

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
		fmt.Println("------ direct commission user: ", h.Inviters[len(h.Inviters)-1].InviteeID, "; appConfigID: ", h.AppConfig.EntID)
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   uuid.Nil.String(),
			CommissionConfigType: types.CommissionConfigType_WithoutCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			Amount:               "0",
		})
		return _comms, nil
	}

	commLast, commLastUseful, err := handler.getAppCommLevelConf(h.Inviters[len(h.Inviters)-1].InviteeID)
	fmt.Println("------ order user: ", h.Inviters[len(h.Inviters)-1].InviteeID, "; commission config: ", commLast)
	if err != nil {
		return nil, err
	}

	if commLast == nil {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, err
	}
	if !commLastUseful {
		percent = decimal.NewFromInt(0)
	}

	amount := h.PaymentAmount
	if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValue
	}

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_AppCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		Amount:               amountLast,
	})

	return _comms, nil
}

//nolint:dupl,funlen
func (h *Handler) CalculateByAppGoodCommConfig(ctx context.Context) ([]*Commission, error) {
	fmt.Println("=========== calculate by app good commission config ================")
	_comms := []*Commission{}
	handler := &calculateHandler{
		Handler: h,
	}

	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, comm1Useful, err := handler.getAppGoodCommLevelConf(inviter.InviteeID)
		if err != nil {
			return nil, err
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm1: ", comm1)

		if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
			percent1 = decimal.NewFromInt(0)
		}

		comm2, comm2Useful, err := handler.getAppGoodCommLevelConf(inviter.InviterID)
		if err != nil {
			return nil, err
		}
		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}
		fmt.Println("------comm2: ", comm2)

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			fmt.Println("----------- percent2 == percent1 ----------")
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm2.EntID,
				CommissionConfigType:    types.CommissionConfigType_AppGoodCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				Amount:                  "0",
			})
			continue
		}

		amount := h.PaymentAmount
		if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
			amount = h.GoodValue
		}

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppGoodCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
		fmt.Println("------ direct commission user: ", h.Inviters[len(h.Inviters)-1].InviteeID, "; appConfigID: ", h.AppConfig.EntID)
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   uuid.Nil.String(),
			CommissionConfigType: types.CommissionConfigType_WithoutCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			Amount:               "0",
		})
		return _comms, nil
	}

	commLast, commLastUseful, err := handler.getAppGoodCommLevelConf(h.Inviters[len(h.Inviters)-1].InviteeID)
	if err != nil {
		return nil, err
	}
	fmt.Println("------ order user: ", h.Inviters[len(h.Inviters)-1].InviteeID, "; commission config: ", commLast)
	if commLast == nil {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, err
	}
	if !commLastUseful {
		percent = decimal.NewFromInt(0)
	}

	amount := h.PaymentAmount
	if h.AppConfig.SettleMode == types.SettleMode_SettleWithGoodValue {
		amount = h.GoodValue
	}

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_AppGoodCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		Amount:               amountLast,
	})

	return _comms, nil
}
