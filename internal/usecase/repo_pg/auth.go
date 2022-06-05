package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type AuthPgRepo struct {
	*BasePgRepo
}

func NewAuthPgRepo(pg *postgres.Postgres, l logger.Interface) *AuthPgRepo {
	return &AuthPgRepo{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *AuthPgRepo) GetUserInfo(email string) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where email = $1", email)
	if err != nil {
		err := fmt.Errorf("AuthPgRepo->r.Conn.Get(): %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &auth, nil
}
