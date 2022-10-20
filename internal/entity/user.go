package entity

import "time"

type UserRole string

const (
	student    UserRole = "st"
	supervisor          = "sv"
)

type UserFull struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	PhotoUrl  string    `db:"photo_url"`
	Role      UserRole  `db:"role"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	LastLogin time.Time `db:"last_login"`
	CreatedAt time.Time `db:"created_at"`
}

type User struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	PhotoUrl  string `db:"photo_url"`
}

type UserShort struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
