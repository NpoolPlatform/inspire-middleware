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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/goodcoinachievement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// GoodCoinAchievementQuery is the builder for querying GoodCoinAchievement entities.
type GoodCoinAchievementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodCoinAchievement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodCoinAchievementQuery builder.
func (gcaq *GoodCoinAchievementQuery) Where(ps ...predicate.GoodCoinAchievement) *GoodCoinAchievementQuery {
	gcaq.predicates = append(gcaq.predicates, ps...)
	return gcaq
}

// Limit adds a limit step to the query.
func (gcaq *GoodCoinAchievementQuery) Limit(limit int) *GoodCoinAchievementQuery {
	gcaq.limit = &limit
	return gcaq
}

// Offset adds an offset step to the query.
func (gcaq *GoodCoinAchievementQuery) Offset(offset int) *GoodCoinAchievementQuery {
	gcaq.offset = &offset
	return gcaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gcaq *GoodCoinAchievementQuery) Unique(unique bool) *GoodCoinAchievementQuery {
	gcaq.unique = &unique
	return gcaq
}

// Order adds an order step to the query.
func (gcaq *GoodCoinAchievementQuery) Order(o ...OrderFunc) *GoodCoinAchievementQuery {
	gcaq.order = append(gcaq.order, o...)
	return gcaq
}

// First returns the first GoodCoinAchievement entity from the query.
// Returns a *NotFoundError when no GoodCoinAchievement was found.
func (gcaq *GoodCoinAchievementQuery) First(ctx context.Context) (*GoodCoinAchievement, error) {
	nodes, err := gcaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodcoinachievement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) FirstX(ctx context.Context) *GoodCoinAchievement {
	node, err := gcaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodCoinAchievement ID from the query.
// Returns a *NotFoundError when no GoodCoinAchievement ID was found.
func (gcaq *GoodCoinAchievementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gcaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodcoinachievement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := gcaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodCoinAchievement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodCoinAchievement entity is found.
// Returns a *NotFoundError when no GoodCoinAchievement entities are found.
func (gcaq *GoodCoinAchievementQuery) Only(ctx context.Context) (*GoodCoinAchievement, error) {
	nodes, err := gcaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodcoinachievement.Label}
	default:
		return nil, &NotSingularError{goodcoinachievement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) OnlyX(ctx context.Context) *GoodCoinAchievement {
	node, err := gcaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodCoinAchievement ID in the query.
// Returns a *NotSingularError when more than one GoodCoinAchievement ID is found.
// Returns a *NotFoundError when no entities are found.
func (gcaq *GoodCoinAchievementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gcaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodcoinachievement.Label}
	default:
		err = &NotSingularError{goodcoinachievement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := gcaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodCoinAchievements.
func (gcaq *GoodCoinAchievementQuery) All(ctx context.Context) ([]*GoodCoinAchievement, error) {
	if err := gcaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gcaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) AllX(ctx context.Context) []*GoodCoinAchievement {
	nodes, err := gcaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodCoinAchievement IDs.
func (gcaq *GoodCoinAchievementQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := gcaq.Select(goodcoinachievement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := gcaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcaq *GoodCoinAchievementQuery) Count(ctx context.Context) (int, error) {
	if err := gcaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gcaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) CountX(ctx context.Context) int {
	count, err := gcaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcaq *GoodCoinAchievementQuery) Exist(ctx context.Context) (bool, error) {
	if err := gcaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gcaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gcaq *GoodCoinAchievementQuery) ExistX(ctx context.Context) bool {
	exist, err := gcaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodCoinAchievementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcaq *GoodCoinAchievementQuery) Clone() *GoodCoinAchievementQuery {
	if gcaq == nil {
		return nil
	}
	return &GoodCoinAchievementQuery{
		config:     gcaq.config,
		limit:      gcaq.limit,
		offset:     gcaq.offset,
		order:      append([]OrderFunc{}, gcaq.order...),
		predicates: append([]predicate.GoodCoinAchievement{}, gcaq.predicates...),
		// clone intermediate query.
		sql:    gcaq.sql.Clone(),
		path:   gcaq.path,
		unique: gcaq.unique,
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
//	client.GoodCoinAchievement.Query().
//		GroupBy(goodcoinachievement.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gcaq *GoodCoinAchievementQuery) GroupBy(field string, fields ...string) *GoodCoinAchievementGroupBy {
	grbuild := &GoodCoinAchievementGroupBy{config: gcaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gcaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gcaq.sqlQuery(ctx), nil
	}
	grbuild.label = goodcoinachievement.Label
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
//	client.GoodCoinAchievement.Query().
//		Select(goodcoinachievement.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (gcaq *GoodCoinAchievementQuery) Select(fields ...string) *GoodCoinAchievementSelect {
	gcaq.fields = append(gcaq.fields, fields...)
	selbuild := &GoodCoinAchievementSelect{GoodCoinAchievementQuery: gcaq}
	selbuild.label = goodcoinachievement.Label
	selbuild.flds, selbuild.scan = &gcaq.fields, selbuild.Scan
	return selbuild
}

func (gcaq *GoodCoinAchievementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gcaq.fields {
		if !goodcoinachievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gcaq.path != nil {
		prev, err := gcaq.path(ctx)
		if err != nil {
			return err
		}
		gcaq.sql = prev
	}
	if goodcoinachievement.Policy == nil {
		return errors.New("ent: uninitialized goodcoinachievement.Policy (forgotten import ent/runtime?)")
	}
	if err := goodcoinachievement.Policy.EvalQuery(ctx, gcaq); err != nil {
		return err
	}
	return nil
}

func (gcaq *GoodCoinAchievementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodCoinAchievement, error) {
	var (
		nodes = []*GoodCoinAchievement{}
		_spec = gcaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodCoinAchievement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodCoinAchievement{config: gcaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gcaq.modifiers) > 0 {
		_spec.Modifiers = gcaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gcaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gcaq *GoodCoinAchievementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gcaq.querySpec()
	if len(gcaq.modifiers) > 0 {
		_spec.Modifiers = gcaq.modifiers
	}
	_spec.Node.Columns = gcaq.fields
	if len(gcaq.fields) > 0 {
		_spec.Unique = gcaq.unique != nil && *gcaq.unique
	}
	return sqlgraph.CountNodes(ctx, gcaq.driver, _spec)
}

func (gcaq *GoodCoinAchievementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gcaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gcaq *GoodCoinAchievementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcoinachievement.Table,
			Columns: goodcoinachievement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoinachievement.FieldID,
			},
		},
		From:   gcaq.sql,
		Unique: true,
	}
	if unique := gcaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gcaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodcoinachievement.FieldID)
		for i := range fields {
			if fields[i] != goodcoinachievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gcaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gcaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gcaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gcaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gcaq *GoodCoinAchievementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gcaq.driver.Dialect())
	t1 := builder.Table(goodcoinachievement.Table)
	columns := gcaq.fields
	if len(columns) == 0 {
		columns = goodcoinachievement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gcaq.sql != nil {
		selector = gcaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gcaq.unique != nil && *gcaq.unique {
		selector.Distinct()
	}
	for _, m := range gcaq.modifiers {
		m(selector)
	}
	for _, p := range gcaq.predicates {
		p(selector)
	}
	for _, p := range gcaq.order {
		p(selector)
	}
	if offset := gcaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gcaq *GoodCoinAchievementQuery) ForUpdate(opts ...sql.LockOption) *GoodCoinAchievementQuery {
	if gcaq.driver.Dialect() == dialect.Postgres {
		gcaq.Unique(false)
	}
	gcaq.modifiers = append(gcaq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gcaq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gcaq *GoodCoinAchievementQuery) ForShare(opts ...sql.LockOption) *GoodCoinAchievementQuery {
	if gcaq.driver.Dialect() == dialect.Postgres {
		gcaq.Unique(false)
	}
	gcaq.modifiers = append(gcaq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gcaq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gcaq *GoodCoinAchievementQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodCoinAchievementSelect {
	gcaq.modifiers = append(gcaq.modifiers, modifiers...)
	return gcaq.Select()
}

// GoodCoinAchievementGroupBy is the group-by builder for GoodCoinAchievement entities.
type GoodCoinAchievementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcagb *GoodCoinAchievementGroupBy) Aggregate(fns ...AggregateFunc) *GoodCoinAchievementGroupBy {
	gcagb.fns = append(gcagb.fns, fns...)
	return gcagb
}

// Scan applies the group-by query and scans the result into the given value.
func (gcagb *GoodCoinAchievementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gcagb.path(ctx)
	if err != nil {
		return err
	}
	gcagb.sql = query
	return gcagb.sqlScan(ctx, v)
}

func (gcagb *GoodCoinAchievementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gcagb.fields {
		if !goodcoinachievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gcagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gcagb *GoodCoinAchievementGroupBy) sqlQuery() *sql.Selector {
	selector := gcagb.sql.Select()
	aggregation := make([]string, 0, len(gcagb.fns))
	for _, fn := range gcagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gcagb.fields)+len(gcagb.fns))
		for _, f := range gcagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gcagb.fields...)...)
}

// GoodCoinAchievementSelect is the builder for selecting fields of GoodCoinAchievement entities.
type GoodCoinAchievementSelect struct {
	*GoodCoinAchievementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gcas *GoodCoinAchievementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gcas.prepareQuery(ctx); err != nil {
		return err
	}
	gcas.sql = gcas.GoodCoinAchievementQuery.sqlQuery(ctx)
	return gcas.sqlScan(ctx, v)
}

func (gcas *GoodCoinAchievementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gcas.sql.Query()
	if err := gcas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gcas *GoodCoinAchievementSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodCoinAchievementSelect {
	gcas.modifiers = append(gcas.modifiers, modifiers...)
	return gcas
}