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

func (repo *Supervisor) GetSupervisorShort(userID int) (*entity.SupervisorShort, error) {
	const query = `
	select * from supervisors where user_id = $1
	`

	supervisorShort := entity.SupervisorShort{}

	err := repo.Conn.Get(&supervisorShort, query, userID)
	if err != nil {
		err := fmt.Errorf("SupervisorShort->Get->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &supervisorShort, nil
}

func (repo *Supervisor) GetSupervisor(userID int) (*entity.Supervisor, error) {
	const query = `
	select 
	    s.about, 
	    s.department_id, 
	    s.birthdate,
	    u.email as "user.email", 
	    u.first_name as "user.first_name", 
	    u.last_name as "user.last_name", 
	    u.photo_url as "user.photo_url", 
	    u.id as "user.id"
	from users u 
		join supervisors s on s.user_id = u.id
	where u.id = $1
	`

	supervisor := entity.Supervisor{}

	err := repo.Conn.Get(&supervisor, query, userID)
	if err != nil {
		err := fmt.Errorf("SupervisorShort->GetFull->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &supervisor, nil
}

func (repo *Supervisor) GetSupervisorsByWorkID(workID int) ([]*entity.WorkSupervisor, error) {
	const query = `
	select 
		sv.about as "sv.about",
		sv.birthdate as "sv.birthdate",   	        
		sv.department_id as "sv.department_id",

	    u.id as "sv.user.id",
		u.email as "sv.user.email",
		u.first_name as "sv.user.first_name",
		u.last_name as "sv.user.last_name",
		u.photo_url as "sv.user.photo_url",
		sw.is_full as is_full, 
		sw.is_head as is_head
	from supervisors sv
		join supervisor_work sw on sv.user_id = sw.supervisor_id
		join users u on sv.user_id = u.id
	where sw.work_id = $1;
	`

	var supervisors []*entity.WorkSupervisor

	err := repo.Conn.Select(&supervisors, query, workID)
	if err != nil {
		err := fmt.Errorf("Work->GetSupervisorsByWorkID->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return supervisors, nil
}
