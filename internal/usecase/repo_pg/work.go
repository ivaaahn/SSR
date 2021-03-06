package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type WorkPgRepo struct {
	*BasePgRepo
}

func NewWorkPgRepo(pg *postgres.Postgres, l logger.Interface) *WorkPgRepo {
	return &WorkPgRepo{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *WorkPgRepo) GetWorksByStudentID(studentID int) ([]*entity.Work, error) {
	const query = `
	with const (st_year, st_department_id, curr_month) as (
		select s.year,
			   s.department_id,
			   extract('month' from current_date)
		from students s
		where student_id = $1
	)
	select w.*,
	   	wk.name as work_kind_name,
	   	subj.name as subject_name
	from works w
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
		join const c on true
	where (((curr_month between 2 and 8) and (semester = st_year * 2))
		or (semester = st_year * 2 - 1))
	  and subj.department_id = c.st_department_id;
	`

	var works []*entity.Work

	err := r.Conn.Select(&works, query, studentID)
	if err != nil {
		err := fmt.Errorf("WorkPgRepo->GetWorksByStudentID->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (r *WorkPgRepo) GetWorksBySupervisorID(supervisorID int) ([]*entity.WorkOfSupervisor, error) {
	const query = `
	select w.*, 
	   	subj.name as subject_name, 
	   	wk.name as work_kind_name, 
	   	sw.is_head as head
	from works w
		join supervisor_work sw using (work_id)
		join supervisors s using (supervisor_id)
		join subjects subj using (subject_id)
		join work_kinds wk using (work_kind_id)
	where s.supervisor_id = $1;
	`
	var works []*entity.WorkOfSupervisor

	err := r.Conn.Select(&works, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("WorkPgRepo->GetWorksBySupervisorID->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (r *WorkPgRepo) GetSupervisorsByWorkID(workID int) ([]*entity.SupervisorOfWork, error) {
	const query = `
	select 
	       u.*, 
	       sv.*, 
	       sw.is_full as "full", 
	       sw.is_head as head
	from supervisors sv
		join supervisor_work sw using (supervisor_id)
		join users u using (user_id)
	where sw.work_id = $1;
	`

	var supervisors []*entity.SupervisorOfWork

	err := r.Conn.Select(&supervisors, query, workID)
	if err != nil {
		err := fmt.Errorf("WorkPgRepo->GetSupervisorsByWorkID->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return supervisors, nil
}
