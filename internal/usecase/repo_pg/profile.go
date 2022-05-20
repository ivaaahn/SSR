package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type ProfilePgRepo struct {
	*BasePgRepo
}

func NewProfilePgRepo(pg *postgres.Postgres, l logger.Interface) *ProfilePgRepo {
	return &ProfilePgRepo{
		BasePgRepo: NewPGRepo(pg, l),
	}
}

func (r *ProfilePgRepo) GetStudentProfile(email string) (*entity.StudentProfile, error) {
	const query = `
select *
from users u 
    join students s using (user_id)
where email = $1
`
	student := entity.StudentProfile{}

	err := r.Conn.Get(&student, query, email)
	if err != nil {
		return nil, fmt.Errorf("ProfilePgRepo - GetStudentProfile - r.Conn.GetStudentProfile: %w", err)
	}

	return &student, nil
}

func (r *ProfilePgRepo) GetSupervisorProfile(email string) (*entity.SupervisorProfile, error) {
	const query = `
select *
from users u 
    join supervisors s using (user_id)
where email = $1
`
	supervisor := entity.SupervisorProfile{}

	err := r.Conn.Get(&supervisor, query, email)
	if err != nil {
		return nil, fmt.Errorf("ProfilePgRepo - GetSupervisorProfile - r.Conn.GetSupervisorProfile: %w", err)
	}

	return &supervisor, nil
}
