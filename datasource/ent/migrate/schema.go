// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "editor", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(1,0)", "postgres": "numeric"}},
		{Name: "open_code", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "user_accounts", Type: field.TypeUUID, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_accounts",
				Columns:    []*schema.Column{AccountsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "account_id_user_accounts",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[0], AccountsColumns[8]},
			},
		},
	}
	// GroupColumns holds the columns for the "group" table.
	GroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// GroupTable holds the schema information for the "group" table.
	GroupTable = &schema.Table{
		Name:       "group",
		Columns:    GroupColumns,
		PrimaryKey: []*schema.Column{GroupColumns[0]},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "editor", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(1,0)", "postgres": "numeric"}},
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
		{Name: "category", Type: field.TypeInt},
		{Name: "url", Type: field.TypeInt},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "editor", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(1,0)", "postgres": "numeric"}},
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "editor", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(1,0)", "postgres": "numeric"}},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "city", Type: field.TypeString, Nullable: true},
		{Name: "introduction", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Size: 11},
		{Name: "password", Type: field.TypeString},
		{Name: "state", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "editor", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(1,0)", "postgres": "numeric"}},
		{Name: "parent_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
	}
	// RolePermissionsColumns holds the columns for the "role_permissions" table.
	RolePermissionsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "permission_id", Type: field.TypeUUID},
	}
	// RolePermissionsTable holds the schema information for the "role_permissions" table.
	RolePermissionsTable = &schema.Table{
		Name:       "role_permissions",
		Columns:    RolePermissionsColumns,
		PrimaryKey: []*schema.Column{RolePermissionsColumns[0], RolePermissionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permissions_role_id",
				Columns:    []*schema.Column{RolePermissionsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_permissions_permission_id",
				Columns:    []*schema.Column{RolePermissionsColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupUsersColumns holds the columns for the "user_group_users" table.
	UserGroupUsersColumns = []*schema.Column{
		{Name: "user_group_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// UserGroupUsersTable holds the schema information for the "user_group_users" table.
	UserGroupUsersTable = &schema.Table{
		Name:       "user_group_users",
		Columns:    UserGroupUsersColumns,
		PrimaryKey: []*schema.Column{UserGroupUsersColumns[0], UserGroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_group_users_user_group_id",
				Columns:    []*schema.Column{UserGroupUsersColumns[0]},
				RefColumns: []*schema.Column{UserGroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_group_users_user_id",
				Columns:    []*schema.Column{UserGroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupRolesColumns holds the columns for the "user_group_roles" table.
	UserGroupRolesColumns = []*schema.Column{
		{Name: "user_group_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
	}
	// UserGroupRolesTable holds the schema information for the "user_group_roles" table.
	UserGroupRolesTable = &schema.Table{
		Name:       "user_group_roles",
		Columns:    UserGroupRolesColumns,
		PrimaryKey: []*schema.Column{UserGroupRolesColumns[0], UserGroupRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_group_roles_user_group_id",
				Columns:    []*schema.Column{UserGroupRolesColumns[0]},
				RefColumns: []*schema.Column{UserGroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_group_roles_role_id",
				Columns:    []*schema.Column{UserGroupRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		GroupTable,
		PermissionsTable,
		RolesTable,
		UsersTable,
		UserGroupsTable,
		RolePermissionsTable,
		UserRolesTable,
		UserGroupUsersTable,
		UserGroupRolesTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	AccountsTable.Annotation = &entsql.Annotation{
		Table: "accounts",
	}
	GroupTable.Annotation = &entsql.Annotation{
		Table: "group",
	}
	PermissionsTable.Annotation = &entsql.Annotation{
		Table: "permissions",
	}
	RolesTable.Annotation = &entsql.Annotation{
		Table: "roles",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
	UserGroupsTable.Annotation = &entsql.Annotation{
		Table: "user_groups",
	}
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
	UserGroupUsersTable.ForeignKeys[0].RefTable = UserGroupsTable
	UserGroupUsersTable.ForeignKeys[1].RefTable = UsersTable
	UserGroupRolesTable.ForeignKeys[0].RefTable = UserGroupsTable
	UserGroupRolesTable.ForeignKeys[1].RefTable = RolesTable
}
