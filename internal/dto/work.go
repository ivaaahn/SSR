package dto

type Work struct {
	WorkID      int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Semester    int8        `json:"semester"`
	Subject     SubjectResp `json:"subject"`
}

type StWork struct {
	WorkID      int    `json:"id"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
	IsStarted   bool   `json:"is_started"`
}

type StWorks struct {
	StudentID int       `json:"studentID"`
	Works     []*StWork `json:"works"`
}

type SvWork struct {
	WorkID      int    `json:"id"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
	Head        bool   `json:"head"`
}

type SvWorkPlenty struct {
	SupervisorID int       `json:"supervisorID"`
	Works        []*SvWork `json:"works"`
}

type WorkSv struct {
	SvProfile
	Head bool `json:"head"`
	Full bool `json:"full"`
}

type WorkSvPlenty struct {
	WorkID      int       `json:"workID"`
	Supervisors []*WorkSv `json:"supervisors"`
}
