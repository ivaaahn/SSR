package entity

import "time"

const (
	bidPending   = "заявка ожидает ответа"
	bidDeclined  = "заявка отклонена"
	bidCancelled = "заявка отозвана"
	bidAccepted  = "заявка принята"
	inProgress   = "в работе"
	completed    = "завершено"
)

type StudentBid struct {
	BidID     int       `db:"ssr_id"`
	CreatedAt time.Time `db:"created_at"`
	Status    string
	*SupervisorProfile
	*Work
	*Subject
}

type SupervisorBid struct {
	BidID     int       `db:"ssr_id"`
	CreatedAt time.Time `db:"created_at"`
	Status    string
	*StudentProfile
	*Work
	*Subject
}
