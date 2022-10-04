package entity

import "time"

type StatusSSR string

const (
	Pending   StatusSSR = "pending"
	Rejected            = "rejected"
	Cancelled           = "cancelled"
	Accepted            = "accepted"
	Wip                 = "wip"
	Completed           = "done"
)

type WaypointRelation struct {
	*Waypoint
	Status string `db:"waypoint_status"`
	SsrID  int    `db:"ssr_id"`
}

type Relation struct {
	RelationID int       `db:"ssr_id"`
	CreatedAt  time.Time `db:"created_at"`
	Status     string    `db:"status"`
	Supervisor `db:"sv"`
	Student    `db:"st"`
	Work       `db:"work"`
}

type RelationShort struct {
	RelationID      int    `db:"ssr_id"`
	Status          string `db:"status"`
	SupervisorShort `db:"sv"`
	StudentShort    `db:"st"`
	Work            `db:"work"`
}
