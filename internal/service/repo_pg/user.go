package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type User struct {
	*BasePgRepo
}

func NewUser(pg *postgres.Postgres, l logger.Interface) *User {
	return &User{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *User) CreateUser(email, password, firstName, lastName, photoUrl string, role entity.UserRole) error {
	query := `
	insert into users (email, password, first_name, last_name, photo_url, role) 
	values ($1, $2, $3, $4, $5)
	`

	res, err := r.Conn.Exec(query, email, password, firstName, lastName, photoUrl, role)

	fmt.Println(res) //TODO

	if err != nil {
		err := fmt.Errorf("Auth->Create->repo.Conn.Exec: %w", err)
		r.l.Error(err) // TODO
		return err
	}

	return nil
}

func (r *User) GetUserByEmail(email string) (*entity.User, error) {
	user := entity.User{}

	err := r.Conn.Get(&user, "select * from users where email = $1", email)
	if err != nil {
		err := fmt.Errorf("User->r.Conn.Get(): %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &user, nil
}

func (r *User) GetUserByID(userID int) (*entity.User, error) {
	auth := entity.User{}

	err := r.Conn.Get(&auth, "select * from user where user_id = $1", userID)
	if err != nil {
		err := fmt.Errorf("User->r.Conn.Get(): %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &auth, nil
}
