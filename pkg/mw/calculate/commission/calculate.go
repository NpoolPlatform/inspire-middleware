package commission

import (
	"context"
	"fmt"

	achievementuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	achievementusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
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

func (h *Handler) Calculate(ctx context.Context) ([]*Commission, error) {
	commMap := map[string]*npool.Commission{}
	for _, comm := range h.Commissions {
		commMap[comm.UserID] = comm
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

		comm2, ok := commMap[inviter.InviterID]
		if ok {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v (%v)", inviter.InviterID, percent2, inviter.InviteeID, percent1, comm1.GetGoodID())
		}

		if percent2.Cmp(percent1) == 0 {
			continue
		}

		amount := h.PaymentAmount
		if comm2.SettleMode == types.SettleMode_SettleWithGoodValue {
			amount = h.GoodValue
		}

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return _comms, nil
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
		return _comms, nil
	}

	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   commLast.EntID,
			CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			Amount:               amount.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	return _comms, nil
}

func (h *calculateHandler) getInvites(ctx context.Context, userID string) (uint32, error) {
	handler, err := achievementuser1.NewHandler(
		ctx,
		achievementuser1.WithConds(&achievementusermwpb.Conds{
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppConfig.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: userID},
		}),
	)
	if err != nil {
		return 0, nil
	}
	achivmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return uint32(0), err
	}
	if len(achivmentUsers) == 0 {
		return uint32(0), nil
	}

	return achivmentUsers[0].DirectInvites + achivmentUsers[0].IndirectInvites, nil
}

//nolint:dupl
func (h *calculateHandler) getAppGoodCommLevelConf(ctx context.Context, userID string) (*appgoodcommissionconfig.AppGoodCommissionConfig, bool, error) {
	invites, err := h.getInvites(ctx, userID)
	if err != nil {
		return nil, false, err
	}
	_comm := &appgoodcommissionconfig.AppGoodCommissionConfig{}
	useful := false
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
		if h.PaymentAmount.Cmp(thresholdAmount) < 0 {
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
func (h *calculateHandler) getAppCommLevelConf(ctx context.Context, userID string) (*appcommissionconfig.AppCommissionConfig, bool, error) {
	invites, err := h.getInvites(ctx, userID)
	if err != nil {
		return nil, false, err
	}

	_comm := &appcommissionconfig.AppCommissionConfig{}
	useful := false
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
		if h.PaymentAmount.Cmp(thresholdAmount) < 0 {
			continue
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

		comm1, comm1Useful, err := handler.getAppCommLevelConf(ctx, inviter.InviteeID)
		if err != nil {
			return nil, err
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		comm2, comm2Useful, err := handler.getAppCommLevelConf(ctx, inviter.InviterID)
		if err != nil {
			return nil, err
		}

		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm1.EntID,
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
			CommissionConfigID:      comm1.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, commLastUseful, err := handler.getAppCommLevelConf(ctx, h.Inviters[len(h.Inviters)-1].InviteeID)
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

		comm1, comm1Useful, err := handler.getAppGoodCommLevelConf(ctx, inviter.InviteeID)
		if err != nil {
			return nil, err
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		comm2, comm2Useful, err := handler.getAppGoodCommLevelConf(ctx, inviter.InviterID)
		if err != nil {
			return nil, err
		}
		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, err
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, fmt.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm1.EntID,
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
			CommissionConfigID:      comm1.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppGoodCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, commLastUseful, err := handler.getAppGoodCommLevelConf(ctx, h.Inviters[len(h.Inviters)-1].InviteeID)
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

	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   commLast.EntID,
			CommissionConfigType: types.CommissionConfigType_AppGoodCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			Amount:               amount.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	return _comms, nil
}
