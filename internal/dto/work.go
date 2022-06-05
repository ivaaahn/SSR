package dto

type Work struct {
	WorkID      int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Semester    int8        `json:"semester"`
	Subject     SubjectResp `json:"subject"`
}

type StudentWork struct {
	WorkID      int    `json:"id"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
	IsStarted   bool   `json:"is_started"`
}

type StudentWorks struct {
	StudentID int            `json:"studentID"`
	Works     []*StudentWork `json:"works"`
}

type SupervisorWork struct {
	WorkID      int    `json:"id"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Subject     string `json:"subject"`
	Head        bool   `json:"head"`
}

type SupervisorWorkPlenty struct {
	SupervisorID int               `json:"supervisorID"`
	Works        []*SupervisorWork `json:"works"`
}

type WorkSupervisor struct {
	SupervisorProfile
	Head bool `json:"head"`
	Full bool `json:"full"`
}

type WorkSupervisorPlenty struct {
	WorkID      int               `json:"workID"`
	Supervisors []*WorkSupervisor `json:"supervisors"`
}
