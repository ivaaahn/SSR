package entity

type Subject struct {
	SubjectID    int    `db:"subject_id"`
	SubjectName  string `db:"subject_name"`
	DepartmentID string `db:"subject_department_id"`
}
