package dto

import (
	"time"
)

type StRelationResp struct {
	BidID      int       `json:"id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	Supervisor SvProfile `json:"supervisor"`
	Work       WorkResp  `json:"work"`
}

type StRelationPlenty struct {
	Relations []*StRelationResp `json:"relations"`
}

type SvBid struct {
	BidID     int       `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	Student   StProfile `json:"student"`
	Work      WorkResp  `json:"work"`
}

type SvBids struct {
	Bids []*SvBid `json:"bids"`
}

type ApplyBid struct {
	StudentID    int `json:"studentID"`
	SupervisorID int `json:"supervisorID"`
	WorkID       int `json:"workID"`
}

type ApplyBidResp struct {
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
