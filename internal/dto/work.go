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

type StWork struct {
	Work      *WorkResp `json:"work"`
	IsStarted bool      `json:"is_started"`
}

type StWorkPlenty struct {
	Works []*StWork `json:"works"`
}

type SvWork struct {
	Work *WorkSv `json:"work"`
	Head bool    `json:"head"`
}

type SvWorkPlenty struct {
	Works []*SvWork `json:"works"`
}

type WorkSv struct {
	*SvProfile
	Head bool `json:"head"`
	Full bool `json:"full"`
}

type WorkSvPlenty struct {
	Supervisors []*WorkSv `json:"supervisors"`
}
