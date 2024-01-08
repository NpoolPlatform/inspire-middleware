// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponcoin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// CouponCoinQuery is the builder for querying CouponCoin entities.
type CouponCoinQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CouponCoin
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CouponCoinQuery builder.
func (ccq *CouponCoinQuery) Where(ps ...predicate.CouponCoin) *CouponCoinQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit adds a limit step to the query.
func (ccq *CouponCoinQuery) Limit(limit int) *CouponCoinQuery {
	ccq.limit = &limit
	return ccq
}

// Offset adds an offset step to the query.
func (ccq *CouponCoinQuery) Offset(offset int) *CouponCoinQuery {
	ccq.offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *CouponCoinQuery) Unique(unique bool) *CouponCoinQuery {
	ccq.unique = &unique
	return ccq
}

// Order adds an order step to the query.
func (ccq *CouponCoinQuery) Order(o ...OrderFunc) *CouponCoinQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// First returns the first CouponCoin entity from the query.
// Returns a *NotFoundError when no CouponCoin was found.
func (ccq *CouponCoinQuery) First(ctx context.Context) (*CouponCoin, error) {
	nodes, err := ccq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{couponcoin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *CouponCoinQuery) FirstX(ctx context.Context) *CouponCoin {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CouponCoin ID from the query.
// Returns a *NotFoundError when no CouponCoin ID was found.
func (ccq *CouponCoinQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ccq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{couponcoin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *CouponCoinQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CouponCoin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CouponCoin entity is found.
// Returns a *NotFoundError when no CouponCoin entities are found.
func (ccq *CouponCoinQuery) Only(ctx context.Context) (*CouponCoin, error) {
	nodes, err := ccq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{couponcoin.Label}
	default:
		return nil, &NotSingularError{couponcoin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *CouponCoinQuery) OnlyX(ctx context.Context) *CouponCoin {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CouponCoin ID in the query.
// Returns a *NotSingularError when more than one CouponCoin ID is found.
// Returns a *NotFoundError when no entities are found.
func (ccq *CouponCoinQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ccq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{couponcoin.Label}
	default:
		err = &NotSingularError{couponcoin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *CouponCoinQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CouponCoins.
func (ccq *CouponCoinQuery) All(ctx context.Context) ([]*CouponCoin, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ccq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ccq *CouponCoinQuery) AllX(ctx context.Context) []*CouponCoin {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CouponCoin IDs.
func (ccq *CouponCoinQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := ccq.Select(couponcoin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *CouponCoinQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *CouponCoinQuery) Count(ctx context.Context) (int, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ccq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *CouponCoinQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *CouponCoinQuery) Exist(ctx context.Context) (bool, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ccq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *CouponCoinQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CouponCoinQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *CouponCoinQuery) Clone() *CouponCoinQuery {
	if ccq == nil {
		return nil
	}
	return &CouponCoinQuery{
		config:     ccq.config,
		limit:      ccq.limit,
		offset:     ccq.offset,
		order:      append([]OrderFunc{}, ccq.order...),
		predicates: append([]predicate.CouponCoin{}, ccq.predicates...),
		// clone intermediate query.
		sql:    ccq.sql.Clone(),
		path:   ccq.path,
		unique: ccq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CouponCoin.Query().
//		GroupBy(couponcoin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ccq *CouponCoinQuery) GroupBy(field string, fields ...string) *CouponCoinGroupBy {
	grbuild := &CouponCoinGroupBy{config: ccq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ccq.sqlQuery(ctx), nil
	}
	grbuild.label = couponcoin.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.CouponCoin.Query().
//		Select(couponcoin.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ccq *CouponCoinQuery) Select(fields ...string) *CouponCoinSelect {
	ccq.fields = append(ccq.fields, fields...)
	selbuild := &CouponCoinSelect{CouponCoinQuery: ccq}
	selbuild.label = couponcoin.Label
	selbuild.flds, selbuild.scan = &ccq.fields, selbuild.Scan
	return selbuild
}

func (ccq *CouponCoinQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ccq.fields {
		if !couponcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	if couponcoin.Policy == nil {
		return errors.New("ent: uninitialized couponcoin.Policy (forgotten import ent/runtime?)")
	}
	if err := couponcoin.Policy.EvalQuery(ctx, ccq); err != nil {
		return err
	}
	return nil
}

func (ccq *CouponCoinQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CouponCoin, error) {
	var (
		nodes = []*CouponCoin{}
		_spec = ccq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CouponCoin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CouponCoin{config: ccq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ccq.modifiers) > 0 {
		_spec.Modifiers = ccq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccq *CouponCoinQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	if len(ccq.modifiers) > 0 {
		_spec.Modifiers = ccq.modifiers
	}
	_spec.Node.Columns = ccq.fields
	if len(ccq.fields) > 0 {
		_spec.Unique = ccq.unique != nil && *ccq.unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *CouponCoinQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ccq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ccq *CouponCoinQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponcoin.Table,
			Columns: couponcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: couponcoin.FieldID,
			},
		},
		From:   ccq.sql,
		Unique: true,
	}
	if unique := ccq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ccq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponcoin.FieldID)
		for i := range fields {
			if fields[i] != couponcoin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *CouponCoinQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(couponcoin.Table)
	columns := ccq.fields
	if len(columns) == 0 {
		columns = couponcoin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.unique != nil && *ccq.unique {
		selector.Distinct()
	}
	for _, m := range ccq.modifiers {
		m(selector)
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ccq *CouponCoinQuery) ForUpdate(opts ...sql.LockOption) *CouponCoinQuery {
	if ccq.driver.Dialect() == dialect.Postgres {
		ccq.Unique(false)
	}
	ccq.modifiers = append(ccq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ccq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ccq *CouponCoinQuery) ForShare(opts ...sql.LockOption) *CouponCoinQuery {
	if ccq.driver.Dialect() == dialect.Postgres {
		ccq.Unique(false)
	}
	ccq.modifiers = append(ccq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ccq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ccq *CouponCoinQuery) Modify(modifiers ...func(s *sql.Selector)) *CouponCoinSelect {
	ccq.modifiers = append(ccq.modifiers, modifiers...)
	return ccq.Select()
}

// CouponCoinGroupBy is the group-by builder for CouponCoin entities.
type CouponCoinGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *CouponCoinGroupBy) Aggregate(fns ...AggregateFunc) *CouponCoinGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ccgb *CouponCoinGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ccgb.path(ctx)
	if err != nil {
		return err
	}
	ccgb.sql = query
	return ccgb.sqlScan(ctx, v)
}

func (ccgb *CouponCoinGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ccgb.fields {
		if !couponcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ccgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ccgb *CouponCoinGroupBy) sqlQuery() *sql.Selector {
	selector := ccgb.sql.Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ccgb.fields)+len(ccgb.fns))
		for _, f := range ccgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ccgb.fields...)...)
}

// CouponCoinSelect is the builder for selecting fields of CouponCoin entities.
type CouponCoinSelect struct {
	*CouponCoinQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *CouponCoinSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	ccs.sql = ccs.CouponCoinQuery.sqlQuery(ctx)
	return ccs.sqlScan(ctx, v)
}

func (ccs *CouponCoinSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ccs.sql.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ccs *CouponCoinSelect) Modify(modifiers ...func(s *sql.Selector)) *CouponCoinSelect {
	ccs.modifiers = append(ccs.modifiers, modifiers...)
	return ccs
}
