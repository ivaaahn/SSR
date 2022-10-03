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
	workRepo     WorkRepo
	relationRepo RelationRepo
	stRepo       StudentRepo
	svRepo       SupervisorRepo
	waypointRepo WaypointRepo
}

func NewWork(
	workRepo WorkRepo,
	ssrRepo RelationRepo,
	stRepo StudentRepo,
	svRepo SupervisorRepo,
	waypointRepo WaypointRepo,
	l logger.Interface,
) *Work {
	return &Work{
		Base:         NewBase(l),
		workRepo:     workRepo,
		relationRepo: ssrRepo,
		stRepo:       stRepo,
		svRepo:       svRepo,
		waypointRepo: waypointRepo,
	}
}

func (service *Work) checkIfBegin(relations []*entity.Relation, workID int) bool {
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

func (service *Work) GetStudentWorks(studentID int) (*dto.StudentViewWorkPlenty, error) {
	studentData, err := service.stRepo.GetStudentShort(studentID)
	if err != nil {
		return nil, err
	}

	semester := service.recognizeSemester(studentData.Year)

	worksData, err := service.workRepo.GetStudentWorks(studentData.DepartmentID, semester)
	if err != nil {
		return nil, err
	}

	var resp []*dto.StudentViewWorkShortResp

	for _, work := range worksData {
		resp = append(resp, &dto.StudentViewWorkShortResp{
			Work: dto.WorkShortResp{
				WorkID: work.WorkID,
				Subject: dto.SubjectResp{
					ID:         work.Subject.SubjectID,
					Name:       work.Subject.Name,
					Department: work.Subject.DepartmentID,
				},
				Kind: dto.WorkKindResp{
					ID:   work.WorkKind.WorkKindID,
					Name: work.WorkKind.Name,
				},
			},
		})
	}

	return &dto.StudentViewWorkPlenty{
		Works: resp,
	}, nil
}

func (service *Work) GetSupervisorWorks(supervisorID int) (*dto.SupervisorViewWorkPlenty, error) {
	worksData, err := service.workRepo.GetSupervisorWorks(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SupervisorViewWorkShortResp

	for _, db := range worksData {
		resp = append(resp, &dto.SupervisorViewWorkShortResp{
			Work: dto.WorkShortResp{
				WorkID: db.WorkID,
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

	return &dto.SupervisorViewWorkPlenty{
		Works: resp,
	}, nil
}

func (service *Work) Get(workID int) (*dto.WorkFullResp, error) {
	workData, err := service.workRepo.Get(workID)
	if err != nil {
		return nil, err
	}

	waypointsData, err := service.waypointRepo.GetPlenty(workID)
	if err != nil {
		return nil, err
	}

	var waypoints []*dto.WaypointResp
	for _, waypointData := range waypointsData {
		waypoints = append(waypoints, &dto.WaypointResp{
			Title:       waypointData.Title,
			Description: waypointData.Description,
			Deadline:    misc.Date{Time: waypointData.Deadline},
		})
	}

	supervisorsData, err := service.svRepo.GetSupervisorsByWorkID(workID)
	if err != nil {
		return nil, err
	}
	var supervisors []*dto.WorkSupervisorShort
	for _, db := range supervisorsData {
		supervisors = append(supervisors, &dto.WorkSupervisorShort{
			SupervisorShort: dto.SupervisorShort{
				UserID:    db.Supervisor.User.UserID,
				FirstName: db.User.FirstName,
				LastName:  db.User.LastName,
			},
			IsHead: db.IsHead,
			IsFull: db.IsFull,
		})
	}

	return &dto.WorkFullResp{
		WorkID:      workData.WorkID,
		Description: workData.Description,
		Semester:    workData.Semester,
		Subject: dto.SubjectResp{
			ID:         workData.Subject.SubjectID,
			Name:       workData.Subject.Name,
			Department: workData.Subject.DepartmentID,
		},
		Kind: dto.WorkKindResp{
			ID:   workData.WorkKind.WorkKindID,
			Name: workData.WorkKind.Name,
		},
		Waypoints:   waypoints,
		Supervisors: supervisors,
	}, nil
}

//
//func (service *Work) GetPlenty(workID int) (*dto.WorkFullResp, error) {
//	workData, err := service.workRepo.Get(workID)
//	if err != nil {
//		return nil, err
//	}
//
//	waypointsData, err := service.waypointRepo.GetPlenty(workID)
//
//	var waypoints []*dto.WaypointResp
//
//	for _, waypointData := range waypointsData {
//		waypoints = append(waypoints, &dto.WaypointResp{
//			Title:       waypointData.Title,
//			Description: waypointData.Description,
//			Deadline:    misc.Date{Time: waypointData.Deadline},
//		})
//	}
//
//	return &dto.WorkFullResp{
//		WorkID:      workData.WorkID,
//		Description: workData.Description,
//		Semester:    workData.Semester,
//		Subject: dto.SubjectResp{
//			ID:         workData.Subject.SubjectID,
//			Name:       workData.Subject.Name,
//			Department: workData.Subject.DepartmentID,
//		},
//		Kind: dto.WorkKindResp{
//			ID:   workData.WorkKind.WorkKindID,
//			Name: workData.WorkKind.Name,
//		},
//		Waypoints: waypoints,
//	}, nil
//}
