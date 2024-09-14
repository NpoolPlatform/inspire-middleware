package config

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	configcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/coin/config"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"
	coinallocated1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/coin/allocated"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinallocatedmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coin/allocated"
	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCoinConfig(ctx context.Context, cli *ent.Client) error {
	if _, err := configcrud.UpdateSet(
		cli.CoinConfig.UpdateOneID(*h.ID),
		&configcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *deleteHandler) checkCoinAllocated(ctx context.Context) error {
	conds := &coinallocatedmwpb.Conds{
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
		CoinConfigID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID.String()},
	}
	handler, err := coinallocated1.NewHandler(
		ctx,
		coinallocated1.WithConds(conds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	exist, err := handler.ExistCoinAllocatedConds(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if exist {
		return wlog.Errorf("exist coinallocated")
	}
	return nil
}

func (h *Handler) DeleteCoinConfig(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCoinConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}
	h.AppID = func() *uuid.UUID { s := uuid.MustParse(info.AppID); return &s }()
	h.EntID = func() *uuid.UUID { s := uuid.MustParse(info.EntID); return &s }()
	if err := handler.checkCoinAllocated(ctx); err != nil {
		return err
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteCoinConfig(_ctx, cli)
	})
}
