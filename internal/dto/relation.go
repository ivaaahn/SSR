package dto

import "time"

type CreateSSR struct {
	StudentID int `json:"studentID"`
	BidID     int `json:"bidID"`
}

type StViewRelation struct {
	RelID      int       `json:"id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	Supervisor SvProfile `json:"supervisor"`
	Work       Work      `json:"work"`
}
