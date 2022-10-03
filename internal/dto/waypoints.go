package dto

import "ssr/pkg/misc"

type WaypointResp struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    misc.Date `json:"deadline" swaggertype:"string"`
}

type WaypointPlenty struct {
	Waypoints []*WaypointResp `json:"waypoints"`
}
