package orderstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entgoodachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodachievement"
	entgoodcoinachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodcoinachievement"
)

type achievementQueryHandler struct {
	*Handler
	entGoodAchievement     *ent.GoodAchievement
	entGoodCoinAchievement *ent.GoodCoinAchievement
}

func (h *achievementQueryHandler) getGoodAchievement(ctx context.Context, cli *ent.Client, must bool) (err error) {
	h.entGoodAchievement, err = cli.
		GoodAchievement.
		Query().
		Where(
			entgoodachievement.UserID(*h.UserID),
			entgoodachievement.AppGoodID(*h.AppGoodID),
		).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) && !must {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *achievementQueryHandler) getGoodCoinAchievement(ctx context.Context, cli *ent.Client, must bool) (err error) {
	h.entGoodCoinAchievement, err = cli.
		GoodCoinAchievement.
		Query().
		Where(
			entgoodcoinachievement.UserID(*h.UserID),
			entgoodcoinachievement.GoodCoinTypeID(*h.GoodCoinTypeID),
		).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) && !must {
		return nil
	}
	return err
}

func (h *achievementQueryHandler) _getAchievement(ctx context.Context, must bool) error {
	if h.AppGoodID == nil || h.UserID == nil || h.GoodCoinTypeID == nil {
		return wlog.Errorf("invalid goodid")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getGoodAchievement(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		return h.getGoodCoinAchievement(_ctx, cli, must)
	})
}

func (h *achievementQueryHandler) getAchievement(ctx context.Context) error {
	return h._getAchievement(ctx, false)
}

func (h *achievementQueryHandler) requireAchievement(ctx context.Context) error {
	return h._getAchievement(ctx, true)
}
