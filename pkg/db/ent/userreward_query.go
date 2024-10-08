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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/userreward"
)

// UserRewardQuery is the builder for querying UserReward entities.
type UserRewardQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserReward
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserRewardQuery builder.
func (urq *UserRewardQuery) Where(ps ...predicate.UserReward) *UserRewardQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit adds a limit step to the query.
func (urq *UserRewardQuery) Limit(limit int) *UserRewardQuery {
	urq.limit = &limit
	return urq
}

// Offset adds an offset step to the query.
func (urq *UserRewardQuery) Offset(offset int) *UserRewardQuery {
	urq.offset = &offset
	return urq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urq *UserRewardQuery) Unique(unique bool) *UserRewardQuery {
	urq.unique = &unique
	return urq
}

// Order adds an order step to the query.
func (urq *UserRewardQuery) Order(o ...OrderFunc) *UserRewardQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// First returns the first UserReward entity from the query.
// Returns a *NotFoundError when no UserReward was found.
func (urq *UserRewardQuery) First(ctx context.Context) (*UserReward, error) {
	nodes, err := urq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userreward.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UserRewardQuery) FirstX(ctx context.Context) *UserReward {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserReward ID from the query.
// Returns a *NotFoundError when no UserReward ID was found.
func (urq *UserRewardQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = urq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userreward.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UserRewardQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserReward entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserReward entity is found.
// Returns a *NotFoundError when no UserReward entities are found.
func (urq *UserRewardQuery) Only(ctx context.Context) (*UserReward, error) {
	nodes, err := urq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userreward.Label}
	default:
		return nil, &NotSingularError{userreward.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UserRewardQuery) OnlyX(ctx context.Context) *UserReward {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserReward ID in the query.
// Returns a *NotSingularError when more than one UserReward ID is found.
// Returns a *NotFoundError when no entities are found.
func (urq *UserRewardQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = urq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userreward.Label}
	default:
		err = &NotSingularError{userreward.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UserRewardQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserRewards.
func (urq *UserRewardQuery) All(ctx context.Context) ([]*UserReward, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return urq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (urq *UserRewardQuery) AllX(ctx context.Context) []*UserReward {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserReward IDs.
func (urq *UserRewardQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := urq.Select(userreward.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UserRewardQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UserRewardQuery) Count(ctx context.Context) (int, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return urq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UserRewardQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UserRewardQuery) Exist(ctx context.Context) (bool, error) {
	if err := urq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return urq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UserRewardQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserRewardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UserRewardQuery) Clone() *UserRewardQuery {
	if urq == nil {
		return nil
	}
	return &UserRewardQuery{
		config:     urq.config,
		limit:      urq.limit,
		offset:     urq.offset,
		order:      append([]OrderFunc{}, urq.order...),
		predicates: append([]predicate.UserReward{}, urq.predicates...),
		// clone intermediate query.
		sql:    urq.sql.Clone(),
		path:   urq.path,
		unique: urq.unique,
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
//	client.UserReward.Query().
//		GroupBy(userreward.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (urq *UserRewardQuery) GroupBy(field string, fields ...string) *UserRewardGroupBy {
	grbuild := &UserRewardGroupBy{config: urq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return urq.sqlQuery(ctx), nil
	}
	grbuild.label = userreward.Label
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
//	client.UserReward.Query().
//		Select(userreward.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (urq *UserRewardQuery) Select(fields ...string) *UserRewardSelect {
	urq.fields = append(urq.fields, fields...)
	selbuild := &UserRewardSelect{UserRewardQuery: urq}
	selbuild.label = userreward.Label
	selbuild.flds, selbuild.scan = &urq.fields, selbuild.Scan
	return selbuild
}

func (urq *UserRewardQuery) prepareQuery(ctx context.Context) error {
	for _, f := range urq.fields {
		if !userreward.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	if userreward.Policy == nil {
		return errors.New("ent: uninitialized userreward.Policy (forgotten import ent/runtime?)")
	}
	if err := userreward.Policy.EvalQuery(ctx, urq); err != nil {
		return err
	}
	return nil
}

func (urq *UserRewardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserReward, error) {
	var (
		nodes = []*UserReward{}
		_spec = urq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserReward).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserReward{config: urq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (urq *UserRewardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	_spec.Node.Columns = urq.fields
	if len(urq.fields) > 0 {
		_spec.Unique = urq.unique != nil && *urq.unique
	}
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UserRewardQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := urq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (urq *UserRewardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userreward.Table,
			Columns: userreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: userreward.FieldID,
			},
		},
		From:   urq.sql,
		Unique: true,
	}
	if unique := urq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := urq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userreward.FieldID)
		for i := range fields {
			if fields[i] != userreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urq *UserRewardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(userreward.Table)
	columns := urq.fields
	if len(columns) == 0 {
		columns = userreward.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urq.unique != nil && *urq.unique {
		selector.Distinct()
	}
	for _, m := range urq.modifiers {
		m(selector)
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector)
	}
	if offset := urq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (urq *UserRewardQuery) ForUpdate(opts ...sql.LockOption) *UserRewardQuery {
	if urq.driver.Dialect() == dialect.Postgres {
		urq.Unique(false)
	}
	urq.modifiers = append(urq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return urq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (urq *UserRewardQuery) ForShare(opts ...sql.LockOption) *UserRewardQuery {
	if urq.driver.Dialect() == dialect.Postgres {
		urq.Unique(false)
	}
	urq.modifiers = append(urq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return urq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (urq *UserRewardQuery) Modify(modifiers ...func(s *sql.Selector)) *UserRewardSelect {
	urq.modifiers = append(urq.modifiers, modifiers...)
	return urq.Select()
}

// UserRewardGroupBy is the group-by builder for UserReward entities.
type UserRewardGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UserRewardGroupBy) Aggregate(fns ...AggregateFunc) *UserRewardGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the group-by query and scans the result into the given value.
func (urgb *UserRewardGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := urgb.path(ctx)
	if err != nil {
		return err
	}
	urgb.sql = query
	return urgb.sqlScan(ctx, v)
}

func (urgb *UserRewardGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range urgb.fields {
		if !userreward.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := urgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (urgb *UserRewardGroupBy) sqlQuery() *sql.Selector {
	selector := urgb.sql.Select()
	aggregation := make([]string, 0, len(urgb.fns))
	for _, fn := range urgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(urgb.fields)+len(urgb.fns))
		for _, f := range urgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(urgb.fields...)...)
}

// UserRewardSelect is the builder for selecting fields of UserReward entities.
type UserRewardSelect struct {
	*UserRewardQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (urs *UserRewardSelect) Scan(ctx context.Context, v interface{}) error {
	if err := urs.prepareQuery(ctx); err != nil {
		return err
	}
	urs.sql = urs.UserRewardQuery.sqlQuery(ctx)
	return urs.sqlScan(ctx, v)
}

func (urs *UserRewardSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := urs.sql.Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (urs *UserRewardSelect) Modify(modifiers ...func(s *sql.Selector)) *UserRewardSelect {
	urs.modifiers = append(urs.modifiers, modifiers...)
	return urs
}
