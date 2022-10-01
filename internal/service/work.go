package service

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type Work struct {
	*Base
	repoWork WorkRepo
	repoSsr  RelationRepo
}

func NewWork(rWork WorkRepo, rSsr RelationRepo, l logger.Interface) *Work {
	return &Work{
		Base:     NewBase(l),
		repoWork: rWork,
		repoSsr:  rSsr,
	}
}

func checkIfBegin(relations []*entity.StRelation, workID int) bool {
	for _, rel := range relations {
		if rel.Work.WorkID == workID {
			return true
		}
	}

	return false
}

func (service *Work) GetStudentWorks(studentID int) (*dto.StWorks, error) {
	dbData, err := service.repoWork.GetWorksByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	relations, err := service.repoSsr.GetStudentRelations(studentID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.StWork

	for _, db := range dbData {
		resp = append(resp, &dto.StWork{
			WorkID:      db.WorkID,
			Kind:        db.WorkKindName,
			Description: db.Description,
			Subject:     db.SubjectName,
			IsStarted:   checkIfBegin(relations, db.WorkID),
		})
	}

	return &dto.StWorks{
		StudentID: studentID,
		Works:     resp,
	}, nil
}

func (service *Work) GetSupervisorWorks(supervisorID int) (*dto.SvWorkPlenty, error) {
	dbData, err := service.repoWork.GetWorksBySupervisorID(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SvWork

	for _, db := range dbData {
		resp = append(resp, &dto.SvWork{
			WorkID:      db.WorkID,
			Kind:        db.WorkKindName,
			Description: db.Description,
			Subject:     db.SubjectName,
			Head:        db.Head,
		})
	}

	return &dto.SvWorkPlenty{
		SupervisorID: supervisorID,
		Works:        resp,
	}, nil
}

func (service *Work) GetWorkSupervisors(workID int) (*dto.WorkSvPlenty, error) {
	dbData, err := service.repoWork.GetSupervisorsByWorkID(workID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.WorkSv

	for _, db := range dbData {
		resp = append(resp, &dto.WorkSv{
			SvProfile: dto.SvProfile{
				SupervisorID: db.SupervisorID,
				Email:        db.Email,
				FirstName:    db.FirstName,
				LastName:     db.LastName,
				About:        db.About,
				Birthdate:    misc.Date{Time: db.Birthdate},
				AvatarUrl:    misc.NullString(db.PhotoUrl),
				Department:   db.DepartmentID,
			},
			Head: db.Head,
			Full: db.Full,
		})
	}

	return &dto.WorkSvPlenty{
		WorkID:      workID,
		Supervisors: resp,
	}, nil
}
