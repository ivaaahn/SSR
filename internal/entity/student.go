package entity

type Student struct {
	UserID       int    `db:"user_id"`
	StudentCard  string `db:"student_card"`
	Year         int
	DepartmentID string `db:"department_id"`
}

type StudentFull struct {
	User         *User
	StudentCard  string `db:"student_card"`
	Year         int
	DepartmentID string `db:"department_id"`
}
