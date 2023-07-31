package achievement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/achievement"

	achievementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AchievementSelect
	infos     []*npool.Achievement
	total     uint32
}

func (h *queryHandler) selectAchievement(stm *ent.AchievementQuery) {
	h.stmSelect = stm.Select(
		entachievement.FieldID,
		entachievement.FieldAppID,
		entachievement.FieldUserID,
		entachievement.FieldGoodID,
		entachievement.FieldCoinTypeID,
		entachievement.FieldTotalAmount,
		entachievement.FieldSelfAmount,
		entachievement.FieldTotalUnitsV1,
		entachievement.FieldSelfUnitsV1,
		entachievement.FieldTotalCommission,
		entachievement.FieldSelfCommission,
		entachievement.FieldCreatedAt,
		entachievement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAchievement(cli *ent.Client) {
	h.selectAchievement(
		cli.Achievement.
			Query().
			Where(
				entachievement.ID(*h.ID),
				entachievement.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryAchievements(ctx context.Context, cli *ent.Client) error {
	stm, err := achievementcrud.SetQueryConds(cli.Achievement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAchievement(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.TotalAmount)
		if err != nil {
			info.TotalAmount = decimal.NewFromInt(0).String()
		} else {
			info.TotalAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfAmount)
		if err != nil {
			info.SelfAmount = decimal.NewFromInt(0).String()
		} else {
			info.SelfAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.TotalUnits)
		if err != nil {
			info.TotalUnits = decimal.NewFromInt(0).String()
		} else {
			info.TotalUnits = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfUnits)
		if err != nil {
			info.SelfUnits = decimal.NewFromInt(0).String()
		} else {
			info.SelfUnits = amount.String()
		}
		amount, err = decimal.NewFromString(info.TotalCommission)
		if err != nil {
			info.TotalCommission = decimal.NewFromInt(0).String()
		} else {
			info.TotalCommission = amount.String()
		}
		amount, err = decimal.NewFromString(info.SelfCommission)
		if err != nil {
			info.SelfCommission = decimal.NewFromInt(0).String()
		} else {
			info.SelfCommission = amount.String()
		}
	}
}

func (h *Handler) GetAchievement(ctx context.Context) (*npool.Achievement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAchievement(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetAchievements(ctx context.Context) ([]*npool.Achievement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Achievement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAchievements(_ctx, cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
