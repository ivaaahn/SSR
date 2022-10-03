package http

import (
	"ssr/internal/dto"
)

type (
	AuthService interface {
		Login(email, password string) (*dto.LoginResponse, error)
	}
	ProfileService interface {
		GetStudentProfile(userID int) (*dto.Student, error)
		GetSupervisorProfile(userID int) (*dto.Supervisor, error)
	}
	RelationsService interface {
		GetPlenty(studentID, supervisorID int) (*dto.RelationPlenty, error)
		Create(data *dto.RelationCreateReq) (*dto.RelationCreateResp, error)
		Update(data *dto.RelationUpdateReq) (*dto.RelationResp, error)
		Get(RelationID int) (*dto.RelationResp, error)
	}
	WorkService interface {
		GetStudentWorks(studentID int) (*dto.StudentViewWorkPlenty, error)
		GetSupervisorWorks(supervisorID int) (*dto.SupervisorViewWorkPlenty, error)
		GetWorkSupervisors(workID int) (*dto.WorkSupervisorPlenty, error)
	}
	FeedbackService interface {
		Add(data *dto.FeedbackReq) (int, error)
		GetOnSupervisor(supervisorID int) (*dto.FeedbackPlenty, error)
	}
)
