package goodcoinachievement

import (
	"fmt"
	entgoodcoinachievement "github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodcoinachievement"
	"time"
)

//nolint:goconst
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "insert into good_coin_achievements "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "good_coin_type_id"
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
	_sql += fmt.Sprintf("%v'%v' as good_coin_type_id", comma, *h.GoodCoinTypeID)
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
	_sql += "select 1 from good_coin_achievements "
	_sql += fmt.Sprintf(
		"where user_id = '%v' and good_coin_type_id = '%v' ",
		*h.UserID,
		*h.GoodCoinTypeID,
	)
	_sql += "limit 1)"
	return _sql
}

func (h *Handler) ConstructUpdateSQL() string {
	sql := fmt.Sprintf(
		`update %v set total_units = total_units + %v`,
		entgoodcoinachievement.Table,
		*h.TotalUnits,
	)
	sql += fmt.Sprintf(
		`, self_units = self_units + %v`,
		*h.SelfUnits,
	)

	sql += fmt.Sprintf(
		`, total_amount_usd = total_amount_usd + %v`,
		*h.TotalAmountUSD,
	)
	sql += fmt.Sprintf(
		`, self_amount_usd = self_amount_usd + %v`,
		*h.SelfAmountUSD,
	)
	sql += fmt.Sprintf(
		`, total_commission_usd = total_commission_usd + %v`,
		*h.TotalCommissionUSD,
	)
	sql += fmt.Sprintf(
		`, self_commission_usd = self_commission_usd + %v`,
		h.SelfCommissionUSD,
	)

	sql += fmt.Sprintf(
		" where user_id = '%v' and good_coin_type_id = '%v' and deleted_at = 0 ",
		h.UserID.String(), h.GoodCoinTypeID.String(),
	)
	return sql
}
