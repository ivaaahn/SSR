package entity

type WorkKind struct {
	WorkKindID   int    `db:"work_kind_id"`
	WorkKindName string `db:"work_kind_name"`
}

type Work struct {
	*WorkKind
	*Subject
	WorkID      int `db:"work_id"`
	Description string
	Semester    int8
}

type SvWork struct {
	*Work
	Head bool `db:"head"`
}

type WorkSv struct {
	*SvProfile
	Head bool `db:"head"`
	Full bool `db:"full"`
}
