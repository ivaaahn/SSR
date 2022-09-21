package entity

import "time"

type Feedback struct {
	FeedbackID      int       `db:"feedback_id"`
	WorkID          int       `db:"work_id"`
	WorkKind        string    `db:"work_kind"`
	WorkSubject     string    `db:"work_subject"`
	StudentID       int       `db:"student_id"`
	StudentFullName string    `db:"student_full_name"`
	SupervisorID    int       `db:"supervisor_id"`
	CreatedAt       time.Time `db:"created_at"`
	Content         string    `db:"content"`
}
