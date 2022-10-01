package service

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type Relation struct {
	*Base
	repo RelationRepo
}

func NewRelation(r RelationRepo, l logger.Interface) *Relation {
	return &Relation{
		Base: NewBase(l),
		repo: r,
	}
}

func (service *Relation) CheckIfStudentBeginWork(studentID, workID int) (bool, error) {
	relations, err := service.repo.GetStudentRelations(studentID)
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

func (service *Relation) Create(data *dto.CreateSSR) (*dto.StViewRelation, error) {
	ssrID, err := service.repo.UpdateStatus(data.BidID, "wip")
	if err != nil {
		return nil, err
	}

	ssr, err := service.repo.GetStudentRelation(data.StudentID, ssrID)
	if err != nil {
		return nil, err
	}

	return &dto.StViewRelation{
		RelID:     ssr.BidID,
		Status:    ssr.Status,
		CreatedAt: ssr.CreatedAt,
		Supervisor: dto.SvProfile{
			Email:      ssr.SvProfile.Email,
			FirstName:  ssr.SvProfile.FirstName,
			LastName:   ssr.SvProfile.LastName,
			About:      ssr.SvProfile.About,
			Birthdate:  misc.Date{Time: ssr.Birthdate},
			PhotoUrl:   ssr.PhotoUrl,
			Department: ssr.DepartmentID,
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
