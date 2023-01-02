//nolint:dupl
package registration

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entreg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/registration"

	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
)

func CreateSubordinateProcedure(ctx context.Context) error {
	const procedure = `
		CREATE PROCEDURE get_subordinates (IN inviters TEXT)
		BEGIN
		  DECLARE subordinates TEXT;
		  DECLARE invitees TEXT;
		  SET subordinates = 'N/A';
		  SET invitees = inviters;
		  WHILE invitees is not null DO
		    if subordinates = 'N/A' THEN
			  SET subordinates = invitees;
			else
			  SET subordinates = CONCAT(subordinates, ',', invitees);
			END if;
		    SELECT GROUP_CONCAT(invitee_id) INTO invitees FROM registrations WHERE FIND_IN_SET(inviter_id, invitees);
		  END WHILE;
		  SELECT subordinates;
		END
	`

	conn, err := mysql.GetConn()
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, procedure, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetSubordinates(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Registration, uint32, error) {
	var infos []*mgrpb.Registration
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		sel := stm.
			Select(
				entreg.FieldID,
				entreg.FieldAppID,
				entreg.FieldInviterID,
				entreg.FieldInviteeID,
				entreg.FieldCreatedAt,
				entreg.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				// TODO: get child recursively
			})

		_total, err := sel.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return sel.
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
