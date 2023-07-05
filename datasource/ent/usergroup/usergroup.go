// Code generated by ent, DO NOT EDIT.

package usergroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usergroup type in the database.
	Label = "user_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the usergroup in the database.
	Table = "user_group"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_group_user"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// RoleTable is the table that holds the role relation/edge. The primary key declared below.
	RoleTable = "user_group_role"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "role"
)

// Columns holds all SQL columns for usergroup fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldName,
	FieldCode,
	FieldIntro,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_group_id", "user_id"}
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"user_group_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the UserGroup queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByIntro orders the results by the intro field.
func ByIntro(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntro, opts...).ToFunc()
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoleCount orders the results by role count.
func ByRoleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoleStep(), opts...)
	}
}

// ByRole orders the results by role terms.
func ByRole(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
	)
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RoleTable, RolePrimaryKey...),
	)
}
