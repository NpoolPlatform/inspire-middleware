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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodordervaluepercent"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodOrderValuePercentQuery is the builder for querying GoodOrderValuePercent entities.
type GoodOrderValuePercentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodOrderValuePercent
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodOrderValuePercentQuery builder.
func (govpq *GoodOrderValuePercentQuery) Where(ps ...predicate.GoodOrderValuePercent) *GoodOrderValuePercentQuery {
	govpq.predicates = append(govpq.predicates, ps...)
	return govpq
}

// Limit adds a limit step to the query.
func (govpq *GoodOrderValuePercentQuery) Limit(limit int) *GoodOrderValuePercentQuery {
	govpq.limit = &limit
	return govpq
}

// Offset adds an offset step to the query.
func (govpq *GoodOrderValuePercentQuery) Offset(offset int) *GoodOrderValuePercentQuery {
	govpq.offset = &offset
	return govpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (govpq *GoodOrderValuePercentQuery) Unique(unique bool) *GoodOrderValuePercentQuery {
	govpq.unique = &unique
	return govpq
}

// Order adds an order step to the query.
func (govpq *GoodOrderValuePercentQuery) Order(o ...OrderFunc) *GoodOrderValuePercentQuery {
	govpq.order = append(govpq.order, o...)
	return govpq
}

// First returns the first GoodOrderValuePercent entity from the query.
// Returns a *NotFoundError when no GoodOrderValuePercent was found.
func (govpq *GoodOrderValuePercentQuery) First(ctx context.Context) (*GoodOrderValuePercent, error) {
	nodes, err := govpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodordervaluepercent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) FirstX(ctx context.Context) *GoodOrderValuePercent {
	node, err := govpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodOrderValuePercent ID from the query.
// Returns a *NotFoundError when no GoodOrderValuePercent ID was found.
func (govpq *GoodOrderValuePercentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = govpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodordervaluepercent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := govpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodOrderValuePercent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodOrderValuePercent entity is found.
// Returns a *NotFoundError when no GoodOrderValuePercent entities are found.
func (govpq *GoodOrderValuePercentQuery) Only(ctx context.Context) (*GoodOrderValuePercent, error) {
	nodes, err := govpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodordervaluepercent.Label}
	default:
		return nil, &NotSingularError{goodordervaluepercent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) OnlyX(ctx context.Context) *GoodOrderValuePercent {
	node, err := govpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodOrderValuePercent ID in the query.
// Returns a *NotSingularError when more than one GoodOrderValuePercent ID is found.
// Returns a *NotFoundError when no entities are found.
func (govpq *GoodOrderValuePercentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = govpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodordervaluepercent.Label}
	default:
		err = &NotSingularError{goodordervaluepercent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := govpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodOrderValuePercents.
func (govpq *GoodOrderValuePercentQuery) All(ctx context.Context) ([]*GoodOrderValuePercent, error) {
	if err := govpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return govpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) AllX(ctx context.Context) []*GoodOrderValuePercent {
	nodes, err := govpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodOrderValuePercent IDs.
func (govpq *GoodOrderValuePercentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := govpq.Select(goodordervaluepercent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := govpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (govpq *GoodOrderValuePercentQuery) Count(ctx context.Context) (int, error) {
	if err := govpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return govpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) CountX(ctx context.Context) int {
	count, err := govpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (govpq *GoodOrderValuePercentQuery) Exist(ctx context.Context) (bool, error) {
	if err := govpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return govpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (govpq *GoodOrderValuePercentQuery) ExistX(ctx context.Context) bool {
	exist, err := govpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodOrderValuePercentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (govpq *GoodOrderValuePercentQuery) Clone() *GoodOrderValuePercentQuery {
	if govpq == nil {
		return nil
	}
	return &GoodOrderValuePercentQuery{
		config:     govpq.config,
		limit:      govpq.limit,
		offset:     govpq.offset,
		order:      append([]OrderFunc{}, govpq.order...),
		predicates: append([]predicate.GoodOrderValuePercent{}, govpq.predicates...),
		// clone intermediate query.
		sql:    govpq.sql.Clone(),
		path:   govpq.path,
		unique: govpq.unique,
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
//	client.GoodOrderValuePercent.Query().
//		GroupBy(goodordervaluepercent.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (govpq *GoodOrderValuePercentQuery) GroupBy(field string, fields ...string) *GoodOrderValuePercentGroupBy {
	grbuild := &GoodOrderValuePercentGroupBy{config: govpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := govpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return govpq.sqlQuery(ctx), nil
	}
	grbuild.label = goodordervaluepercent.Label
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
//	client.GoodOrderValuePercent.Query().
//		Select(goodordervaluepercent.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (govpq *GoodOrderValuePercentQuery) Select(fields ...string) *GoodOrderValuePercentSelect {
	govpq.fields = append(govpq.fields, fields...)
	selbuild := &GoodOrderValuePercentSelect{GoodOrderValuePercentQuery: govpq}
	selbuild.label = goodordervaluepercent.Label
	selbuild.flds, selbuild.scan = &govpq.fields, selbuild.Scan
	return selbuild
}

func (govpq *GoodOrderValuePercentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range govpq.fields {
		if !goodordervaluepercent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if govpq.path != nil {
		prev, err := govpq.path(ctx)
		if err != nil {
			return err
		}
		govpq.sql = prev
	}
	if goodordervaluepercent.Policy == nil {
		return errors.New("ent: uninitialized goodordervaluepercent.Policy (forgotten import ent/runtime?)")
	}
	if err := goodordervaluepercent.Policy.EvalQuery(ctx, govpq); err != nil {
		return err
	}
	return nil
}

func (govpq *GoodOrderValuePercentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodOrderValuePercent, error) {
	var (
		nodes = []*GoodOrderValuePercent{}
		_spec = govpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodOrderValuePercent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodOrderValuePercent{config: govpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(govpq.modifiers) > 0 {
		_spec.Modifiers = govpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, govpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (govpq *GoodOrderValuePercentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := govpq.querySpec()
	if len(govpq.modifiers) > 0 {
		_spec.Modifiers = govpq.modifiers
	}
	_spec.Node.Columns = govpq.fields
	if len(govpq.fields) > 0 {
		_spec.Unique = govpq.unique != nil && *govpq.unique
	}
	return sqlgraph.CountNodes(ctx, govpq.driver, _spec)
}

func (govpq *GoodOrderValuePercentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := govpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (govpq *GoodOrderValuePercentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodordervaluepercent.Table,
			Columns: goodordervaluepercent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodordervaluepercent.FieldID,
			},
		},
		From:   govpq.sql,
		Unique: true,
	}
	if unique := govpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := govpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodordervaluepercent.FieldID)
		for i := range fields {
			if fields[i] != goodordervaluepercent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := govpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := govpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := govpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := govpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (govpq *GoodOrderValuePercentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(govpq.driver.Dialect())
	t1 := builder.Table(goodordervaluepercent.Table)
	columns := govpq.fields
	if len(columns) == 0 {
		columns = goodordervaluepercent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if govpq.sql != nil {
		selector = govpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if govpq.unique != nil && *govpq.unique {
		selector.Distinct()
	}
	for _, m := range govpq.modifiers {
		m(selector)
	}
	for _, p := range govpq.predicates {
		p(selector)
	}
	for _, p := range govpq.order {
		p(selector)
	}
	if offset := govpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := govpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (govpq *GoodOrderValuePercentQuery) ForUpdate(opts ...sql.LockOption) *GoodOrderValuePercentQuery {
	if govpq.driver.Dialect() == dialect.Postgres {
		govpq.Unique(false)
	}
	govpq.modifiers = append(govpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return govpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (govpq *GoodOrderValuePercentQuery) ForShare(opts ...sql.LockOption) *GoodOrderValuePercentQuery {
	if govpq.driver.Dialect() == dialect.Postgres {
		govpq.Unique(false)
	}
	govpq.modifiers = append(govpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return govpq
}

// GoodOrderValuePercentGroupBy is the group-by builder for GoodOrderValuePercent entities.
type GoodOrderValuePercentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (govpgb *GoodOrderValuePercentGroupBy) Aggregate(fns ...AggregateFunc) *GoodOrderValuePercentGroupBy {
	govpgb.fns = append(govpgb.fns, fns...)
	return govpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (govpgb *GoodOrderValuePercentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := govpgb.path(ctx)
	if err != nil {
		return err
	}
	govpgb.sql = query
	return govpgb.sqlScan(ctx, v)
}

func (govpgb *GoodOrderValuePercentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range govpgb.fields {
		if !goodordervaluepercent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := govpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := govpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (govpgb *GoodOrderValuePercentGroupBy) sqlQuery() *sql.Selector {
	selector := govpgb.sql.Select()
	aggregation := make([]string, 0, len(govpgb.fns))
	for _, fn := range govpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(govpgb.fields)+len(govpgb.fns))
		for _, f := range govpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(govpgb.fields...)...)
}

// GoodOrderValuePercentSelect is the builder for selecting fields of GoodOrderValuePercent entities.
type GoodOrderValuePercentSelect struct {
	*GoodOrderValuePercentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (govps *GoodOrderValuePercentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := govps.prepareQuery(ctx); err != nil {
		return err
	}
	govps.sql = govps.GoodOrderValuePercentQuery.sqlQuery(ctx)
	return govps.sqlScan(ctx, v)
}

func (govps *GoodOrderValuePercentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := govps.sql.Query()
	if err := govps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
