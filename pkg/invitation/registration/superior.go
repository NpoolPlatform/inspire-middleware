//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"strings"

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

	const procedure = `
	DROP PROCEDURE IF EXISTS get_superiores;
	SET SESSION GROUP_CONCAT_MAX_LEN = 102400;
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
		    SELECT GROUP_CONCAT(DISTINCT inviter_id) INTO my_invitees FROM registrations WHERE FIND_IN_SET(invitee_id, my_invitees);
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

	rawClient, err := db.Client()
	if err != nil {
		return nil, 0, err
	}

	inviteeIDs := strings.Join(conds.GetInviteeIDs().GetValue(), ",")
	rows, err := rawClient.QueryContext(ctx, fmt.Sprintf("CALL get_superiores(\"%v\")", inviteeIDs))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	superiores := ""
	for rows.Next() {
		if err := rows.Scan(&superiores); err != nil {
			return nil, 0, err
		}
	}

	_inviteeIDs := strings.Split(superiores, ",")
	conds.InviteeIDs.Value = _inviteeIDs

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return stm.
			Offset(int(offset)).
			Limit(int(limit)).
			Select(
				entreg.FieldID,
				entreg.FieldAppID,
				entreg.FieldInviterID,
				entreg.FieldInviteeID,
				entreg.FieldCreatedAt,
				entreg.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
