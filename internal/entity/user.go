package entity

type UserRole string

const (
	student    UserRole = "student"
	supervisor          = "supervisor"
)

type User struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Avatar    string
	Role      UserRole
}

type Auth struct {
	Email    string
	Password string
}
