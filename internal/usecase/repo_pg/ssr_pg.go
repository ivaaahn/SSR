package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type SSRPGRepo struct {
	*BasePGRepo
}

func NewSSRPGRepo(pg *postgres.Postgres, l logger.Interface) *SSRPGRepo {
	return &SSRPGRepo{
		BasePGRepo: NewPGRepo(pg, l),
	}
}

func (r *SSRPGRepo) GetBidsByStudentID(studentID int) ([]*entity.StudentBid, error) {
	query := `
select 
    b.ssr_id,
    b.status,
    b.created_at,
	sv.*,
   	u.*,
    w.*,
	wk.name as work_kind_name,
    subj.name as subject_name,
    subj.department_id as subject_department_id
from ssr as b 
	join supervisors sv using (supervisor_id)
    join users u using (user_id)
	join works w using (work_id)
	join work_kinds wk using (work_kind_id)
	join subjects subj using (subject_id)
where b.status in (
       'заявка ожидает ответа', 
       'заявка отклонена', 
       'заявка отозвана', 
       'заявка принята'
	) and b.student_id = $1
`

	var bids []*entity.StudentBid

	err := r.Conn.Select(&bids, query, studentID)
	if err != nil {
		return nil, fmt.Errorf("SSRPGRepo - GetBidsByStID - r.Conn.Select: %w", err)
	}

	return bids, nil
}

func (r *SSRPGRepo) GetBidsBySupervisorID(supervisorID int) ([]*entity.SupervisorBid, error) {
	query := `
select 
    b.ssr_id,
    b.status,
    b.created_at,
	st.*,
   	u.*,
    w.*,
	wk.name as work_kind_name,
    subj.name as subject_name,
    subj.department_id as subject_department_id
from ssr as b 
	join students st using (student_id)
    join users u using (user_id)
	join works w using (work_id)
	join work_kinds wk using (work_kind_id)
	join subjects subj using (subject_id)
where b.status in (
       'заявка ожидает ответа', 
       'заявка отклонена', 
       'заявка отозвана', 
       'заявка принята'
	) and b.supervisor_id = $1
`

	var bids []*entity.SupervisorBid

	err := r.Conn.Select(&bids, query, supervisorID)
	if err != nil {
		return nil, fmt.Errorf("SSRPGRepo - GetBidsBySvID - r.Conn.Select: %w", err)
	}

	return bids, nil
}

func (r *SSRPGRepo) CreateBid(studentID, supervisorID, workID int) (int, error) {
	query := `
insert into ssr (student_id, supervisor_id, work_id) 
values ($1, $2, $3)
returning ssr_id;
`
	var bidID int
	err := r.Conn.QueryRowx(query, studentID, supervisorID, workID).Scan(&bidID)
	if err != nil {
		return 0, err
	}

	return bidID, nil
}
