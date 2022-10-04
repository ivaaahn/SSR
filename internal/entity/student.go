package entity

type StudentShort struct {
	User         UserShort `db:"user"`
	Year         int
	DepartmentID string `db:"department_id"`
}

type Student struct {
	User         User
	StudentCard  string `db:"student_card"`
	Year         int
	DepartmentID string `db:"department_id"`
}
