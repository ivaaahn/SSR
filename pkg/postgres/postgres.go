package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Postgres struct {
	Conn *sqlx.DB
}

func New(dsn string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5 && (err != nil || i == 0); i++ {
		err = db.Ping()
		time.Sleep(5 * time.Second)
	}

	if err != nil {
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
