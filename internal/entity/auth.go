package entity

type Auth struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
