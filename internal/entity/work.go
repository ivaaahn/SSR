package entity

type WorkKind struct {
	WorkKindID int `db:"work_kind_id"`
	Name       string
}

type Work struct {
	WorkKind    `db:"work_kind"`
	Subject     `db:"subject"`
	WorkID      int `db:"work_id"`
	Description string
	Semester    int8
}

type SupervisorViewWork struct {
	Work
	IsHead bool `db:"is_head"`
	IsFull bool `db:"is_full"`
}

type WorkSupervisor struct {
	Supervisor `db:"sv"`
	IsHead     bool `db:"is_head"`
	IsFull     bool `db:"is_full"`
}
