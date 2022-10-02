package entity

import "time"

type StatusSSR string

const (
	bidPending   StatusSSR = "pending"
	bidDeclined            = "rejected"
	bidCancelled           = "cancelled"
	bidAccepted            = "accepted"
	inProgress             = "wip"
	completed              = "done"
)

type WaypointRelation struct {
	*Waypoint
	Status string `db:"waypoint_status"`
	SsrID  int    `db:"ssr_id"`
}

type StRelation struct {
	RelationID     int       `db:"ssr_id"`
	CreatedAt      time.Time `db:"created_at"`
	Status         string    `db:"status"`
	SupervisorFull `db:"sv"`
	Work           `db:"work"`
}

type SvRelation struct {
	RelationID int       `db:"ssr_id"`
	CreatedAt  time.Time `db:"created_at"`
	Status     string    `db:"ssr_status"`
	*StProfile
	*Work
}
