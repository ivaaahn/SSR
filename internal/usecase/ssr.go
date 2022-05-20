package usecase

import (
	"fmt"
	"ssr/internal/dto"
	"ssr/pkg/misc"
)

type SSRUseCase struct {
	repo IRelRepo
}

func NewSsrUC(r IRelRepo) *SSRUseCase {
	return &SSRUseCase{
		repo: r,
	}
}

func (uc *SSRUseCase) Create(data *dto.CreateSSR) (*dto.StudentViewSSR, error) {
	ssrID, err := uc.repo.UpdateStatus(data.BidID, "wip")
	if err != nil {
		return nil, fmt.Errorf("SSRUseCase - Create - repo.UpdateStatus %w", err)
	}

	ssr, err := uc.repo.GetStudentViewSSR(data.StudentID, ssrID)
	if err != nil {
		return nil, fmt.Errorf("SSRUseCase - GetStudentViewSSR - repo.GetStudentViewSSR %w", err)
	}

	return &dto.StudentViewSSR{
		RelID:     ssr.BidID,
		Status:    ssr.Status,
		CreatedAt: ssr.CreatedAt,
		Supervisor: dto.SupervisorProfile{
			SupervisorID: ssr.SupervisorID,
			Email:        ssr.SupervisorProfile.Email,
			FirstName:    ssr.SupervisorProfile.FirstName,
			LastName:     ssr.SupervisorProfile.LastName,
			About:        ssr.SupervisorProfile.About,
			Birthdate:    misc.Date{Time: ssr.Birthdate},
			AvatarUrl:    misc.NullString(ssr.Avatar),
			Department:   ssr.DepartmentID,
		},
		Work: dto.Work{
			WorkID:      ssr.WorkID,
			Name:        ssr.Work.WorkKindName,
			Description: ssr.Work.Description,
			Semester:    ssr.Work.Semester,
			Subject: dto.SubjectResp{
				SubjectID:  ssr.SubjectID,
				Name:       ssr.SubjectName,
				Department: ssr.DepartmentID,
			},
		},
	}, nil
}
