// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// Editor holds the value of the "editor" field.
	Editor string `json:"editor,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted float64 `json:"deleted,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Introduction holds the value of the "introduction" field.
	Introduction string `json:"introduction,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// State holds the value of the "state" field.
	State int `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserGroups holds the value of the user_groups edge.
	UserGroups []*UserGroup `json:"user_groups,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Accounts holds the value of the accounts edge.
	Accounts []*Account `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserGroupsOrErr returns the UserGroups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserGroupsOrErr() ([]*UserGroup, error) {
	if e.loadedTypes[0] {
		return e.UserGroups, nil
	}
	return nil, &NotLoadedError{edge: "user_groups"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[1] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[2] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldDeleted:
			values[i] = new(sql.NullFloat64)
		case user.FieldAge, user.FieldState:
			values[i] = new(sql.NullInt64)
		case user.FieldCreator, user.FieldEditor, user.FieldUsername, user.FieldNickname, user.FieldAvatar, user.FieldCity, user.FieldIntroduction, user.FieldEmail, user.FieldPhone, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				u.Creator = value.String
			}
		case user.FieldEditor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field editor", values[i])
			} else if value.Valid {
				u.Editor = value.String
			}
		case user.FieldDeleted:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				u.Deleted = value.Float64
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				u.Age = int(value.Int64)
			}
		case user.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				u.City = value.String
			}
		case user.FieldIntroduction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field introduction", values[i])
			} else if value.Valid {
				u.Introduction = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				u.State = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserGroups queries the "user_groups" edge of the User entity.
func (u *User) QueryUserGroups() *UserGroupQuery {
	return NewUserClient(u.config).QueryUserGroups(u)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return NewUserClient(u.config).QueryRoles(u)
}

// QueryAccounts queries the "accounts" edge of the User entity.
func (u *User) QueryAccounts() *AccountQuery {
	return NewUserClient(u.config).QueryAccounts(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("creator=")
	builder.WriteString(u.Creator)
	builder.WriteString(", ")
	builder.WriteString("editor=")
	builder.WriteString(u.Editor)
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", u.Deleted))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(u.City)
	builder.WriteString(", ")
	builder.WriteString("introduction=")
	builder.WriteString(u.Introduction)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
