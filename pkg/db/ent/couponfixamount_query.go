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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/couponfixamount"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CouponFixAmountQuery is the builder for querying CouponFixAmount entities.
type CouponFixAmountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CouponFixAmount
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CouponFixAmountQuery builder.
func (cfaq *CouponFixAmountQuery) Where(ps ...predicate.CouponFixAmount) *CouponFixAmountQuery {
	cfaq.predicates = append(cfaq.predicates, ps...)
	return cfaq
}

// Limit adds a limit step to the query.
func (cfaq *CouponFixAmountQuery) Limit(limit int) *CouponFixAmountQuery {
	cfaq.limit = &limit
	return cfaq
}

// Offset adds an offset step to the query.
func (cfaq *CouponFixAmountQuery) Offset(offset int) *CouponFixAmountQuery {
	cfaq.offset = &offset
	return cfaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cfaq *CouponFixAmountQuery) Unique(unique bool) *CouponFixAmountQuery {
	cfaq.unique = &unique
	return cfaq
}

// Order adds an order step to the query.
func (cfaq *CouponFixAmountQuery) Order(o ...OrderFunc) *CouponFixAmountQuery {
	cfaq.order = append(cfaq.order, o...)
	return cfaq
}

// First returns the first CouponFixAmount entity from the query.
// Returns a *NotFoundError when no CouponFixAmount was found.
func (cfaq *CouponFixAmountQuery) First(ctx context.Context) (*CouponFixAmount, error) {
	nodes, err := cfaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{couponfixamount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) FirstX(ctx context.Context) *CouponFixAmount {
	node, err := cfaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CouponFixAmount ID from the query.
// Returns a *NotFoundError when no CouponFixAmount ID was found.
func (cfaq *CouponFixAmountQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cfaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{couponfixamount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cfaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CouponFixAmount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CouponFixAmount entity is found.
// Returns a *NotFoundError when no CouponFixAmount entities are found.
func (cfaq *CouponFixAmountQuery) Only(ctx context.Context) (*CouponFixAmount, error) {
	nodes, err := cfaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{couponfixamount.Label}
	default:
		return nil, &NotSingularError{couponfixamount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) OnlyX(ctx context.Context) *CouponFixAmount {
	node, err := cfaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CouponFixAmount ID in the query.
// Returns a *NotSingularError when more than one CouponFixAmount ID is found.
// Returns a *NotFoundError when no entities are found.
func (cfaq *CouponFixAmountQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cfaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{couponfixamount.Label}
	default:
		err = &NotSingularError{couponfixamount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cfaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CouponFixAmounts.
func (cfaq *CouponFixAmountQuery) All(ctx context.Context) ([]*CouponFixAmount, error) {
	if err := cfaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cfaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) AllX(ctx context.Context) []*CouponFixAmount {
	nodes, err := cfaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CouponFixAmount IDs.
func (cfaq *CouponFixAmountQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cfaq.Select(couponfixamount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cfaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cfaq *CouponFixAmountQuery) Count(ctx context.Context) (int, error) {
	if err := cfaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cfaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) CountX(ctx context.Context) int {
	count, err := cfaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cfaq *CouponFixAmountQuery) Exist(ctx context.Context) (bool, error) {
	if err := cfaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cfaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cfaq *CouponFixAmountQuery) ExistX(ctx context.Context) bool {
	exist, err := cfaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CouponFixAmountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cfaq *CouponFixAmountQuery) Clone() *CouponFixAmountQuery {
	if cfaq == nil {
		return nil
	}
	return &CouponFixAmountQuery{
		config:     cfaq.config,
		limit:      cfaq.limit,
		offset:     cfaq.offset,
		order:      append([]OrderFunc{}, cfaq.order...),
		predicates: append([]predicate.CouponFixAmount{}, cfaq.predicates...),
		// clone intermediate query.
		sql:    cfaq.sql.Clone(),
		path:   cfaq.path,
		unique: cfaq.unique,
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
//	client.CouponFixAmount.Query().
//		GroupBy(couponfixamount.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cfaq *CouponFixAmountQuery) GroupBy(field string, fields ...string) *CouponFixAmountGroupBy {
	grbuild := &CouponFixAmountGroupBy{config: cfaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cfaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cfaq.sqlQuery(ctx), nil
	}
	grbuild.label = couponfixamount.Label
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
//	client.CouponFixAmount.Query().
//		Select(couponfixamount.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (cfaq *CouponFixAmountQuery) Select(fields ...string) *CouponFixAmountSelect {
	cfaq.fields = append(cfaq.fields, fields...)
	selbuild := &CouponFixAmountSelect{CouponFixAmountQuery: cfaq}
	selbuild.label = couponfixamount.Label
	selbuild.flds, selbuild.scan = &cfaq.fields, selbuild.Scan
	return selbuild
}

func (cfaq *CouponFixAmountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cfaq.fields {
		if !couponfixamount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cfaq.path != nil {
		prev, err := cfaq.path(ctx)
		if err != nil {
			return err
		}
		cfaq.sql = prev
	}
	if couponfixamount.Policy == nil {
		return errors.New("ent: uninitialized couponfixamount.Policy (forgotten import ent/runtime?)")
	}
	if err := couponfixamount.Policy.EvalQuery(ctx, cfaq); err != nil {
		return err
	}
	return nil
}

func (cfaq *CouponFixAmountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CouponFixAmount, error) {
	var (
		nodes = []*CouponFixAmount{}
		_spec = cfaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CouponFixAmount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CouponFixAmount{config: cfaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cfaq.modifiers) > 0 {
		_spec.Modifiers = cfaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cfaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cfaq *CouponFixAmountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cfaq.querySpec()
	if len(cfaq.modifiers) > 0 {
		_spec.Modifiers = cfaq.modifiers
	}
	_spec.Node.Columns = cfaq.fields
	if len(cfaq.fields) > 0 {
		_spec.Unique = cfaq.unique != nil && *cfaq.unique
	}
	return sqlgraph.CountNodes(ctx, cfaq.driver, _spec)
}

func (cfaq *CouponFixAmountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cfaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cfaq *CouponFixAmountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponfixamount.Table,
			Columns: couponfixamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponfixamount.FieldID,
			},
		},
		From:   cfaq.sql,
		Unique: true,
	}
	if unique := cfaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cfaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponfixamount.FieldID)
		for i := range fields {
			if fields[i] != couponfixamount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cfaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cfaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cfaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cfaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cfaq *CouponFixAmountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cfaq.driver.Dialect())
	t1 := builder.Table(couponfixamount.Table)
	columns := cfaq.fields
	if len(columns) == 0 {
		columns = couponfixamount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cfaq.sql != nil {
		selector = cfaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cfaq.unique != nil && *cfaq.unique {
		selector.Distinct()
	}
	for _, m := range cfaq.modifiers {
		m(selector)
	}
	for _, p := range cfaq.predicates {
		p(selector)
	}
	for _, p := range cfaq.order {
		p(selector)
	}
	if offset := cfaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cfaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cfaq *CouponFixAmountQuery) ForUpdate(opts ...sql.LockOption) *CouponFixAmountQuery {
	if cfaq.driver.Dialect() == dialect.Postgres {
		cfaq.Unique(false)
	}
	cfaq.modifiers = append(cfaq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cfaq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cfaq *CouponFixAmountQuery) ForShare(opts ...sql.LockOption) *CouponFixAmountQuery {
	if cfaq.driver.Dialect() == dialect.Postgres {
		cfaq.Unique(false)
	}
	cfaq.modifiers = append(cfaq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cfaq
}

// CouponFixAmountGroupBy is the group-by builder for CouponFixAmount entities.
type CouponFixAmountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cfagb *CouponFixAmountGroupBy) Aggregate(fns ...AggregateFunc) *CouponFixAmountGroupBy {
	cfagb.fns = append(cfagb.fns, fns...)
	return cfagb
}

// Scan applies the group-by query and scans the result into the given value.
func (cfagb *CouponFixAmountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cfagb.path(ctx)
	if err != nil {
		return err
	}
	cfagb.sql = query
	return cfagb.sqlScan(ctx, v)
}

func (cfagb *CouponFixAmountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cfagb.fields {
		if !couponfixamount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cfagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cfagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cfagb *CouponFixAmountGroupBy) sqlQuery() *sql.Selector {
	selector := cfagb.sql.Select()
	aggregation := make([]string, 0, len(cfagb.fns))
	for _, fn := range cfagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cfagb.fields)+len(cfagb.fns))
		for _, f := range cfagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cfagb.fields...)...)
}

// CouponFixAmountSelect is the builder for selecting fields of CouponFixAmount entities.
type CouponFixAmountSelect struct {
	*CouponFixAmountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cfas *CouponFixAmountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cfas.prepareQuery(ctx); err != nil {
		return err
	}
	cfas.sql = cfas.CouponFixAmountQuery.sqlQuery(ctx)
	return cfas.sqlScan(ctx, v)
}

func (cfas *CouponFixAmountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cfas.sql.Query()
	if err := cfas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
