package repo

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type UserPGRepo struct {
	*postgres.Postgres
	l logger.Interface
}

func NewUserPGRepo(pg *postgres.Postgres, l logger.Interface) *UserPGRepo {
	return &UserPGRepo{
		Postgres: pg,
		l:        l,
	}
}

func (r *UserPGRepo) Get(email string) (*entity.User, error) {
	user := entity.User{}

	err := r.Conn.Get(&user, "select * from users where email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("UserPGRepo - Get - r.Conn.Get: %w", err)
	}

	return &user, nil
}
