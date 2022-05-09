package dto

import (
	"ssr/internal/entity"
	"ssr/internal/misc"
)

type UserResponseDTO struct {
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	AvatarUrl misc.NullString `json:"avatarUrl"`
	Role      entity.UserRole `json:"role"`
}
