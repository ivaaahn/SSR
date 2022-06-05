package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type SsrPgRepo struct {
	*BasePgRepo
}

func NewSSRPgRepo(pg *postgres.Postgres, l logger.Interface) *SsrPgRepo {
	return &SsrPgRepo{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (r *SsrPgRepo) GetStudentRelation(studentID, ssrID int) (*entity.StudentSsr, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status as ssr_status,
		ssr.created_at,
		sv.*,
		u.*,
		w.*,
		wk.name as work_kind_name,
		subj.name as subject_name,
		subj.department_id as subject_department_id
	from ssr 
		join supervisors sv using (supervisor_id)
		join users u using (user_id)
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.ssr_id = $1 and ssr.student_id = $2;
	`
	ssr := entity.StudentSsr{}

	err := r.Conn.Get(&ssr, query, ssrID, studentID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->GetStudentRelation->r.Conn.Get: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return &ssr, nil
}

func (r *SsrPgRepo) GetStudentBids(studentID int) ([]*entity.StudentSsr, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status as ssr_status,
		ssr.created_at,
		sv.*,
		u.*,
		w.*,
		wk.name as work_kind_name,
		subj.name as subject_name,
		subj.department_id as subject_department_id
	from ssr 
		join supervisors sv using (supervisor_id)
		join users u using (user_id)
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.status in ('pending','rejected', 'cancelled','accepted') and ssr.student_id = $1;
	`

	var bids []*entity.StudentSsr

	err := r.Conn.Select(&bids, query, studentID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->GetStudentBids->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (r *SsrPgRepo) GetStudentRelations(studentID int) ([]*entity.StudentSsr, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status as ssr_status,
		ssr.created_at,
		sv.*,
		u.*,
		w.*,
		wk.name as work_kind_name,
		subj.name as subject_name,
		subj.department_id as subject_department_id
	from ssr 
		join supervisors sv using (supervisor_id)
		join users u using (user_id)
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.status in ('wip', 'done') and ssr.student_id = $1;
	`

	var bids []*entity.StudentSsr

	err := r.Conn.Select(&bids, query, studentID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->GetStudentBids->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (r *SsrPgRepo) GetSupervisorBids(supervisorID int) ([]*entity.SupervisorSsr, error) {
	query := `
	select 
		ssr.ssr_id,
		ssr.status as ssr_status,
		ssr.created_at,
		st.*,
		u.*,
		w.*,
		wk.name as work_kind_name,
		subj.name as subject_name,
		subj.department_id as subject_department_id
	from ssr 
		join students st using (student_id)
		join users u using (user_id)
		join works w using (work_id)
		join work_kinds wk using (work_kind_id)
		join subjects subj using (subject_id)
	where ssr.status in ('pending','rejected', 'cancelled','accepted') and ssr.supervisor_id = $1;
	`

	var bids []*entity.SupervisorSsr

	err := r.Conn.Select(&bids, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->GetSupervisorBids->r.Conn.Select: %w", err)
		r.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (r *SsrPgRepo) Create(studentID, supervisorID, workID int) (int, error) {
	query := `
	insert into ssr (student_id, supervisor_id, work_id) 
	values ($1, $2, $3)
	returning ssr_id;
	`

	var bidID int
	err := r.Conn.QueryRowx(query, studentID, supervisorID, workID).Scan(&bidID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->Create->r.Conn.QueryRowx: %w", err)
		r.l.Error(err)
		return 0, err
	}

	return bidID, nil
}

func (r *SsrPgRepo) UpdateStatus(id int, newStatus entity.StatusSSR) (int, error) {
	query := `
	update ssr set status = $1
	where ssr_id = $2
	returning ssr_id;
	`

	var bidID int
	err := r.Conn.QueryRowx(query, newStatus, id).Scan(&bidID)
	if err != nil {
		err := fmt.Errorf("SsrPgRepo->UpdateStatus->r.Conn.QueryRowx: %w", err)
		r.l.Error(err)
		return 0, err
	}

	return bidID, nil
}
