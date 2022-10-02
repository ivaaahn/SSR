package service

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
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

	var resp []*dto.StWorkResp

	for _, work := range worksData {
		resp = append(resp, &dto.StWorkResp{
			Work: dto.WorkResp{
				WorkID:      work.WorkID,
				Description: work.Description,
				Semester:    work.Semester,
				Subject: dto.SubjectResp{
					ID:         work.Subject.SubjectID,
					Name:       work.Subject.Name,
					Department: work.Subject.DepartmentID,
				},
				Kind: dto.WorkKindResp{
					ID:   work.WorkKind.WorkKindID,
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

func (service *Work) GetSupervisorWorks(supervisorID int) (*dto.SvWorkPlenty, error) {
	worksData, err := service.workRepo.GetSupervisorWorks(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SvWorkResp

	for _, db := range worksData {
		resp = append(resp, &dto.SvWorkResp{
			Work: dto.WorkResp{
				WorkID:      db.WorkID,
				Description: db.Description,
				Semester:    db.Semester,
				Subject: dto.SubjectResp{
					ID:         db.Subject.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
				Kind: dto.WorkKindResp{
					ID:   db.WorkKind.WorkKindID,
					Name: db.WorkKind.Name,
				},
			},
			IsHead: db.IsHead,
			IsFull: db.IsFull,
		})
	}

	return &dto.SvWorkPlenty{
		Works: resp,
	}, nil
}

func (service *Work) GetWorkSupervisors(workID int) (*dto.WorkSvPlenty, error) {
	supervisorsData, err := service.workRepo.GetWorkSupervisors(workID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.WorkSv

	for _, db := range supervisorsData {
		resp = append(resp, &dto.WorkSv{
			SvProfile: dto.SvProfile{
				Email:      db.User.Email,
				FirstName:  db.User.FirstName,
				LastName:   db.User.LastName,
				About:      db.SupervisorFull.About,
				Birthdate:  misc.Date{Time: db.SupervisorFull.Birthdate},
				PhotoUrl:   db.User.PhotoUrl,
				Department: db.DepartmentID,
			},
			IsHead: db.IsHead,
			IsFull: db.IsFull,
		})
	}

	return &dto.WorkSvPlenty{
		Supervisors: resp,
	}, nil
}
