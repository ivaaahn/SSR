package dto

type LoginResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user-id"`
	Role   string `json:"role"`
}
