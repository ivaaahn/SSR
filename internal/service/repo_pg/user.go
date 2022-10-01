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

func New(pg *postgres.Postgres, l logger.Interface) *User {
	return &User{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *Auth) Create(email, firstName, LastName, photoUrl string, role entity.UserRole) error {
	query := `
	insert into users (email, first_name, last_name, photo_url, role) 
	values ($1, $2, $3, $4, $5)
	`

	res, err := r.Conn.Exec(query, email, password, userID)

	fmt.Println(res) //TODO

	if err != nil {
		err := fmt.Errorf("Auth->Create->repo.Conn.Exec: %w", err)
		r.l.Error(err) // TODO
		return err
	}

	return nil
}
