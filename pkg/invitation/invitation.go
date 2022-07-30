package invitation

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
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

	infos = []*npool.Invitation{}

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

		return stm.
			Select(
				registrationinvitation.FieldInviterID,
				registrationinvitation.FieldInviteeID,
				registrationinvitation.FieldCreateAt,
			).
			Order(ent.Desc(registrationinvitation.FieldUpdateAt)).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify().
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	logger.Sugar().Infow("GetInvitees", "infos", infos, "total", total, "inviters", inviters, "app", appID, "error", err)

	return infos, total, nil
}

func GetInviters(ctx context.Context, appID, userID string, offset, limit int32) ([]*npool.Invitation, uint32, error) {
	return nil, 0, fmt.Errorf("NOT IMPLEMENTED")
}
