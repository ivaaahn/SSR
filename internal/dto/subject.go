package dto

type SubjectResp struct {
	SubjectID  int    `json:"subjectID"`
	Name       string `json:"name"`
	Department string `json:"department"`
}
