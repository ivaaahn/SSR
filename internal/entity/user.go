package entity

import "database/sql"

type User struct {
	UserID    int            `db:"user_id"`
	Email     string         `db:"email"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Avatar    sql.NullString `db:"avatar_url"`
}
