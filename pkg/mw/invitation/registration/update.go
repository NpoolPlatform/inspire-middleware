package registration

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	achievementusercrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/achievement/user"
	registrationcrud "github.com/NpoolPlatform/inspire-middleware/pkg/crud/invitation/registration"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/invitation/registration"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	registration *npool.Registration
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
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		return nil
	}

	_req := &achievementusercrud.Req{}

	invites := uint32(1)
	directInvites := info.DirectInvites
	indirectInvites := info.IndirectInvites
	if h.InviterID == req.InviterID && directInvites != uint32(0) {
		directInvites -= invites
		_req.DirectInvites = &directInvites
	}
	if h.InviterID != req.InviterID && indirectInvites != uint32(0) {
		indirectInvites -= invites
		_req.IndirectInvites = &indirectInvites
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

//nolint:dupl
func (h *updateHandler) createOrAddInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
	key := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateInspireAchievement,
		*req.AppID,
		*req.InviterID,
	)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
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
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
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
			return err
		}
		return nil
	}

	directInvites := info.DirectInvites
	indirectInvites := info.IndirectInvites
	if h.InviterID == req.InviterID {
		invites += directInvites
		_req.DirectInvites = &invites
	} else {
		invites += indirectInvites
		_req.IndirectInvites = &invites
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

//nolint:dupl
func (h *updateHandler) addInvites(ctx context.Context, tx *ent.Tx) error {
	handler, err := NewHandler(ctx)
	if err != nil {
		return err
	}

	inviteeID := uuid.MustParse(h.registration.InviteeID)
	handler.AppID = h.AppID
	handler.InviteeID = &inviteeID

	inviters, _, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return err
	}

	for _, inviter := range inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter.InviterID)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.createOrAddInvites(ctx, tx, req); err != nil {
			return err
		}
	}

	return nil
}

//nolint:dupl
func (h *updateHandler) subInvites(ctx context.Context, tx *ent.Tx) error {
	handler, err := NewHandler(ctx)
	if err != nil {
		return err
	}

	inviteeID := uuid.MustParse(h.registration.InviteeID)
	handler.AppID = h.AppID
	handler.InviteeID = &inviteeID

	inviters, _, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return err
	}

	for _, inviter := range inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter.InviterID)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.subAchievementInvites(ctx, tx, req); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) UpdateRegistration(ctx context.Context) (*npool.Registration, error) {
	info, err := h.GetRegistration(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("registration not found")
	}
	if info.InviterID == h.InviterID.String() || info.InviteeID == h.InviterID.String() {
		return nil, fmt.Errorf("invalid inviterid")
	}

	if err := h.validateInvitationCode(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler:      h,
		registration: info,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := registrationcrud.UpdateSet(
			tx.Registration.UpdateOneID(*h.ID),
			&registrationcrud.Req{
				InviterID: h.InviterID,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if err := handler.subInvites(ctx, tx); err != nil {
			return nil
		}
		if err := handler.addInvites(ctx, tx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRegistration(ctx)
}
