package dto

type SubjectResp struct {
	SubjectID  int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}
