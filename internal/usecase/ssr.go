package usecase

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type SSR struct {
	*Base
	repo IRepoSSR
}

func NewSSR(r IRepoSSR, l logger.Interface) *SSR {
	return &SSR{
		Base: NewBase(l),
		repo: r,
	}
}

func (uc *SSR) CheckIfStudentBeginWork(studentID, workID int) (bool, error) {
	relations, err := uc.repo.GetStudentRelations(studentID)
	if err != nil {
		return false, err
	}

	for _, rel := range relations {
		if rel.Work.WorkID == workID {
			return true, nil
		}
	}

	return false, nil
}

func (uc *SSR) Create(data *dto.CreateSSR) (*dto.StudentViewSSR, error) {
	ssrID, err := uc.repo.UpdateStatus(data.BidID, "wip")
	if err != nil {
		return nil, err
	}

	ssr, err := uc.repo.GetStudentRelation(data.StudentID, ssrID)
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
