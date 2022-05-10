package dto

type WorkDTO struct {
	WorkID      int        `json:"workID"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Semester    int8       `json:"semester"`
	Subject     SubjectDTO `json:"subject"`
}
