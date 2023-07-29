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

func (h *createHandler) createStatement(ctx context.Context, tx *ent.Tx, req *statementcrud.Req) error {
	if _, err := statementcrud.CreateSet(
		tx.ArchivementDetail.Create(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createOrAddArchivement(ctx context.Context, tx *ent.Tx, req *statementcrud.Req) error {
	key := fmt.Sprintf(
		"%v:%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateInspireArchivement,
		*req.AppID,
		*req.UserID,
		*req.GoodID,
		*req.CoinTypeID,
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
			AppID:      &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID:     &cruder.Cond{Op: cruder.EQ, Val: *req.UserID},
			GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *req.GoodID},
			CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *req.CoinTypeID},
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

	_req := &archivementcrud.Req{
		AppID:           req.AppID,
		UserID:          req.UserID,
		GoodID:          req.GoodID,
		CoinTypeID:      req.CoinTypeID,
		TotalAmount:     req.Amount,
		TotalUnits:      req.Units,
		TotalCommission: req.Commission,
	}
	if req.SelfOrder != nil && *req.SelfOrder {
		_req.SelfAmount = req.Amount
		_req.SelfUnits = req.Units
		_req.SelfCommission = req.Commission
	}

	if info == nil {
		if _, err = archivementcrud.CreateSet(
			tx.ArchivementGeneral.Create(),
			_req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	// TODO: add amount

	if _, err := archivementcrud.UpdateSet(
		tx.ArchivementGeneral.UpdateOneID(info.ID),
		_req,
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
		if err := handler.createStatement(_ctx, tx, &handler.Req); err != nil {
			return err
		}
		if err := handler.createOrAddArchivement(_ctx, tx, &handler.Req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStatement(ctx)
}

func (h *Handler) CreateStatements(ctx context.Context) ([]*npool.Statement, error) {
	ids := []uuid.UUID{}

	handler := &createHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID == nil {
				req.ID = &id
			}
			if err := handler.createStatement(_ctx, tx, req); err != nil {
				return err
			}
			if err := handler.createOrAddArchivement(_ctx, tx, req); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &statementcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err := h.GetStatements(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
