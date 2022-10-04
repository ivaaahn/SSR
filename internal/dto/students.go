package dto

type Student struct {
	UserID      int    `json:"user_id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhotoUrl    string `json:"photo_url"`
	Year        int    `json:"year"`
	StudentCard string `json:"student_card"`
	Department  string `json:"department"`
}

type StudentShort struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
