// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/predicate"
	"SWTAC/datasource/ent/role"
	"SWTAC/datasource/ent/user"
	"SWTAC/datasource/ent/usergroup"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserGroupQuery is the builder for querying UserGroup entities.
type UserGroupQuery struct {
	config
	ctx        *QueryContext
	order      []usergroup.OrderOption
	inters     []Interceptor
	predicates []predicate.UserGroup
	withUsers  *UserQuery
	withRoles  *RoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserGroupQuery builder.
func (ugq *UserGroupQuery) Where(ps ...predicate.UserGroup) *UserGroupQuery {
	ugq.predicates = append(ugq.predicates, ps...)
	return ugq
}

// Limit the number of records to be returned by this query.
func (ugq *UserGroupQuery) Limit(limit int) *UserGroupQuery {
	ugq.ctx.Limit = &limit
	return ugq
}

// Offset to start from.
func (ugq *UserGroupQuery) Offset(offset int) *UserGroupQuery {
	ugq.ctx.Offset = &offset
	return ugq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ugq *UserGroupQuery) Unique(unique bool) *UserGroupQuery {
	ugq.ctx.Unique = &unique
	return ugq
}

// Order specifies how the records should be ordered.
func (ugq *UserGroupQuery) Order(o ...usergroup.OrderOption) *UserGroupQuery {
	ugq.order = append(ugq.order, o...)
	return ugq
}

// QueryUsers chains the current query on the "users" edge.
func (ugq *UserGroupQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: ugq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ugq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usergroup.Table, usergroup.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usergroup.UsersTable, usergroup.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ugq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoles chains the current query on the "roles" edge.
func (ugq *UserGroupQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: ugq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ugq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ugq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usergroup.Table, usergroup.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usergroup.RolesTable, usergroup.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ugq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserGroup entity from the query.
// Returns a *NotFoundError when no UserGroup was found.
func (ugq *UserGroupQuery) First(ctx context.Context) (*UserGroup, error) {
	nodes, err := ugq.Limit(1).All(setContextOp(ctx, ugq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usergroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ugq *UserGroupQuery) FirstX(ctx context.Context) *UserGroup {
	node, err := ugq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserGroup ID from the query.
// Returns a *NotFoundError when no UserGroup ID was found.
func (ugq *UserGroupQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ugq.Limit(1).IDs(setContextOp(ctx, ugq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usergroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ugq *UserGroupQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ugq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserGroup entity is found.
// Returns a *NotFoundError when no UserGroup entities are found.
func (ugq *UserGroupQuery) Only(ctx context.Context) (*UserGroup, error) {
	nodes, err := ugq.Limit(2).All(setContextOp(ctx, ugq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usergroup.Label}
	default:
		return nil, &NotSingularError{usergroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ugq *UserGroupQuery) OnlyX(ctx context.Context) *UserGroup {
	node, err := ugq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserGroup ID in the query.
// Returns a *NotSingularError when more than one UserGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (ugq *UserGroupQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ugq.Limit(2).IDs(setContextOp(ctx, ugq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usergroup.Label}
	default:
		err = &NotSingularError{usergroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ugq *UserGroupQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ugq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserGroups.
func (ugq *UserGroupQuery) All(ctx context.Context) ([]*UserGroup, error) {
	ctx = setContextOp(ctx, ugq.ctx, "All")
	if err := ugq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserGroup, *UserGroupQuery]()
	return withInterceptors[[]*UserGroup](ctx, ugq, qr, ugq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ugq *UserGroupQuery) AllX(ctx context.Context) []*UserGroup {
	nodes, err := ugq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserGroup IDs.
func (ugq *UserGroupQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ugq.ctx.Unique == nil && ugq.path != nil {
		ugq.Unique(true)
	}
	ctx = setContextOp(ctx, ugq.ctx, "IDs")
	if err = ugq.Select(usergroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ugq *UserGroupQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ugq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ugq *UserGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ugq.ctx, "Count")
	if err := ugq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ugq, querierCount[*UserGroupQuery](), ugq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ugq *UserGroupQuery) CountX(ctx context.Context) int {
	count, err := ugq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ugq *UserGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ugq.ctx, "Exist")
	switch _, err := ugq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ugq *UserGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := ugq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ugq *UserGroupQuery) Clone() *UserGroupQuery {
	if ugq == nil {
		return nil
	}
	return &UserGroupQuery{
		config:     ugq.config,
		ctx:        ugq.ctx.Clone(),
		order:      append([]usergroup.OrderOption{}, ugq.order...),
		inters:     append([]Interceptor{}, ugq.inters...),
		predicates: append([]predicate.UserGroup{}, ugq.predicates...),
		withUsers:  ugq.withUsers.Clone(),
		withRoles:  ugq.withRoles.Clone(),
		// clone intermediate query.
		sql:  ugq.sql.Clone(),
		path: ugq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (ugq *UserGroupQuery) WithUsers(opts ...func(*UserQuery)) *UserGroupQuery {
	query := (&UserClient{config: ugq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ugq.withUsers = query
	return ugq
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (ugq *UserGroupQuery) WithRoles(opts ...func(*RoleQuery)) *UserGroupQuery {
	query := (&RoleClient{config: ugq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ugq.withRoles = query
	return ugq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserGroup.Query().
//		GroupBy(usergroup.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ugq *UserGroupQuery) GroupBy(field string, fields ...string) *UserGroupGroupBy {
	ugq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupGroupBy{build: ugq}
	grbuild.flds = &ugq.ctx.Fields
	grbuild.label = usergroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserGroup.Query().
//		Select(usergroup.FieldCreatedAt).
//		Scan(ctx, &v)
func (ugq *UserGroupQuery) Select(fields ...string) *UserGroupSelect {
	ugq.ctx.Fields = append(ugq.ctx.Fields, fields...)
	sbuild := &UserGroupSelect{UserGroupQuery: ugq}
	sbuild.label = usergroup.Label
	sbuild.flds, sbuild.scan = &ugq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserGroupSelect configured with the given aggregations.
func (ugq *UserGroupQuery) Aggregate(fns ...AggregateFunc) *UserGroupSelect {
	return ugq.Select().Aggregate(fns...)
}

func (ugq *UserGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ugq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ugq); err != nil {
				return err
			}
		}
	}
	for _, f := range ugq.ctx.Fields {
		if !usergroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ugq.path != nil {
		prev, err := ugq.path(ctx)
		if err != nil {
			return err
		}
		ugq.sql = prev
	}
	return nil
}

func (ugq *UserGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserGroup, error) {
	var (
		nodes       = []*UserGroup{}
		_spec       = ugq.querySpec()
		loadedTypes = [2]bool{
			ugq.withUsers != nil,
			ugq.withRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserGroup{config: ugq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ugq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ugq.withUsers; query != nil {
		if err := ugq.loadUsers(ctx, query, nodes,
			func(n *UserGroup) { n.Edges.Users = []*User{} },
			func(n *UserGroup, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := ugq.withRoles; query != nil {
		if err := ugq.loadRoles(ctx, query, nodes,
			func(n *UserGroup) { n.Edges.Roles = []*Role{} },
			func(n *UserGroup, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ugq *UserGroupQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*UserGroup, init func(*UserGroup), assign func(*UserGroup, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*UserGroup)
	nids := make(map[uuid.UUID]map[*UserGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(usergroup.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(usergroup.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(usergroup.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(usergroup.UsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*UserGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ugq *UserGroupQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*UserGroup, init func(*UserGroup), assign func(*UserGroup, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*UserGroup)
	nids := make(map[uuid.UUID]map[*UserGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(usergroup.RolesTable)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(usergroup.RolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(usergroup.RolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(usergroup.RolesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*UserGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Role](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ugq *UserGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ugq.querySpec()
	_spec.Node.Columns = ugq.ctx.Fields
	if len(ugq.ctx.Fields) > 0 {
		_spec.Unique = ugq.ctx.Unique != nil && *ugq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ugq.driver, _spec)
}

func (ugq *UserGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usergroup.Table, usergroup.Columns, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID))
	_spec.From = ugq.sql
	if unique := ugq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ugq.path != nil {
		_spec.Unique = true
	}
	if fields := ugq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usergroup.FieldID)
		for i := range fields {
			if fields[i] != usergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ugq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ugq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ugq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ugq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ugq *UserGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ugq.driver.Dialect())
	t1 := builder.Table(usergroup.Table)
	columns := ugq.ctx.Fields
	if len(columns) == 0 {
		columns = usergroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ugq.sql != nil {
		selector = ugq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ugq.ctx.Unique != nil && *ugq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ugq.predicates {
		p(selector)
	}
	for _, p := range ugq.order {
		p(selector)
	}
	if offset := ugq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ugq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupGroupBy is the group-by builder for UserGroup entities.
type UserGroupGroupBy struct {
	selector
	build *UserGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uggb *UserGroupGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupGroupBy {
	uggb.fns = append(uggb.fns, fns...)
	return uggb
}

// Scan applies the selector query and scans the result into the given value.
func (uggb *UserGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uggb.build.ctx, "GroupBy")
	if err := uggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserGroupQuery, *UserGroupGroupBy](ctx, uggb.build, uggb, uggb.build.inters, v)
}

func (uggb *UserGroupGroupBy) sqlScan(ctx context.Context, root *UserGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uggb.fns))
	for _, fn := range uggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uggb.flds)+len(uggb.fns))
		for _, f := range *uggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserGroupSelect is the builder for selecting fields of UserGroup entities.
type UserGroupSelect struct {
	*UserGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ugs *UserGroupSelect) Aggregate(fns ...AggregateFunc) *UserGroupSelect {
	ugs.fns = append(ugs.fns, fns...)
	return ugs
}

// Scan applies the selector query and scans the result into the given value.
func (ugs *UserGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugs.ctx, "Select")
	if err := ugs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserGroupQuery, *UserGroupSelect](ctx, ugs.UserGroupQuery, ugs, ugs.inters, v)
}

func (ugs *UserGroupSelect) sqlScan(ctx context.Context, root *UserGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ugs.fns))
	for _, fn := range ugs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ugs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
