//nolint:dupl
package registration

import (
	"context"
	"fmt"
	"strings"

	mgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"

	"github.com/NpoolPlatform/inspire-manager/pkg/db"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"

	entreg "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/registration"

	crud "github.com/NpoolPlatform/inspire-manager/pkg/crud/invitation/registration"
)

func CreateSubordinateProcedure(ctx context.Context) error {
	conn, err := mysql.GetConn()
	if err != nil {
		return err
	}

	const procedure = `
		DROP PROCEDURE IF EXISTS get_subordinates;
		SET SESSION GROUP_CONCAT_MAX_LEN = 1024000;
		CREATE PROCEDURE get_subordinates (IN inviters TEXT)
		BEGIN
		  DECLARE subordinates TEXT;
		  DECLARE my_inviters TEXT;
		  SET subordinates = 'N/A';
		  SET my_inviters = inviters;
		  WHILE my_inviters is not null DO
		    if subordinates = 'N/A' THEN
			  SET subordinates = my_inviters;
			else
			  SET subordinates = CONCAT(subordinates, ',', inviters);
			END if;
		    SELECT GROUP_CONCAT(DISTINCT invitee_id) INTO my_inviters FROM registrations WHERE FIND_IN_SET(inviter_id, my_inviters);
		  END WHILE;
		  SELECT subordinates;
		END
	`
	_, err = conn.ExecContext(ctx, procedure)
	if err != nil {
		return err
	}

	return nil
}

func GetSubordinates(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.Registration, uint32, error) {
	var infos []*mgrpb.Registration
	var total uint32

	rawClient, err := db.Client()
	if err != nil {
		return nil, 0, err
	}

	inviterIDs := strings.Join(conds.GetInviterIDs().GetValue(), ",")
	rows, err := rawClient.QueryContext(ctx, fmt.Sprintf("CALL get_subordinates(\"%v\")", inviterIDs))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	subordinates := ""
	for rows.Next() {
		if err := rows.Scan(&subordinates); err != nil {
			return nil, 0, err
		}
	}

	_inviterIDs := strings.Split(subordinates, ",")
	conds.InviterIDs.Value = _inviterIDs

	logger.Sugar().Infow("XXXXXX", "Inviters", inviterIDs)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
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
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
