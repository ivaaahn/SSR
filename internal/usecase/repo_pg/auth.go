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
		BasePgRepo: NewPGRepo(pg, l),
	}
}

func (r *AuthPgRepo) Get(email string) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("AuthPgRepo - GetStudentProfile - r.Conn.GetStudentProfile: %w", err)
	}

	return &auth, nil
}
