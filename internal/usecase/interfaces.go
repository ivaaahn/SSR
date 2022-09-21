package usecase

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
)

type (
	IRepoAuth interface {
		GetUserInfo(email string) (*entity.Auth, error)
	}
	IUsecaseAuth interface {
		Login(email, password string) (*dto.LoginResponse, error)
	}

	IRepoProfile interface {
		GetStudentProfile(email string) (*entity.StudentProfile, error)
		GetSupervisorProfile(email string) (*entity.SupervisorProfile, error)
	}
	IUsecaseProfile interface {
		GetStudentProfile(email string) (*dto.StudentProfile, error)
		GetSupervisorProfile(email string) (*dto.SupervisorProfile, error)
	}

	IRepoSSR interface {
		Create(studentID, supervisorID, workID int) (int, error)
		GetStudentBids(studentID int) ([]*entity.StudentSsr, error)
		GetSupervisorBids(studentID int) ([]*entity.SupervisorSsr, error)
		GetStudentRelations(studentID int) ([]*entity.StudentSsr, error)
		GetStudentRelation(studentID, ssrID int) (*entity.StudentSsr, error)
		UpdateStatus(id int, newStatus entity.StatusSSR) (int, error)
	}

	IUsecaseStudentBid interface {
		GetStudentBids(studentID int) (*dto.StudentBids, error)
		Apply(data *dto.ApplyBid) (*dto.ApplyBidResponse, error)
	}
	IUseCaseSupervisorBid interface {
		GetSupervisorBids(supervisorID int) (*dto.SupervisorBids, error)
		Resolve(data *dto.ResolveBid) error
	}

	IUseCaseStudentRelation interface {
		Create(data *dto.CreateSSR) (*dto.StudentViewSSR, error)
	}

	IRepoWork interface {
		GetWorksByStudentID(studentID int) ([]*entity.Work, error)
		GetWorksBySupervisorID(supervisorID int) ([]*entity.WorkOfSupervisor, error)
		GetSupervisorsByWorkID(workID int) ([]*entity.SupervisorOfWork, error)
	}
	IStudentWorkUC interface {
		GetStudentWorks(studentID int) (*dto.StudentWorks, error)
		GetWorkSupervisors(workID int) (*dto.WorkSupervisorPlenty, error)
	}
	ISupervisorWorkUC interface {
		GetSupervisorWorks(supervisorID int) (*dto.SupervisorWorkPlenty, error)
	}

	IUsecaseFeedback interface {
		Add(data *dto.FeedbackReq) (int, error)
		GetOnSupervisor(supervisorID int) (*dto.FeedbackPlenty, error)
	}
	IRepoFeedback interface {
		Create(studentID, supervisorID, workID int, content string) (int, error)
		GetBySupervisorID(supervisorID int) ([]*entity.Feedback, error)
	}
)
