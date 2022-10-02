package entity

import "time"

type Supervisor struct {
	UserID       int `db:"user_id"`
	Birthdate    string
	About        int
	DepartmentID string `db:"department_id"`
}

type SupervisorFull struct {
	User         *User `db:"user"`
	Birthdate    time.Time
	About        string
	DepartmentID string `db:"department_id"`
}
