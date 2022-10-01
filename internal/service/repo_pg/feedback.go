package repo_pg

import (
	"fmt"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

type Feedback struct {
	*BasePgRepo
}

func NewFeedback(pg *postgres.Postgres, l logger.Interface) *Feedback {
	return &Feedback{
		BasePgRepo: NewPgRepo(pg, l),
	}
}

func (f *Feedback) Create(studentID, supervisorID, workID int, content string) (int, error) {
	query := `
	insert into feedbacks (student_id, supervisor_id, work_id, content)
	values ($1, $2, $3, $4)
	returning feedback_id;
	`

	var feedbackID int
	err := f.Conn.QueryRowx(query, studentID, supervisorID, workID, content).Scan(&feedbackID)
	if err != nil {
		err := fmt.Errorf("repo_pg.Feedback->Create->r.Conn.QueryRowx: %w", err)
		f.l.Error(err)
		return 0, err
	}

	return feedbackID, nil
}

func (f *Feedback) GetBySupervisorID(supervisorID int) ([]*entity.Feedback, error) {
	query := `
    select 
        f.*, 
        concat (u.first_name, ' ', u.last_name) AS student_full_name, 
        wk.name as work_kind,
        subj.name as work_subject
    from feedbacks f
        join students s using (student_id)
        join users u using (user_id)
        join works w using (work_id)
        join work_kinds wk using (work_kind_id)
        join subjects subj using (subject_id)
    where f.supervisor_id = $1
    `

	var feedbacks []*entity.Feedback

	if err := f.Conn.Select(&feedbacks, query, supervisorID); err != nil {
		err := fmt.Errorf("repo_pg.Feedback->GetBySupervisorID->r.Conn.Select: %w", err)
		f.l.Error(err)
		return nil, err
	}

	return feedbacks, nil
}
