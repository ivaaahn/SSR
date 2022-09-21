package dto

import (
	"time"
)

type StudentBid struct {
	BidID      int               `json:"id"`
	Status     string            `json:"status"`
	CreatedAt  time.Time         `json:"createdAt"`
	Supervisor SupervisorProfile `json:"supervisor"`
	Work       Work              `json:"work"`
}

type StudentBids struct {
	Bids []*StudentBid `json:"bids"`
}

type SupervisorBid struct {
	BidID     int            `json:"id"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	Student   StudentProfile `json:"student"`
	Work      Work           `json:"work"`
}

type SupervisorBids struct {
	Bids []*SupervisorBid `json:"bids"`
}

type ApplyBid struct {
	StudentID    int `json:"studentID"`
	SupervisorID int `json:"supervisorID"`
	WorkID       int `json:"workID"`
}

type ApplyBidResponse struct {
	BidID int `json:"bidID"`
}

type ResolveBid struct {
	SupervisorID int  `json:"supervisorID"`
	BidID        int  `json:"bidID"`
	Accept       bool `json:"accept"`
}

type ResolveBidResp struct {
	NewStatus string `json:"new_status"`
}
