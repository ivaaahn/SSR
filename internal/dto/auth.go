package dto

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
