// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_name", Type: field.TypeString, Unique: true, Size: 256},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "password_digest", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "nick_name", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "status", Type: field.TypeString, Nullable: true, Size: 256},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "money", Type: field.TypeString, Nullable: true, Size: 256},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
