package goodachievement

import (
	"fmt"
	"time"
)

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
	_sql += "where not exist ("
	_sql += "select 1 from order_statements "
	_sql += fmt.Sprintf(
		"where user_id = '%v' and app_good_id = '%v' ",
		*h.UserID,
		*h.AppGoodID,
	)
	_sql += "limit 1)"
	return _sql
}
