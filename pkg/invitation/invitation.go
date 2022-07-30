package invitation

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/inspire-middleware/pkg/db"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/inspire/invitation"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"

	"github.com/google/uuid"
)

func GetInvitees(
	ctx context.Context,
	appID string, inviterIDs []string,
	offset, limit int32,
) (
	infos []*npool.Invitation,
	total uint32,
	err error,
) {
	inviters := []uuid.UUID{}
	for _, inviter := range inviterIDs {
		inviters = append(inviters, uuid.MustParse(inviter))
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Debug().
			RegistrationInvitation.
			Query().
			Select(
				registrationinvitation.FieldInviterID,
				registrationinvitation.FieldInviterID,
				registrationinvitation.FieldCreateAt,
			).
			Where(
				registrationinvitation.AppID(uuid.MustParse(appID)),
				registrationinvitation.InviterIDIn(inviters...),
			)
		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return stm.
			Order(ent.Desc(registrationinvitation.FieldUpdateAt)).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify().
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func GetInviters(ctx context.Context, appID, userID string, offset, limit int32) ([]*npool.Invitation, uint32, error) {
	return nil, 0, fmt.Errorf("NOT IMPLEMENTED")
}
