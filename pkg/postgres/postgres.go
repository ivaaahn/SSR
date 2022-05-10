package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Conn *sqlx.DB
}

func New(dsn string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{Conn: db}, nil
}

func (pg *Postgres) Close() error {
	return pg.Conn.Close()
}

func (pg *Postgres) Check() error {
	return pg.Conn.Ping()
}
