package entity

type WorkKind struct {
	WorkKindID int    `db:"work_kind_id"`
	Name       string `db:"work_kind_name"`
}

type Work struct {
	*WorkKind
	WorkID      int `db:"work_id"`
	Description string
	Semester    int8
	Subject     Subject
}
