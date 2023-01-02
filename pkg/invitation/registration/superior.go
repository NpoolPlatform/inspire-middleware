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

func CreateSuperiorProcedure(ctx context.Context) error {
	conn, err := mysql.GetConn()
	if err != nil {
		return err
	}

	const drop_procedure = `DROP PROCEDURE IF EXISTS get_superiores`
	_, err = conn.ExecContext(ctx, drop_procedure)
	if err != nil {
		return err
	}

	const procedure = `
		CREATE PROCEDURE get_superiores (IN invitees TEXT)
		BEGIN
		  DECLARE superiores TEXT;
		  DECLARE my_invitees TEXT;
		  SET superiores = 'N/A';
		  SET my_invitees = invitees;
		  WHILE my_invitees is not null DO
		    if superiores = 'N/A' THEN
			  SET superiores = my_invitees;
			else
			  SET superiores = CONCAT(superiores, ',', my_invitees);
			END if;
		    SELECT GROUP_CONCAT(inviter_id) INTO my_invitees FROM registrations WHERE FIND_IN_SET(invitee_id, my_invitees);
		  END WHILE;
		  SELECT superiores;
		END
	`
	_, err = conn.ExecContext(ctx, procedure)
	if err != nil {
		return err
	}

	return nil
}

func GetSuperiores(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Registration, uint32, error) {
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
				// TODO: get parent recursively
				inviteeIDs := strings.Join(conds.GetInviteeIDs().GetValue(), ",")
				callProcedure := fmt.Sprintf("CALL get_superiores(\"%v\")", inviteeIDs)
				s.
					QueryContext(callProcedure)
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
