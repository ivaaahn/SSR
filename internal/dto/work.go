package dto

type WorkKindResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WorkResp struct {
	WorkID      int          `json:"id"`
	Description string       `json:"description"`
	Semester    int8         `json:"semester"`
	Subject     SubjectResp  `json:"subject"`
	Kind        WorkKindResp `json:"kind"`
}

type StWorkResp struct {
	Work      WorkResp `json:"work"`
	IsStarted bool     `json:"is_started"`
}

type StWorkPlenty struct {
	Works []*StWorkResp `json:"works"`
}

type SvWorkResp struct {
	Work   WorkResp `json:"work"`
	IsHead bool     `json:"is_head"`
	IsFull bool     `json:"is_full"`
}

type SvWorkPlenty struct {
	Works []*SvWorkResp `json:"works"`
}

type WorkSv struct {
	SvProfile
	Head bool `json:"head"`
	Full bool `json:"full"`
}

type WorkSvPlenty struct {
	Supervisors []*WorkSv `json:"supervisors"`
}
