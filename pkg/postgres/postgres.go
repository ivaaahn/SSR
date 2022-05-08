package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func New(dsn string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// TODO: closeResource
	return &Postgres{db: db}, nil
}

func (pg *Postgres) Close() error {
	return pg.db.Close()
}

func (pg *Postgres) Check() error {
	return pg.db.Ping()
}
