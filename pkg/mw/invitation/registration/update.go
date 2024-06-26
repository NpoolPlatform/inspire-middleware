package registration

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	achievementusercrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/user"
	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	achievementuser1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user"
	common1 "github.com/NpoolPlatform/inspire-middleware/pkg/mw/achievement/user/common"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	achievementusermwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/user"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	registration   *npool.Registration
	subInviterIDs  []string
	addInviterIDs  []string
	inviteeInvites uint32
}

func (h *updateHandler) subAchievementInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
	stm, err := achievementusercrud.SetQueryConds(
		tx.AchievementUser.Query(),
		&achievementusercrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *req.InviterID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}
	if info == nil {
		return nil
	}

	_req := &achievementusercrud.Req{}

	invites := uint32(1)
	directInvites := info.DirectInvites
	indirectInvites := info.IndirectInvites
	oldInviterID := uuid.MustParse(h.registration.InviterID)
	if oldInviterID == *req.InviterID {
		if directInvites != uint32(0) {
			directInvites -= invites
			_req.DirectInvites = &directInvites
		}
		if indirectInvites != uint32(0) {
			indirectInvites -= h.inviteeInvites
			_req.IndirectInvites = &indirectInvites
		}
	}
	if oldInviterID != *req.InviterID && indirectInvites != uint32(0) {
		indirectInvites = indirectInvites - invites - h.inviteeInvites
		_req.IndirectInvites = &indirectInvites
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *updateHandler) getTotalInvites(ctx context.Context) error {
	handler, err := achievementuser1.NewHandler(
		ctx,
		common1.WithConds(&achievementusermwpb.Conds{
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.registration.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: h.registration.InviteeID},
		}),
		common1.WithLimit(int32(1)),
	)
	if err != nil {
		return nil
	}
	achivmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(achivmentUsers) == 0 {
		return nil
	}
	h.inviteeInvites = achivmentUsers[0].DirectInvites + achivmentUsers[0].IndirectInvites
	return nil
}

func (h *updateHandler) createOrAddInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
	key := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateInspireAchievement,
		*req.AppID,
		*req.InviterID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return wlog.WrapError(err)
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	stm, err := achievementusercrud.SetQueryConds(
		tx.AchievementUser.Query(),
		&achievementusercrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *req.InviterID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}

	_req := &achievementusercrud.Req{
		AppID:  req.AppID,
		UserID: req.InviterID,
	}

	invites := uint32(1)
	if h.InviterID == req.InviterID {
		_req.DirectInvites = &invites
	} else {
		_req.IndirectInvites = &invites
	}

	if info == nil {
		if _, err = achievementusercrud.CreateSet(
			tx.AchievementUser.Create(),
			_req,
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	}

	directInvites := info.DirectInvites
	indirectInvites := info.IndirectInvites

	if h.InviterID.String() == req.InviterID.String() {
		directInvites += invites
		_req.DirectInvites = &directInvites
		indirectInvites += h.inviteeInvites
		_req.IndirectInvites = &indirectInvites
	} else {
		indirectInvites = indirectInvites + invites + h.inviteeInvites
		_req.IndirectInvites = &indirectInvites
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *updateHandler) addInvites(ctx context.Context, tx *ent.Tx) error {
	for _, inviter := range h.addInviterIDs {
		if inviter == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.createOrAddInvites(ctx, tx, req); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) subInvites(ctx context.Context, tx *ent.Tx) error {
	for _, inviter := range h.subInviterIDs {
		if inviter == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.subAchievementInvites(ctx, tx, req); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) getInviters(ctx context.Context, inviterID *uuid.UUID) ([]string, error) {
	handler, err := NewHandler(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.AppID = h.AppID
	handler.InviteeID = inviterID

	_, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return inviterIDs, nil
}

func (h *Handler) UpdateRegistration(ctx context.Context) (*npool.Registration, error) {
	info, err := h.GetRegistration(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("registration not found")
	}
	if info.InviterID == h.InviterID.String() || info.InviteeID == h.InviterID.String() {
		return nil, wlog.Errorf("invalid inviterid")
	}

	if err := h.validateInvitationCode(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &updateHandler{
		Handler:        h,
		registration:   info,
		subInviterIDs:  []string{},
		addInviterIDs:  []string{},
		inviteeInvites: uint32(0),
	}

	if err := handler.getTotalInvites(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	inviteeID := uuid.MustParse(handler.registration.InviterID)
	subInviters, err := handler.getInviters(ctx, &inviteeID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	addInviters, err := handler.getInviters(ctx, h.InviterID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.subInviterIDs = subInviters
	handler.addInviterIDs = addInviters

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.subInvites(ctx, tx); err != nil {
			return nil
		}

		if _, err := registrationcrud.UpdateSet(
			tx.Registration.UpdateOneID(*h.ID),
			&registrationcrud.Req{
				InviterID: h.InviterID,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.addInvites(ctx, tx); err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetRegistration(ctx)
}
