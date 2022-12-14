//nolint:dupl
package invitation

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"

	"github.com/google/uuid"
)

func GetInvitees(
	ctx context.Context, appID string, inviterIDs []string, offset, limit int32,
) (
	infos []*Invitation, total uint32, err error,
) {
	inviters := []uuid.UUID{}
	for _, inviter := range inviterIDs {
		inviters = append(inviters, uuid.MustParse(inviter))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			RegistrationInvitation.
			Query().
			Where(
				registrationinvitation.AppID(uuid.MustParse(appID)),
				registrationinvitation.InviterIDIn(inviters...),
			)
		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return join(stm, offset, limit).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func GetInviters(
	ctx context.Context, appID string, inviteeIDs []string, offset, limit int32,
) (
	infos []*Invitation, total uint32, err error,
) {
	invitees := []uuid.UUID{}
	for _, invitee := range inviteeIDs {
		invitees = append(invitees, uuid.MustParse(invitee))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			RegistrationInvitation.
			Query().
			Where(
				registrationinvitation.AppID(uuid.MustParse(appID)),
				registrationinvitation.InviteeIDIn(invitees...),
			)
		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return join(stm, offset, limit).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.RegistrationInvitationQuery, offset, limit int32) *ent.RegistrationInvitationSelect {
	return stm.
		Select(
			registrationinvitation.FieldInviterID,
			registrationinvitation.FieldInviteeID,
			registrationinvitation.FieldCreateAt,
		).
		Order(ent.Desc(registrationinvitation.FieldUpdateAt)).
		Offset(int(offset)).
		Limit(int(limit)).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(userinvitationcode.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(registrationinvitation.FieldInviteeID),
					t1.C(userinvitationcode.FieldUserID),
				).
				AppendSelect(
					sql.As(t1.C(userinvitationcode.FieldInvitationCode), "invitation_code"),
				)
		})
}
