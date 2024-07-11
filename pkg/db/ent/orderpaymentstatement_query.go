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
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/orderpaymentstatement"
	"github.com/NpoolPlatform/inspire-middleware/pkg/db/ent/predicate"
)

// OrderPaymentStatementQuery is the builder for querying OrderPaymentStatement entities.
type OrderPaymentStatementQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderPaymentStatement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderPaymentStatementQuery builder.
func (opsq *OrderPaymentStatementQuery) Where(ps ...predicate.OrderPaymentStatement) *OrderPaymentStatementQuery {
	opsq.predicates = append(opsq.predicates, ps...)
	return opsq
}

// Limit adds a limit step to the query.
func (opsq *OrderPaymentStatementQuery) Limit(limit int) *OrderPaymentStatementQuery {
	opsq.limit = &limit
	return opsq
}

// Offset adds an offset step to the query.
func (opsq *OrderPaymentStatementQuery) Offset(offset int) *OrderPaymentStatementQuery {
	opsq.offset = &offset
	return opsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opsq *OrderPaymentStatementQuery) Unique(unique bool) *OrderPaymentStatementQuery {
	opsq.unique = &unique
	return opsq
}

// Order adds an order step to the query.
func (opsq *OrderPaymentStatementQuery) Order(o ...OrderFunc) *OrderPaymentStatementQuery {
	opsq.order = append(opsq.order, o...)
	return opsq
}

// First returns the first OrderPaymentStatement entity from the query.
// Returns a *NotFoundError when no OrderPaymentStatement was found.
func (opsq *OrderPaymentStatementQuery) First(ctx context.Context) (*OrderPaymentStatement, error) {
	nodes, err := opsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderpaymentstatement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) FirstX(ctx context.Context) *OrderPaymentStatement {
	node, err := opsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderPaymentStatement ID from the query.
// Returns a *NotFoundError when no OrderPaymentStatement ID was found.
func (opsq *OrderPaymentStatementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = opsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderpaymentstatement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := opsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderPaymentStatement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderPaymentStatement entity is found.
// Returns a *NotFoundError when no OrderPaymentStatement entities are found.
func (opsq *OrderPaymentStatementQuery) Only(ctx context.Context) (*OrderPaymentStatement, error) {
	nodes, err := opsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderpaymentstatement.Label}
	default:
		return nil, &NotSingularError{orderpaymentstatement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) OnlyX(ctx context.Context) *OrderPaymentStatement {
	node, err := opsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderPaymentStatement ID in the query.
// Returns a *NotSingularError when more than one OrderPaymentStatement ID is found.
// Returns a *NotFoundError when no entities are found.
func (opsq *OrderPaymentStatementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = opsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderpaymentstatement.Label}
	default:
		err = &NotSingularError{orderpaymentstatement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := opsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderPaymentStatements.
func (opsq *OrderPaymentStatementQuery) All(ctx context.Context) ([]*OrderPaymentStatement, error) {
	if err := opsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return opsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) AllX(ctx context.Context) []*OrderPaymentStatement {
	nodes, err := opsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderPaymentStatement IDs.
func (opsq *OrderPaymentStatementQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := opsq.Select(orderpaymentstatement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := opsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opsq *OrderPaymentStatementQuery) Count(ctx context.Context) (int, error) {
	if err := opsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return opsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) CountX(ctx context.Context) int {
	count, err := opsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opsq *OrderPaymentStatementQuery) Exist(ctx context.Context) (bool, error) {
	if err := opsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return opsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (opsq *OrderPaymentStatementQuery) ExistX(ctx context.Context) bool {
	exist, err := opsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderPaymentStatementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opsq *OrderPaymentStatementQuery) Clone() *OrderPaymentStatementQuery {
	if opsq == nil {
		return nil
	}
	return &OrderPaymentStatementQuery{
		config:     opsq.config,
		limit:      opsq.limit,
		offset:     opsq.offset,
		order:      append([]OrderFunc{}, opsq.order...),
		predicates: append([]predicate.OrderPaymentStatement{}, opsq.predicates...),
		// clone intermediate query.
		sql:    opsq.sql.Clone(),
		path:   opsq.path,
		unique: opsq.unique,
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
//	client.OrderPaymentStatement.Query().
//		GroupBy(orderpaymentstatement.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (opsq *OrderPaymentStatementQuery) GroupBy(field string, fields ...string) *OrderPaymentStatementGroupBy {
	grbuild := &OrderPaymentStatementGroupBy{config: opsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := opsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return opsq.sqlQuery(ctx), nil
	}
	grbuild.label = orderpaymentstatement.Label
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
//	client.OrderPaymentStatement.Query().
//		Select(orderpaymentstatement.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (opsq *OrderPaymentStatementQuery) Select(fields ...string) *OrderPaymentStatementSelect {
	opsq.fields = append(opsq.fields, fields...)
	selbuild := &OrderPaymentStatementSelect{OrderPaymentStatementQuery: opsq}
	selbuild.label = orderpaymentstatement.Label
	selbuild.flds, selbuild.scan = &opsq.fields, selbuild.Scan
	return selbuild
}

func (opsq *OrderPaymentStatementQuery) prepareQuery(ctx context.Context) error {
	for _, f := range opsq.fields {
		if !orderpaymentstatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opsq.path != nil {
		prev, err := opsq.path(ctx)
		if err != nil {
			return err
		}
		opsq.sql = prev
	}
	if orderpaymentstatement.Policy == nil {
		return errors.New("ent: uninitialized orderpaymentstatement.Policy (forgotten import ent/runtime?)")
	}
	if err := orderpaymentstatement.Policy.EvalQuery(ctx, opsq); err != nil {
		return err
	}
	return nil
}

func (opsq *OrderPaymentStatementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderPaymentStatement, error) {
	var (
		nodes = []*OrderPaymentStatement{}
		_spec = opsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrderPaymentStatement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrderPaymentStatement{config: opsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(opsq.modifiers) > 0 {
		_spec.Modifiers = opsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (opsq *OrderPaymentStatementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opsq.querySpec()
	if len(opsq.modifiers) > 0 {
		_spec.Modifiers = opsq.modifiers
	}
	_spec.Node.Columns = opsq.fields
	if len(opsq.fields) > 0 {
		_spec.Unique = opsq.unique != nil && *opsq.unique
	}
	return sqlgraph.CountNodes(ctx, opsq.driver, _spec)
}

func (opsq *OrderPaymentStatementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := opsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (opsq *OrderPaymentStatementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymentstatement.Table,
			Columns: orderpaymentstatement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymentstatement.FieldID,
			},
		},
		From:   opsq.sql,
		Unique: true,
	}
	if unique := opsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := opsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpaymentstatement.FieldID)
		for i := range fields {
			if fields[i] != orderpaymentstatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := opsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opsq *OrderPaymentStatementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opsq.driver.Dialect())
	t1 := builder.Table(orderpaymentstatement.Table)
	columns := opsq.fields
	if len(columns) == 0 {
		columns = orderpaymentstatement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opsq.sql != nil {
		selector = opsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opsq.unique != nil && *opsq.unique {
		selector.Distinct()
	}
	for _, m := range opsq.modifiers {
		m(selector)
	}
	for _, p := range opsq.predicates {
		p(selector)
	}
	for _, p := range opsq.order {
		p(selector)
	}
	if offset := opsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (opsq *OrderPaymentStatementQuery) ForUpdate(opts ...sql.LockOption) *OrderPaymentStatementQuery {
	if opsq.driver.Dialect() == dialect.Postgres {
		opsq.Unique(false)
	}
	opsq.modifiers = append(opsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return opsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (opsq *OrderPaymentStatementQuery) ForShare(opts ...sql.LockOption) *OrderPaymentStatementQuery {
	if opsq.driver.Dialect() == dialect.Postgres {
		opsq.Unique(false)
	}
	opsq.modifiers = append(opsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return opsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opsq *OrderPaymentStatementQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderPaymentStatementSelect {
	opsq.modifiers = append(opsq.modifiers, modifiers...)
	return opsq.Select()
}

// OrderPaymentStatementGroupBy is the group-by builder for OrderPaymentStatement entities.
type OrderPaymentStatementGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opsgb *OrderPaymentStatementGroupBy) Aggregate(fns ...AggregateFunc) *OrderPaymentStatementGroupBy {
	opsgb.fns = append(opsgb.fns, fns...)
	return opsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (opsgb *OrderPaymentStatementGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := opsgb.path(ctx)
	if err != nil {
		return err
	}
	opsgb.sql = query
	return opsgb.sqlScan(ctx, v)
}

func (opsgb *OrderPaymentStatementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range opsgb.fields {
		if !orderpaymentstatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := opsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (opsgb *OrderPaymentStatementGroupBy) sqlQuery() *sql.Selector {
	selector := opsgb.sql.Select()
	aggregation := make([]string, 0, len(opsgb.fns))
	for _, fn := range opsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(opsgb.fields)+len(opsgb.fns))
		for _, f := range opsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(opsgb.fields...)...)
}

// OrderPaymentStatementSelect is the builder for selecting fields of OrderPaymentStatement entities.
type OrderPaymentStatementSelect struct {
	*OrderPaymentStatementQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (opss *OrderPaymentStatementSelect) Scan(ctx context.Context, v interface{}) error {
	if err := opss.prepareQuery(ctx); err != nil {
		return err
	}
	opss.sql = opss.OrderPaymentStatementQuery.sqlQuery(ctx)
	return opss.sqlScan(ctx, v)
}

func (opss *OrderPaymentStatementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := opss.sql.Query()
	if err := opss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opss *OrderPaymentStatementSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderPaymentStatementSelect {
	opss.modifiers = append(opss.modifiers, modifiers...)
	return opss
}
