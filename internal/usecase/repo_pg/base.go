package repo_pg

import (
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type BasePgRepo struct {
	*postgres.Postgres
	l logger.Interface
}

func NewPgRepo(pg *postgres.Postgres, l logger.Interface) *BasePgRepo {
	return &BasePgRepo{
		Postgres: pg,
		l:        l,
	}
}
