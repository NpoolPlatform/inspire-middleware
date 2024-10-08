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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/eventcoin"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// EventCoinQuery is the builder for querying EventCoin entities.
type EventCoinQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EventCoin
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventCoinQuery builder.
func (ecq *EventCoinQuery) Where(ps ...predicate.EventCoin) *EventCoinQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit adds a limit step to the query.
func (ecq *EventCoinQuery) Limit(limit int) *EventCoinQuery {
	ecq.limit = &limit
	return ecq
}

// Offset adds an offset step to the query.
func (ecq *EventCoinQuery) Offset(offset int) *EventCoinQuery {
	ecq.offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EventCoinQuery) Unique(unique bool) *EventCoinQuery {
	ecq.unique = &unique
	return ecq
}

// Order adds an order step to the query.
func (ecq *EventCoinQuery) Order(o ...OrderFunc) *EventCoinQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// First returns the first EventCoin entity from the query.
// Returns a *NotFoundError when no EventCoin was found.
func (ecq *EventCoinQuery) First(ctx context.Context) (*EventCoin, error) {
	nodes, err := ecq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventcoin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EventCoinQuery) FirstX(ctx context.Context) *EventCoin {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventCoin ID from the query.
// Returns a *NotFoundError when no EventCoin ID was found.
func (ecq *EventCoinQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ecq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventcoin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EventCoinQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventCoin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EventCoin entity is found.
// Returns a *NotFoundError when no EventCoin entities are found.
func (ecq *EventCoinQuery) Only(ctx context.Context) (*EventCoin, error) {
	nodes, err := ecq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventcoin.Label}
	default:
		return nil, &NotSingularError{eventcoin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EventCoinQuery) OnlyX(ctx context.Context) *EventCoin {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventCoin ID in the query.
// Returns a *NotSingularError when more than one EventCoin ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EventCoinQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ecq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventcoin.Label}
	default:
		err = &NotSingularError{eventcoin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EventCoinQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventCoins.
func (ecq *EventCoinQuery) All(ctx context.Context) ([]*EventCoin, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ecq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EventCoinQuery) AllX(ctx context.Context) []*EventCoin {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventCoin IDs.
func (ecq *EventCoinQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := ecq.Select(eventcoin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EventCoinQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EventCoinQuery) Count(ctx context.Context) (int, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ecq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EventCoinQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EventCoinQuery) Exist(ctx context.Context) (bool, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ecq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EventCoinQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventCoinQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EventCoinQuery) Clone() *EventCoinQuery {
	if ecq == nil {
		return nil
	}
	return &EventCoinQuery{
		config:     ecq.config,
		limit:      ecq.limit,
		offset:     ecq.offset,
		order:      append([]OrderFunc{}, ecq.order...),
		predicates: append([]predicate.EventCoin{}, ecq.predicates...),
		// clone intermediate query.
		sql:    ecq.sql.Clone(),
		path:   ecq.path,
		unique: ecq.unique,
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
//	client.EventCoin.Query().
//		GroupBy(eventcoin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ecq *EventCoinQuery) GroupBy(field string, fields ...string) *EventCoinGroupBy {
	grbuild := &EventCoinGroupBy{config: ecq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ecq.sqlQuery(ctx), nil
	}
	grbuild.label = eventcoin.Label
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
//	client.EventCoin.Query().
//		Select(eventcoin.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ecq *EventCoinQuery) Select(fields ...string) *EventCoinSelect {
	ecq.fields = append(ecq.fields, fields...)
	selbuild := &EventCoinSelect{EventCoinQuery: ecq}
	selbuild.label = eventcoin.Label
	selbuild.flds, selbuild.scan = &ecq.fields, selbuild.Scan
	return selbuild
}

func (ecq *EventCoinQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ecq.fields {
		if !eventcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	if eventcoin.Policy == nil {
		return errors.New("ent: uninitialized eventcoin.Policy (forgotten import ent/runtime?)")
	}
	if err := eventcoin.Policy.EvalQuery(ctx, ecq); err != nil {
		return err
	}
	return nil
}

func (ecq *EventCoinQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EventCoin, error) {
	var (
		nodes = []*EventCoin{}
		_spec = ecq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*EventCoin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &EventCoin{config: ecq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ecq *EventCoinQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	if len(ecq.modifiers) > 0 {
		_spec.Modifiers = ecq.modifiers
	}
	_spec.Node.Columns = ecq.fields
	if len(ecq.fields) > 0 {
		_spec.Unique = ecq.unique != nil && *ecq.unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EventCoinQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ecq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ecq *EventCoinQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventcoin.Table,
			Columns: eventcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: eventcoin.FieldID,
			},
		},
		From:   ecq.sql,
		Unique: true,
	}
	if unique := ecq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ecq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventcoin.FieldID)
		for i := range fields {
			if fields[i] != eventcoin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EventCoinQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(eventcoin.Table)
	columns := ecq.fields
	if len(columns) == 0 {
		columns = eventcoin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.unique != nil && *ecq.unique {
		selector.Distinct()
	}
	for _, m := range ecq.modifiers {
		m(selector)
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ecq *EventCoinQuery) ForUpdate(opts ...sql.LockOption) *EventCoinQuery {
	if ecq.driver.Dialect() == dialect.Postgres {
		ecq.Unique(false)
	}
	ecq.modifiers = append(ecq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ecq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ecq *EventCoinQuery) ForShare(opts ...sql.LockOption) *EventCoinQuery {
	if ecq.driver.Dialect() == dialect.Postgres {
		ecq.Unique(false)
	}
	ecq.modifiers = append(ecq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ecq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecq *EventCoinQuery) Modify(modifiers ...func(s *sql.Selector)) *EventCoinSelect {
	ecq.modifiers = append(ecq.modifiers, modifiers...)
	return ecq.Select()
}

// EventCoinGroupBy is the group-by builder for EventCoin entities.
type EventCoinGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EventCoinGroupBy) Aggregate(fns ...AggregateFunc) *EventCoinGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ecgb *EventCoinGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ecgb.path(ctx)
	if err != nil {
		return err
	}
	ecgb.sql = query
	return ecgb.sqlScan(ctx, v)
}

func (ecgb *EventCoinGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ecgb.fields {
		if !eventcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ecgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ecgb *EventCoinGroupBy) sqlQuery() *sql.Selector {
	selector := ecgb.sql.Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ecgb.fields)+len(ecgb.fns))
		for _, f := range ecgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ecgb.fields...)...)
}

// EventCoinSelect is the builder for selecting fields of EventCoin entities.
type EventCoinSelect struct {
	*EventCoinQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EventCoinSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	ecs.sql = ecs.EventCoinQuery.sqlQuery(ctx)
	return ecs.sqlScan(ctx, v)
}

func (ecs *EventCoinSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ecs.sql.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ecs *EventCoinSelect) Modify(modifiers ...func(s *sql.Selector)) *EventCoinSelect {
	ecs.modifiers = append(ecs.modifiers, modifiers...)
	return ecs
}
