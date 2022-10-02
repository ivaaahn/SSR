package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Work struct {
	*BasePgRepo
}

func NewWork(pg *postgres.Postgres, l logger.Interface) *Work {
	return &Work{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *Work) GetStudentWorks(departmentID string, semester int) ([]*entity.Work, error) {
	const query = `
	select w.work_id, w.description, w.semester,
	   	wk.name as "work_kind.name",
	   	wk.work_kind_id as "work_kind.work_kind_id",
	   	subj.subject_id as "subject.subject_id",
	   	subj.name as "subject.name",
		subj.department_id as "subject.department_id"
	from works w
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where subj.department_id = $1 and w.semester = $2;
	`

	var works []*entity.Work

	err := r.Conn.Select(&works, query, departmentID, semester)
	if err != nil {
		err := fmt.Errorf("Work->GetStudentWorks->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (r *Work) GetSupervisorWorks(supervisorID int) ([]*entity.SvWork, error) {
	const query = `
	select w.work_id, w.description, w.semester, 
	   	wk.name as "work_kind.name",
		wk.work_kind_id as "work_kind.work_kind_id",
	   	subj.subject_id as "subject.subject_id",
	   	subj.name as "subject.name",
		subj.department_id as "subject.department_id",
	   	sw.is_head,
	   	sw.is_full 
	from works w
		join supervisor_work sw using (work_id)
		join subjects subj using (subject_id)
		join work_kinds wk using (work_kind_id)
	where sw.supervisor_id = $1;
	`
	var works []*entity.SvWork

	err := r.Conn.Select(&works, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("Work->GetSupervisorWorks->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (r *Work) GetSupervisorsByWorkID(workID int) ([]*entity.WorkSv, error) {
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

	var supervisors []*entity.WorkSv

	err := r.Conn.Select(&supervisors, query, workID)
	if err != nil {
		err := fmt.Errorf("Work->GetSupervisorsByWorkID->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return supervisors, nil
}
