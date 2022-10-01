package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Relation struct {
	*BasePgRepo
}

func NewRelation(pg *postgres.Postgres, l logger.Interface) *Relation {
	return &Relation{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (repo *Relation) GetStudentRelation(studentID, ssrID int) (*entity.StRelation, error) {
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
	ssr := entity.StRelation{}

	err := repo.Conn.Get(&ssr, query, ssrID, studentID)
	if err != nil {
		err := fmt.Errorf("Relation->GetStudentRelation->repo.Conn.Get: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return &ssr, nil
}

func (repo *Relation) GetStudentBids(studentID int) ([]*entity.StRelation, error) {
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

	var bids []*entity.StRelation

	err := repo.Conn.Select(&bids, query, studentID)
	if err != nil {
		err := fmt.Errorf("Relation->GetStudentBids->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (repo *Relation) GetStudentRelations(studentID int) ([]*entity.StRelation, error) {
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

	var bids []*entity.StRelation

	err := repo.Conn.Select(&bids, query, studentID)
	if err != nil {
		err := fmt.Errorf("Relation->GetStudentBids->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (repo *Relation) GetSupervisorBids(supervisorID int) ([]*entity.SvRelation, error) {
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

	var bids []*entity.SvRelation

	err := repo.Conn.Select(&bids, query, supervisorID)
	if err != nil {
		err := fmt.Errorf("Relation->GetSupervisorBids->repo.Conn.Select: %w", err)
		repo.l.Error(err)
		return nil, err
	}

	return bids, nil
}

func (repo *Relation) Create(studentID, supervisorID, workID int) (int, error) {
	query := `
	insert into ssr (student_id, supervisor_id, work_id) 
	values ($1, $2, $3)
	returning ssr_id;
	`

	var bidID int
	err := repo.Conn.QueryRowx(query, studentID, supervisorID, workID).Scan(&bidID)
	if err != nil {
		err := fmt.Errorf("Relation->Create->repo.Conn.QueryRowx: %w", err)
		repo.l.Error(err)
		return 0, err
	}

	return bidID, nil
}

func (repo *Relation) UpdateStatus(id int, newStatus entity.StatusSSR) (int, error) {
	query := `
	update ssr set status = $1
	where ssr_id = $2
	returning ssr_id;
	`

	var bidID int
	err := repo.Conn.QueryRowx(query, newStatus, id).Scan(&bidID)
	if err != nil {
		err := fmt.Errorf("Relation->UpdateStatus->repo.Conn.QueryRowx: %w", err)
		repo.l.Error(err)
		return 0, err
	}

	return bidID, nil
}