package dto

import (
	"time"
)

type StudentBidDTO struct {
	BidID      int                  `json:"BidID"`
	Status     string               `json:"status"`
	CreatedAt  time.Time            `json:"createdAt"`
	Supervisor SupervisorProfileDTO `json:"supervisor"`
	Work       WorkDTO              `json:"work"`
}

type SupervisorBidDTO struct {
	BidID     int               `json:"BidID"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	Student   StudentProfileDTO `json:"student"`
	Work      WorkDTO           `json:"work"`
}

type StudentBidsDTO struct {
	Bids []*StudentBidDTO
}

type SupervisorBidsDTO struct {
	Bids []*SupervisorBidDTO
}

type StudentApplyBidDTO struct {
	StudentID    int `json:"studentID"`
	SupervisorID int `json:"supervisorID"`
	WorkID       int `json:"workID"`
}

type StudentApplyBidResponseDTO struct {
	BidID int `json:"bidID"`
}
