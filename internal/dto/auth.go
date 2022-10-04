package dto

type LoginResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
	UserID    int    `json:"user_id"`
	Role      string `json:"role"`
}
