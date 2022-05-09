package repo

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type AuthPGRepo struct {
	*postgres.Postgres
	l logger.Interface
}

func NewAuthPGRepo(pg *postgres.Postgres, l logger.Interface) *AuthPGRepo {
	return &AuthPGRepo{
		Postgres: pg,
		l:        l,
	}
}

func (r *AuthPGRepo) Get(email string) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("AuthPGRepo - Get - r.Conn.Get: %w", err)
	}

	return &auth, nil
}
