package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Auth struct {
	*BasePgRepo
}

func NewAuth(pg *postgres.Postgres, l logger.Interface) *Auth {
	return &Auth{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *Auth) Create(email, password string, userID int) error {
	query := `
	insert into auth (email, password, user_id) 
	values ($1, $2, $3)
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

func (r *Auth) GetUserByEmail(email string) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where email = $1", email)
	if err != nil {
		err := fmt.Errorf("Auth->r.Conn.Get(): %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &auth, nil
}

func (r *Auth) GetUserByID(userID int) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where user_id = $1", userID)
	if err != nil {
		err := fmt.Errorf("Auth->r.Conn.Get(): %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &auth, nil
}
