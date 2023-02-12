// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/claims"
	"license-manager/pkg/repositories/ent-fw/ent/credentials"
	"license-manager/pkg/repositories/ent-fw/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClaimsQuery is the builder for querying Claims entities.
type ClaimsQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Claims
	withClaimer *CredentialsQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClaimsQuery builder.
func (cq *ClaimsQuery) Where(ps ...predicate.Claims) *ClaimsQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ClaimsQuery) Limit(limit int) *ClaimsQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ClaimsQuery) Offset(offset int) *ClaimsQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ClaimsQuery) Unique(unique bool) *ClaimsQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ClaimsQuery) Order(o ...OrderFunc) *ClaimsQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryClaimer chains the current query on the "claimer" edge.
func (cq *ClaimsQuery) QueryClaimer() *CredentialsQuery {
	query := (&CredentialsClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(claims.Table, claims.FieldID, selector),
			sqlgraph.To(credentials.Table, credentials.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, claims.ClaimerTable, claims.ClaimerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Claims entity from the query.
// Returns a *NotFoundError when no Claims was found.
func (cq *ClaimsQuery) First(ctx context.Context) (*Claims, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{claims.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ClaimsQuery) FirstX(ctx context.Context) *Claims {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Claims ID from the query.
// Returns a *NotFoundError when no Claims ID was found.
func (cq *ClaimsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{claims.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ClaimsQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Claims entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Claims entity is found.
// Returns a *NotFoundError when no Claims entities are found.
func (cq *ClaimsQuery) Only(ctx context.Context) (*Claims, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{claims.Label}
	default:
		return nil, &NotSingularError{claims.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ClaimsQuery) OnlyX(ctx context.Context) *Claims {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Claims ID in the query.
// Returns a *NotSingularError when more than one Claims ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ClaimsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{claims.Label}
	default:
		err = &NotSingularError{claims.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ClaimsQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClaimsSlice.
func (cq *ClaimsQuery) All(ctx context.Context) ([]*Claims, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Claims, *ClaimsQuery]()
	return withInterceptors[[]*Claims](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ClaimsQuery) AllX(ctx context.Context) []*Claims {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Claims IDs.
func (cq *ClaimsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err := cq.Select(claims.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ClaimsQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ClaimsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ClaimsQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ClaimsQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ClaimsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ClaimsQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClaimsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ClaimsQuery) Clone() *ClaimsQuery {
	if cq == nil {
		return nil
	}
	return &ClaimsQuery{
		config:      cq.config,
		ctx:         cq.ctx.Clone(),
		order:       append([]OrderFunc{}, cq.order...),
		inters:      append([]Interceptor{}, cq.inters...),
		predicates:  append([]predicate.Claims{}, cq.predicates...),
		withClaimer: cq.withClaimer.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithClaimer tells the query-builder to eager-load the nodes that are connected to
// the "claimer" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClaimsQuery) WithClaimer(opts ...func(*CredentialsQuery)) *ClaimsQuery {
	query := (&CredentialsClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withClaimer = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Claims.Query().
//		GroupBy(claims.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ClaimsQuery) GroupBy(field string, fields ...string) *ClaimsGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClaimsGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = claims.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//	}
//
//	client.Claims.Query().
//		Select(claims.FieldKey).
//		Scan(ctx, &v)
func (cq *ClaimsQuery) Select(fields ...string) *ClaimsSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ClaimsSelect{ClaimsQuery: cq}
	sbuild.label = claims.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClaimsSelect configured with the given aggregations.
func (cq *ClaimsQuery) Aggregate(fns ...AggregateFunc) *ClaimsSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ClaimsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !claims.ValidColumn(f) {
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
	return nil
}

func (cq *ClaimsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Claims, error) {
	var (
		nodes       = []*Claims{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withClaimer != nil,
		}
	)
	if cq.withClaimer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, claims.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Claims).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Claims{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := cq.withClaimer; query != nil {
		if err := cq.loadClaimer(ctx, query, nodes, nil,
			func(n *Claims, e *Credentials) { n.Edges.Claimer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ClaimsQuery) loadClaimer(ctx context.Context, query *CredentialsQuery, nodes []*Claims, init func(*Claims), assign func(*Claims, *Credentials)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Claims)
	for i := range nodes {
		if nodes[i].credentials_claims == nil {
			continue
		}
		fk := *nodes[i].credentials_claims
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(credentials.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "credentials_claims" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ClaimsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ClaimsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   claims.Table,
			Columns: claims.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: claims.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, claims.FieldID)
		for i := range fields {
			if fields[i] != claims.FieldID {
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
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
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

func (cq *ClaimsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(claims.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = claims.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClaimsGroupBy is the group-by builder for Claims entities.
type ClaimsGroupBy struct {
	selector
	build *ClaimsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ClaimsGroupBy) Aggregate(fns ...AggregateFunc) *ClaimsGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ClaimsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClaimsQuery, *ClaimsGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ClaimsGroupBy) sqlScan(ctx context.Context, root *ClaimsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClaimsSelect is the builder for selecting fields of Claims entities.
type ClaimsSelect struct {
	*ClaimsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ClaimsSelect) Aggregate(fns ...AggregateFunc) *ClaimsSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ClaimsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClaimsQuery, *ClaimsSelect](ctx, cs.ClaimsQuery, cs, cs.inters, v)
}

func (cs *ClaimsSelect) sqlScan(ctx context.Context, root *ClaimsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
