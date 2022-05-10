package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type AuthPGRepo struct {
	*BasePGRepo
}

func NewAuthPGRepo(pg *postgres.Postgres, l logger.Interface) *AuthPGRepo {
	return &AuthPGRepo{
		BasePGRepo: NewPGRepo(pg, l),
	}
}

func (r *AuthPGRepo) Get(email string) (*entity.Auth, error) {
	auth := entity.Auth{}

	err := r.Conn.Get(&auth, "select * from auth where email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("AuthPGRepo - GetStudentProfile - r.Conn.GetStudentProfile: %w", err)
	}

	return &auth, nil
}
