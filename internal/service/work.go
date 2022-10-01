package service

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"time"
)

type Work struct {
	*Base
	workRepo WorkRepo
	ssrRepo  RelationRepo
	stRepo   StudentRepo
}

func NewWork(workRepo WorkRepo, ssrRepo RelationRepo, stRepo StudentRepo, l logger.Interface) *Work {
	return &Work{
		Base:     NewBase(l),
		workRepo: workRepo,
		ssrRepo:  ssrRepo,
		stRepo:   stRepo,
	}
}

func (service *Work) checkIfBegin(relations []*entity.StRelation, workID int) bool {
	for _, rel := range relations {
		if rel.Work.WorkID == workID {
			return true
		}
	}

	return false
}

func (service *Work) recognizeSemester(studentYear int) int {
	month := time.Now().Month()
	if time.February <= month && month <= time.August {
		return studentYear * 2
	} else {
		return studentYear*2 - 1
	}
}

func (service *Work) GetStudentWorks(studentID int) (*dto.StWorkPlenty, error) {
	studentData, err := service.stRepo.GetStudent(studentID)
	if err != nil {
		return nil, err
	}

	semester := service.recognizeSemester(studentData.Year)

	worksData, err := service.workRepo.GetStudentWorks(studentData.DepartmentID, semester)
	if err != nil {
		return nil, err
	}

	//relationsData, err := service.ssrRepo.GetStudentRelations(studentID)
	//if err != nil {
	//	return nil, err
	//}

	var resp []*dto.StWork

	for _, work := range worksData {
		resp = append(resp, &dto.StWork{
			Work: &dto.WorkResp{
				WorkID:      work.WorkID,
				Description: work.Description,
				Semester:    work.Semester,
				Subject: dto.SubjectResp{
					Name:       work.Subject.Name,
					Department: work.Subject.DepartmentID,
				},
				Kind: dto.WorkKindResp{
					Name: work.WorkKind.Name,
				},
				//IsStarted: service.checkIfBegin(relations, work.WorkID), TODO
			},
		})
	}

	return &dto.StWorkPlenty{
		Works: resp,
	}, nil
}

//
//func (service *Work) GetSupervisorWorks(supervisorID int) (*dto.SvWorkPlenty, error) {
//	dbData, err := service.workRepo.GetWorksBySupervisorID(supervisorID)
//	if err != nil {
//		return nil, err
//	}
//
//	var resp []*dto.SvWork
//
//	for _, db := range dbData {
//		resp = append(resp, &dto.SvWork{
//			WorkID:      db.WorkID,
//			Kind:        db.WorkKindName,
//			Description: db.Description,
//			Subject:     db.SubjectName,
//			Head:        db.Head,
//		})
//	}
//
//	return &dto.SvWorkPlenty{
//		SupervisorID: supervisorID,
//		Works:        resp,
//	}, nil
//}
//
//func (service *Work) GetWorkSupervisors(workID int) (*dto.WorkSvPlenty, error) {
//	dbData, err := service.workRepo.GetSupervisorsByWorkID(workID)
//	if err != nil {
//		return nil, err
//	}
//
//	var resp []*dto.WorkSv
//
//	for _, db := range dbData {
//		resp = append(resp, &dto.WorkSv{
//			SvProfile: dto.SvProfile{
//				Email:      db.Email,
//				FirstName:  db.FirstName,
//				LastName:   db.LastName,
//				About:      db.About,
//				Birthdate:  misc.Date{Time: db.Birthdate},
//				PhotoUrl:   db.PhotoUrl,
//				Department: db.DepartmentID,
//			},
//			Head: db.Head,
//			Full: db.Full,
//		})
//	}
//
//	return &dto.WorkSvPlenty{
//		WorkID:      workID,
//		Supervisors: resp,
//	}, nil
//}
