package achievement

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	constant "github.com/NpoolPlatform/inspire-middleware/pkg/const"
	achievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement"
	statement1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/statement"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"
	statementmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type expropriateHandler struct {
	*Handler
	statements   []*statementmwpb.Statement
	achievements map[string]*npool.Achievement
}

func (h *expropriateHandler) getStatements(ctx context.Context) error {
	handler, err := statement1.NewHandler(
		ctx,
		statement1.WithConds(&statementmwpb.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID.String()},
		}),
		statement1.WithLimit(constant.DefaultRowLimit),
	)
	if err != nil {
		return err
	}

	for {
		statements, _, err := handler.GetStatements(ctx)
		if err != nil {
			return err
		}
		if len(statements) == 0 {
			break
		}
		h.statements = append(h.statements, statements...)
		handler.Offset += handler.Limit
	}

	appMap := map[string]struct{}{}
	goodMap := map[string]struct{}{}
	coinMap := map[string]struct{}{}

	for _, statement := range h.statements {
		appMap[statement.AppID] = struct{}{}
		goodMap[statement.GoodID] = struct{}{}
		coinMap[statement.CoinTypeID] = struct{}{}
	}
	if len(appMap) > 1 || len(goodMap) > 1 || len(coinMap) > 1 {
		return fmt.Errorf("invalid orderid")
	}

	return nil
}

func (h *expropriateHandler) getAchievements(ctx context.Context) error {
	h.Conds = &achievementcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(h.statements[0].AppID)},
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(h.statements[0].GoodID)},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: uuid.MustParse(h.statements[0].CoinTypeID)},
	}
	ids := []uuid.UUID{}
	for _, statement := range h.statements {
		ids = append(ids, uuid.MustParse(statement.UserID))
	}
	h.Conds.UserIDs = &cruder.Cond{Op: cruder.IN, Val: ids}
	h.Offset = 0
	h.Limit = int32(len(ids))

	achievements, _, err := h.GetAchievements(ctx)
	if err != nil {
		return err
	}

	for _, achievement := range achievements {
		h.achievements[achievement.UserID] = achievement
	}

	return nil
}

func (h *expropriateHandler) expropriate(ctx context.Context, tx *ent.Tx) error { //nolint
	for _, statement := range h.statements {
		achievement, ok := h.achievements[statement.UserID]
		if !ok {
			continue
		}
		orderAmount, err := decimal.NewFromString(statement.USDAmount)
		if err != nil {
			return err
		}
		totalAmount, err := decimal.NewFromString(achievement.TotalAmount)
		if err != nil {
			return err
		}
		if orderAmount.Cmp(totalAmount) > 0 {
			return fmt.Errorf("invalid amount")
		}
		currency, err := decimal.NewFromString(statement.PaymentCoinUSDCurrency)
		if err != nil {
			return err
		}
		orderCommission, err := decimal.NewFromString(statement.Commission)
		if err != nil {
			return err
		}
		orderCommission = orderCommission.Mul(currency)
		totalCommission, err := decimal.NewFromString(achievement.TotalCommission)
		if err != nil {
			return err
		}
		if orderCommission.Cmp(totalCommission) > 0 {
			return fmt.Errorf("invalid commission")
		}
		orderUnits, err := decimal.NewFromString(statement.Units)
		if err != nil {
			return err
		}
		totalUnits, err := decimal.NewFromString(achievement.TotalUnits)
		if err != nil {
			return err
		}
		if orderUnits.Cmp(totalUnits) > 0 {
			return fmt.Errorf("invalid units")
		}
		selfAmount, err := decimal.NewFromString(achievement.SelfAmount)
		if err != nil {
			return err
		}
		if statement.SelfOrder && orderAmount.Cmp(selfAmount) > 0 {
			return fmt.Errorf("invalid amount")
		}
		selfCommission, err := decimal.NewFromString(achievement.SelfCommission)
		if err != nil {
			return err
		}
		if statement.SelfOrder && orderCommission.Cmp(selfCommission) > 0 {
			return fmt.Errorf("invalid commission")
		}
		selfUnits, err := decimal.NewFromString(achievement.SelfUnits)
		if err != nil {
			return err
		}
		if statement.SelfOrder && orderUnits.Cmp(selfUnits) > 0 {
			return fmt.Errorf("invalid units")
		}

		totalAmount = totalAmount.Sub(orderAmount)
		totalCommission = totalCommission.Sub(orderCommission)
		totalUnits = totalUnits.Sub(orderUnits)
		if statement.SelfOrder {
			selfAmount = selfAmount.Sub(orderAmount)
			selfCommission = selfCommission.Sub(orderCommission)
			selfUnits = selfUnits.Sub(orderUnits)
		}

		id, err := uuid.Parse(achievement.ID)
		if err != nil {
			return err
		}

		if _, err := achievementcrud.UpdateSet(
			tx.Achievement.UpdateOneID(id),
			&achievementcrud.Req{
				TotalAmount:     &totalAmount,
				TotalCommission: &totalCommission,
				TotalUnits:      &totalUnits,
				SelfAmount:      &selfAmount,
				SelfCommission:  &selfCommission,
				SelfUnits:       &selfUnits,
			},
		).Save(ctx); err != nil {
			return err
		}

		id1, err := uuid.Parse(statement.ID)
		if err != nil {
			return err
		}

		if _, err := tx.
			Statement.
			UpdateOneID(id1).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) ExpropriateAchievement(ctx context.Context) error {
	handler := &expropriateHandler{
		Handler:      h,
		achievements: map[string]*npool.Achievement{},
	}
	if err := handler.getStatements(ctx); err != nil {
		return err
	}
	if len(handler.statements) == 0 {
		return nil
	}
	if err := handler.getAchievements(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.expropriate(_ctx, tx)
	})
}
