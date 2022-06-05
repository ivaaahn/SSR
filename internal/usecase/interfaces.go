package usecase

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
)

type (
	IAuthRepo interface {
		GetUserInfo(email string) (*entity.Auth, error)
	}
	IAuthUC interface {
		Login(email, password string) (*dto.LoginResponse, error)
	}

	IProfileRepo interface {
		GetStudentProfile(email string) (*entity.StudentProfile, error)
		GetSupervisorProfile(email string) (*entity.SupervisorProfile, error)
	}
	IProfileUC interface {
		GetStudentProfile(email string) (*dto.StudentProfile, error)
		GetSupervisorProfile(email string) (*dto.SupervisorProfile, error)
	}

	IRelRepo interface {
		Create(studentID, supervisorID, workID int) (int, error)
		GetStudentViewBidPlenty(studentID int) ([]*entity.StudentViewSsr, error)
		GetSupervisorViewBidPlenty(studentID int) ([]*entity.SupervisorViewSSR, error)
		GetStudentViewSSR(studentID, ssrID int) (*entity.StudentViewSsr, error)
		UpdateStatus(id int, newStatus entity.StatusSSR) (int, error)
	}

	IStudentBidUC interface {
		GetStudentBids(studentID int) (*dto.StudentBids, error)
		Apply(data *dto.ApplyBid) (*dto.ApplyBidResponse, error)
	}
	ISupervisorBidUC interface {
		GetSupervisorBids(supervisorID int) (*dto.SupervisorBids, error)
		Resolve(data *dto.ResolveBid) error
	}

	IStudentRelUC interface {
		Create(data *dto.CreateSSR) (*dto.StudentViewSSR, error)
	}

	IWorkRepo interface {
		GetWorksByStudentID(studentID int) ([]*entity.Work, error)
		GetWorksBySupervisorID(supervisorID int) ([]*entity.WorkOfSupervisor, error)
		GetSupervisorsByWorkID(workID int) ([]*entity.SupervisorOfWork, error)
	}
	IStudentWorkUC interface {
		GetStudentWorks(studentID int) (*dto.StudentWorkPlenty, error)
		GetWorkSupervisors(workID int) (*dto.WorkSupervisorPlenty, error)
	}
	ISupervisorWorkUC interface {
		GetSupervisorWorks(supervisorID int) (*dto.SupervisorWorkPlenty, error)
	}
)
