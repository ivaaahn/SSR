package dto

type LoginResponseDTO struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
