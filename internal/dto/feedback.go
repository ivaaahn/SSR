package dto

type FeedbackReq struct {
	StudentID    int    `json:"studentID"`
	SupervisorID int    `json:"supervisorID"`
	WorkID       int    `json:"workID"`
	Content      string `json:"content"`
}

type FeedbackResp struct {
	StudentID       int    `json:"studentID"`
	StudentFullName string `json:"student_full_name"`
	SupervisorID    int    `json:"supervisorID"`
	WorkID          int    `json:"workID"`
	WorkKind        string `json:"work_kind"`
	WorkSubject     string `json:"work_subject"`
	Content         string `json:"content"`
}

type FeedbackAddResp struct {
	FeedbackID int `json:"feedback_id"`
}

type FeedbackPlenty struct {
	Feedbacks []*FeedbackResp `json:"feedbacks"`
}
