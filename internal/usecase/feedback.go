package usecase

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
)

type Feedback struct {
	*Base
	repo IRepoFeedback
}

func NewFeedback(r IRepoFeedback, l logger.Interface) *Feedback {
	return &Feedback{
		Base: NewBase(l),
		repo: r,
	}
}

func (uc *Feedback) Add(data *dto.FeedbackReq) (int, error) {
	return uc.repo.Create(data.StudentID, data.SupervisorID, data.WorkID, data.Content)
}

func (uc *Feedback) GetOnSupervisor(supervisorID int) (*dto.FeedbackPlenty, error) {
	dbData, err := uc.repo.GetBySupervisorID(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.FeedbackResp

	for _, db := range dbData {
		resp = append(resp, &dto.FeedbackResp{
			StudentID:       db.StudentID,
			StudentFullName: db.StudentFullName,
			SupervisorID:    db.SupervisorID,
			WorkID:          db.WorkID,
			WorkKind:        db.WorkKind,
			WorkSubject:     db.WorkSubject,
			Content:         db.Content,
		})
	}

	return &dto.FeedbackPlenty{Feedbacks: resp}, nil
}
