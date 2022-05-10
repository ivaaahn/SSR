package repo_pg

import (
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type BasePGRepo struct {
	*postgres.Postgres
	l logger.Interface
}

func NewPGRepo(pg *postgres.Postgres, l logger.Interface) *BasePGRepo {
	return &BasePGRepo{
		Postgres: pg,
		l:        l,
	}
}
