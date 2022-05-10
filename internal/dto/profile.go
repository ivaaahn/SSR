package dto

import (
	"ssr/pkg/misc"
)

type UserInfoDTO struct {
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	AvatarUrl misc.NullString `json:"avatarUrl"`
}

type StudentProfileDTO struct {
	StudentID   int             `json:"studentID"`
	Email       string          `json:"email"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	AvatarUrl   misc.NullString `json:"avatarUrl"`
	Year        int             `json:"year"`
	StudentCard string          `json:"studentCard"`
	Department  string          `json:"department"`
}

type SupervisorProfileDTO struct {
	SupervisorID int             `json:"supervisorID"`
	Email        string          `json:"email"`
	FirstName    string          `json:"firstName"`
	LastName     string          `json:"lastName"`
	About        string          `json:"about"`
	Birthdate    misc.BirthDate  `json:"birthdate"`
	AvatarUrl    misc.NullString `json:"avatarUrl"`
	Department   string          `json:"department"`
}
