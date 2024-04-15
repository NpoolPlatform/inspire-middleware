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

type createHandler struct {
	*Handler
}

func (h *createHandler) createOrAddInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
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

func (h *createHandler) addInvites(ctx context.Context, tx *ent.Tx) error {
	handler, err := NewHandler(ctx)
	if err != nil {
		return err
	}

	handler.AppID = h.AppID
	handler.InviteeID = h.InviterID

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
	req := &registrationcrud.Req{
		AppID:     h.AppID,
		InviterID: h.InviterID,
	}
	if err := h.createOrAddInvites(ctx, tx, req); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateRegistration(ctx context.Context) (*npool.Registration, error) {
	if err := h.validateInvitationCode(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRegistration, *h.AppID, *h.InviteeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &registrationcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		InviteeID: &cruder.Cond{Op: cruder.EQ, Val: *h.InviteeID},
	}
	exist, err := h.ExistRegistrationConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	handler := &createHandler{
		Handler: h,
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := registrationcrud.CreateSet(
			tx.Registration.Create(),
			&registrationcrud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				InviterID: h.InviterID,
				InviteeID: h.InviteeID,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if err := handler.addInvites(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRegistration(ctx)
}
