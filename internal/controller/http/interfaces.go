package http

import (
	"ssr/internal/dto"
)

type (
	AuthService interface {
		Login(email, password string) (*dto.LoginResponse, error)
	}
	ProfileService interface {
		GetStudentProfile(userID int) (*dto.StProfile, error)
		GetSupervisorProfile(userID int) (*dto.SvProfile, error)
	}
	SvProfileService interface {
		GetSupervisorProfile(userID int) (*dto.SvProfile, error)
	}
	SvBidService interface {
		GetSupervisorBids(supervisorID int) (*dto.SvBids, error)
		Resolve(data *dto.ResolveBid) error
	}
	RelationsService interface {
		GetStudentRelations(studentID int) (*dto.StRelationPlenty, error)
		Create(data *dto.RelationCreateReq) (*dto.RelationCreateResp, error)
	}
	WorkService interface {
		GetStudentWorks(studentID int) (*dto.StWorkPlenty, error)
		GetSupervisorWorks(supervisorID int) (*dto.SvWorkPlenty, error)
		GetWorkSupervisors(workID int) (*dto.WorkSvPlenty, error)
	}
	FeedbackService interface {
		Add(data *dto.FeedbackReq) (int, error)
		GetOnSupervisor(supervisorID int) (*dto.FeedbackPlenty, error)
	}
)
