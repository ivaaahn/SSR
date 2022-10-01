package http

import (
	"ssr/internal/dto"
)

type (
	AuthService interface {
		Login(email, password string) (*dto.LoginResponse, error)
	}
	StProfileService interface {
		GetStudentProfile(userID int) (*dto.StProfile, error)
	}
	SvProfileService interface {
		GetSupervisorProfile(userID int) (*dto.SvProfile, error)
	}
	StBidService interface {
		GetStudentBids(studentID int) (*dto.StBids, error)
		Apply(data *dto.ApplyBid) (*dto.ApplyBidResp, error)
	}
	SvBidService interface {
		GetSupervisorBids(supervisorID int) (*dto.SvBids, error)
		Resolve(data *dto.ResolveBid) error
	}
	StRelationService interface {
		Create(data *dto.CreateSSR) (*dto.StViewRelation, error)
	}
	StWorkService interface {
		GetStudentWorks(studentID int) (*dto.StWorks, error)
		GetWorkSupervisors(workID int) (*dto.WorkSvPlenty, error)
	}
	SvWorkService interface {
		GetSupervisorWorks(supervisorID int) (*dto.SvWorkPlenty, error)
	}
	FeedbackService interface {
		Add(data *dto.FeedbackReq) (int, error)
		GetOnSupervisor(supervisorID int) (*dto.FeedbackPlenty, error)
	}
)
