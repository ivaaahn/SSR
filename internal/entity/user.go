package entity

import "database/sql"

type UserRole string

const (
	student    UserRole = "st"
	supervisor          = "sv"
)

type User struct {
	UserID    int            `db:"user_id"`
	Email     string         `db:"email"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	PhotoUrl  sql.NullString `db:"photo_url"`
	Role      UserRole       `db:"role"`
}
