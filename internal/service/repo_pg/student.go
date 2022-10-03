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

func NewStudent(pg *postgres.Postgres, l logger.Interface) *Student {
	return &Student{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Student) GetStudentShort(userID int) (*entity.StudentShort, error) {
	const query = `
	select 
	     s.year, 
	     s.department_id,
		 u.last_name as "user.last_name",
		 u.first_name as "user.first_name",
		 u.user_id as "user.user_id"
	from students s
		join users u on s.user_id = u.user_id
	where u.user_id = $1
	`

	student := entity.StudentShort{}

	err := repo.Conn.Get(&student, query, userID)
	if err != nil {
		err := fmt.Errorf("Student->Get->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &student, nil
}

func (repo *Student) GetStudent(userID int) (*entity.Student, error) {
	const query = `
	select 
	    s.student_card, 
	    s.department_id, 
	    s.year,
	    u.email as "user.email", 
	    u.first_name as "user.first_name", 
	    u.last_name as "user.last_name", 
	    u.photo_url as "user.photo_url", 
	    u.user_id as "user.user_id"
	from users u 
		join students s using (user_id)
	where user_id = $1
	`

	studentFull := entity.Student{}

	err := repo.Conn.Get(&studentFull, query, userID)
	if err != nil {
		err := fmt.Errorf("Student->GetFull->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &studentFull, nil
}
