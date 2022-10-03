package dto

type SupervisorViewWorkResp struct {
	Work   WorkResp `json:"work"`
	IsHead bool     `json:"is_head"`
	IsFull bool     `json:"is_full"`
}

type SupervisorViewWorkShortResp struct {
	Work   WorkShortResp `json:"work"`
	IsHead bool          `json:"is_head"`
	IsFull bool          `json:"is_full"`
}

type SupervisorViewWorkPlenty struct {
	Works []*SupervisorViewWorkShortResp `json:"works"`
}
