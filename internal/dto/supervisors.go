package dto

import "ssr/pkg/misc"

type Supervisor struct {
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	About      string    `json:"about"`
	Birthdate  misc.Date `json:"birthdate" swaggertype:"string"`
	PhotoUrl   string    `json:"photo_url"`
	Department string    `json:"department"`
}

type SupervisorShort struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type WorkSupervisorShort struct {
	SupervisorShort
	IsHead bool `json:"head"`
	IsFull bool `json:"full"`
}
