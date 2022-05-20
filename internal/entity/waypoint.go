package entity

import (
	"time"
)

type Waypoint struct {
	WorkID      int
	Deadline    time.Time
	Description string
}
