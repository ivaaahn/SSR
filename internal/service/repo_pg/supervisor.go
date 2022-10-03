package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Supervisor struct {
	*BasePgRepo
}

func NewSupervisor(pg *postgres.Postgres, l logger.Interface) *Supervisor {
	return &Supervisor{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Supervisor) GetSupervisor(userID int) (*entity.SupervisorShort, error) {
	const query = `
	select * from supervisors where user_id = $1
	`

	supervisor := entity.SupervisorShort{}

	err := repo.Conn.Get(&supervisor, query, userID)
	if err != nil {
		err := fmt.Errorf("SupervisorShort->Get->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &supervisor, nil
}

func (repo *Supervisor) GetFullSupervisor(userID int) (*entity.Supervisor, error) {
	const query = `
	select 
	    s.about, 
	    s.department_id, 
	    s.birthdate,
	    u.email as "user.email", 
	    u.first_name as "user.first_name", 
	    u.last_name as "user.last_name", 
	    u.photo_url as "user.photo_url", 
	    u.user_id as "user.user_id"
	from users u 
		join supervisors s using (user_id)
	where user_id = $1
	`

	supervisorFull := entity.Supervisor{}

	err := repo.Conn.Get(&supervisorFull, query, userID)
	if err != nil {
		err := fmt.Errorf("SupervisorShort->GetFull->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &supervisorFull, nil
}
