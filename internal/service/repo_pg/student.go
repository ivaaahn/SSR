package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Student struct {
	*BasePgRepo
}

func New(pg *postgres.Postgres, l logger.Interface) *Student {
	return &Student{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Student) Get(userID int) (*entity.Student, error) {
	const query = `
	select * from users where user_id = $1
	`

	student := entity.Student{}

	err := repo.Conn.Get(&student, query, userID)
	if err != nil {
		err := fmt.Errorf("Student->Get->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &student, nil
}

func (repo *Student) GetFull(userID int) (*entity.StudentFull, error) {
	const query = `
	select *
	from users u 
		join students s using (user_id)
	where user_id = $1
	`

	studentFull := entity.StudentFull{}

	err := repo.Conn.Get(&studentFull, query, userID)
	if err != nil {
		err := fmt.Errorf("Student->GetFull->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &studentFull, nil
}