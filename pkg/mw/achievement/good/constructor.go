package goodachievement

import (
	"fmt"
	"time"

	entgoodachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodachievement"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into good_achievements "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "good_id"
	_sql += comma + "app_good_id"
	_sql += comma + "total_units"
	_sql += comma + "self_units"
	_sql += comma + "total_amount_usd"
	_sql += comma + "self_amount_usd"
	_sql += comma + "total_commission_usd"
	_sql += comma + "self_commission_usd"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	_sql += fmt.Sprintf("%v'%v' as total_units", comma, *h.TotalUnits)
	_sql += fmt.Sprintf("%v'%v' as self_units", comma, *h.SelfUnits)
	_sql += fmt.Sprintf("%v'%v' as total_amount_usd", comma, *h.TotalAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as self_amount_usd", comma, *h.SelfAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as total_commission_usd", comma, *h.TotalCommissionUSD)
	_sql += fmt.Sprintf("%v'%v' as self_commission_usd", comma, *h.SelfCommissionUSD)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_achievements "
	_sql += fmt.Sprintf(
		"where app_id = '%v' and user_id = '%v' and good_id = '%v' and app_good_id = '%v' ",
		h.AppID.String(),
		h.UserID.String(),
		h.GoodID.String(),
		h.AppGoodID.String(),
	)
	_sql += "limit 1)"
	return _sql
}

func (h *Handler) ConstructUpdateSQL() string {
	now := time.Now().Unix()
	sql := fmt.Sprintf(
		`update %v set updated_at = %v`,
		entgoodachievement.Table,
		now,
	)
	if h.TotalAmountUSD != nil {
		sql += fmt.Sprintf(
			`, total_amount_usd = total_amount_usd + %v`,
			*h.TotalAmountUSD,
		)
	}
	if h.SelfAmountUSD != nil {
		sql += fmt.Sprintf(
			`, self_amount_usd = self_amount_usd + %v`,
			*h.SelfAmountUSD,
		)
	}
	if h.TotalUnits != nil {
		sql += fmt.Sprintf(
			`, total_units = total_units + %v`,
			*h.TotalUnits,
		)
	}
	if h.SelfUnits != nil {
		sql += fmt.Sprintf(
			`, self_units = self_units + %v`,
			*h.SelfUnits,
		)
	}
	if h.TotalCommissionUSD != nil {
		sql += fmt.Sprintf(
			`, total_commission_usd = total_commission_usd + %v`,
			*h.TotalCommissionUSD,
		)
	}
	if h.SelfCommissionUSD != nil {
		sql += fmt.Sprintf(
			`, self_commission_usd = self_commission_usd + %v`,
			h.SelfCommissionUSD,
		)
	}

	sql += fmt.Sprintf(
		" where app_id = '%v' and user_id = '%v' and good_id = '%v' and app_good_id = '%v' and deleted_at = 0 ",
		h.AppID.String(),
		h.UserID.String(),
		h.GoodID.String(),
		h.AppGoodID.String(),
	)
	return sql
}
