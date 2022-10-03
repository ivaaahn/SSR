package dto

type WorkResp struct {
	WorkID      int          `json:"id"`
	Description string       `json:"description"`
	Semester    int8         `json:"semester"`
	Subject     SubjectResp  `json:"subject"`
	Kind        WorkKindResp `json:"kind"`
}

type WorkFullResp struct {
	WorkID      int                    `json:"id"`
	Description string                 `json:"description"`
	Semester    int8                   `json:"semester"`
	Subject     SubjectResp            `json:"subject"`
	Kind        WorkKindResp           `json:"kind"`
	Waypoints   []*WaypointResp        `json:"waypoints"`
	Supervisors []*WorkSupervisorShort `json:"supervisors"`
}

type WorkShortResp struct {
	WorkID  int          `json:"id"`
	Subject SubjectResp  `json:"subject"`
	Kind    WorkKindResp `json:"kind"`
}

type WorkPlenty struct {
	Works []*WorkShortResp `json:"works"`
}
