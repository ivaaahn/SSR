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

type StudentProfile struct {
	StudentID   int             `json:"studentID"`
	Email       string          `json:"email"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	AvatarUrl   misc.NullString `json:"avatarUrl"swaggertype:"string"`
	Year        int             `json:"year"`
	StudentCard string          `json:"studentCard"`
	Department  string          `json:"department"`
}

type SupervisorProfile struct {
	SupervisorID int             `json:"supervisorID"`
	Email        string          `json:"email"`
	FirstName    string          `json:"firstName"`
	LastName     string          `json:"lastName"`
	About        string          `json:"about"`
	Birthdate    misc.Date       `json:"birthdate"`
	AvatarUrl    misc.NullString `json:"avatarUrl"swaggertype:"string"`
	Department   string          `json:"department"`
}
