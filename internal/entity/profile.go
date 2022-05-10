package entity

import "time"

type StudentProfile struct {
	*User
	StudentID    int    `db:"student_id"`
	StudentCard  string `db:"student_card"`
	Year         int
	DepartmentID string `db:"department_id"`
}

type SupervisorProfile struct {
	*User
	SupervisorID int       `db:"supervisor_id"`
	Birthdate    time.Time `db:"birthdate"`
	About        string    `db:"about"`
	DepartmentID string    `db:"department_id"`
}
