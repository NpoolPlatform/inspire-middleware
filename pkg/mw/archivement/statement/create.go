package statement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	archivementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/archivement"
	statementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/archivement/statement"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/statement"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	if _, err := statementcrud.CreateSet(
		tx.ArchivementDetail.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createOrAddArchivement(ctx context.Context, tx *ent.Tx) error {
	key := fmt.Sprintf(
		"%v:%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateInspireArchivement,
		*h.AppID,
		*h.UserID,
		*h.GoodID,
		*h.CoinTypeID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	stm, err := archivementcrud.SetQueryConds(
		tx.ArchivementGeneral.Query(),
		&archivementcrud.Conds{
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		},
	)
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &archivementcrud.Req{
		AppID:           h.AppID,
		UserID:          h.UserID,
		GoodID:          h.GoodID,
		CoinTypeID:      h.CoinTypeID,
		TotalAmount:     h.Amount,
		TotalUnits:      h.Units,
		TotalCommission: h.Commission,
	}
	if h.SelfOrder != nil && *h.SelfOrder {
		req.SelfAmount = h.Amount
		req.SelfUnits = h.Units
		req.SelfCommission = h.Commission
	}

	if info == nil {
		if _, err = archivementcrud.CreateSet(
			tx.ArchivementGeneral.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := archivementcrud.UpdateSet(
		tx.ArchivementGeneral.UpdateOneID(info.ID),
		req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateStatement(ctx context.Context) (*npool.Statement, error) {
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	if h.UserID == nil {
		return nil, fmt.Errorf("invalid userid")
	}
	if h.DirectContributorID == nil {
		return nil, fmt.Errorf("invalid directcontributorid")
	}
	if h.GoodID == nil {
		return nil, fmt.Errorf("invalid goodid")
	}
	if h.OrderID == nil {
		return nil, fmt.Errorf("invalid orderid")
	}
	if h.PaymentID == nil {
		return nil, fmt.Errorf("invalid paymentid")
	}
	if h.CoinTypeID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}
	if h.PaymentCoinTypeID == nil {
		return nil, fmt.Errorf("invalid paymentcointypeid")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateInspireArchivementStatement, *h.AppID, *h.OrderID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	// TODO: Check order
	h.Conds = &statementcrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		OrderID: &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID},
	}
	exist, err := h.ExistStatementConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createStatement(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createOrAddArchivement(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStatement(ctx)
}
