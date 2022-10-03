package entity

import "time"

type SupervisorShort struct {
	User         UserShort `db:"user"`
	DepartmentID string    `db:"department_id"`
}

type Supervisor struct {
	User         User `db:"user"`
	Birthdate    time.Time
	About        string
	DepartmentID string `db:"department_id"`
}
