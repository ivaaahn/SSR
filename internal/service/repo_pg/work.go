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

func (repo *Work) GetStudentWorks(departmentID string, semester int) ([]*entity.Work, error) {
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

	err := repo.Conn.Select(&works, query, departmentID, semester)
	if err != nil {
		err := fmt.Errorf("Work->GetStudentWorks->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (repo *Work) GetSupervisorWorks(supervisorID int) ([]*entity.SupervisorViewWork, error) {
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
	var works []*entity.SupervisorViewWork

	err := repo.Conn.Select(&works, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("Work->GetSupervisorWorks->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return works, nil
}

func (repo *Work) Get(workID int) (*entity.Work, error) {
	const query = `
	select 
	    w.work_id, w.description, w.semester,
		wk.work_kind_id as "work_kind.work_kind_id",
		wk.name as "work_kind.name",
		s.subject_id as "subject.subject_id",
		s.name as "subject.name",
		s.department_id as "subject.department_id"
	from works w
		join work_kinds  wk using ("work_kind_id")
		join subjects s using ("subject_id")
	where work_id = $1
	`

	work := entity.Work{}

	err := repo.Conn.Get(&work, query, workID)
	if err != nil {
		err := fmt.Errorf("Work->Get->repo.Conn.Get(): %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &work, nil
}
