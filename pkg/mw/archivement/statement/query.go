package statement

import (
	"context"
	"fmt"

	statementcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/archivement/statement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	entarchivementdetail "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/archivementdetail"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/statement"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ArchivementDetailSelect
	infos     []*npool.Statement
	total     uint32
}

func (h *queryHandler) queryStatement(cli *ent.Client) {
	h.stmSelect = cli.
		ArchivementDetail.
		Query().
		Where(
			entarchivementdetail.ID(*h.ID),
			entarchivementdetail.DeletedAt(0),
		).
		Select()
}

func (h *queryHandler) queryStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := statementcrud.SetQueryConds(cli.ArchivementDetail.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.stmSelect = stm.Select()
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetStatement(ctx context.Context) (*npool.Statement, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryStatement(cli)
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

	return handler.infos[0], nil
}

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(ctx, cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
