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

type RelationCreateReq struct {
	StudentID    int `json:"student_id"`
	SupervisorID int `json:"supervisor_id"`
	WorkID       int `json:"work_id"`
}

type RelationCreateResp struct {
	RelationID int `json:"relation_id"`
}

type ResolveBid struct {
	SupervisorID int  `json:"supervisorID"`
	BidID        int  `json:"bidID"`
	Accept       bool `json:"accept"`
}

type ResolveBidResp struct {
	NewStatus string `json:"new_status"`
}
