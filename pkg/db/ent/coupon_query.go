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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/coupon"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// CouponQuery is the builder for querying Coupon entities.
type CouponQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Coupon
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CouponQuery builder.
func (cq *CouponQuery) Where(ps ...predicate.Coupon) *CouponQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CouponQuery) Limit(limit int) *CouponQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CouponQuery) Offset(offset int) *CouponQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CouponQuery) Unique(unique bool) *CouponQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CouponQuery) Order(o ...OrderFunc) *CouponQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Coupon entity from the query.
// Returns a *NotFoundError when no Coupon was found.
func (cq *CouponQuery) First(ctx context.Context) (*Coupon, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coupon.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CouponQuery) FirstX(ctx context.Context) *Coupon {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Coupon ID from the query.
// Returns a *NotFoundError when no Coupon ID was found.
func (cq *CouponQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coupon.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CouponQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Coupon entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Coupon entity is found.
// Returns a *NotFoundError when no Coupon entities are found.
func (cq *CouponQuery) Only(ctx context.Context) (*Coupon, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coupon.Label}
	default:
		return nil, &NotSingularError{coupon.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CouponQuery) OnlyX(ctx context.Context) *Coupon {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Coupon ID in the query.
// Returns a *NotSingularError when more than one Coupon ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CouponQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coupon.Label}
	default:
		err = &NotSingularError{coupon.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CouponQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Coupons.
func (cq *CouponQuery) All(ctx context.Context) ([]*Coupon, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CouponQuery) AllX(ctx context.Context) []*Coupon {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Coupon IDs.
func (cq *CouponQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := cq.Select(coupon.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CouponQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CouponQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CouponQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CouponQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CouponQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CouponQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CouponQuery) Clone() *CouponQuery {
	if cq == nil {
		return nil
	}
	return &CouponQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		predicates: append([]predicate.Coupon{}, cq.predicates...),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
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
//	client.Coupon.Query().
//		GroupBy(coupon.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CouponQuery) GroupBy(field string, fields ...string) *CouponGroupBy {
	grbuild := &CouponGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = coupon.Label
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
//	client.Coupon.Query().
//		Select(coupon.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (cq *CouponQuery) Select(fields ...string) *CouponSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &CouponSelect{CouponQuery: cq}
	selbuild.label = coupon.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *CouponQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !coupon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	if coupon.Policy == nil {
		return errors.New("ent: uninitialized coupon.Policy (forgotten import ent/runtime?)")
	}
	if err := coupon.Policy.EvalQuery(ctx, cq); err != nil {
		return err
	}
	return nil
}

func (cq *CouponQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Coupon, error) {
	var (
		nodes = []*Coupon{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Coupon).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Coupon{config: cq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *CouponQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CouponQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cq *CouponQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coupon.Table,
			Columns: coupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coupon.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coupon.FieldID)
		for i := range fields {
			if fields[i] != coupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CouponQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(coupon.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = coupon.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cq *CouponQuery) ForUpdate(opts ...sql.LockOption) *CouponQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cq *CouponQuery) ForShare(opts ...sql.LockOption) *CouponQuery {
	if cq.driver.Dialect() == dialect.Postgres {
		cq.Unique(false)
	}
	cq.modifiers = append(cq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CouponQuery) Modify(modifiers ...func(s *sql.Selector)) *CouponSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CouponGroupBy is the group-by builder for Coupon entities.
type CouponGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CouponGroupBy) Aggregate(fns ...AggregateFunc) *CouponGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CouponGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *CouponGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !coupon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CouponGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// CouponSelect is the builder for selecting fields of Coupon entities.
type CouponSelect struct {
	*CouponQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CouponSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CouponQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *CouponSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CouponSelect) Modify(modifiers ...func(s *sql.Selector)) *CouponSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
