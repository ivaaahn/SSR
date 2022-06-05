package usecase

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type SSRUseCase struct {
	*BaseUC
	repo IRelRepo
}

func NewSsrUC(r IRelRepo, l logger.Interface) *SSRUseCase {
	return &SSRUseCase{
		BaseUC: NewUC(l),
		repo:   r,
	}
}

func (uc *SSRUseCase) Create(data *dto.CreateSSR) (*dto.StudentViewSSR, error) {
	ssrID, err := uc.repo.UpdateStatus(data.BidID, "wip")
	if err != nil {
		return nil, err
	}

	ssr, err := uc.repo.GetStudentViewSSR(data.StudentID, ssrID)
	if err != nil {
		return nil, err
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
