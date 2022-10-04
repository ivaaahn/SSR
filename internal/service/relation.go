package service

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
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

func (service *Relation) GetPlenty(studentID, supervisorID int) (*dto.RelationPlenty, error) {
	var relations []*entity.RelationShort
	var err error

	if studentID != 0 {
		relations, err = service.repo.GetRelationsByStudentID(studentID)
	} else {
		relations, err = service.repo.GetRelationsBySupervisorID(supervisorID)
	}

	if err != nil {
		return nil, err
	}

	var resp []*dto.RelationShortResp

	for _, db := range relations {
		resp = append(resp, &dto.RelationShortResp{
			RelationID: db.RelationID,
			Status:     db.Status,
			Supervisor: dto.SupervisorShort{
				UserID:    db.SupervisorShort.User.ID,
				FirstName: db.SupervisorShort.User.FirstName,
				LastName:  db.SupervisorShort.User.LastName,
			},
			Student: dto.StudentShort{
				UserID:    db.StudentShort.User.ID,
				FirstName: db.StudentShort.User.FirstName,
				LastName:  db.StudentShort.User.LastName,
			},
			Work: dto.WorkShortResp{
				WorkID: db.WorkID,
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

	return &dto.RelationPlenty{Relations: resp}, nil
}

func (service *Relation) Get(id int) (*dto.RelationResp, error) {
	rel, err := service.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return &dto.RelationResp{
		RelationID: rel.RelationID,
		Status:     rel.Status,
		Supervisor: dto.Supervisor{
			UserID:     rel.Supervisor.User.ID,
			Email:      rel.Supervisor.User.Email,
			FirstName:  rel.Supervisor.User.FirstName,
			LastName:   rel.Supervisor.User.LastName,
			About:      rel.Supervisor.About,
			Birthdate:  misc.Date{Time: rel.Supervisor.Birthdate},
			PhotoUrl:   rel.Supervisor.User.PhotoUrl,
			Department: rel.Supervisor.DepartmentID,
		},
		Work: dto.WorkResp{
			WorkID: rel.WorkID,
			Kind: dto.WorkKindResp{
				ID:   rel.WorkKind.WorkKindID,
				Name: rel.WorkKind.Name,
			},
			Subject: dto.SubjectResp{
				ID:         rel.SubjectID,
				Name:       rel.Subject.Name,
				Department: rel.Subject.DepartmentID,
			},
			Description: rel.Work.Description,
			Semester:    rel.Work.Semester,
		},
		Student: dto.Student{
			UserID:      rel.Student.User.ID,
			Email:       rel.Student.User.Email,
			FirstName:   rel.Student.User.FirstName,
			LastName:    rel.Student.User.LastName,
			PhotoUrl:    rel.Student.User.PhotoUrl,
			Year:        rel.Student.Year,
			StudentCard: rel.Student.StudentCard,
			Department:  rel.Student.DepartmentID,
		},
	}, nil
}

func (service *Relation) Create(data *dto.RelationCreateReq) (*dto.RelationCreateResp, error) {
	relationID, err := service.repo.Create(data.StudentID, data.SupervisorID, data.WorkID)
	if err != nil {
		return nil, err
	}

	return &dto.RelationCreateResp{RelationID: relationID}, nil
}

func (service *Relation) Update(data *dto.RelationUpdateReq) (*dto.RelationResp, error) {
	relationID, err := service.repo.Update(data.RelationID, data.Status)
	if err != nil {
		return nil, err
	}

	relation, err := service.repo.Get(relationID)

	return &dto.RelationResp{
		RelationID: relation.RelationID,
		Work: dto.WorkResp{
			WorkID:      relation.Work.WorkID,
			Description: relation.Work.Description,
			Semester:    relation.Semester,
			Subject: dto.SubjectResp{
				ID:         relation.Subject.SubjectID,
				Name:       relation.Subject.Name,
				Department: relation.Subject.DepartmentID,
			},
			Kind: dto.WorkKindResp{
				ID:   relation.WorkKind.WorkKindID,
				Name: relation.WorkKind.Name,
			},
		},
		Student: dto.Student{
			UserID:      relation.Student.User.ID,
			Email:       relation.Student.User.Email,
			FirstName:   relation.Student.User.FirstName,
			LastName:    relation.Student.User.LastName,
			PhotoUrl:    relation.Student.User.PhotoUrl,
			Year:        relation.Student.Year,
			StudentCard: relation.Student.StudentCard,
			Department:  relation.Student.DepartmentID,
		},
		Status: relation.Status,
	}, nil
}
