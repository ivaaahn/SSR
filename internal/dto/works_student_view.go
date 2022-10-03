package dto

type StudentViewWorkResp struct {
	Work      WorkResp `json:"work"`
	IsStarted bool     `json:"is_started"` // TODO
}

type StudentViewWorkShortResp struct {
	Work WorkShortResp `json:"work"`
}

type StudentViewWorkPlenty struct {
	Works []*StudentViewWorkShortResp `json:"works"`
}
