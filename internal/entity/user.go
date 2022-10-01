package entity

type UserRole string

const (
	student    UserRole = "st"
	supervisor          = "sv"
)

type User struct {
	UserID    int      `db:"user_id"`
	Email     string   `db:"email"`
	FirstName string   `db:"first_name"`
	LastName  string   `db:"last_name"`
	PhotoUrl  string   `db:"photo_url"`
	Role      UserRole `db:"role"`
	Password  string   `db:"password"`
}
