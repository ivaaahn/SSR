package entity

import (
	"time"
)

type Waypoint struct {
	WaypointID  int `db:"waypoint_id"`
	WorkID      int `db:"work_id"`
	Deadline    time.Time
	Title       string
	Description string
}
