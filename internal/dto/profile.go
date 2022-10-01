package dto

import (
	"ssr/pkg/misc"
)

type UserInfo struct {
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	AvatarUrl misc.NullString `json:"avatarUrl"`
}

type StProfile struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhotoUrl    string `json:"photoUrl"`
	Year        int    `json:"year"`
	StudentCard string `json:"studentCard"`
	Department  string `json:"department"`
}

type SvProfile struct {
	Email      string    `json:"email"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	About      string    `json:"about"`
	Birthdate  misc.Date `json:"birthdate" swaggertype:"string"`
	PhotoUrl   string    `json:"avatarUrl"`
	Department string    `json:"department"`
}
