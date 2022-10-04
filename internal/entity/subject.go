package entity

type Subject struct {
	SubjectID    int `db:"subject_id"`
	Name         string
	DepartmentID string `db:"department_id"`
}
