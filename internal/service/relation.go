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

func (service *Relation) GetStudentRelations(studentID int) (*dto.StRelationPlenty, error) {
	dbData, err := service.repo.GetStudentRelations(studentID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.StRelationResp

	for _, db := range dbData {
		resp = append(resp, &dto.StRelationResp{
			BidID:     db.RelationID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Supervisor: dto.SvProfile{
				Email:     db.User.Email,
				FirstName: db.User.FirstName,
				LastName:  db.User.LastName,
				About:     db.SupervisorFull.About,
				Birthdate: misc.Date{
					Time: db.Birthdate,
				},
				PhotoUrl:   db.User.PhotoUrl,
				Department: db.SupervisorFull.DepartmentID,
			},
			Work: dto.WorkResp{
				WorkID:      db.WorkID,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Kind: dto.WorkKindResp{
					ID:   db.WorkKind.WorkKindID,
					Name: db.WorkKind.Name,
				},
				Subject: dto.SubjectResp{
					ID:         db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.StRelationPlenty{Relations: resp}, nil
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

func (service *Relation) Create(data *dto.RelationCreateReq) (*dto.RelationCreateResp, error) {
	relationID, err := service.repo.Create(data.StudentID, data.SupervisorID, data.WorkID)
	if err != nil {
		return nil, err
	}

	return &dto.RelationCreateResp{RelationID: relationID}, nil
}

//func (service *Relation) Accept(data *dto.CreateSSR) (*dto.StViewRelation, error) {
//	ssrID, err := service.repo.UpdateStatus(data.BidID, "wip")
//	if err != nil {
//		return nil, err
//	}
//
//	ssr, err := service.repo.GetStudentRelation(data.StudentID, ssrID)
//	if err != nil {
//		return nil, err
//	}
//
//	return &dto.StViewRelation{
//		RelID:     ssr.RelationID,
//		Status:    ssr.Status,
//		CreatedAt: ssr.CreatedAt,
//		Supervisor: dto.SvProfile{
//			Email:      ssr.SvProfile.Email,
//			FirstName:  ssr.SvProfile.FirstName,
//			LastName:   ssr.SvProfile.LastName,
//			About:      ssr.SvProfile.About,
//			Birthdate:  misc.Date{Time: ssr.Birthdate},
//			PhotoUrl:   ssr.PhotoUrl,
//			Department: ssr.DepartmentID,
//		},
//		Work: dto.WorkResp{
//			WorkID:      ssr.WorkID,
//			Description: ssr.Work.Description,
//			Semester:    ssr.Work.Semester,
//			Subject: dto.SubjectResp{
//				ID:         ssr.ID,
//				Name:       ssr.Subject.Name,
//				Department: ssr.DepartmentID,
//			},
//		},
//	}, nil
//}
