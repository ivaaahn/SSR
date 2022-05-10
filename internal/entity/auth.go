package entity

type UserRole string

const (
	student    UserRole = "student"
	supervisor          = "supervisor"
)

type Auth struct {
	Email    string   `db:"email"`
	Password string   `db:"password"`
	Role     UserRole `db:"role"`
}
