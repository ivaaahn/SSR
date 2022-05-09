package entity

import "database/sql"

type UserRole string

const (
	student    UserRole = "student"
	supervisor          = "supervisor"
)

type User struct {
	Id        int
	Email     string
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Avatar    sql.NullString `db:"avatar_url"`
	Role      UserRole
}
